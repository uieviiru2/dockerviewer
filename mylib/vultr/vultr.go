package vultr

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/Equanox/gotron"
	"github.com/uieviiru/test/mylib/docker"
	"github.com/uieviiru/test/mylib/jsonedit"
	"github.com/uieviiru/test/mylib/socket/config"
)

//curl -H 'API-Key: YOURKEY' https://api.vultr.com/v1/server/list
func Detail(id string) string {
	configData := config.LoadConfig()

	req, _ := http.NewRequest("GET", "https://api.vultr.com/v1/server/list?SUBID="+id, nil)
	req.Header.Set("API-Key", configData.VultrApiKey)
	client := new(http.Client)
	resp, err := client.Do(req)
	fmt.Println(err)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	respStr := string(byteArray)
	respStr = jsonedit.StripQ(respStr)
	//fmt.Println(respStr)
	output := jsonedit.Key("vultrDetail", respStr) // htmlをstringで取得

	return output
}
func List(window *gotron.BrowserWindow) string {
	configData := config.LoadConfig()
	if configData.VultrApiKey == "" {
		output := jsonedit.Key("vultr", "") // htmlをstringで取得
		return output
	}
	req, _ := http.NewRequest("GET", "https://api.vultr.com/v1/server/list", nil)
	req.Header.Set("API-Key", configData.VultrApiKey)
	client := new(http.Client)
	resp, err := client.Do(req)
	fmt.Println(err)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	respStr := string(byteArray)
	respStr = jsonedit.StripQ(respStr)
	fmt.Println(respStr)
	output := jsonedit.Key("vultr", respStr) // htmlをstringで取得

	docker.OutLog("■■■Vultr list■■■", window)
	docker.OutLog(respStr, window)
	return output
}

func Create(region, plan, os, sshkey, networkid, tag string, window *gotron.BrowserWindow) {
	configData := config.LoadConfig()

	values := url.Values{}
	values.Add("DCID", region)
	values.Add("VPSPLANID", plan)
	values.Add("OSID", os)
	values.Add("SSHKEYID", sshkey)
	if networkid != "" {
		values.Add("NETWORKID[]", networkid)
	}
	if tag != "" {
		values.Add("tag", tag)
	}

	docker.OutLog("■■■INPUT■■■", window)
	docker.OutLog("https://api.vultr.com/v1/server/create", window)
	docker.OutLog(fmt.Sprintf("%s", values), window)

	req, _ := http.NewRequest("POST", "https://api.vultr.com/v1/server/create", strings.NewReader(values.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("API-Key", configData.VultrApiKey)
	client := new(http.Client)
	resp, err := client.Do(req)
	fmt.Println(err)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	respStr := string(byteArray)
	respStr = jsonedit.StripQ(respStr)
	fmt.Println(respStr)
	if err != nil {
		docker.OutLog("■■■OUTPUT ERROR■■■", window)
		docker.OutLog(err.Error(), window)
	} else {
		docker.OutLog("■■■OUTPUT■■■", window)
		docker.OutLog(respStr, window)
	}

}
func Destroy(id string, window *gotron.BrowserWindow) {
	configData := config.LoadConfig()

	values := url.Values{}
	values.Add("SUBID", id)

	docker.OutLog("■■■INPUT■■■", window)
	docker.OutLog("https://api.vultr.com/v1/server/destroy", window)
	docker.OutLog(fmt.Sprintf("%s", values), window)
	req, _ := http.NewRequest("POST", "https://api.vultr.com/v1/server/destroy", strings.NewReader(values.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("API-Key", configData.VultrApiKey)
	client := new(http.Client)
	resp, err := client.Do(req)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	respStr := string(byteArray)
	//respStr = jsonedit.StripQ(respStr) なぜかここでこける
	if err != nil {
		docker.OutLog("■■■OUTPUT ERROR■■■", window)
		docker.OutLog(err.Error(), window)
	} else {
		docker.OutLog("■■■OUTPUT■■■", window)
		docker.OutLog(respStr, window)
	}

}

func Sshkey() string {
	configData := config.LoadConfig()

	if configData.VultrApiKey == "" {
		output := jsonedit.Key("sshkey", "") // htmlをstringで取得
		return output
	}
	req, _ := http.NewRequest("GET", "https://api.vultr.com/v1/sshkey/list", nil)
	req.Header.Set("API-Key", configData.VultrApiKey)
	client := new(http.Client)
	resp, err := client.Do(req)
	fmt.Println(err)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	respStr := string(byteArray)
	respStr = jsonedit.StripQ(respStr)
	//fmt.Println(respStr)
	output := jsonedit.Key("sshkey", respStr) // htmlをstringで取得

	return output
}

func Region() string {

	req, _ := http.NewRequest("GET", "https://api.vultr.com/v1/regions/list", nil)
	client := new(http.Client)
	resp, err := client.Do(req)
	fmt.Println(err)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	respStr := string(byteArray)
	respStr = jsonedit.StripQ(respStr)
	//fmt.Println(respStr)
	output := jsonedit.Key("region", respStr) // htmlをstringで取得

	return output
}

func Plan() string {

	req, _ := http.NewRequest("GET", "https://api.vultr.com/v1/plans/list", nil)
	client := new(http.Client)
	resp, err := client.Do(req)
	fmt.Println(err)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	respStr := string(byteArray)
	respStr = jsonedit.StripQ(respStr)
	//fmt.Println(respStr)
	output := jsonedit.Key("plan", respStr) // htmlをstringで取得

	return output
}
func Os() string {

	req, _ := http.NewRequest("GET", "https://api.vultr.com/v1/os/list", nil)
	client := new(http.Client)
	resp, err := client.Do(req)
	fmt.Println(err)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	respStr := string(byteArray)
	respStr = jsonedit.StripQ(respStr)
	//fmt.Println(respStr)
	output := jsonedit.Key("os", respStr) // htmlをstringで取得

	return output
}
func Account() string {
	configData := config.LoadConfig()
	if configData.VultrApiKey == "" {
		output := jsonedit.Key("account", "") // htmlをstringで取得
		return output
	}
	req, _ := http.NewRequest("GET", "https://api.vultr.com/v1/account/info", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("API-Key", configData.VultrApiKey)
	client := new(http.Client)
	resp, err := client.Do(req)
	fmt.Println(err)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	respStr := string(byteArray)
	fmt.Println(respStr)
	output := jsonedit.On("account", respStr) // htmlをstringで取得

	return output
}
func Network() string {
	configData := config.LoadConfig()
	if configData.VultrApiKey == "" {
		output := jsonedit.Key("network", "") // htmlをstringで取得
		return output
	}
	req, _ := http.NewRequest("GET", "https://api.vultr.com/v1/network/list", nil)
	req.Header.Set("API-Key", configData.VultrApiKey)
	client := new(http.Client)
	resp, err := client.Do(req)
	fmt.Println(err)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	respStr := string(byteArray)
	respStr = jsonedit.StripQ(respStr)
	fmt.Println(respStr)
	output := jsonedit.Key("network", respStr) // htmlをstringで取得

	return output
}
