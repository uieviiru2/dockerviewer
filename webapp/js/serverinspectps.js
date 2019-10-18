var ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");
//通信が接続された場合
ws.onopen = function(e) { 
    console.log("Setup web socket:", ws);
    ws.send(JSON.stringify({
        "event": "serverinspectps",
        "id": getParam('id'),
        "ip": getParam('ip'),
        "v": getParam('v')
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

        $('#docker-inspect').jsonViewer(data.inspect);
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
    $("#host_title").html("Server Inspect " + getParam('ip'))
    $("#submit_button2").click(function() {
        var fromData = formToData($("#frm"));
        fromData['event'] = "serverinspectps-remove"
        fromData['id'] =  getParam('id');
        fromData['ip'] =  getParam('ip');
        fromData['v'] =  getParam('v');
        ws.send(JSON.stringify(fromData));
        console.log("Send success!!");
    });
    $("#submit_button3").click(function() {
        var fromData = formToData($("#frm"));
        fromData['id'] =  getParam('id');
        fromData['ip'] =  getParam('ip');
        fromData['v'] =  getParam('v');
        fromData['event'] = "serverinspectps-bash"

        ws.send(JSON.stringify(fromData))
        console.log("Send success!!");
    });
});
