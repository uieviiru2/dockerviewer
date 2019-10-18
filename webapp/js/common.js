function formToData(form) {
    var output = {};
    var formList = form.serializeArray();
    for (var i = 0; i < formList.length; i++) {
        output[formList[i]['name']] = formList[i]['value'];
    }
    return output;
}
function replaceInnerView(elm, data) {
    elm.empty();
    elm.append(data);
}
function upload(file, emitTarget) {
    var fileReader = new FileReader();
    var send_file = file;
    var type = send_file.type;
    var data = {};
    fileReader.readAsBinaryString(send_file);
    fileReader.onload = function (event) {
        data.file = event.target.result;
        data.type = type;
        data.name = file.name;
        socket.emit(emitTarget, data);
    }
}
function getParam(name, url) {
    if (!url) url = window.location.href;
    name = name.replace(/[\[\]]/g, "\\$&");
    var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
        results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, " "));
}
function include(filename, afterfunc) {

    include.seq = (include.seq)? include.seq + 1: 1;
  
    var id = new Date().getTime() + "-" + include.seq;
    var inc = document.createElement("iframe");
  
    inc.id = "inc-" + id;
    inc.src = filename;
    inc.style.display = "none";
    document.write("<span id=\"" + id + "\"></span>");
      
    var incfunc = function() {
      
      var s = (function() {
        var suffix = (n = filename.lastIndexOf(".")) >= 0 ? filename.substring(n): "default";
        if (suffix == ".html") return inc.contentWindow.document.body.innerHTML;
        if (suffix == ".txt") return inc.contentWindow.document.body.firstChild.innerHTML;
        if (suffix == "default") return inc.contentWindow.document.body.innerHTML;
      })();
  
      var span = document.getElementById(id);
  
      var insertBeforeHTML = function(htmlStr, refNode) {
        if (document.createRange) {
          var range = document.createRange();
          range.setStartBefore(refNode);
          refNode.parentNode.insertBefore(range.createContextualFragment(htmlStr), refNode);
        } else {
          refNode.insertAdjacentHTML('BeforeBegin', htmlStr);
        }
      };
  
      insertBeforeHTML(s.split("&gt;").join(">").split("&lt;").join("<"), span);
      document.body.removeChild(inc);
      span.parentNode.removeChild(span);
      if (afterfunc) afterfunc();
    };
  
    if (window.attachEvent) {
      window.attachEvent('onload', 
        function() {
          document.body.appendChild(inc); 
          inc.onreadystatechange = function() { if (this.readyState == "complete") incfunc(); };
        });
    }
    else {
      document.body.appendChild(inc);
      inc.onload = incfunc;
    }
  }

  function explorer(path) {
    ws.send(JSON.stringify({
      "event": "explorer",
      "path":path
    }));

  }
  function webbrowser(url) {
    ws.send(JSON.stringify({
        "event": "index-webbrowser",
        "url":url
    }));
  }
  function view(viewFile) {
    ws.send(JSON.stringify({
        "event": "index-view",
        "page":viewFile
    }));
  }
  function NNSH_decodeHTML(str) {
    return str .replace(/&lt;/g, '<') .replace(/&gt;/g, '>') .replace(/&quot;/g, '"') .replace(/&#039;/g, '\'') .replace(/&#044;/g, ',') .replace(/&amp;/g, '&');
  }