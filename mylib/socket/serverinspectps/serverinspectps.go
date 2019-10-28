package serverinspectps

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Equanox/gotron"
	"github.com/uieviiru2/mylib/docker"
	"github.com/uieviiru2/mylib/jsonedit"
	"github.com/uieviiru2/mylib/socket/config"
	"github.com/uieviiru2/mylib/socket/configserver"
)

type dockerID struct {
	ID string `json:"id"`
	Ip string `json:"ip"`
	V  string `json:"v"`
}

func Use(window *gotron.BrowserWindow) {

	window.On(&gotron.Event{Event: "serverinspectps"}, func(bin []byte) {
		var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(string(b))
		fmt.Println(d.ID)
		output := jsonedit.Val("eventName", "serverinspectps")

		if d.V == "1" {
			configData := config.LoadConfig()
			jDocker := docker.ServerInspect(d.Ip, "root", configData.VultrPem, d.ID, window)
			output = jsonedit.Con(output, jDocker)
			window.Send(&gotron.Event{Event: jsonedit.End(output)})
		} else {

			jServer := docker.LoadServerfile(d.Ip)
			var d2 configserver.ConfigServer
			if err2 := json.Unmarshal([]byte(jServer), &d2); err2 != nil {
				// ...
			}

			jDocker := docker.ServerInspect(d2.Ip, d2.User, d2.ServerPem, d.ID, window)
			output = jsonedit.Con(output, jDocker)
			window.Send(&gotron.Event{Event: jsonedit.End(output)})
		}

	})

	window.On(&gotron.Event{Event: "serverinspectps-remove"}, func(bin []byte) {
		//var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		var d dockerID
		if err := json.Unmarshal(b, &d); err != nil {

		}
		jServer := docker.LoadServerfile(d.Ip)
		var d2 configserver.ConfigServer
		if err2 := json.Unmarshal([]byte(jServer), &d2); err2 != nil {
			// ...
		}
		if d.V == "1" {
			configData := config.LoadConfig()
			/*
				fmt.Println("ssh -oStrictHostKeyChecking=no -i " + configData.VultrPem + " root@" + d.Ip + " docker stop " + d.ID)
				docker.SshCommand("-oStrictHostKeyChecking=no", "-i", configData.VultrPem, "root@"+d.Ip, "docker", "stop", d.ID)
				fmt.Println("ssh -oStrictHostKeyChecking=no -i " + configData.VultrPem + " root@" + d.Ip + " docker rm " + d.ID)
				docker.SshCommand("-oStrictHostKeyChecking=no", "-i", configData.VultrPem, "root@"+d.Ip, "docker", "rm", d.ID)
			*/
			docker.ServerRemove(d.Ip, "root", configData.VultrPem, d.ID, window)
		} else {
			docker.ServerRemove(d2.Ip, d2.User, d2.ServerPem, d.ID, window)
		}
		//fmt.Println(output)
		//window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "serverinspectps-bash"}, func(bin []byte) {
		//var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		var d dockerID
		if err := json.Unmarshal(b, &d); err != nil {

		}
		jServer := docker.LoadServerfile(d.Ip)
		var d2 configserver.ConfigServer
		if err2 := json.Unmarshal([]byte(jServer), &d2); err2 != nil {
			// ...
		}
		//ssh  -o StrictHostKeyChecking=no -i C:\Users\fg\.ssh\vultr.pem -t core@192.168.10.191 docker exec -it 1b9c80884253 bash
		if d.V == "1" {
			configData := config.LoadConfig()
			docker.OpenServerDockerEnter("ssh  -o StrictHostKeyChecking=no -i " + configData.VultrPem + " -t " + "root@" + d.Ip + " docker exec -it " + d.ID + " bash")
		} else {
			docker.OpenServerDockerEnter("ssh  -o StrictHostKeyChecking=no -i " + d2.ServerPem + " -t " + d2.User + "@" + d2.Ip + " docker exec -it " + d.ID + " bash")
		}

		//fmt.Println(output)
		//window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "serverinspectps-start"}, func(bin []byte) {
		//var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		var d dockerID
		if err := json.Unmarshal(b, &d); err != nil {

		}
		jServer := docker.LoadServerfile(d.Ip)
		var d2 configserver.ConfigServer
		if err2 := json.Unmarshal([]byte(jServer), &d2); err2 != nil {
			// ...
		}
		//ssh  -o StrictHostKeyChecking=no -i C:\Users\fg\.ssh\vultr.pem -t core@192.168.10.191 docker exec -it 1b9c80884253 bash
		if d.V == "1" {
			configData := config.LoadConfig()
			docker.OpenServerDockerEnter("ssh  -o StrictHostKeyChecking=no -i " + configData.VultrPem + " -t " + "root@" + d.Ip + " docker start " + d.ID)
		} else {
			docker.OpenServerDockerEnter("ssh  -o StrictHostKeyChecking=no -i " + d2.ServerPem + " -t " + d2.User + "@" + d2.Ip + " docker start " + d.ID)
		}

		//fmt.Println(output)
		//window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "serverinspectps-stop"}, func(bin []byte) {
		//var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		var d dockerID
		if err := json.Unmarshal(b, &d); err != nil {

		}
		jServer := docker.LoadServerfile(d.Ip)
		var d2 configserver.ConfigServer
		if err2 := json.Unmarshal([]byte(jServer), &d2); err2 != nil {
			// ...
		}
		//ssh  -o StrictHostKeyChecking=no -i C:\Users\fg\.ssh\vultr.pem -t core@192.168.10.191 docker exec -it 1b9c80884253 bash
		if d.V == "1" {
			configData := config.LoadConfig()
			docker.OpenServerDockerEnter("ssh  -o StrictHostKeyChecking=no -i " + configData.VultrPem + " -t " + "root@" + d.Ip + " docker stop " + d.ID)
		} else {
			docker.OpenServerDockerEnter("ssh  -o StrictHostKeyChecking=no -i " + d2.ServerPem + " -t " + d2.User + "@" + d2.Ip + " docker stop " + d.ID)
		}

		//fmt.Println(output)
		//window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
}
