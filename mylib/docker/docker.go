package docker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"runtime"
	"strings"
	"sync"

	"golang.org/x/crypto/ssh"

	"github.com/Equanox/gotron"
	"github.com/uieviiru2/mylib/jsonedit"
	"github.com/uieviiru2/mylib/socket/config"
)

func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}

func DockerMachineIp(name string) string {
	output := ExecMachine("ip", name)
	output = strings.Replace(output, "\r", "", -1)
	output = strings.Replace(output, "\n", "", -1)
	return output
}

func Dockerfiles(window *gotron.BrowserWindow) string {
	var fileNameJson = ""
	configData := config.LoadConfig()
	if configData.ConfigDir+"/Dockerfile" == "" {
		return fileNameJson
	}
	files, err := ioutil.ReadDir(configData.ConfigDir + "/dockerfile")
	if err != nil {
		OutLog("Config error ->rm .docker.bin ", window)
		panic(err)
	}
	firstFlg := true
	for _, file := range files {

		tmpFileName := "{" + jsonedit.Val("name", file.Name()) + "," + jsonedit.Val("created_at", file.ModTime().String()) + "}"
		if !firstFlg {
			fileNameJson = jsonedit.Con(fileNameJson, tmpFileName)
		} else {
			fileNameJson = tmpFileName
			firstFlg = false
		}
	}
	OutLog("■■■Dockerfile■■■", window)
	OutLog(fileNameJson, window)

	return jsonedit.List("dockerfiles", fileNameJson)
}

func SaveDockerfile(name, script string, window *gotron.BrowserWindow) string {
	script = strings.Replace(script, "\r\n", "\n", -1)
	script = strings.Replace(script, "\r", "\n", -1)
	configData := config.LoadConfig()
	fmt.Println(configData.ConfigDir + "/Dockerfile/" + name)
	if err := os.MkdirAll(configData.ConfigDir+"/Dockerfile/"+name, 0777); err != nil {
		fmt.Println(err)
	}
	file, err := os.Create(configData.ConfigDir + "/Dockerfile/" + name + "/Dockerfile")
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()

	file.Write(([]byte)(script))
	OutLog("■■■Save Dockerfile■■■", window)
	OutLog(name, window)

	return "OK"
}
func DeleteDockerfile(name string, window *gotron.BrowserWindow) {
	configData := config.LoadConfig()
	if err := os.RemoveAll(configData.ConfigDir + "/Dockerfile/" + name); err != nil {
		fmt.Println(err)
	}
	OutLog("■■■Delete Dockerfile■■■", window)
	OutLog(name, window)
}
func LoadDockerfile(name string, window *gotron.BrowserWindow) string {
	configData := config.LoadConfig()
	fmt.Println(configData.ConfigDir + "/Dockerfile/" + name + "/Dockerfile")
	file, err := os.Open(configData.ConfigDir + "/Dockerfile/" + name + "/Dockerfile")
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()

	buf := make([]byte, 1024)
	out := ""
	for {
		// nはバイト数を示す
		n, err := file.Read(buf)
		// バイト数が0になることは、読み取り終了を示す
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		// バイト型スライスを文字列型に変換してファイルの内容を出力
		out += string(buf[:n])
	}
	OutLog("■■■Load Dockerfile■■■", window)
	OutLog(name, window)
	return out
}
func BuildDockerfiles(name string, window *gotron.BrowserWindow) {
	configData := config.LoadConfig()
	OutLog("■■■INPUT■■■", window)
	OutLog("docker build -t "+name+" "+configData.ConfigDir+"//Dockerfile//"+name, window)
	///p, _ := os.Getwd()
	//os.Chdir(configData.ConfigDir + "//Dockerfile//" + name)
	//output, err := ExecCommand2("build", "-t", name, ".")
	output, err := ExecCommand2("build", "-t", name, configData.ConfigDir+"//Dockerfile//"+name)
	//os.Chdir(p) //重要カレントディレクトリ変更後、戻らないとhtmlファイルが参照できなくなる
	if err != "" {
		OutLog("■■■OUTPUT ERROR■■■", window)
		OutLog(err, window)
		return
	}
	OutLog("■■■OUTPUT■■■", window)
	OutLog(output, window)
	return
}
func UpDockerCompose(name string, window *gotron.BrowserWindow) {
	configData := config.LoadConfig()
	//os.Chdir(configData.ConfigDir + "//docker-compose//" + name)
	OutLog("■■■INPUT■■■", window)
	OutLog("docker-compose -f "+configData.ConfigDir+"//docker-compose//"+name+"//docker-compose.yml up -d", window)
	out, err := ExecCompose("-f", configData.ConfigDir+"//docker-compose//"+name+"//docker-compose.yml", "up", "-d")
	if err != "" {
		OutLog("■■■OUTPUT ERROR■■■", window)
		OutLog(err, window)
		return
	}
	OutLog("■■■OUTPUT■■■", window)
	OutLog(out, window)
	return
}
func DockerCompose(window *gotron.BrowserWindow) string {
	var fileNameJson = ""
	configData := config.LoadConfig()
	if configData.ConfigDir+"/docker-compose" == "" {
		return fileNameJson
	}
	files, err := ioutil.ReadDir(configData.ConfigDir + "/docker-compose")
	if err != nil {
		OutLog("Config error ->rm .docker.bin ", window)
		panic(err)
	}
	firstFlg := true
	for _, file := range files {

		tmpFileName := "{" + jsonedit.Val("name", file.Name()) + "," + jsonedit.Val("created_at", file.ModTime().String()) + "}"
		if !firstFlg {
			fileNameJson = jsonedit.Con(fileNameJson, tmpFileName)
		} else {
			fileNameJson = tmpFileName
			firstFlg = false
		}
	}
	OutLog("■■■DockerCompose■■■", window)
	OutLog(fileNameJson, window)
	return jsonedit.List("dockercompose", fileNameJson)
}

