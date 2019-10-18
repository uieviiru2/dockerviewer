if(ws==undefined)var ws;
if(ws)ws.close();
ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");
//通信が接続された場合
ws.onopen = function(e) { 
    console.log("Setup web socket:", ws);
    ws.send(JSON.stringify({
        "event": "config",
    }))  
    ws.onmessage = function (event) {

        const eObj = JSON.parse(event.data);
        console.log(eObj);
        const data = JSON.parse(eObj.event);
        if(data.eventName == "index-view") {
            $(".content-wrapper").html(NNSH_decodeHTML(data.view.view_data));
            return
        }
        else if(data.eventName == "log") {
            document.getElementById("console-output").innerHTML += "\n" + data.log;
            $('#console-output').animate({scrollTop: $('#console-output')[0].scrollHeight}, 'fast');
            return
        }
        else if(data.eventName == "config-save") {
            location.href="./index.html";
            return
        }
        console.log(data.config);
        $("#docker_exe").val(data.config.docker_exe);
        $("#docker_tmp_save_path").val(data.config.docker_tmp_save_path);
        $("#dockerfile_save_path").val(data.config.dockerfile_save_path);
        $("#vultr_api_key").val(data.config.vultr_api_key);
        $("#vultr_pem").val(data.config.vultr_pem);
        $("#config_dir").val(data.config.config_dir);
        $("#shell_exe").val(data.config.shell_exe);

    };
};
 
//エラーが発生した場合
ws.onerror = function(error) { };
 
//メッセージを受け取った場合
ws.onmessage = function(e) { };
  
//通信が切断された場合
ws.onclose = function() { };

//ページが読み込み完了した場合
$(document).ready(function () {

    $("#submit_button").click(function() {
        var fromData = formToData($("#frm"));
        fromData['event'] = "config-save"

        ws.send(JSON.stringify(fromData))
        console.log("Send success!!");
    });
    $("#select_dir").on('change', function() {
        for (i = 0; i < this.files.length; i++) {
            $("#docker_exe").val(this.files[i].path)
          }
    });

    $("#select_dir3").on('change', function() {
        for (i = 0; i < this.files.length; i++) {
            $("#vultr_pem").val(this.files[i].path)
          }
    });
    $("#select_dir5").on('change', function() {
        for (i = 0; i < this.files.length; i++) {
            $("#config_dir").val(this.files[i].path)
          }
    });
    $("#select_dir6").on('change', function() {
        for (i = 0; i < this.files.length; i++) {
            $("#shell_exe").val(this.files[i].path)
          }
    });
});
