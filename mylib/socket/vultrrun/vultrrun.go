package vultrrun

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Equanox/gotron"
	"github.com/uieviiru2/mylib/jsonedit"
	"github.com/uieviiru2/mylib/vultr"
)

type vultrCreateStruct struct {
	Region    string `json:"region"`
	Plan      string `json:"plan"`
	Os        string `json:"os"`
	Sshkey    string `json:"sshkey"`
	Networkid string `json:"networkid"`
	Tag       string `json:"tag"`
}

func Use(window *gotron.BrowserWindow) {

	window.On(&gotron.Event{Event: "vultrrun"}, func(bin []byte) {

		output := jsonedit.Val("eventName", "vultrrun")
		output = jsonedit.Con(output, vultr.Region())
		output = jsonedit.Con(output, vultr.Plan())
		output = jsonedit.Con(output, vultr.Os())
		output = jsonedit.Con(output, vultr.Sshkey())
		output = jsonedit.Con(output, vultr.Account())
		output = jsonedit.Con(output, vultr.Network())

		window.Send(&gotron.Event{Event: jsonedit.End(output)})

	})
	window.On(&gotron.Event{Event: "vultrrun-create"}, func(bin []byte) {

		//var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		var d vultrCreateStruct
		if err := json.Unmarshal(b, &d); err != nil {

		}
		fmt.Println(d.Region)
		fmt.Println(d.Plan)
		fmt.Println(d.Os)
		fmt.Println(d.Sshkey)
		fmt.Println(d.Networkid)
		fmt.Println(d.Tag)

		vultr.Create(d.Region, d.Plan, d.Os, d.Sshkey, d.Networkid, d.Tag, window)

	})

}
