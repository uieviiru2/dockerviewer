package vultrdetail

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/uieviiru/test/mylib/jsonedit"
	"github.com/uieviiru/test/mylib/vultr"

	"github.com/Equanox/gotron"
)

type vultrID struct {
	ID string `json:"id"`
}

func Use(window *gotron.BrowserWindow) {

	window.On(&gotron.Event{Event: "vultrdetail"}, func(bin []byte) {
		var d vultrID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		if err := json.Unmarshal(b, &d); err != nil {
			// ...
		}
		fmt.Println(string(b))
		fmt.Println(d.ID)
		output := jsonedit.Val("eventName", "vultrdetail")

		jVultrDetail := vultr.Detail(d.ID)
		output = jsonedit.Con(output, jVultrDetail)

		fmt.Println(jVultrDetail)
		window.Send(&gotron.Event{Event: jsonedit.End(output)})
	})
	window.On(&gotron.Event{Event: "vultrdetail-destroy"}, func(bin []byte) {
		//var d dockerID
		b := []byte(bin)
		buf := bytes.NewBuffer(b)
		fmt.Println(buf)
		var d vultrID
		if err := json.Unmarshal(b, &d); err != nil {

		}
		vultr.Destroy(d.ID, window)
	})
}
