package dockercompose

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Equanox/gotron"
	"github.com/uieviiru2/mylib/docker"
	"github.com/uieviiru2/mylib/jsonedit"
	"github.com/uieviiru2/mylib/socket/config"
)

type dockercompose struct {
	Name   string `json:"name"`
	Script string `json:"script"`
}
type userPath struct {
	Path string `json:"path"`
}

func Use(window *gotron.BrowserWindow) {

	window.On(&gotron.Event{Event: "dockercompose"}, func(bin []byte) {
		configData := config.LoadConfig()
		output := jsonedit.Val("eventName", "dockercompose")

		jDockerfiles := docker.DockerCompose(window)
		output = jsonedit.Con(output, jDockerfiles)
		path := configData.ConfigDir + "/docker-compose"
		var up userPath
		up.Path = path
		bytes, _ := json.Marshal(up)
		output = jsonedit.Con(output, jsonedit.StripQ(string(bytes)))
		fmt.Println("output")
		fmt.Println(output)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "dockercompose-save"}, func(bin []byte) {
		configData := config.LoadConfig()
		var d dockercompose
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(d.Name)
		fmt.Println(d.Script)
		docker.SaveDockerCompose(d.Name, d.Script)

		output := jsonedit.Val("eventName", "dockercompose-save")
		jDockerfiles := docker.DockerCompose(window)
		output = jsonedit.Con(output, jDockerfiles)
		path := configData.ConfigDir + "/docker-compose"
		var up userPath
		up.Path = path
		bytes, _ := json.Marshal(up)
		output = jsonedit.Con(output, jsonedit.StripQ(string(bytes)))
		fmt.Println("output")
		fmt.Println(output)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "dockercompose-test"}, func(bin []byte) {
		configData := config.LoadConfig()
		var d dockercompose
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		docker.SaveDockerCompose(d.Name, d.Script)
		output := jsonedit.Val("eventName", "dockercompose-test")
		jDockerfiles := docker.DockerCompose(window)
		output = jsonedit.Con(output, jDockerfiles)
		path := configData.ConfigDir + "/docker-compose"
		var up userPath
		up.Path = path
		bytes, _ := json.Marshal(up)
		output = jsonedit.Con(output, jsonedit.StripQ(string(bytes)))
		fmt.Println("output")
		fmt.Println(output)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})

		docker.UpDockerCompose(d.Name, window)
	})
	window.On(&gotron.Event{Event: "dockercompose-delete"}, func(bin []byte) {
		configData := config.LoadConfig()
		var d dockercompose
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(d.Name)
		docker.DeleteDockerCompose(d.Name)

		output := jsonedit.Val("eventName", "dockercompose-delete")
		jDockerfiles := docker.DockerCompose(window)
		output = jsonedit.Con(output, jDockerfiles)
		path := configData.ConfigDir + "/docker-compose"
		var up userPath
		up.Path = path
		bytes, _ := json.Marshal(up)
		output = jsonedit.Con(output, jsonedit.StripQ(string(bytes)))
		fmt.Println("output")
		fmt.Println(output)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "dockercompose-load"}, func(bin []byte) {
		var d dockercompose
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(d.Name)
		script := docker.LoadDockerCompose(d.Name)
		bytes, _ := json.Marshal(script)

		scriptSafeString := string(bytes)
		output := jsonedit.Val("eventName", "dockercompose-load")
		output = jsonedit.Con(output, jsonedit.On("script", scriptSafeString))
		output = jsonedit.Con(output, jsonedit.Val("name", d.Name))
		fmt.Println("output")
		fmt.Println(output)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})

	})
}
