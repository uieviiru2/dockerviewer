if(ws==undefined)var ws;
if(ws)ws.close();
ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");
//通信が接続された場合
ws.onopen = function(e) { 
    console.log("Setup web socket:", ws);
    ws.send(JSON.stringify({
        "event": "configserver",
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
        else if(data.eventName == "configserver-save") {
            location.reload();
            return
        }
        if(data.eventName == "configserver-delete") {
            location.reload();
            return
        }
        if(data.eventName == "configserver-load") {
            $("#ip").val(data.serverfile.ip);
            $("#user").val(data.serverfile.user)
            $("#server_pem").val(data.serverfile.server_pem)
            return
        }
   


        if(data.serverlist) {
            $('#serverlist').columns({data:data.serverlist,
                schema: [
                    {"header":"IP OR HOST", "key":"ip", "template":'<a href="#" onclick="loadServerFile(\'{{ip}}\')">{{ip}}</a>'},
                    {"header":"USER", "key":"user"},
                    {"header":"PEM FILE", "key":"server_pem"},
                    {"header":" ", "key":"", "template":'<input type="button" onclick="deleteServerFile(\'{{ip}}\')" value="DELETE" />'},
            ]
            });
        }
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
        if($("#ip").val() == "") {
            alert("require host or ip")
            return;
        }
        if($("#user").val() == "") {
            alert("require user")
            return;
        }
        if($("#server_pem").val() == "") {
            alert("require PEM FILE")
            return;
        }
        fromData['event'] = "configserver-save"

        ws.send(JSON.stringify(fromData))
        console.log("Send success!!");
    });
    $("#select_dir").on('change', function() {
        for (i = 0; i < this.files.length; i++) {
            $("#server_pem").val(this.files[i].path)
          }
    });

});
function deleteServerFile(ip) {

    ws.send(JSON.stringify({
        "event": "configserver-delete",
        "ip":ip
    }));
}
function loadServerFile(ip) {

    ws.send(JSON.stringify({
        "event": "configserver-load",
        "ip":ip
    }));
}
