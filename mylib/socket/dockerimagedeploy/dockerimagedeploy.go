package dockerimagedeploy

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

type DockerDeploy struct {
	ID        string `json:"image_id"`
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

	window.On(&gotron.Event{Event: "dockerimagedeploy"}, func(bin []byte) {
		output := jsonedit.Val("eventName", "dockerimagedeploy")

		jDockerImage := docker.Image(window)
		output = jsonedit.Con(output, jDockerImage)

		jServerfiles := docker.Serverfiles(window)
		output = jsonedit.Con(output, jServerfiles)

		jVultr := vultr.List(window)
		output = jsonedit.Con(output, jVultr)

		jRunInputfiles := docker.DeployInputfiles()
		output = jsonedit.Con(output, jRunInputfiles)

		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "dockerimagedeploy-deploy"}, func(bin []byte) {
		//var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		var d DockerDeploy
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
		go docker.ServerDeploy(d.ID, d.Name, d.Dit, d2.ServerPem, d2.User, d.ServerIP, d.Port, d.DirName, d.DirName2, d.DirNameA, d.DirNameA2, d.DirNameB, d.DirNameB2, d.DirNameC, d.DirNameC2, d.DirNameD, d.DirNameD2, d.Option, d.Option2, window)
		//fmt.Println(output)
		//window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "dockerimagedeploy-saveinput"}, func(bin []byte) {
		var d DockerDeploy

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

		output := jsonedit.Val("eventName", "dockerimagedeploy-saveinput")
		jRunInputfiles := docker.DeployInputfiles()
		output = jsonedit.Con(output, jRunInputfiles)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})

	})
	window.On(&gotron.Event{Event: "dockerimagedeploy-load"}, func(bin []byte) {
		var d DockerDeploy
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(d.Name)

		output := jsonedit.Val("eventName", "dockerimagedeploy-load")
		jServer := docker.LoadDeployInputfile(d.Name)
		output = jsonedit.Con(output, jsonedit.On("deployInputData", jServer))

		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "dockerimagedeploy-delete"}, func(bin []byte) {
		var d DockerDeploy
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(d.Name)
		docker.DeleteDeployInputfile(d.Name, window)

		output := jsonedit.Val("eventName", "dockerimagedeploy-delete")
		jRunInputfiles := docker.DeployInputfiles()
		output = jsonedit.Con(output, jRunInputfiles)
		fmt.Println(output)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
}

func makeRunStr(d DockerDeploy) string {
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
	cmd += " " + d.ID

	if d.Option2 != "" {
		cmd += " " + d.Option2
	}
	return cmd
}
