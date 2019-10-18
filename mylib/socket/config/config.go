package config

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/Equanox/gotron"
	"github.com/uieviiru/test/mylib/jsonedit"
)

type ConfigStruct struct {
	DockerExe   string `json:"docker_exe"`
	VultrApiKey string `json:"vultr_api_key"`
	VultrPem    string `json:"vultr_pem"`
	ConfigDir   string `json:"config_dir"`
}

func Use(window *gotron.BrowserWindow) {

	window.On(&gotron.Event{Event: "config"}, func(bin []byte) {
		//fmt.Println(runtime.GOOS)
		userData, _ := user.Current()

		if _, err := os.Stat(userData.HomeDir + "\\.docker.bin"); os.IsNotExist(err) {
		} else {
			f, err := os.Open(userData.HomeDir + "\\.docker.bin")
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			var d ConfigStruct
			dec := gob.NewDecoder(f)
			if err := dec.Decode(&d); err != nil {
				log.Fatal("decode error:", err)
			}
			output := jsonedit.Val("eventName", "config")
			bytes, _ := json.Marshal(&d)
			output = jsonedit.Con(output, jsonedit.On("config", string(bytes)))

			window.Send(&gotron.Event{Event: jsonedit.End(output)})
		}
	})

	window.On(&gotron.Event{Event: "config-save"}, func(bin []byte) {
		//var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		var d ConfigStruct
		if err := json.Unmarshal(b, &d); err != nil {

		}
		if f, err := os.Stat(d.ConfigDir + "\\server"); os.IsNotExist(err) || !f.IsDir() {
			os.Mkdir(d.ConfigDir+"\\server", 0777)
		}
		if f, err := os.Stat(d.ConfigDir + "\\Dockerfile"); os.IsNotExist(err) || !f.IsDir() {
			os.Mkdir(d.ConfigDir+"\\Dockerfile", 0777)
		}
		if f, err := os.Stat(d.ConfigDir + "\\docker-compose"); os.IsNotExist(err) || !f.IsDir() {
			os.Mkdir(d.ConfigDir+"\\docker-compose", 0777)
		}
		if f, err := os.Stat(d.ConfigDir + "\\runinput"); os.IsNotExist(err) || !f.IsDir() {
			os.Mkdir(d.ConfigDir+"\\runinput", 0777)
		}
		if f, err := os.Stat(d.ConfigDir + "\\deployinput"); os.IsNotExist(err) || !f.IsDir() {
			os.Mkdir(d.ConfigDir+"\\deployinput", 0777)
		}
		if f, err := os.Stat(d.ConfigDir + "\\tmp"); os.IsNotExist(err) || !f.IsDir() {
			os.Mkdir(d.ConfigDir+"\\tmp", 0777)
		}
		userData, _ := user.Current()
		f, err := os.Create(userData.HomeDir + "\\.docker.bin")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		enc := gob.NewEncoder(f)

		if err := enc.Encode(d); err != nil {
			log.Fatal(err)
		}
		//docker.OutLog("Saved", window)
		output := jsonedit.Val("eventName", "config-save")
		window.Send(&gotron.Event{Event: jsonedit.End(output)})

	})
}

func LoadConfig() ConfigStruct {

	userData, _ := user.Current()
	fmt.Println(userData.HomeDir + "\\.docker.bin")
	f, err := os.Open(userData.HomeDir + "\\.docker.bin")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var d ConfigStruct
	dec := gob.NewDecoder(f)
	if err := dec.Decode(&d); err != nil {
		log.Fatal("decode error:", err)
	}
	return d
}
