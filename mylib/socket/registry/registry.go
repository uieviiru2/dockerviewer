package registry

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

type dockerSearchStruct struct {
	Search string `json:"search"`
}
type dockerPullStruct struct {
	PullName string `json:"pull_name"`
}
type RegisterDeploy struct {
	ID       string `json:"image_id"`
	Name     string `json:"name"`
	ServerIP string `json:"server_ip"`
	Port     string `json:"port"`
}

func Use(window *gotron.BrowserWindow) {

	window.On(&gotron.Event{Event: "registry"}, func(bin []byte) {
		var d dockerSearchStruct
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		output := jsonedit.Val("eventName", "registry")
		output = jsonedit.Con(output, docker.RegistryTag())
		output = jsonedit.Con(output, docker.Serverfiles(window))
		output = jsonedit.Con(output, docker.Registryfiles(window))

		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "registry-deploy"}, func(bin []byte) {
		//var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		var d RegisterDeploy
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
		bytesData, _ := json.Marshal(d)
		docker.SaveRegistryfile(d.Name, string(bytesData), window)
		//fmt.Println(output)
		//window.Send(&gotron.Event{Event: jsonedit.End(output)})
		output := jsonedit.Val("eventName", "registry-deploy")
		output = jsonedit.Con(output, docker.Registryfiles(window))
		window.Send(&gotron.Event{Event: jsonedit.End(output)})
		go docker.RegistryDeploy(d.ID, d.Name, d2.ServerPem, d2.User, d.ServerIP, d.Port, window)
	})
	window.On(&gotron.Event{Event: "registry-delete"}, func(bin []byte) {

		var d RegisterDeploy
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(d.Name)
		docker.DeleteRegistryfile(d.Name, window)

		output := jsonedit.Val("eventName", "registry-delete")
		output = jsonedit.Con(output, docker.Registryfiles(window))
		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
}
