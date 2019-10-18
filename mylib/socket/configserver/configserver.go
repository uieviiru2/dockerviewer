package configserver

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Equanox/gotron"
	"github.com/uieviiru/test/mylib/docker"
	"github.com/uieviiru/test/mylib/jsonedit"
)

type ConfigServer struct {
	Ip        string `json:"ip"`
	User      string `json:"user"`
	ServerPem string `json:"server_pem"`
}

func Use(window *gotron.BrowserWindow) {

	window.On(&gotron.Event{Event: "configserver"}, func(bin []byte) {

		output := jsonedit.Val("eventName", "configserver")

		jServerfiles := docker.Serverfiles(window)
		output = jsonedit.Con(output, jServerfiles)
		fmt.Println("output")
		fmt.Println(output)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})

	window.On(&gotron.Event{Event: "configserver-save"}, func(bin []byte) {
		var d ConfigServer
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {

		}
		docker.SaveServerfile(d.Ip, string(b), window)
		output := jsonedit.Val("eventName", "configserver-save")
		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})

	window.On(&gotron.Event{Event: "configserver-delete"}, func(bin []byte) {
		var d ConfigServer
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(d.Ip)
		docker.DeleteServerfile(d.Ip)
		output := jsonedit.Val("eventName", "configserver-delete")
		window.Send(&gotron.Event{Event: jsonedit.End(output)})

	})
	window.On(&gotron.Event{Event: "configserver-load"}, func(bin []byte) {
		var d ConfigServer
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(d.Ip)
		jServer := docker.LoadServerfile(d.Ip)

		output := jsonedit.Val("eventName", "configserver-load")
		output = jsonedit.Con(output, jsonedit.On("serverfile", jServer))
		fmt.Println("output")
		fmt.Println(output)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})

	})
}
