package dockerfile

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Equanox/gotron"
	"github.com/uieviiru2/mylib/docker"
	"github.com/uieviiru2/mylib/jsonedit"
	"github.com/uieviiru2/mylib/socket/config"
)

type dockerfile struct {
	Name   string `json:"name"`
	Script string `json:"script"`
}
type userPath struct {
	Path string `json:"path"`
}

func Use(window *gotron.BrowserWindow) {

	window.On(&gotron.Event{Event: "dockerfile"}, func(bin []byte) {
		configData := config.LoadConfig()
		output := jsonedit.Val("eventName", "dockerfile")

		jDockerfiles := docker.Dockerfiles(window)
		output = jsonedit.Con(output, jDockerfiles)
		path := configData.ConfigDir + "\\Dockerfile"
		var up userPath
		up.Path = path
		bytes, _ := json.Marshal(up)
		output = jsonedit.Con(output, jsonedit.StripQ(string(bytes)))
		fmt.Println("output")
		fmt.Println(output)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "dockerfile-save"}, func(bin []byte) {
		configData := config.LoadConfig()
		var d dockerfile
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(d.Name)
		fmt.Println(d.Script)
		docker.SaveDockerfile(d.Name, d.Script, window)

		output := jsonedit.Val("eventName", "Saved")

		jDockerfiles := docker.Dockerfiles(window)
		output = jsonedit.Con(output, jDockerfiles)
		path := configData.ConfigDir + "\\Dockerfile"
		var up userPath
		up.Path = path
		bytes, _ := json.Marshal(up)
		output = jsonedit.Con(output, jsonedit.StripQ(string(bytes)))
		window.Send(&gotron.Event{Event: jsonedit.End(output)})

	})
	window.On(&gotron.Event{Event: "dockerfile-test"}, func(bin []byte) {
		configData := config.LoadConfig()
		var d dockerfile
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		docker.SaveDockerfile(d.Name, d.Script, window)

		output := jsonedit.Val("eventName", "dockerfile-test")
		jDockerfiles := docker.Dockerfiles(window)
		output = jsonedit.Con(output, jDockerfiles)
		path := configData.ConfigDir + "\\Dockerfile"
		var up userPath
		up.Path = path
		bytes, _ := json.Marshal(up)
		output = jsonedit.Con(output, jsonedit.StripQ(string(bytes)))
		window.Send(&gotron.Event{Event: jsonedit.End(output)})

		docker.BuildDockerfiles(d.Name, window)
	})
	window.On(&gotron.Event{Event: "dockerfile-delete"}, func(bin []byte) {
		configData := config.LoadConfig()
		var d dockerfile
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(d.Name)
		docker.DeleteDockerfile(d.Name, window)

		output := jsonedit.Val("eventName", "dockerfile-delete")

		jDockerfiles := docker.Dockerfiles(window)
		output = jsonedit.Con(output, jDockerfiles)
		path := configData.ConfigDir + "\\Dockerfile"
		var up userPath
		up.Path = path
		bytes, _ := json.Marshal(up)
		output = jsonedit.Con(output, jsonedit.StripQ(string(bytes)))
		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "dockerfile-load"}, func(bin []byte) {
		var d dockerfile
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(d.Name)
		script := docker.LoadDockerfile(d.Name, window)
		bytes, _ := json.Marshal(script)

		scriptSafeString := string(bytes)
		output := jsonedit.Val("eventName", "dockerfile-load")
		output = jsonedit.Con(output, jsonedit.On("script", scriptSafeString))
		output = jsonedit.Con(output, jsonedit.Val("name", d.Name))
		fmt.Println("output")
		fmt.Println(output)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})

	})
}
