package serverinspect

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Equanox/gotron"

	"github.com/uieviiru/test/mylib/docker"
	"github.com/uieviiru/test/mylib/jsonedit"
	"github.com/uieviiru/test/mylib/socket/config"
	"github.com/uieviiru/test/mylib/socket/configserver"
)

type dockerID struct {
	Ip string `json:"ip"`
	V  string `json:"v"`
}

func Use(window *gotron.BrowserWindow) {

	window.On(&gotron.Event{Event: "serverinspect"}, func(bin []byte) {
		var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		output := jsonedit.Val("eventName", "serverinspect")
		fmt.Println("eventName" + ":" + "serverinspect")
		fmt.Println(d.Ip)
		fmt.Println(d.V)

		if d.V == "1" {
			configData := config.LoadConfig()
			jDocker := docker.GetServerPs(d.Ip, "root", configData.VultrPem, window)
			output = jsonedit.Con(output, jDocker)
			jDockerImage := docker.ServerImage(d.Ip, "root", configData.VultrPem, window)
			output = jsonedit.Con(output, jDockerImage)

			fmt.Println(output)
			window.Send(&gotron.Event{Event: jsonedit.End(output)})
		} else {
			var d2 configserver.ConfigServer
			jServer := docker.LoadServerfile(d.Ip)

			if err2 := json.Unmarshal([]byte(jServer), &d2); err2 != nil {
				// ...
			}
			jDocker := docker.GetServerPs(d2.Ip, d2.User, d2.ServerPem, window)
			output = jsonedit.Con(output, jDocker)
			jDockerImage := docker.ServerImage(d2.Ip, d2.User, d2.ServerPem, window)
			output = jsonedit.Con(output, jDockerImage)

			fmt.Println(output)
			window.Send(&gotron.Event{Event: jsonedit.End(output)})
		}

	})
	window.On(&gotron.Event{Event: "serverinspect-removeall"}, func(bin []byte) {
		var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println("eventName" + ":" + "serverinspect-removeall")
		fmt.Println(d.Ip)
		fmt.Println(d.V)

		if d.V == "1" {
			configData := config.LoadConfig()

			docker.ServerRemoveall(d.Ip, "root", configData.VultrPem, window)

		} else {
			var d2 configserver.ConfigServer
			jServer := docker.LoadServerfile(d.Ip)

			if err2 := json.Unmarshal([]byte(jServer), &d2); err2 != nil {
				// ...
			}
			docker.ServerRemoveall(d2.Ip, d2.User, d2.ServerPem, window)
		}

	})
	window.On(&gotron.Event{Event: "serverinspect-bash"}, func(bin []byte) {
		var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println("eventName" + ":" + "serverinspect-bash")
		fmt.Println(d.Ip)
		fmt.Println(d.V)
		//ssh  -o StrictHostKeyChecking=no -i C:\Users\fg\.ssh\vultr.pem -t core@192.168.10.191
		if d.V == "1" {
			configData := config.LoadConfig()

			docker.OpenServer("ssh  -o StrictHostKeyChecking=no -i " + configData.VultrPem + " -t root@" + d.Ip)

		} else {
			var d2 configserver.ConfigServer
			jServer := docker.LoadServerfile(d.Ip)

			if err2 := json.Unmarshal([]byte(jServer), &d2); err2 != nil {
				// ...
			}
			docker.OpenServer("ssh  -o StrictHostKeyChecking=no -i " + d2.ServerPem + " -t " + d2.User + "@" + d2.Ip)
		}

	})
}
