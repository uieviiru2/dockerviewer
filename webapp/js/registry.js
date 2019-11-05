if(ws==undefined)var ws;
if(ws)ws.close();
ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");
//通信が接続された場合
ws.onopen = function(e) {
    ws.send(JSON.stringify({
        "event": "registry"
    }));
    ws.onmessage = function (event) {
        console.log(event.data);
        const eObj = JSON.parse(event.data);
        console.log(eObj);
        const data = JSON.parse(eObj.event);
        console.log(data)
        if(data.eventName == "index-view") {
            $(".content-wrapper").html(NNSH_decodeHTML(data.view.view_data));
            return
        }
        else if(data.eventName == "log") {
            document.getElementById("console-output").innerHTML += "\n" + data.log;
            $('#console-output').animate({scrollTop: $('#console-output')[0].scrollHeight}, 'fast');
            return
        }
        else if(data.eventName == "registry") {


            $.each(data.tags, function (key, val) {
                $('#image_id').append($('<option>').val("registry:" +val.name).text("registry:" +val.name));
            });
            if(data.serverlist) {
                $.each(data.serverlist, function (key, val) {
                    $('#server_id').append($('<option>').val(val.ip).text("server:"+val.user+"@"+val.ip));
                });
            }
            if(data.vultr) {
                $.each(data.vultr, function (key, val) {
                    $('#server_id').append($('<option>').val(val.main_ip).text("vultr:"+key+":"+val.main_ip+"("+val.os+")"));
                });
            }
        }
        if(data.eventName == "registry" || data.eventName == "registry-deploy" || data.eventName == "registry-delete") {
            if(data.registrylist) {
                try {$('#registry-list').columns('destroy');} catch(e) {};
                $('#registry-list').columns({data:data.registrylist,
                    schema: [
                        {"header":"NAME", "key":"name", "template":'<a href="#" onclick="loadDockerFile(\'{{name}}\')">{{name}}</a>'},
                        {"header":"SEVER IP", "key":"server_ip"},
                        {"header":"DELETE", "key":"", "template":'<input type="button" onclick="deleteRegistryFile(\'{{name}}\')" value="DELETE" />'},
                ]
                });
            }
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
        fromData['event'] = "registry-deploy"

        ws.send(JSON.stringify(fromData))
        console.log("Send success!!");
    });
});
function deleteRegistryFile(name) {

    ws.send(JSON.stringify({
        "event": "registry-delete",
        "name":name
    }));
}