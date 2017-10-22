setlocal
del hello.exe
set GOOS=windows
::set GOARCH=amd64
set GOARCH=386
go build hello.go
cygcheck32 hello.exe
dir *.exe
endlocal