package dockerrun

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Equanox/gotron"
	"github.com/uieviiru2/mylib/docker"
	"github.com/uieviiru2/mylib/jsonedit"
)

type dockerRunStruct struct {
	Docker    string `json:"docker"`
	Name      string `json:"name"`
	Dit       string `json:"dit"`
	Port      string `json:"port"`
	DirName   string `json:"dirname"`
	DirName2  string `json:"dirname2"`
	DirNameA  string `json:"dirname_a"`
	DirNameA2 string `json:"dirname_a2"`
	DirNameB  string `json:"dirname_b"`
	DirNameB2 string `json:"dirname_b2"`
	DirNameC  string `json:"dirname_c"`
	DirNameC2 string `json:"dirname_c2"`
	DirNameD  string `json:"dirname_d"`
	DirNameD2 string `json:"dirname_d2"`
	Option    string `json:"option"`
	Option2   string `json:"option2"`
	Cmd       string `json:"cmd"`
}

func Use(window *gotron.BrowserWindow) {

	window.On(&gotron.Event{Event: "dockerrun"}, func(bin []byte) {

		output := jsonedit.Val("eventName", "dockerrun")
		jDockerImage := docker.Image(window)
		output = jsonedit.Con(output, jDockerImage)
		jServerfiles := docker.RunInputfiles()
		output = jsonedit.Con(output, jServerfiles)

		window.Send(&gotron.Event{Event: jsonedit.End(output)})

	})
	window.On(&gotron.Event{Event: "dockerrun-excute"}, func(bin []byte) {

		var d dockerRunStruct

		b := []byte(bin)
		//buf := bytes.NewBuffer(b)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(string(b))
		fmt.Println(d.Docker)
		fmt.Println(d.Name)
		fmt.Println(d.Port)
		fmt.Println(d.DirName)
		fmt.Println(d.DirName2)
		fmt.Println(d.DirNameA)
		fmt.Println(d.DirNameA2)
		fmt.Println(d.DirNameB)
		fmt.Println(d.DirNameB2)
		fmt.Println(d.DirNameC)
		fmt.Println(d.DirNameC2)
		fmt.Println(d.DirNameD)
		fmt.Println(d.DirNameD2)
		fmt.Println(d.Option)
		if docker.IsWindows() {
			docker.Run(d.Docker, d.Name, d.Dit, d.Port, d.DirName, d.DirName2, d.DirNameA, d.DirNameA2, d.DirNameB, d.DirNameB2, d.DirNameC, d.DirNameC2, d.DirNameD, d.DirNameD2, d.Option, d.Option2, window)
		} else {
			docker.RunMac(d.Docker, d.Name, d.Dit, d.Port, d.DirName, d.DirName2, d.DirNameA, d.DirNameA2, d.DirNameB, d.DirNameB2, d.DirNameC, d.DirNameC2, d.DirNameD, d.DirNameD2, d.Option, d.Option2, window)
		}
	})
	window.On(&gotron.Event{Event: "dockerrun-saveinput"}, func(bin []byte) {

		var d dockerRunStruct

		b := []byte(bin)
		//buf := bytes.NewBuffer(b)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(string(b))
		fmt.Println(d.Docker)
		fmt.Println(d.Name)
		fmt.Println(d.Port)
		fmt.Println(d.DirName)
		fmt.Println(d.DirName2)
		fmt.Println(d.DirNameA)
		fmt.Println(d.DirNameA2)
		fmt.Println(d.DirNameB)
		fmt.Println(d.DirNameB2)
		fmt.Println(d.DirNameC)
		fmt.Println(d.DirNameC2)
		fmt.Println(d.DirNameD)
		fmt.Println(d.DirNameD2)
		fmt.Println(d.Option)
		d.Cmd = makeRunStr(d)

		bytes, _ := json.Marshal(d)
		docker.SaveRunInputfile(d.Name, string(bytes), window)

		output := jsonedit.Val("eventName", "dockerrun-saveinput")
		jServerfiles := docker.RunInputfiles()
		output = jsonedit.Con(output, jServerfiles)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})

	})
	window.On(&gotron.Event{Event: "dockerrun-load"}, func(bin []byte) {
		var d dockerRunStruct
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(d.Name)
		jServer := docker.LoadRunInputfile(d.Name)

		output := jsonedit.Val("eventName", "dockerrun-load")
		output = jsonedit.Con(output, jsonedit.On("runInputData", jServer))
		fmt.Println("output")
		fmt.Println(output)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "dockerrun-delete"}, func(bin []byte) {
		var d dockerRunStruct
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(d.Name)
		docker.DeleteRunInputfile(d.Name, window)

		output := jsonedit.Val("eventName", "dockerrun-delete")
		jServerfiles := docker.RunInputfiles()
		output = jsonedit.Con(output, jServerfiles)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})

	})
}
func makeRunStr(d dockerRunStruct) string {
	cmd := "docker run " + d.Dit + " --name " + d.Name
	if d.DirName != "" {
		cmd += " -v " + docker.ChangeDockerPath(d.DirName) + ":/" + d.DirName2
	}
	if d.DirNameA != "" {
		cmd += " -v " + docker.ChangeDockerPath(d.DirNameA) + ":/" + d.DirNameA2
	}
	if d.DirNameB != "" {
		cmd += " -v " + docker.ChangeDockerPath(d.DirNameB) + ":/" + d.DirNameB2
	}
	if d.DirNameC != "" {
		cmd += " -v " + docker.ChangeDockerPath(d.DirNameC) + ":/" + d.DirNameC2
	}
	if d.DirNameD != "" {
		cmd += " -v " + docker.ChangeDockerPath(d.DirNameD) + ":/" + d.DirNameD2
	}
	if d.Port != "" {
		cmd += " -p " + d.Port
	}
	if d.Option != "" {
		cmd += " " + d.Option
	}
	cmd += " " + d.Docker

	if d.Option2 != "" {
		cmd += " " + d.Option2
	}
	return cmd
}
