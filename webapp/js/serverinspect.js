var ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");
//通信が接続された場合
ws.onopen = function(e) { 
    console.log("Setup web socket:", ws);
    ws.send(JSON.stringify({
        "event": "serverinspect",
        "ip": getParam('ip'),
        "v": getParam('v'),
    }))
    ws.onmessage = function (event) {
        const eObj = JSON.parse(event.data);
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
        if(data.ps && data.ps.length) {
            $('#docker-ps').columns({data:data.ps,
            schema: [
                {"header":"CONTAINER ID", "key":"ID", "template":'<a href="./serverinspectps.html?id={{ID}}&ip='+getParam('ip')+'&v='+ getParam('v') +'">{{ID}}</a>'},
                {"header":"IMAGE", "key":"Image"},
                {"header":"COMMAND", "key":"Command"},
                {"header":"CREATED", "key":"CreatedAt"},
                {"header":"STATUS", "key":"Status"},
                {"header":"PORTS", "key":"Ports"},
                {"header":"NAMES", "key":"Names"},
            ]
            });
        }
        if(data.image && data.image.length) {
            $('#docker-image').columns({data:data.image,
                schema: [
                    {"header":"REPOSITORY", "key":"Repository"},
                    {"header":"TAG", "key":"Tag"},
                    {"header":"IMAGE ID", "key":"ID", "template":'<a href="./serverinspectimage.html?id={{ID}}&ip='+getParam('ip')+'&v='+ getParam('v') +'">{{ID}}</a>'},
                    {"header":"CREATED", "key":"CreatedAt"},
                    {"header":"SIZE", "key":"Size"},
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
    $("#host_title").html("Server Inspect " + getParam('ip'));
    $("#submit_button").click(function() {
        var fromData = formToData($("#frm"));
        fromData['event'] = "serverinspect-bash"
        fromData['ip'] =  getParam('ip');
        fromData['v'] =  getParam('v');
        ws.send(JSON.stringify(fromData));
        console.log("Send success!!");
    });
});
function removeall() {
    if(confirm("DLETE ALL?")) {
        ws.send(JSON.stringify({
            "event": "serverinspect-removeall",
            "ip": getParam('ip'),
            "v": getParam('v'),
        }))
    }
}