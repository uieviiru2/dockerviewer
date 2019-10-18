package serverinspectimage

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
	ID    string `json:"id"`
	Ip    string `json:"ip"`
	V     string `json:"v"`
	Force string `json:"force"`
}

func Use(window *gotron.BrowserWindow) {

	window.On(&gotron.Event{Event: "serverinspectimage"}, func(bin []byte) {
		var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(string(b))
		fmt.Println(d.ID)
		output := jsonedit.Val("eventName", "serverinspectimage")

		if d.V == "1" {
			configData := config.LoadConfig()
			jDocker := docker.ServerImageInspect(d.Ip, "root", configData.VultrPem, d.ID, window)
			output = jsonedit.Con(output, jDocker)
			window.Send(&gotron.Event{Event: jsonedit.End(output)})
		} else {

			jServer := docker.LoadServerfile(d.Ip)
			var d2 configserver.ConfigServer
			if err2 := json.Unmarshal([]byte(jServer), &d2); err2 != nil {
				// ...
			}

			jDocker := docker.ServerImageInspect(d2.Ip, d2.User, d2.ServerPem, d.ID, window)
			output = jsonedit.Con(output, jDocker)

			window.Send(&gotron.Event{Event: jsonedit.End(output)})
		}
	})

	window.On(&gotron.Event{Event: "serverinspectimage-remove"}, func(bin []byte) {
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
			//fmt.Println("ssh -oStrictHostKeyChecking=no -i " + configData.VultrPem + " root@" + d2.Ip + " docker rmi " + d.ID)
			//docker.SshCommand("-oStrictHostKeyChecking=no", "-i", configData.VultrPem, "root@"+d2.Ip, "docker", "rmi", d.ID)
			docker.ServerImageRemove(d.Ip, "root", configData.VultrPem, d.ID, d.ID, window)

		} else {
			docker.ServerImageRemove(d2.Ip, d2.User, d2.ServerPem, d.ID, d.Force, window)
		}

		//fmt.Println(output)
		//window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})

}
