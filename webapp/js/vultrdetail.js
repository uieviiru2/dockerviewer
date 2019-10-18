if(ws==undefined)var ws;
if(ws)ws.close();
ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");
//通信が接続された場合
ws.onopen = function(e) { 

    console.log("Setup web socket:", ws);
    ws.send(JSON.stringify({
        "event": "vultrdetail",
        "id": getParam('id'),
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
        $('#vultr-detail').jsonViewer(data.vultrDetail);

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
    $("#id").val(getParam('id'))

    $("#submit_button").click(function() {
        var fromData = formToData($("#frm"));
        fromData['event'] = "vultrdetail-destroy"

        ws.send(JSON.stringify(fromData))
        console.log("Send success!!");
    });
});
