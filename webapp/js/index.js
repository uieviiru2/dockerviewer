if(ws==undefined)var ws;
if(ws)ws.close();
ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");
//通信が接続された場合
ws.onopen = function(e) { 

    console.log("Setup web socket:", ws);
    ws.send(JSON.stringify({
        "event": "index",
    }))  
    ws.onmessage = function (event) {
        const eObj = JSON.parse(event.data);
        console.log(event.data);
        
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
        else if(data.eventName == "index-config") {
            location.href="./index.html?page=config";
            return
        }
        if(data.configData.vultr_api_key =="" || data.configData.vultr_pem =="") {
            $(".level2").hide();
        } else {
            $(".level2").show();
        }
        if(data.configData.docker_exe =="" || data.configData.config_dir == "") {
            $(".level1").hide();
            view('config')
        } else {
            $(".level1").show();
        }


        if(data.machine && data.machine.length) {
            $('#docker-machine').columns({data:data.machine,
                schema: [
                    {"header":"DOCKER", "key":"DOCKER"},
                    {"header":"RESPONSE", "key":"RESPONSE"},
                    {"header":"NAME", "key":"NAME"},
                    {"header":"ACTIVE_SWARM", "key":"ACTIVE_SWARM"},
                    {"header":"STATE", "key":"STATE"},
                    {"header":"SWARM_OPTIONS", "key":"SWARM_OPTIONS"},
                    {"header":"ERRORS", "key":"ERRORS"},
                    {"header":"ACTIVE", "key":"ACTIVE"},
                    {"header":"ACTIVE_HOST", "key":"ACTIVE_HOST"},
                    {"header":"DRIVER", "key":"DRIVER"},
                    {"header":"URL", "key":"URL"},
                    {"header":" ", "key":"NAME", "template":'<button class="machinereloadbutton" name="reload" onclick="machinereload(\'{{NAME}}\')";>RESTART</button><br /><button class="machinereloadbutton" name="ssh" onclick="machinessh(\'{{NAME}}\')";>SSH</button>'},
                ]
                });
        }
        if(data.ps && data.ps.length) {
            $('#docker-ps').columns({data:data.ps,
            schema: [
                {"header":"CONTAINER ID", "key":"ID", "template":'<a onclick="history.replaceState(\'\',\'\',\'index.html?id={{ID}}\');view(\'dockerdetail\');" href="javascript:void(0);">{{ID}}</a>'},
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
                    {"header":"IMAGE ID", "key":"ID", "template":'<a onclick="history.replaceState(\'\',\'\',\'index.html?id={{ID}}\');view(\'dockerimagedetail\');" href="javascript:void(0);">{{ID}}</a>'},
                    {"header":"CREATED", "key":"CreatedAt"},
                    {"header":"SIZE", "key":"Size"},
            ]
            });
        }

        if(data.dockerfiles) {
            $('#dockerfile').columns({data:data.dockerfiles,
                schema: [
                    {"header":"NAME", "key":"name", "template":'<a href="javascript:void(0);" onclick="explorer(\''+ data.path1.replace(/\\/g, "\\\\") + '\\\\{{name}}\')" >{{name}}</a>'},
                    {"header":"CREATED", "key":"created_at"},
            ]
            });
        }
        if(data.dockercompose) {
            $('#dockercompose').columns({data:data.dockercompose,
                schema: [
                    {"header":"NAME", "key":"name", "template":'<a href="javascript:void(0);" onclick="explorer(\''+ data.path2.replace(/\\/g, "\\\\") + '\\\\{{name}}\')" >{{name}}</a>'},
                    {"header":"CREATED", "key":"created_at"},
            ]
            });
        }
        if(data.serverlist) {
            $('#server-list').columns({data:data.serverlist,
                schema: [
                    {"header":"IP OR HOST", "key":"ip", "template":'<a href="serverinspect.html?ip={{ip}}&v=0" target="_blank">{{ip}}</a>'},
                    {"header":"USER", "key":"user"},
                    {"header":"PEM FILE", "key":"server_pem"},
            ]
            });
        }

        //if(data.vultr && data.vultr.length) {
            //console.log(data.vultr);
            var vultr_list = [];
            var vlutr_data = data.vultr;
            console.log(vlutr_data.length);
            for (key in vlutr_data) {
                console.log('key:' + key + ' value:' + vlutr_data[key]);
                vultr_list.push(vlutr_data[key]);
            }
            console.log(vultr_list)
            $('#vultr-list').columns({data:vultr_list,
                schema: [
                    {"header":"SUBID", "key":"SUBID", "template":'<a href="./index.html?page=vultrdetail&id={{SUBID}}">{{SUBID}}</a>'},
                    {"header":"OS", "key":"os"},
                    {"header":"RAM", "key":"ram"},
                    {"header":"DISK", "key":"disk"},
                    {"header":"MAIN IP", "key":"main_ip", "template":'<a href="serverinspect.html?ip={{main_ip}}&v=1" target="_blank">{{main_ip}}</a>'},
                    {"header":"CREATED", "key":"date_created"},
                    {"header":"STATUS", "key":"status"},
            ]
            });

            
        //}

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


});

function machinereload(machine) {
    var formData = {};
    formData['event'] = "index-macinereload";
    formData['machine_name'] = machine;
    
    ws.send(JSON.stringify(formData))
    console.log("Send success!!");
}
function machinessh(machine) {
    var formData = {};
    formData['event'] = "index-machinessh";
    formData['machine_name'] = machine;
    
    ws.send(JSON.stringify(formData))
    console.log("Send success!!");
}
function machinestart() {
    var formData = {};
    formData['event'] = "index-machinestart";
    ws.send(JSON.stringify(formData))
    console.log("Send success!!");
}
function machinecreate() {
    var formData = {};
    formData['event'] = "index-machinecreate";
    ws.send(JSON.stringify(formData))
    console.log("Send success!!");
}