func SaveDockerCompose(name, script string) string {
	script = strings.Replace(script, "\r\n", "\n", -1)
	script = strings.Replace(script, "\r", "\n", -1)
	configData := config.LoadConfig()
	fmt.Println(configData.ConfigDir + "/docker-compose/" + name)
	if err := os.MkdirAll(configData.ConfigDir+"/docker-compose/"+name, 0777); err != nil {
		fmt.Println(err)
	}
	file, err := os.Create(configData.ConfigDir + "/docker-compose/" + name + "/docker-compose.yml")
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()

	file.Write(([]byte)(script))
	return "OK"
}
func DeleteDockerCompose(name string) {
	configData := config.LoadConfig()
	if err := os.RemoveAll(configData.ConfigDir + "/docker-compose/" + name); err != nil {
		fmt.Println(err)
	}
}
func LoadDockerCompose(name string) string {
	configData := config.LoadConfig()
	fmt.Println(configData.ConfigDir + "/docker-compose/" + name + "/docker-compose.yml")
	file, err := os.Open(configData.ConfigDir + "/docker-compose/" + name + "/docker-compose.yml")
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()

	buf := make([]byte, 1024)
	out := ""
	for {
		// nはバイト数を示す
		n, err := file.Read(buf)
		// バイト数が0になることは、読み取り終了を示す
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		// バイト型スライスを文字列型に変換してファイルの内容を出力
		out += string(buf[:n])
	}
	return out
}
func Serverfiles(window *gotron.BrowserWindow) string {
	var serverListJson = ""
	configData := config.LoadConfig()
	if configData.ConfigDir == "" {
		return serverListJson
	}

	files, err := ioutil.ReadDir(configData.ConfigDir + "/server")
	if err != nil {
		OutLog("Config error ->rm .docker.bin ", window)
		panic(err)
	}
	firstFlg := true

	for _, file := range files {

		if !firstFlg {
			serverListJson += "," + FileRead(configData.ConfigDir+"/server/"+file.Name())
		} else {
			serverListJson += FileRead(configData.ConfigDir + "/server/" + file.Name())
			firstFlg = false
		}
	}
	OutLog("■■■Server list■■■", window)
	OutLog(serverListJson, window)
	return jsonedit.List("serverlist", serverListJson)
}
func RunInputfiles() string {
	var serverListJson = ""
	configData := config.LoadConfig()
	if configData.ConfigDir == "" {
		return serverListJson
	}
	files, err := ioutil.ReadDir(configData.ConfigDir + "/runinput")
	if err != nil {
		panic(err)
	}
	firstFlg := true

	for _, file := range files {

		if !firstFlg {
			serverListJson += "," + FileRead(configData.ConfigDir+"/runinput/"+file.Name())
		} else {
			serverListJson += FileRead(configData.ConfigDir + "/runinput/" + file.Name())
			firstFlg = false
		}
	}

	return jsonedit.List("runInputlist", serverListJson)
}
func DeployInputfiles() string {
	var serverListJson = ""
	configData := config.LoadConfig()
	if configData.ConfigDir == "" {
		return serverListJson
	}
	files, err := ioutil.ReadDir(configData.ConfigDir + "/deployinput")
	if err != nil {
		panic(err)
	}
	firstFlg := true

	for _, file := range files {

		if !firstFlg {
			serverListJson += "," + FileRead(configData.ConfigDir+"/deployinput/"+file.Name())
		} else {
			serverListJson += FileRead(configData.ConfigDir + "/deployinput/" + file.Name())
			firstFlg = false
		}
	}

	return jsonedit.List("deployInputlist", serverListJson)
}
func FileRead(path string) string {
	fmt.Println("ファイル読み取り処理を開始します")
	// ファイルをOpenする
	f, err := os.Open(path)
	// 読み取り時の例外処理
	if err != nil {
		fmt.Println("error")
	}
	// 関数が終了した際に確実に閉じるようにする
	defer f.Close()

	// バイト型スライスの作成
	buf := make([]byte, 1024)
	out := ""
	for {
		// nはバイト数を示す
		n, err := f.Read(buf)
		// バイト数が0になることは、読み取り終了を示す
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		// バイト型スライスを文字列型に変換してファイルの内容を出力
		out += string(buf[:n])
	}
	return out
}

