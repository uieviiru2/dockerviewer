if(ws==undefined)var ws;
if(ws)ws.close();
ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");
//通信が接続された場合
ws.onopen = function(e) { 
    console.log("Setup web socket:", ws);
    ws.send(JSON.stringify({
        "event": "vultrrun",
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
        $.each(data.region, function (key, val) {
                $('#region').append($('<option>').val(key).text(val.country + ":" + val.name));
        });
        $('#region').val(25);
        $.each(data.plan, function (key, val) {
            $('#plan').append($('<option>').val(key).text(val.plan_type+":"+val.name+":"+ val.price_per_month + "💲/month"));
        });
        $('#plan').val(400);
        $.each(data.os, function (key, val) {
            $('#os').append($('<option>').val(key).text(val.name));
        });
        $('#os').val(179);
        $.each(data.sshkey, function (key, val) {
            $('#sshkey').append($('<option>').val(key).text(val.name));
        });
        $('#sshkey').val("5d79caa41797b");
        $.each(data.network, function (key, val) {
            $('#networkid').append($('<option>').val(key).text(val.description+"(" + val.v4_subnet + "/"+ val.v4_subnet_mask + ")"));
        });
        $('#networkid').val("net5d81b9b42edd5");

        
        alert("balance:" + data.account.balance);
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
        fromData['event'] = "vultrrun-create"

        ws.send(JSON.stringify(fromData))
        console.log("Send success!!");
    });
    $("#select_dir").on('change', function() {
        for (i = 0; i < this.files.length; i++) {
            $("#dirname").val(this.files[i].path)
          }
    });
});
