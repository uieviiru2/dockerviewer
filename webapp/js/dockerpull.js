if(ws==undefined)var ws;
if(ws)ws.close();
ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");
//通信が接続された場合
ws.onopen = function(e) { 
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
        $.each(data.search, function(index, value){
            if(value.IsOfficial=="true") {
                data.search[index].Url = "https://hub.docker.com/_/" + value.Name;
            } else {
                data.search[index].Url = "https://hub.docker.com/r/" + value.Name;
            }
        });
        if(data.search && data.search.length) {

            try {$('#docker-search').columns('destroy');} catch(e) {}
            
            $('#docker-search').columns({data:data.search,
                schema: [
                    {"header":"Name", "key":"Url","template":'<a href="{{Url}}" target="_blank">{{Name}}</a>'},
                    {"header":"Description", "key":"Description"},
                    {"header":"IsAutomated", "key":"IsAutomated"},
                    {"header":"IsOfficial", "key":"IsOfficial"},
                    {"header":"StarCount", "key":"StarCount"},
                    {"header":"", "key":"Name", "template":'<form onsubmit="return dockerpullexecute(this);">{{Name}}:<input type="text" name="sub" value="" /><input type="hidden" name="name" value="{{Name}}" /><input type="submit" class="douckerpullbutton" name="pull" value="PULL"></form>'},
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
        fromData['event'] = "dockerpull-search";
        ws.send(JSON.stringify(fromData));
        console.log("Send success!!");
    });
    $("#select_dir").on('change', function() {
        for (i = 0; i < this.files.length; i++) {
            $("#dirname").val(this.files[i].path)
          }
    });
});
/*
function dockerpullexecute(dockername) {
    var formData = {};
    formData['event'] = "dockerpull-excute";
    formData['pull_name'] = dockername;
    
    ws.send(JSON.stringify(formData))
    console.log("Send success!!");
}
*/
function dockerpullexecute(frm) {
    var formData = {};
    formData['event'] = "dockerpull-excute";
    formData['pull_name'] = frm.elements["name"].value;
    if(frm.elements["sub"].value.trim() != "") {
        formData['pull_name'] += ":" + frm.elements["sub"].value.trim();
    }
    ws.send(JSON.stringify(formData))
    return false;
}
document.onkeypress = enter;
function enter(){
  if( window.event.keyCode == 13 ){

    var fromData = formToData($("#frm"));
    fromData['event'] = "dockerpull-search";
    ws.send(JSON.stringify(fromData));
    console.log("Send success!!");
    
    return false;
  }
}