func SaveServerfile(name, json string, window *gotron.BrowserWindow) {
	configData := config.LoadConfig()

	if f, err := os.Stat(configData.ConfigDir + "/server"); os.IsNotExist(err) || !f.IsDir() {
		os.Mkdir(configData.ConfigDir+"/server", 0777)
	}
	fmt.Println(configData.ConfigDir + "/server/" + name)
	file, err := os.Create(configData.ConfigDir + "/server/" + name)
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()
	file.Write(([]byte)(json))
	OutLog("Save Server File "+name+"\n"+json, window)
}
func DeleteServerfile(name string) {
	configData := config.LoadConfig()
	if err := os.Remove(configData.ConfigDir + "/server/" + name); err != nil {
		fmt.Println(err)
	}
}
func LoadServerfile(name string) string {
	configData := config.LoadConfig()
	file, err := os.Open(configData.ConfigDir + "/server/" + name)
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()

	buf := make([]byte, 1024)
	out := ""
	for {
		// nはバイト数を示す
		n, err := file.Read(buf)
		// バイト数が0になることは、読み取り終了を示す
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		// バイト型スライスを文字列型に変換してファイルの内容を出力
		out += string(buf[:n])
	}
	return out
}
func SaveRunInputfile(name, json string, window *gotron.BrowserWindow) {
	configData := config.LoadConfig()

	if f, err := os.Stat(configData.ConfigDir + "/runinput"); os.IsNotExist(err) || !f.IsDir() {
		os.Mkdir(configData.ConfigDir+"/runinput", 0777)
	}
	fmt.Println(configData.ConfigDir + "/runinput/" + name)
	file, err := os.Create(configData.ConfigDir + "/runinput/" + name)
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()
	file.Write(([]byte)(json))
	OutLog("■■■Save Run Input File■■■", window)
	OutLog(name, window)
}
func DeleteRunInputfile(name string, window *gotron.BrowserWindow) {
	configData := config.LoadConfig()
	if err := os.Remove(configData.ConfigDir + "/runinput/" + name); err != nil {
		fmt.Println(err)
	}
	OutLog("■■■Delete Run Input File■■■", window)
	OutLog(name, window)
}
func LoadRunInputfile(name string) string {
	configData := config.LoadConfig()
	fmt.Println(configData.ConfigDir + "/runinput/" + name)
	file, err := os.Open(configData.ConfigDir + "/runinput/" + name)
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()

	buf := make([]byte, 1024)
	out := ""
	for {
		// nはバイト数を示す
		n, err := file.Read(buf)
		// バイト数が0になることは、読み取り終了を示す
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		// バイト型スライスを文字列型に変換してファイルの内容を出力
		out += string(buf[:n])
	}
	return out
}
func SaveDeployInputfile(name, json string, window *gotron.BrowserWindow) {
	configData := config.LoadConfig()

	if f, err := os.Stat(configData.ConfigDir + "/deployinput"); os.IsNotExist(err) || !f.IsDir() {
		os.Mkdir(configData.ConfigDir+"/deployinput", 0777)
	}
	fmt.Println(configData.ConfigDir + "/deployinput/" + name)
	file, err := os.Create(configData.ConfigDir + "/deployinput/" + name)
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()
	file.Write(([]byte)(json))
	OutLog("■■■Save Deploy Input File■■■", window)
	OutLog(name, window)
}
func LoadDeployInputfile(name string) string {
	configData := config.LoadConfig()
	fmt.Println(configData.ConfigDir + "/deployinput/" + name)
	file, err := os.Open(configData.ConfigDir + "/deployinput/" + name)
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()

	buf := make([]byte, 1024)
	out := ""
	for {
		// nはバイト数を示す
		n, err := file.Read(buf)
		// バイト数が0になることは、読み取り終了を示す
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		// バイト型スライスを文字列型に変換してファイルの内容を出力
		out += string(buf[:n])
	}
	return out
}
func DeleteDeployInputfile(name string, window *gotron.BrowserWindow) {
	configData := config.LoadConfig()
	if err := os.Remove(configData.ConfigDir + "deployinput/" + name); err != nil {
		fmt.Println(err)
	}
	OutLog("■■■Delete Deploy Input File■■■", window)
	OutLog(name, window)
}
func SaveImage(containerID, imageName, message string, window *gotron.BrowserWindow) {
	OutLog("■■■INPUT■■■", window)
	OutLog("docker commit -m "+containerID+" "+imageName, window)
	output := ExecCommand("commit", "-m", "\""+message+"\"", containerID, imageName)
	OutLog("■■■OUTPUT■■■", window)
	OutLog(output, window)
}

func StartContainer(containerID string, window *gotron.BrowserWindow) {
	OutLog("■■■INPUT■■■", window)
	OutLog("docker start "+containerID, window)
	output := ExecCommand("start", containerID)
	OutLog("■■■OUTPUT■■■", window)
	OutLog(output, window)
}
func StopContainer(containerID string, window *gotron.BrowserWindow) {
	OutLog("■■■INPUT■■■", window)
	OutLog("docker stop "+containerID, window)
	output := ExecCommand("stop", containerID)
	OutLog("■■■OUTPUT■■■", window)
	OutLog(output, window)
}
func Pull2(pullName string, window *gotron.BrowserWindow) {
	OutLog("■■■INPUT■■■", window)
	OutLog("docker pull "+pullName, window)
	output := ExecCommand("pull", pullName)
	OutLog("■■■OUTPUT■■■", window)
	OutLog(output, window)
}
func MakePem() {
	output := ExecMachine("ssh", "default;ssh-keygen", "-t", "rsa;\n;\n;\n;\n;\n;\n")
	fmt.Println(output)
}

func Machine(window *gotron.BrowserWindow) string {
	OutLog("■■■INPUT■■■", window)
	OutLog("docker-machine ls --format '{\"DOCKER\":\"{{.DockerVersion}}\",\"RESPONSE\":\"{{.ResponseTime}}\",\"NAME\":\"{{.Name}}\",\"ACTIVE_SWARM\":\"{{.ActiveSwarm}}\",\"STATE\":\"{{.State}}\",\"SWARM_OPTIONS\":\"{{.SwarmOptions}}\",\"ERRORS\":\"{{.Error}}\",\"ACTIVE\":\"{{.Active}}\",\"ACTIVE_HOST\":\"{{.ActiveHost}}\",\"DRIVER\":\"{{.DriverName}}\",\"URL\":\"{{.URL}}\"}'", window)
	output, err := ExecMachine2("ls", "--format", "'{\"DOCKER\":\"{{.DockerVersion}}\",\"RESPONSE\":\"{{.ResponseTime}}\",\"NAME\":\"{{.Name}}\",\"ACTIVE_SWARM\":\"{{.ActiveSwarm}}\",\"STATE\":\"{{.State}}\",\"SWARM_OPTIONS\":\"{{.SwarmOptions}}\",\"ERRORS\":\"{{.Error}}\",\"ACTIVE\":\"{{.Active}}\",\"ACTIVE_HOST\":\"{{.ActiveHost}}\",\"DRIVER\":\"{{.DriverName}}\",\"URL\":\"{{.URL}}\"}'")
	if err != "" {
		OutLog("■■■OUTPUT ERROR■■■", window)
		OutLog(err, window)
		return ""
	}
	OutLog("■■■OUTPUT■■■", window)
	OutLog(output, window)
	output = jsonedit.Split2(output, "\r\n|\n\r|\n|\r", "machine")
	return output
}
func ReloadMachine(name string, window *gotron.BrowserWindow) {
	OutLog("■■■INPUT■■■", window)
	OutLog("docker-machine restart", window)
	output := ExecMachine("restart", name)
	OutLog("■■■OUTPUT■■■", window)
	OutLog(output, window)
}
func GetPs(window *gotron.BrowserWindow) string {
	OutLog("■■■INPUT■■■", window)
	OutLog("docker ps -a --format \"{{json . }}\"", window)
	output, err := ExecCommand2("ps", "-a", "--format", "\"{{json . }}\"")
	if err != "" {
		OutLog("■■■OUTPUT ERROR■■■", window)
		OutLog(err, window)
		return ""
	}
	OutLog("■■■OUTPUT■■■", window)
	OutLog(output, window)
	output = jsonedit.Split2(output, "\r\n|\n\r|\n|\r", "ps")

	return output
}

