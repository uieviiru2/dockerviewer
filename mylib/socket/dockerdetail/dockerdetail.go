package dockerdetail

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/uieviiru2/mylib/jsonedit"
	"github.com/uieviiru2/mylib/socket/config"
	"github.com/uieviiru2/mylib/socket/configserver"
	"github.com/uieviiru2/mylib/vultr"

	"github.com/Equanox/gotron"
	"github.com/uieviiru2/mylib/docker"
)

type dockerID struct {
	ID string `json:"id"`
}
type dockerDeploy struct {
	ID        string `json:"id"`
	ServerIP  string `json:"server_ip"`
	Name      string `json:"name"`
	Dit       string `json:"dit"`
	Port      string `json:"port"`
	DirName   string `json:"dirname"`
	DirName2  string `json:"dirname2"`
	DirNameA  string `json:"dirname_a"`
	DirNameA2 string `json:"dirname_a2"`
	DirNameB  string `json:"dirname_b"`
	DirNameB2 string `json:"dirname_b2"`
	DirNameC  string `json:"dirname_c"`
	DirNameC2 string `json:"dirname_c2"`
	DirNameD  string `json:"dirname_d"`
	DirNameD2 string `json:"dirname_d2"`
	Option    string `json:"option"`
	Option2   string `json:"option2"`
	Cmd       string `json:"cmd"`
}

func Use(window *gotron.BrowserWindow) {

	window.On(&gotron.Event{Event: "dockerdetail"}, func(bin []byte) {
		var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(string(b))
		fmt.Println(d.ID)

		output := jsonedit.Val("eventName", "dockerdetail")

		jDocker := docker.Inspect(d.ID)
		output = jsonedit.Con(output, jDocker)

		jServerfiles := docker.Serverfiles(window)
		output = jsonedit.Con(output, jServerfiles)

		jVultr := vultr.List(window)
		output = jsonedit.Con(output, jVultr)

		jRunInputfiles := docker.DeployInputfiles()
		output = jsonedit.Con(output, jRunInputfiles)

		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "dockerdetail-deploy"}, func(bin []byte) {
		//var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		var d dockerDeploy
		if err := json.Unmarshal(b, &d); err != nil {

		}
		configData := config.LoadConfig()
		jServer := docker.LoadServerfile(d.ServerIP)
		var d2 configserver.ConfigServer
		if err2 := json.Unmarshal([]byte(jServer), &d2); err2 != nil {
			//なければvultr
			d2.Ip = d.ServerIP
			d2.ServerPem = configData.VultrPem
			d2.User = "root"
		}
		docker.Deploy(d.ID, d.Name, d.Dit, d2.ServerPem, d2.User, d.ServerIP, d.Port, d.DirName, d.DirName2, d.DirNameA, d.DirNameA2, d.DirNameB, d.DirNameB2, d.DirNameC, d.DirNameC2, d.DirNameD, d.DirNameD2, d.Option, d.Option2, window)
		//fmt.Println(output)
		//window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "dockerdetail-remove"}, func(bin []byte) {
		//var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		var d dockerDeploy
		if err := json.Unmarshal(b, &d); err != nil {

		}

		docker.Remove(d.ID, window)
		//fmt.Println(output)
		//window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "dockerdetail-bash"}, func(bin []byte) {
		//var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		var d dockerDeploy
		if err := json.Unmarshal(b, &d); err != nil {

		}

		docker.OpenDockerEnter("docker exec -it " + d.ID + " bash")

		//fmt.Println(output)
		//window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "dockerdetail-saveinput"}, func(bin []byte) {
		var d dockerDeploy

		b := []byte(bin)
		//buf := bytes.NewBuffer(b)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(string(b))
		fmt.Println(d.ServerIP)
		fmt.Println(d.Name)
		fmt.Println(d.Port)
		fmt.Println(d.DirName)
		fmt.Println(d.DirName2)
		fmt.Println(d.DirNameA)
		fmt.Println(d.DirNameA2)
		fmt.Println(d.DirNameB)
		fmt.Println(d.DirNameB2)
		fmt.Println(d.DirNameC)
		fmt.Println(d.DirNameC2)
		fmt.Println(d.DirNameD)
		fmt.Println(d.DirNameD2)
		fmt.Println(d.Option)
		d.Cmd = makeRunStr(d)

		bytes, _ := json.Marshal(d)
		docker.SaveDeployInputfile(d.Name, string(bytes), window)

		output := jsonedit.Val("eventName", "dockerdetail-saveinput")
		jRunInputfiles := docker.DeployInputfiles()
		output = jsonedit.Con(output, jRunInputfiles)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})

	})
	window.On(&gotron.Event{Event: "dockerdetail-load"}, func(bin []byte) {
		var d dockerDeploy
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(d.Name)
		jServer := docker.LoadDeployInputfile(d.Name)

		output := jsonedit.Val("eventName", "dockerdetail-load")
		output = jsonedit.Con(output, jsonedit.On("deployInputData", jServer))
		fmt.Println("output")
		fmt.Println(output)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "dockerdetail-delete"}, func(bin []byte) {
		var d dockerDeploy
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(d.Name)
		docker.DeleteDeployInputfile(d.Name, window)
		output := jsonedit.Val("eventName", "dockerdetail-delete")
		jRunInputfiles := docker.DeployInputfiles()
		output = jsonedit.Con(output, jRunInputfiles)
		fmt.Println(output)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
}
func makeRunStr(d dockerDeploy) string {
	cmd := "HOST:" + d.ServerIP + " docker run " + d.Dit + " --name " + d.Name
	if d.DirName != "" {
		cmd += " -v " + docker.ChangeDockerPath(d.DirName) + ":/" + d.DirName2
	}
	if d.DirNameA != "" {
		cmd += " -v " + docker.ChangeDockerPath(d.DirNameA) + ":/" + d.DirNameA2
	}
	if d.DirNameB != "" {
		cmd += " -v " + docker.ChangeDockerPath(d.DirNameB) + ":/" + d.DirNameB2
	}
	if d.DirNameC != "" {
		cmd += " -v " + docker.ChangeDockerPath(d.DirNameC) + ":/" + d.DirNameC2
	}
	if d.DirNameD != "" {
		cmd += " -v " + docker.ChangeDockerPath(d.DirNameD) + ":/" + d.DirNameD2
	}
	if d.Port != "" {
		cmd += " -p " + d.Port
	}
	if d.Option != "" {
		cmd += " " + d.Option
	}
	//cmd += " " + d.ID
	cmd += " " + "[image]"
	if d.Option2 != "" {
		cmd += " " + d.Option2
	}
	return cmd
}
