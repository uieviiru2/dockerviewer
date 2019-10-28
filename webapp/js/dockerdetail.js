if(ws==undefined)var ws;
if(ws)ws.close();
ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");
//通信が接続された場合
ws.onopen = function(e) { 
    console.log("Setup web socket:", ws);
    ws.send(JSON.stringify({
        "event": "dockerdetail",
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
        if(data.eventName == "dockerdetail-load") {
            $("#name").val(data.deployInputData.name);
            $("#dit").val(data.deployInputData.dit);
            $("#port").val(data.deployInputData.port);
            $("#option").val(data.deployInputData.option);
            $("#dirname").val(data.deployInputData.dirname);
            $("#dirname2").val(data.deployInputData.dirname2);
            $("#dirname_a").val(data.deployInputData.dirname_a);
            $("#dirname_a2").val(data.deployInputData.dirname_a2);
            $("#dirname_b").val(data.deployInputData.dirname_b);
            $("#dirname_b2").val(data.deployInputData.dirname_b2);
            $("#dirname_c").val(data.deployInputData.dirname_c);
            $("#dirname_c2").val(data.deployInputData.dirname_c2);
            $("#dirname_d").val(data.deployInputData.dirname_d);
            $("#dirname_d2").val(data.deployInputData.dirname_d2);
            $("#option2").val(data.deployInputData.option2);
            //ないかもしれないのでラスト
            $("#server_id").val(data.deployInputData.server_id);
            return
        }
        if(data.eventName == "dockerdetail-saveinput") {
            if(data.deployInputlist) {
                try {$('#inputlist').columns('destroy');} catch(e) {}
                $('#inputlist').columns({data:data.deployInputlist,
                    schema: [
                        {"header":"NAME", "key":"name", "template":'<a href="#" onclick="loadDeployInput(\'{{name}}\')">{{name}}</a>'},
                        {"header":"", "key":"cmd"},
                        {"header":" ", "key":"", "template":'<input type="button" onclick="deleteDeployInput(\'{{name}}\')" value="DELETE" />'},
                ]
                });
            }
            return
        }
        if(data.eventName == "dockerdetail-delete") {
            if(data.deployInputlist) {
                try {$('#inputlist').columns('destroy');} catch(e) {}
                $('#inputlist').columns({data:data.deployInputlist,
                    schema: [
                        {"header":"NAME", "key":"name", "template":'<a href="#" onclick="loadDeployInput(\'{{name}}\')">{{name}}</a>'},
                        {"header":"", "key":"cmd"},
                        {"header":" ", "key":"", "template":'<input type="button" onclick="deleteDeployInput(\'{{name}}\')" value="DELETE" />'},
                ]
                });
            }
            return
        }

        $('#docker-inspect').jsonViewer(data.inspect);
        console.log(data.vultr);
        $.each(data.serverlist, function (key, val) {
            $('#server_id').append($('<option>').val(val.ip).text("server:"+val.user+"@"+val.ip));
        });
        $.each(data.vultr, function (key, val) {
            $('#server_id').append($('<option>').val(val.main_ip).text("vultr:"+key+":"+val.main_ip+"("+val.os+")"));
        });
        console.log(data.inspect);
        if(data.inspect[0] && data.inspect[0].HostConfig && data.inspect[0].HostConfig.Binds) {
            if(data.inspect[0].HostConfig.Binds[0] ) {
                var result = data.inspect[0].HostConfig.Binds[0].split(':');
                if(result[0])$('#dirname').val(convertWindowsDirectoryForm(result[0]));
                if(result[1])$('#dirname2').val(result[1]);
            }
            if(data.inspect[0].HostConfig.Binds[1] ) {
                var result = data.inspect[0].HostConfig.Binds[1].split(':');
                if(result[0])$('#dirname_a').val(convertWindowsDirectoryForm(result[0]));
                if(result[1])$('#dirname_a2').val(result[1]);
            }
            if(data.inspect[0].HostConfig.Binds[2] ) {
                var result = data.inspect[0].HostConfig.Binds[2].split(':');
                if(result[0])$('#dirname_b').val(convertWindowsDirectoryForm(result[0]));
                if(result[1])$('#dirname_b2').val(result[1]);
            }
            if(data.inspect[0].HostConfig.Binds[3] ) {
                var result = data.inspect[0].HostConfig.Binds[3].split(':');
                if(result[0])$('#dirname_c').val(convertWindowsDirectoryForm(result[0]));
                if(result[1])$('#dirname_c2').val(result[1]);
            }
            if(data.inspect[0].HostConfig.Binds[4] ) {
                var result = data.inspect[0].HostConfig.Binds[4].split(':');
                if(result[0])$('#dirname_d').val(convertWindowsDirectoryForm(result[0]));
                if(result[1])$('#dirname_d2').val(result[1]);
            }
        }
        if(data.deployInputlist) {
            $('#inputlist').columns({data:data.deployInputlist,
                schema: [
                    {"header":"NAME", "key":"name", "template":'<a href="#" onclick="loadDeployInput(\'{{name}}\')">{{name}}</a>'},
                    {"header":"", "key":"cmd"},
                    {"header":" ", "key":"", "template":'<input type="button" onclick="deleteDeployInput(\'{{name}}\')" value="DELETE" />'},
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
    $("#id").val(getParam('id'))

    $("#submit_button").click(function() {
        var fromData = formToData($("#frm"));
        fromData['event'] = "dockerdetail-deploy"
        ws.send(JSON.stringify(fromData))
        console.log("Send success!!");
    });
    $("#submit_button2").click(function() {
        var fromData = formToData($("#frm"));
        fromData['event'] = "dockerdetail-remove"
        ws.send(JSON.stringify(fromData))
        console.log("Send success!!");
    });
    $("#submit_button3").click(function() {
        var fromData = formToData($("#frm"));
        fromData['event'] = "dockerdetail-bash"
        ws.send(JSON.stringify(fromData))
        console.log("Send success!!");
    });
    $("#submit_button4").click(function() {
        var fromData = formToData($("#frm"));
        fromData['event'] = "dockerdetail-saveinput"
        ws.send(JSON.stringify(fromData))
        console.log("Send success!!");
    });
    $("#submit_button5").click(function() {
        var fromData = formToData($("#frm2"));
        fromData.id = getParam('id')
        if($("#iamge_name").val() == "") {
            alert("require name")
            return;
        }
        fromData['event'] = "dockerdetail-saveimage"
        ws.send(JSON.stringify(fromData))
        console.log("Send success!!");
    });
    $("#submit_button6").click(function() {
        var fromData = formToData($("#frm"));
        fromData['event'] = "dockerdetail-start"
        ws.send(JSON.stringify(fromData))
        console.log("Send success!!");
    });
    $("#submit_button7").click(function() {
        var fromData = formToData($("#frm"));
        fromData['event'] = "dockerdetail-stop"
        ws.send(JSON.stringify(fromData))
        console.log("Send success!!");
    });
    $("#select_dir").on('change', function() {
        for (i = 0; i < this.files.length; i++) {
            $("#dirname").val(this.files[i].path)
          }
    });
    $("#select_dir2").on('change', function() {
        for (i = 0; i < this.files.length; i++) {
            $("#dirname_a").val(this.files[i].path)
          }
    });
    $("#select_dir3").on('change', function() {
        for (i = 0; i < this.files.length; i++) {
            $("#dirname_b").val(this.files[i].path)
          }
    });
    $("#select_dir4").on('change', function() {
        for (i = 0; i < this.files.length; i++) {
            $("#dirname_c").val(this.files[i].path)
          }
    });
    $("#select_dir5").on('change', function() {
        for (i = 0; i < this.files.length; i++) {
            $("#dirname_d").val(this.files[i].path)
          }
    });
});
function convertWindowsDirectoryForm(path) {
    var initChar = path.charAt(1).toUpperCase();
    path =  path.replace(/\//g, '\\');
    return initChar + ":" + path.substr(2);
}
function loadDeployInput(name) {

    ws.send(JSON.stringify({
        "event": "dockerdetail-load",
        "name":name
    }));
}
function deleteDeployInput(name) {

    ws.send(JSON.stringify({
        "event": "dockerdetail-delete",
        "name":name
    }));
}