func GetServerPs(ip, user, key string, window *gotron.BrowserWindow) string {
	fmt.Println(user + "@" + ip + ":" + key)
	output := GOGetOutput(user, key, ip, "docker ps -a --format \"{{json . }}\"", window)
	fmt.Println("GetServerPs")
	fmt.Println(output)
	output = jsonedit.Split(output, "\r\n|\n\r|\n|\r", "ps")

	return string(output)
}

func Image(window *gotron.BrowserWindow) string {
	OutLog("■■■INPUT■■■", window)
	OutLog("docker image list --format \"{{json . }}\"", window)
	output, err := ExecCommand2("image", "list", "--format", "\"{{json . }}\"")
	if err != "" {
		OutLog("■■■OUTPUT ERROR■■■", window)
		OutLog(err, window)
		return ""
	}
	OutLog("■■■OUTPUT■■■", window)
	OutLog(output, window)
	output = jsonedit.Split2(output, "\r\n|\n\r|\n|\r", "image")
	fmt.Println(output)
	return output
}
func ServerImage(ip, user, key string, window *gotron.BrowserWindow) string {

	output := GOGetOutput(user, key, ip, "docker image list --format \"{{json . }}\"", window)
	fmt.Println("ServerImage")
	fmt.Println(output)
	output = jsonedit.Split(output, "\r\n|\n\r|\n|\r", "image")

	return string(output)
}
func Inspect(id string) string {

	output := ExecCommand("inspect", id)
	output = jsonedit.On("inspect", output)
	fmt.Println(output)
	return output
}
func ServerInspect(ip, user, key, id string, window *gotron.BrowserWindow) string {

	output := GOGetOutput(user, key, ip, "docker inspect "+id, window)
	output = jsonedit.On("inspect", string(output))
	fmt.Println(output)
	return output
}
func ImageInspect(id string) string {
	output := ExecCommand("image", "inspect", id)
	output = jsonedit.On("inspect", output)
	fmt.Println(output)
	return output
}
func ServerImageInspect(ip, user, key, id string, window *gotron.BrowserWindow) string {
	output := GOGetOutput(user, key, ip, "docker image inspect "+id, window)
	output = jsonedit.On("inspect", output)
	fmt.Println(output)
	return output
}
func Deploy(id, name, dit, pem, user, serverip, port, dirname, dirname2, dirnameA, dirnameA2, dirnameB, dirnameB2, dirnameC, dirnameC2, dirnameD, dirnameD2, option, option2 string, window *gotron.BrowserWindow) {
	//専用ディレクトリ作成
	Go(user, pem, serverip, "mkdir -p ~/dockerconfig", window)

	configData := config.LoadConfig()
	if err := os.MkdirAll(configData.ConfigDir+"/tmp/"+id, 0777); err != nil {
		OutLog(err.Error(), window)
		return
	}
	//ローカル作業
	OutLog("docker commit "+id+" "+id, window)
	OutLog(ExecCommand("commit", id, id), window)

	OutLog("docker save "+id+" io "+configData.ConfigDir+"/tmp/"+id+"/"+id, window)
	OutLog(ExecCommand("save", id, "-o", configData.ConfigDir+"/tmp/"+id+"/"+id), window)

	//リモート作業
	OutLog(ScpCommand("-oStrictHostKeyChecking=no", "-i", pem, "-r", configData.ConfigDir+"/tmp/"+id, user+"@"+serverip+":~/dockerconfig/"), window)

	cmd := "cd ~/dockerconfig/" + id
	cmd += "\ndocker load < ./" + id
	cmd += "\ndocker run " + dit + " --name " + name
	if dirname != "" {
		OutLog(ScpCommand("-oStrictHostKeyChecking=no", "-i", pem, "-r", dirname, user+"@"+serverip+":~/dockerconfig/"+id+"/"+id+"_data1"), window)
		cmd += " -v ~/dockerconfig/" + id + "/" + id + "_data1:/" + dirname2
	}
	if dirnameA != "" {
		OutLog(ScpCommand("-oStrictHostKeyChecking=no", "-i", pem, "-r", dirnameA, user+"@"+serverip+":~/dockerconfig/"+id+"/"+id+"_data2"), window)
		cmd += " -v ~/dockerconfig/" + id + "/" + id + "_data2:/" + dirnameA2
	}
	if dirnameB != "" {
		OutLog(ScpCommand("-oStrictHostKeyChecking=no", "-i", pem, "-r", dirnameB, user+"@"+serverip+":~/dockerconfig/"+id+"/"+id+"_data3"), window)
		cmd += " -v ~/dockerconfig/" + id + "/" + id + "_data3:/" + dirnameB2
	}
	if dirnameC != "" {
		OutLog(ScpCommand("-oStrictHostKeyChecking=no", "-i", pem, "-r", dirnameC, user+"@"+serverip+":~/dockerconfig/"+id+"/"+id+"_data4"), window)
		cmd += " -v ~/dockerconfig/" + id + "/" + id + "_data4:/" + dirnameC2
	}
	if dirnameD != "" {
		OutLog(ScpCommand("-oStrictHostKeyChecking=no", "-i", pem, "-r", dirnameD, user+"@"+serverip+":~/dockerconfig/"+id+"/"+id+"_data5"), window)
		cmd += " -v ~/dockerconfig/" + id + "/" + id + "_data5:/" + dirnameD2
	}
	if port != "" {
		cmd += " -p " + port
	}
	if option != "" {
		cmd += " " + option
	}

	cmd += " " + id
	if option2 != "" {
		cmd += " " + option2
	}

	cmd += "\ndocker start " + name
	fmt.Println(cmd)
	Go(user, pem, serverip, cmd, window)
}
func ServerDeploy(id, name, dit, pem, user, serverip, port, dirname, dirname2, dirnameA, dirnameA2, dirnameB, dirnameB2, dirnameC, dirnameC2, dirnameD, dirnameD2, option, option2 string, window *gotron.BrowserWindow) {
	//専用ディレクトリ作成
	Go(user, pem, serverip, "mkdir -p ~/dockerconfig", window)

	configData := config.LoadConfig()
	if err := os.MkdirAll(configData.ConfigDir+"/tmp/"+id, 0777); err != nil {
		OutLog("■■■OUTPUT ERROR■■■", window)
		OutLog(err.Error(), window)
		return
	}
	OutLog("■■■INPUT■■■", window)
	OutLog("docker save "+id+" io "+configData.ConfigDir+"/tmp/"+id+"/"+id, window)
	OutLog("■■■OUTPUT■■■", window)
	OutLog(ExecCommand("save", id, "-o", configData.ConfigDir+"/tmp/"+id+"/"+id), window)

	//リモート作業
	OutLog("■■■INPUT■■■", window)
	OutLog("scp -oStrictHostKeyChecking=no -i "+pem+" -r "+configData.ConfigDir+"/tmp/"+id+" "+user+"@"+serverip+":~/dockerconfig/", window)
	OutLog("■■■OUTPUT■■■", window)
	OutLog(ScpCommand("-oStrictHostKeyChecking=no", "-i", pem, "-r", configData.ConfigDir+"/tmp/"+id, user+"@"+serverip+":~/dockerconfig/"), window)

	cmd := "cd ~/dockerconfig/" + id
	cmd += "\ndocker load < ./" + id
	cmd += "\ndocker run " + dit + " --name " + name
	if dirname != "" {
		OutLog("■■■INPUT■■■", window)
		OutLog("scp -oStrictHostKeyChecking=no -i "+pem+" -r "+dirname+" "+user+"@"+serverip+":~/dockerconfig/"+id+"/"+id+"_data1", window)
		OutLog("■■■OUTPUT■■■", window)
		OutLog(ScpCommand("-oStrictHostKeyChecking=no", "-i", pem, "-r", dirname, user+"@"+serverip+":~/dockerconfig/"+id+"/"+id+"_data1"), window)
		cmd += " -v ~/dockerconfig/" + id + "/" + id + "_data1:/" + dirname2
	}
	if dirnameA != "" {
		OutLog("■■■INPUT■■■", window)
		OutLog("scp -oStrictHostKeyChecking=no -i "+pem+" -r "+dirnameA+" "+user+"@"+serverip+":~/dockerconfig/"+id+"/"+id+"_data1", window)
		OutLog("■■■OUTPUT■■■", window)
		OutLog(ScpCommand("-oStrictHostKeyChecking=no", "-i", pem, "-r", dirnameA, user+"@"+serverip+":~/dockerconfig/"+id+"/"+id+"_data2"), window)
		cmd += " -v ~/dockerconfig/" + id + "/" + id + "_data2:/" + dirnameA2
	}
	if dirnameB != "" {
		OutLog(ScpCommand("-oStrictHostKeyChecking=no", "-i", pem, "-r", dirnameB, user+"@"+serverip+":~/dockerconfig/"+id+"/"+id+"_data3"), window)
		cmd += " -v ~/dockerconfig/" + id + "/" + id + "_data3:/" + dirnameB2
	}
	if dirnameC != "" {
		OutLog(ScpCommand("-oStrictHostKeyChecking=no", "-i", pem, "-r", dirnameC, user+"@"+serverip+":~/dockerconfig/"+id+"/"+id+"_data4"), window)
		cmd += " -v ~/dockerconfig/" + id + "/" + id + "_data4:/" + dirnameC2
	}
	if dirnameD != "" {
		OutLog(ScpCommand("-oStrictHostKeyChecking=no", "-i", pem, "-r", dirnameD, user+"@"+serverip+":~/dockerconfig/"+id+"/"+id+"_data5"), window)
		cmd += " -v ~/dockerconfig/" + id + "/" + id + "_data5:/" + dirnameD2
	}
	if port != "" {
		cmd += " -p " + port
	}
	if option != "" {
		cmd += " " + option
	}

	cmd += " " + id
	if option2 != "" {
		cmd += " " + option2
	}

	cmd += "\ndocker start " + name
	fmt.Println(cmd)
	Go(user, pem, serverip, cmd, window)

}
func RegistryDeploy(id, name, pem, user, serverip, port string, window *gotron.BrowserWindow) {
	//専用ディレクトリ作成
	Go(user, pem, serverip, "mkdir -p ~/dockerconfig", window)
	//リモート作業
	cmd := "\ndocker run -d --name " + name + " -p " + port + " " + id
	Go(user, pem, serverip, cmd, window)

}
func SaveRegistryfile(name, json string, window *gotron.BrowserWindow) {
	configData := config.LoadConfig()

	if f, err := os.Stat(configData.ConfigDir + "/registry"); os.IsNotExist(err) || !f.IsDir() {
		os.Mkdir(configData.ConfigDir+"/registry", 0777)
	}
	fmt.Println(configData.ConfigDir + "/registry/" + name)
	file, err := os.Create(configData.ConfigDir + "/registry/" + name)
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()
	file.Write(([]byte)(json))
	OutLog("■■■Save Registry File■■■", window)
	OutLog(name, window)
}
func Registryfiles(window *gotron.BrowserWindow) string {
	registryListJson := ""
	configData := config.LoadConfig()
	if configData.ConfigDir == "" {
		return registryListJson
	}
	files, err := ioutil.ReadDir(configData.ConfigDir + "/registry")
	if err != nil {
		panic(err)
	}
	firstFlg := true

	for _, file := range files {
		if !firstFlg {
			registryListJson += "," + FileRead(configData.ConfigDir+"/registry/"+file.Name())
		} else {
			registryListJson += FileRead(configData.ConfigDir + "/registry/" + file.Name())
			firstFlg = false
		}
	}
	OutLog("■■■Registry File■■■", window)
	OutLog(registryListJson, window)
	return jsonedit.List("registrylist", registryListJson)
}
func Remove(id string, window *gotron.BrowserWindow) {
	OutLog("■■■INPUT■■■", window)
	OutLog("docker stop "+id, window)
	output, err := ExecCommand2("stop", id)
	if err != "" {
		OutLog("■■■OUTPUT ERROR■■■", window)
		OutLog(err, window)
		return
	}
	OutLog("■■■OUTPUT■■■", window)
	OutLog(output, window)

	OutLog("■■■INPUT■■■", window)
	OutLog("docker rm "+id, window)
	output, err = ExecCommand2("rm", id)
	if err != "" {
		OutLog("■■■OUTPUT ERROR■■■", window)
		OutLog(err, window)
		return
	}
	OutLog("■■■OUTPUT■■■", window)
	OutLog(output, window)
}
func ServerRemove(ip, user, key, id string, window *gotron.BrowserWindow) string {

	output := GOGetOutput(user, key, ip, "docker stop "+id, window)
	output += GOGetOutput(user, key, ip, "docker rm "+id, window)

	fmt.Println(output)
	return output

}
func ImageRemove(id, force string, window *gotron.BrowserWindow) {
	OutLog("■■■INPUT■■■", window)
	OutLog("docker rmi "+id+" "+force, window)
	output, err := ExecCommand2("rmi", id, force)
	if err != "" {
		OutLog("■■■OUTPUT ERROR■■■", window)
		OutLog(err, window)
		return
	}
	OutLog("■■■OUTPUT■■■", window)
	OutLog(output, window)
}
func ServerImageRemove(ip, user, key, id, force string, window *gotron.BrowserWindow) string {
	output := GOGetOutput(user, key, ip, "docker rmi "+id+" "+force, window)
	fmt.Println(output)
	return output

}
func ServerRemoveall(ip, user, key string, window *gotron.BrowserWindow) {
	GOGetOutput(user, key, ip, "docker stop $(docker ps -q)", window)
	GOGetOutput(user, key, ip, "docker rm $(docker ps -qa)", window)
	GOGetOutput(user, key, ip, "docker images -aq | xargs docker rmi", window)
	GOGetOutput(user, key, ip, "rm -rf ~/dockerconfig/*", window) //may be can't delete ↑
}

