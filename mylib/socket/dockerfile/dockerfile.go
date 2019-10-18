package dockerfile

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Equanox/gotron"
	"github.com/uieviiru/test/mylib/docker"
	"github.com/uieviiru/test/mylib/jsonedit"
	"github.com/uieviiru/test/mylib/socket/config"
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

		output := jsonedit.Val("eventName", "log")
		output = jsonedit.Con(output, jsonedit.Val("log", "Saved"))
		fmt.Println("output")
		fmt.Println(output)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "dockerfile-test"}, func(bin []byte) {
		var d dockerfile
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		window.Send(&gotron.Event{Event: jsonedit.End("")})
		docker.BuildDockerfiles(d.Name, window)
	})
	window.On(&gotron.Event{Event: "dockerfile-delete"}, func(bin []byte) {
		var d dockerfile
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(d.Name)
		docker.DeleteDockerfile(d.Name, window)

		/*
			output := jsonedit.Val("eventName", "dockerfile")

			jDockerfiles := docker.Dockerfiles()
			output = jsonedit.Con(output, jDockerfiles)
			fmt.Println("output")
			fmt.Println(output)
			window.Send(&gotron.Event{Event: jsonedit.End(output)})
		*/
		//window.Send(&gotron.Event{Event: jsonedit.End("")})
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
