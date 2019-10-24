if(ws==undefined)var ws;
if(ws)ws.close();
ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");
//通信が接続された場合
ws.onopen = function(e) { 
    console.log("Setup web socket:", ws);
    ws.send(JSON.stringify({
        "event": "dockerrun"
    }))  
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
        else if(data.eventName == "dockerrun-load") {
            
            $("#name").val(data.runInputData.name);
            $("#dit").val(data.runInputData.dit);
            $("#port").val(data.runInputData.port);
            $("#option").val(data.runInputData.option);
            $("#dirname").val(data.runInputData.dirname);
            $("#dirname2").val(data.runInputData.dirname2);
            $("#dirname_a").val(data.runInputData.dirname_a);
            $("#dirname_a2").val(data.runInputData.dirname_a2);
            $("#dirname_b").val(data.runInputData.dirname_b);
            $("#dirname_b2").val(data.runInputData.dirname_b2);
            $("#dirname_c").val(data.runInputData.dirname_c);
            $("#dirname_c2").val(data.runInputData.dirname_c2);
            $("#dirname_d").val(data.runInputData.dirname_d);
            $("#dirname_d2").val(data.runInputData.dirname_d2);
            $("#option2").val(data.runInputData.option2);
            //ないかもしれないのでラスト
            $("#docker").val(data.runInputData.docker);
            return
        }
        if(data.eventName == "dockerrun-saveinput") {

            try {$('#inputlist').columns('destroy');} catch(e) {}

            if(data.runInputlist) {
                $('#inputlist').columns({data:data.runInputlist,
                    schema: [
                        {"header":"NAME", "key":"name", "template":'<a href="#" onclick="loadRunInput(\'{{name}}\')">{{name}}</a>'},
                        {"header":"", "key":"cmd"},
                        {"header":" ", "key":"", "template":'<input type="button" onclick="deleteRunInput(\'{{name}}\')" value="DELETE" />'},
                ]
                });
            }
            return;
        }
        if(data.eventName == "dockerrun-delete") {

            try {$('#inputlist').columns('destroy');} catch(e) {}

            if(data.runInputlist) {
                $('#inputlist').columns({data:data.runInputlist,
                    schema: [
                        {"header":"NAME", "key":"name", "template":'<a href="#" onclick="loadRunInput(\'{{name}}\')">{{name}}</a>'},
                        {"header":"", "key":"cmd"},
                        {"header":" ", "key":"", "template":'<input type="button" onclick="deleteRunInput(\'{{name}}\')" value="DELETE" />'},
                ]
                });
            }
            return;
        }
        $.each(data.image, function (key, val) {
            $('#docker').append($('<option>').val(val.ID).text(val.Repository+"("+val.ID+")   "+ val.CreatedAt));
        });
        console.log(data.runInputlist);
        if(data.runInputlist) {
            $('#inputlist').columns({data:data.runInputlist,
                schema: [
                    {"header":"NAME", "key":"name", "template":'<a href="#" onclick="loadRunInput(\'{{name}}\')">{{name}}</a>'},
                    {"header":"", "key":"cmd"},
                    {"header":" ", "key":"", "template":'<input type="button" onclick="deleteRunInput(\'{{name}}\')" value="DELETE" />'},
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
        fromData['event'] = "dockerrun-excute"

        ws.send(JSON.stringify(fromData))
        console.log("Send success!!");
    });
    $("#submit_button2").click(function() {
        var fromData = formToData($("#frm"));
        if($("#name").val() == "") {
            alert("require name")
            return;
        }


        fromData['event'] = "dockerrun-saveinput"

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
function loadRunInput(name) {

    ws.send(JSON.stringify({
        "event": "dockerrun-load",
        "name":name
    }));
}
function deleteRunInput(name) {

    ws.send(JSON.stringify({
        "event": "dockerrun-delete",
        "name":name
    }));
}