func Search(search string, window *gotron.BrowserWindow) string {
	OutLog("■■■INPUT■■■", window)
	OutLog("docker search --format \"{{json . }}\"", window)
	output := ExecCommand("search", search, "--format", "\"{{json . }}\"")
	OutLog("■■■OUTPUT■■■", window)
	OutLog(output, window)
	output = jsonedit.Split2(output, "\r\n|\n\r|\n|\r", "search")
	fmt.Println(output)
	return output
}

/*
/c/Users/fg/source/aaa://usr/share/nginx/html 必ずこの形式じゃないとダメ /c/小文字
*/
// docker run -d --name simple2 -v /c/Users/fg/source://usr/share/nginx/html -p 8082:80 nginx
func Run(docker, name, dit, port, dirname, dirname2, dirnameA, dirnameA2, dirnameB, dirnameB2, dirnameC, dirnameC2, dirnameD, dirnameD2, option, option2 string, window *gotron.BrowserWindow) {

	cmd := "docker run " + dit + " --name " + name
	if dirname != "" {
		cmd += " -v " + ChangeDockerPath(dirname) + ":/" + dirname2
	}
	if dirnameA != "" {
		cmd += " -v " + ChangeDockerPath(dirnameA) + ":/" + dirnameA2
	}
	if dirnameB != "" {
		cmd += " -v " + ChangeDockerPath(dirnameB) + ":/" + dirnameB2
	}
	if dirnameC != "" {
		cmd += " -v " + ChangeDockerPath(dirnameC) + ":/" + dirnameC2
	}
	if dirnameD != "" {
		cmd += " -v " + ChangeDockerPath(dirnameD) + ":/" + dirnameD2
	}
	if port != "" {
		cmd += " -p " + port
	}
	if option != "" {
		cmd += " " + option
	}
	cmd += " " + docker

	if option2 != "" {
		cmd += " " + option2
	}

	if dit == "-i" {
		cmd += "\n" + "docker start " + name
	}
	ip := DockerMachineIp("default")
	userData, _ := user.Current()
	Go("docker", userData.HomeDir+"/.docker/machine/machines/default/id_rsa", ip, cmd, window)
}
func RunMac(docker, name, dit, port, dirname, dirname2, dirnameA, dirnameA2, dirnameB, dirnameB2, dirnameC, dirnameC2, dirnameD, dirnameD2, option, option2 string, window *gotron.BrowserWindow) {

	cmd := "run " + dit + " --name " + name
	if dirname != "" {
		cmd += " -v " + ChangeDockerPath(dirname) + ":/" + dirname2
	}
	if dirnameA != "" {
		cmd += " -v " + ChangeDockerPath(dirnameA) + ":/" + dirnameA2
	}
	if dirnameB != "" {
		cmd += " -v " + ChangeDockerPath(dirnameB) + ":/" + dirnameB2
	}
	if dirnameC != "" {
		cmd += " -v " + ChangeDockerPath(dirnameC) + ":/" + dirnameC2
	}
	if dirnameD != "" {
		cmd += " -v " + ChangeDockerPath(dirnameD) + ":/" + dirnameD2
	}
	if port != "" {
		cmd += " -p " + port
	}
	if option != "" {
		cmd += " " + option
	}
	cmd += " " + docker

	if option2 != "" {
		cmd += " " + option2
	}
	slice := strings.Split(cmd, " ")
	OutLog("■■■INPUT■■■", window)
	OutLog("docker "+cmd, window)
	output, err := ExecCommand2(slice...)
	if err != "" {
		OutLog("■■■OUTPUT ERROR■■■", window)
		OutLog(err, window)
		return
	}
	OutLog("■■■OUTPUT■■■", window)
	OutLog(output, window)

	if dit == "-i" {
		OutLog("■■■INPUT■■■", window)
		OutLog("docker start "+name, window)
		output2, err2 := ExecCommand2("start", name)
		if err2 != "" {
			OutLog("■■■OUTPUT ERROR■■■", window)
			OutLog(err2, window)
			return
		}
		OutLog("■■■OUTPUT■■■", window)
		OutLog(output2, window)
	}
}
func ExecCommand(option ...string) string {
	fmt.Println(option)
	configData := config.LoadConfig()

	cmdStr := configData.DockerExe + "/docker"
	if IsWindows() {
		cmdStr += ".exe"
	}
	cmd := exec.Command(cmdStr, option...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return fmt.Sprint(err) + ": " + stderr.String()
	}
	fmt.Println("Result: " + out.String())
	return out.String()
}
func ExecCommand2(option ...string) (string, string) {
	configData := config.LoadConfig()
	cmdStr := configData.DockerExe + "/docker"
	if IsWindows() {
		cmdStr += ".exe"
	}
	cmd := exec.Command(cmdStr, option...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", fmt.Sprint(err) + ": " + stderr.String()
	}
	return out.String(), ""
}
func ExecMachine(option ...string) string {
	fmt.Println(option)
	configData := config.LoadConfig()
	cmdStr := configData.DockerExe + "/docker-machine"
	if IsWindows() {
		cmdStr += ".exe"
	}
	cmd := exec.Command(cmdStr, option...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return fmt.Sprint(err) + ": " + stderr.String()
	}
	fmt.Println("Result: " + out.String())
	return out.String()

}
func ExecMachine2(option ...string) (string, string) {
	fmt.Println(option)
	configData := config.LoadConfig()
	cmdStr := configData.DockerExe + "/docker-machine"
	if IsWindows() {
		cmdStr += ".exe"
	}
	cmd := exec.Command(cmdStr, option...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return "", fmt.Sprint(err) + ": " + stderr.String()
	}
	fmt.Println("Result: " + out.String())
	return out.String(), ""

}
func ExecCompose(option ...string) (string, string) {
	fmt.Println(option)
	configData := config.LoadConfig()
	cmdStr := configData.DockerExe + "/docker-compose"
	if IsWindows() {
		cmdStr += ".exe"
	}
	cmd := exec.Command(cmdStr, option...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return "", fmt.Sprint(err) + ": " + stderr.String()
	}
	fmt.Println("Result: " + out.String())
	return out.String(), ""

}

