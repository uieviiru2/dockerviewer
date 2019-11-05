package main

import (
	"github.com/Equanox/gotron"
	"github.com/uieviiru2/mylib/socket/config"
	"github.com/uieviiru2/mylib/socket/configserver"
	"github.com/uieviiru2/mylib/socket/dockercompose"
	"github.com/uieviiru2/mylib/socket/dockerdetail"
	"github.com/uieviiru2/mylib/socket/dockerfile"
	"github.com/uieviiru2/mylib/socket/dockerimagedeploy"
	"github.com/uieviiru2/mylib/socket/dockerimagedetail"
	"github.com/uieviiru2/mylib/socket/dockerpull"
	"github.com/uieviiru2/mylib/socket/dockerrun"
	"github.com/uieviiru2/mylib/socket/index"
	"github.com/uieviiru2/mylib/socket/registry"
	"github.com/uieviiru2/mylib/socket/serverinspect"
	"github.com/uieviiru2/mylib/socket/serverinspectimage"
	"github.com/uieviiru2/mylib/socket/serverinspectps"
	"github.com/uieviiru2/mylib/socket/vultrdetail"
	"github.com/uieviiru2/mylib/socket/vultrrun"
)

func main() {

	// browser window instanceを生成する
	window, err := gotron.New("webapp")
	if err != nil {
		panic(err)
	}
	// デフォルトのwindowサイズとタイトルを変更する
	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 980
	window.WindowOptions.Title = "docker-vltr"
	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	// dev toolsを使う場合は下記をコメントアウトする
	//window.OpenDevTools()

	index.Use(window)
	dockerrun.Use(window)
	dockerdetail.Use(window)
	dockerimagedetail.Use(window)
	dockerpull.Use(window)
	dockerfile.Use(window)
	dockercompose.Use(window)
	dockerimagedeploy.Use(window)
	registry.Use(window)
	vultrrun.Use(window)
	vultrdetail.Use(window)
	config.Use(window)
	configserver.Use(window)
	serverinspect.Use(window)
	serverinspectps.Use(window)
	serverinspectimage.Use(window)
	// アプリケーションがcloseするのを待つ
	<-done
}
