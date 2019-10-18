if(ws==undefined)var ws;
if(ws)ws.close();
ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");
//通信が接続された場合
ws.onopen = function(e) { 

    ws.send(JSON.stringify({
        "event": "dockerfile"
    }));


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
        if(data.eventName == "dockerfile-load") {
            $("#name").val(data.name);
            $("#script").val(data.script)
            return
        }
        if(data.dockerfiles && data.dockerfiles) {
            $('#dockerfile').columns({data:data.dockerfiles,
                schema: [
                    {"header":"NAME", "key":"name", "template":'<a href="#" onclick="loadDockerFile(\'{{name}}\')">{{name}}</a>'},
                    {"header":"CREATED", "key":"created_at"},
                    {"header":"DIRECTORY", "key":"name", "template":'<input type="button" onclick="explorer(\''+ data.path.replace(/\\/g, "\\\\") + '\\\\{{name}}\')" value="DIRECTORY" />'},
                    {"header":"DELETE", "key":"", "template":'<input type="button" onclick="deleteDockerFile(\'{{name}}\')" value="DELETE" />'},
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
        fromData['event'] = "dockerfile-save"

        ws.send(JSON.stringify(fromData))
        console.log("Send success!!");
    });
    $("#submit_button2").click(function() {
        var fromData = formToData($("#frm"));
        fromData['event'] = "dockerfile-test"

        ws.send(JSON.stringify(fromData))
        console.log("Send success!!");
    });
});
function deleteDockerFile(name) {

    ws.send(JSON.stringify({
        "event": "dockerfile-delete",
        "name":name
    }));
}
function loadDockerFile(name) {

    ws.send(JSON.stringify({
        "event": "dockerfile-load",
        "name":name
    }));
}