func ScpCommand(option ...string) string {
	fmt.Println(option)
	out, err := exec.Command("scp", option...).Output()

	if err != nil {
		return fmt.Sprint(err)
	}
	fmt.Println("Result: " + string(out))
	return string(out)
}

func SshCommand(option ...string) {
	fmt.Println(option)
	exec.Command("ssh", option...).Start()
}

//WindowのpathをDockerように変更する
func ChangeDockerPath(path string) string {
	if IsWindows() {
		path = strings.Replace(path, "\\", "/", -1)
		path = strings.Replace(path, ":", "", 1)
		path = strings.ToLower(path[0:1]) + path[1:]
		path = "/" + path
		return path
	} else {
		return path
	}
}

///////////////////////////////////////////////////////////////////////////
//これ以降はシェル関連処理
//相手方のサーバーで
//docker load < 180d04b117c9
//docker run -d --name 180d04b117c9 -v /root/180d04b117c9_data://usr/share/nginx/html -p 80:80 180d04b117c9
func MuxShell(w io.Writer, r io.Reader) (chan<- string, <-chan string) {
	in := make(chan string, 1)
	out := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(1) //for the shell itself
	go func() {
		for cmd := range in {
			wg.Add(1)
			w.Write([]byte(cmd + "\n"))
			wg.Wait()
		}
	}()
	go func() {
		var (
			buf [65 * 1024]byte

			t int
		)
		for {
			n, err := r.Read(buf[t:])
			if err != nil {
				close(in)
				close(out)
				return
			}
			t += n
			if buf[t-2] == '$' && buf[t-1] == ' ' { //assuming the $PS1 == 'sh-4.3$ '
				out <- string(buf[:t])
				t = 0
				wg.Done()

			} else if buf[t-2] == '#' { //root の場合  [root@localhost ~]#
				//assuming the $PS1 == 'sh-4.3$ '
				out <- string(buf[:t])
				t = 0
				wg.Done()
			}
		}
	}()
	return in, out
}
func Go(user, keypath, ipaddress, scripts string, window *gotron.BrowserWindow) {
	fmt.Println("start")

	//鍵設定
	privateKeyBytes, err := ioutil.ReadFile(keypath)
	if err != nil {
		log.Fatal("Failed to load private key (" + keypath + ")")
	}
	privateKey, err := ssh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		log.Fatal("Failed to parse private key")
	}
	/////////////////////////////////
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			//ssh.Password("core"),
			ssh.PublicKeys(privateKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", ipaddress+":22", config)
	if err != nil {
		panic(err)
	}
	fmt.Println("connect")
	defer client.Close()
	session, err := client.NewSession()
	fmt.Println("1")
	if err != nil {
		fmt.Println("2")
		fmt.Println(err)
		log.Fatalf("unable to create session: %s", err)
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		fmt.Println("3")
		fmt.Println(err)
		log.Fatal(err)
	}
	w, err := session.StdinPipe()
	if err != nil {
		fmt.Println("4")
		fmt.Println(err)
		panic(err)
	}
	r, err := session.StdoutPipe()
	if err != nil {
		fmt.Println("5")
		fmt.Println(err)
		panic(err)
	}
	in, out := MuxShell(w, r)
	if err := session.Start("/bin/sh"); err != nil {
		fmt.Println("6")
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println("script start")
	fmt.Println(scripts)
	<-out //ignore the shell output
	for i, v := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(scripts, -1) {
		fmt.Println(i)
		fmt.Println(v)
		if strings.TrimSpace(v) != "" {
			OutLog("■■■INPUT■■■", window)
			OutLog(fmt.Sprintf(v+"\n"), window)
			in <- v
			OutLog("■■■OUTPUT■■■", window)
			OutLog(fmt.Sprintf("%s\n", <-out), window)
		}
	}
	in <- "exit"
	session.Wait()
}
func GOGetOutput(user, keypath, ipaddress, scripts string, window *gotron.BrowserWindow) string {

	fmt.Println("start")

	//鍵設定
	privateKeyBytes, err := ioutil.ReadFile(keypath)
	if err != nil {
		log.Fatal("Failed to load private key (" + keypath + ")")
	}
	privateKey, err := ssh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		log.Fatal("Failed to parse private key")
	}
	/////////////////////////////////
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			//ssh.Password("core"),
			ssh.PublicKeys(privateKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	// SSH connect.
	client, err := ssh.Dial("tcp", ipaddress+":22", config)
	if err != nil {
		//panic(err)
		return string(err.Error())
	}

	session, err := client.NewSession()
	fmt.Sprintln("%s", err)
	defer session.Close()

	sessStdOut, err := session.StdoutPipe()
	if err != nil {
		//panic(err)
		return string(err.Error())
	}
	//go io.Copy(os.Stdout, sessStdOut)
	sessStderr, err := session.StderrPipe()
	if err != nil {
		//panic(err)
		return string(err.Error())
	}
	go io.Copy(os.Stderr, sessStderr)

	//err = session.Run("docker inspect b1db06e28438") // eg., /usr/bin/whoami
	OutLog("■■■INPUT■■■", window)
	OutLog(scripts, window)
	err = session.Run(scripts) // eg., /usr/bin/whoami
	if err != nil {
		OutLog("■■■OUTPUT ERROR■■■", window)
		OutLog(err.Error(), window)
		return err.Error()
	}
	buffer, err := ioutil.ReadAll(sessStdOut)
	if err != nil {
		OutLog("■■■OUTPUT ERROR■■■", window)
		OutLog(err.Error(), window)
		return err.Error()
	}
	OutLog("■■■OUTPUT■■■", window)
	OutLog(string(buffer), window)
	return string(buffer)
}
func OutLog(logData string, window *gotron.BrowserWindow) {
	fmt.Println(logData)
	bytes, _ := json.Marshal(html.UnescapeString(logData))

	logString := string(bytes)

	output := jsonedit.Val("eventName", "log")
	output = jsonedit.Con(output, jsonedit.On("log", logString))
	window.Send(&gotron.Event{Event: jsonedit.End(output)})
	//window.Send(&gotron.Event{Event: jsonedit.End(output)})
}
func OpenExplorer(path string) {
	if IsWindows() {
		exec.Command("C:/Windows/EXPLORER.EXE", path).Output()
	} else {
		exec.Command("open", path).Output()
	}
}
func OpenDockerEnter(path string) {
	configData := config.LoadConfig()
	ioutil.WriteFile(configData.ConfigDir+"/tmp/tmp.vbs", []byte("Dim oShell\nSet oShell = WScript.CreateObject (\"WSCript.shell\")\noShell.run \"cmd /K "+path+"\",1,1\nSet oShell = Nothing"), 0777)

	exec.Command("powershell", "START", configData.ConfigDir+"/tmp/tmp.vbs").Output()

}
func OpenServerDockerEnter(path string) {
	configData := config.LoadConfig()
	ioutil.WriteFile(configData.ConfigDir+"/tmp/tmp.vbs", []byte("Dim oShell\nSet oShell = WScript.CreateObject (\"WSCript.shell\")\noShell.run \"cmd /K "+path+"\",1,1\nSet oShell = Nothing"), 0777)

	exec.Command("powershell", "START", configData.ConfigDir+"/tmp/tmp.vbs").Output()

}
func OpenServer(path string) {
	configData := config.LoadConfig()
	ioutil.WriteFile(configData.ConfigDir+"/tmp/tmp.vbs", []byte("Dim oShell\nSet oShell = WScript.CreateObject (\"WSCript.shell\")\noShell.run \"cmd /K "+path+"\",1,1\nSet oShell = Nothing"), 0777)

	exec.Command("powershell", "START", configData.ConfigDir+"/tmp/tmp.vbs").Output()

}
func OpenMachineSsh(path string) {
	configData := config.LoadConfig()
	ioutil.WriteFile(configData.ConfigDir+"/tmp/tmp.vbs", []byte("Dim oShell\nSet oShell = WScript.CreateObject (\"WSCript.shell\")\noShell.run \"cmd /K "+path+"\",1,1\nSet oShell = Nothing"), 0777)

	exec.Command("powershell", "START", configData.ConfigDir+"/tmp/tmp.vbs").Output()

}
func StartDocker() {
	configData := config.LoadConfig()
	ioutil.WriteFile(configData.ConfigDir+"/tmp/tmp.vbs", []byte("Dim oShell\nSet oShell = WScript.CreateObject (\"WSCript.shell\")\noShell.run \"\"\""+configData.DockerExe+"/start.sh\"\"\""), 0777)

	exec.Command("powershell", "START", configData.ConfigDir+"/tmp/tmp.vbs").Output()
}
func CreateDocker(window *gotron.BrowserWindow) {
	OutLog("■■■INPUT■■■", window)
	OutLog("docker-machine create --driver virtualbox default", window)
	output, err := ExecMachine2("create", "--driver", "virtualbox", "default")
	if err != "" {
		OutLog("■■■OUTPUT ERROR■■■", window)
		OutLog(err, window)
	}
	OutLog("■■■OUTPUT■■■", window)
	OutLog(output, window)
}
func RegistryTag() string {

	req, _ := http.NewRequest("GET", "https://registry.hub.docker.com/v1/repositories/registry/tags", nil)
	client := new(http.Client)
	resp, err := client.Do(req)
	fmt.Println(err)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	respStr := string(byteArray)
	output := jsonedit.On("tags", respStr) // htmlをstringで取得
	fmt.Println(output)
	return output
}
func DeleteRegistryfile(name string, window *gotron.BrowserWindow) {
	configData := config.LoadConfig()
	if err := os.RemoveAll(configData.ConfigDir + "/registry/" + name); err != nil {
		fmt.Println(err)
	}
	OutLog("■■■Delete Registryfile■■■", window)
	OutLog(name, window)
}
