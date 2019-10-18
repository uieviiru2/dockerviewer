package dockerpull

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Equanox/gotron"
	"github.com/uieviiru/test/mylib/docker"
	"github.com/uieviiru/test/mylib/jsonedit"
)

type dockerSearchStruct struct {
	Search string `json:"search"`
}
type dockerPullStruct struct {
	PullName string `json:"pull_name"`
}

func Use(window *gotron.BrowserWindow) {

	window.On(&gotron.Event{Event: "dockerpull-search"}, func(bin []byte) {
		var d dockerSearchStruct
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		output := jsonedit.Val("eventName", "dockerpull-search")
		jDockerSearch := docker.Search(d.Search, window)
		output = jsonedit.Con(output, jDockerSearch)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})

	})

	window.On(&gotron.Event{Event: "dockerpull-excute"}, func(bin []byte) {

		var d dockerPullStruct

		b := []byte(bin)
		//buf := bytes.NewBuffer(b)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}

		fmt.Println(d.PullName)
		go docker.Pull2(d.PullName, window)
		/*
			res := docker.Pull(d.PullName, window)
			output := jsonedit.Val("eventName", "log")
			output = jsonedit.Con(output, jsonedit.Val("log", res))
			fmt.Println("output")
			fmt.Println(output)
			window.Send(&gotron.Event{Event: jsonedit.End(output)})
		*/

	})

}
