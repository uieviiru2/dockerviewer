package dockerimagedetail

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/uieviiru2/mylib/jsonedit"
	"github.com/uieviiru2/mylib/vultr"

	"github.com/Equanox/gotron"
	"github.com/uieviiru2/mylib/docker"
)

type dockerID struct {
	ID    string `json:"id"`
	Force string `json:"force"`
}
type dockerDeploy struct {
	ID       string `json:"id"`
	ServerIP string `json:"server_ip"`
	Port     string `json:"port"`
	DirName  string `json:"dirname"`
	DirName2 string `json:"dirname2"`
}

func Use(window *gotron.BrowserWindow) {

	window.On(&gotron.Event{Event: "dockerimagedetail"}, func(bin []byte) {
		var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(string(b))
		fmt.Println(d.ID)
		output := jsonedit.Val("eventName", "dockerimagedetail")

		jDocker := docker.ImageInspect(d.ID)
		output = jsonedit.Con(output, jDocker)

		jVultr := vultr.List(window)
		output = jsonedit.Con(output, jVultr)

		fmt.Println(jVultr)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	/*
		window.On(&gotron.Event{Event: "dockerdetail-deploy"}, func(bin []byte) {
			//var d dockerID
			b := []byte(bin)
			buf := bytes.NewBuffer(b)
			fmt.Println(buf)
			var d dockerDeploy
			if err := json.Unmarshal(b, &d); err != nil {

			}

			docker.Deploy(d.ID, d.ServerIP, d.Port, d.DirName, d.DirName2)
			//fmt.Println(output)
			//window.Send(&gotron.Event{Event: jsonedit.End(output)})
		})
	*/
	window.On(&gotron.Event{Event: "dockerimagedetail-remove"}, func(bin []byte) {
		//var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		var d dockerID
		if err := json.Unmarshal(b, &d); err != nil {

		}

		docker.ImageRemove(d.ID, d.Force, window)
		//fmt.Println(output)
		//window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})

}
