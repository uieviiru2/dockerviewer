package index

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"

	"github.com/Equanox/gotron"

	"github.com/uieviiru2/mylib/docker"
	"github.com/uieviiru2/mylib/jsonedit"
	"github.com/uieviiru2/mylib/socket/config"
	"github.com/uieviiru2/mylib/vultr"
)

type Explorer struct {
	Path string `json:"path"`
}
type Reload struct {
	MachineName string `json:"machine_name"`
}
type PathList struct {
	Path1 string `json:"path1"`
	Path2 string `json:"path2"`
}
type Url struct {
	Url string `json:"url"`
}
type Page struct {
	Page     string `json:"page"`
	ViewData string `json:"view_data"`
}

func Use(window *gotron.BrowserWindow) {

	window.On(&gotron.Event{Event: "index"}, func(bin []byte) {
		//設定ファイルの存在チェック
		userData, _ := user.Current()
		if _, err := os.Stat(userData.HomeDir + "/.docker.bin"); os.IsNotExist(err) {
			fmt.Println("can't open .docker.bin")
			erroutput := jsonedit.Val("eventName", "index-config")
			window.Send(&gotron.Event{Event: jsonedit.End(erroutput)})
			return
		}

		configData := config.LoadConfig()
		output := jsonedit.Val("eventName", "index")

		path1 := configData.ConfigDir + "/Dockerfile"
		path2 := configData.ConfigDir + "/docker-compose"
		var pl PathList
		pl.Path1 = path1
		pl.Path2 = path2
		bytes, _ := json.Marshal(pl)
		output = jsonedit.Con(output, jsonedit.StripQ(string(bytes)))

		jMachine := docker.Machine(window)
		if jMachine != "" {
			output = jsonedit.Con(output, jMachine)
		}

		jDocker := docker.GetPs(window)
		if jDocker != "" {
			output = jsonedit.Con(output, jDocker)
		}

		jDockerImage := docker.Image(window)
		if jDockerImage != "" {
			output = jsonedit.Con(output, jDockerImage)
		}

		jServerfiles := docker.Serverfiles(window)
		output = jsonedit.Con(output, jServerfiles)

		jVultr := vultr.List(window)
		output = jsonedit.Con(output, jVultr)

		jDockerfiles := docker.Dockerfiles(window)
		output = jsonedit.Con(output, jDockerfiles)

		jDockercompose := docker.DockerCompose(window)
		output = jsonedit.Con(output, jDockercompose)

		configDataByte, _ := json.Marshal(configData)
		configDataStr := jsonedit.On("configData", string(configDataByte))
		output = jsonedit.Con(output, configDataStr)

		fmt.Println("■■■output■■■")
		fmt.Println(output)

		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "index-macinereload"}, func(bin []byte) {
		var d Reload
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		docker.ReloadMachine(d.MachineName, window)
	})
	window.On(&gotron.Event{Event: "index-machinessh"}, func(bin []byte) {
		var d Reload
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		docker.OpenMachineSsh("docker-machine ssh " + d.MachineName)
	})
	window.On(&gotron.Event{Event: "index-machinestart"}, func(bin []byte) {
		var d Reload
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		docker.StartDocker()
	})
	window.On(&gotron.Event{Event: "index-machinecreate"}, func(bin []byte) {
		var d Reload
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		docker.CreateDocker(window)
	})
	window.On(&gotron.Event{Event: "index-webbrowser"}, func(bin []byte) {
		var d Url
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", d.Url).Start()
	})
	window.On(&gotron.Event{Event: "explorer"}, func(bin []byte) {
		var d Explorer
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		docker.OpenExplorer(d.Path)
	})
	window.On(&gotron.Event{Event: "index-view"}, func(bin []byte) {
		var d Page
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println("./webapp/" + d.Page + ".html")
		f, err := os.Open("./webapp/" + d.Page + ".html")
		if err != nil {
			fmt.Println("error")
		}
		defer f.Close()
		viewData, err := ioutil.ReadAll(f)
		d.ViewData = string(viewData)
		jsonData, _ := json.Marshal(d)

		output := jsonedit.Val("eventName", "index-view")
		output = jsonedit.Con(output, jsonedit.On("view", string(jsonData)))
		fmt.Println(output)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})

	})
}
