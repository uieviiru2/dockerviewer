Dim oShell
Set oShell = WScript.CreateObject ("WSCript.shell")
oShell.run "go run main.go",0
Set oShell = Nothing