@echo off
color 0a
title Compile EasyDoc

set goName=encryptTool.go
set execName=encryptTool
set version=1.0

::Cross compile is not support CGO
set CGO_ENABLED=0



@echo Compile windows 64 bit...
set GOOS=windows
set GOARCH=amd64
set GOEXE=.exe
::mkdir %execName%-%version%-%GOOS%-%GOARCH%
go build -ldflags="-H windowsgui"
::去掉符号表后不知为何系统崩溃了go build -ldflags="-s -w" -o %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE% %goName%
::去壳工具报错，不知道是啥意思upx %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE%
::rar a %execName%-%version%-%GOOS%-%GOARCH%.zip -r %execName%-%version%-%GOOS%-%GOARCH%
@echo.

goto end

@echo Compile windows 32 bit...
set GOOS=windows
set GOARCH=386
set GOEXE=.exe
mkdir %execName%-%version%-%GOOS%-%GOARCH%
go build -ldflags="-s -w" -o %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE% %goName%
upx %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE%
::rar a %execName%-%version%-%GOOS%-%GOARCH%.zip -r %execName%-%version%-%GOOS%-%GOARCH%
@echo.

@echo Compile Linux 32 bit...
set GOOS=linux
set GOARCH=386
set GOEXE=
mkdir %execName%-%version%-%GOOS%-%GOARCH%
go build -ldflags="-s -w" -o %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE% %goName%
upx %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE%
::rar a %execName%-%version%-%GOOS%-%GOARCH%.zip -r %execName%-%version%-%GOOS%-%GOARCH%
@echo.

@echo Compile Linux 64 bit...
set GOOS=linux
set GOARCH=amd64
set GOEXE=
mkdir %execName%-%version%-%GOOS%-%GOARCH%
go build -ldflags="-s -w" -o %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE% %goName%
upx %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE%
::rar a %execName%-%version%-%GOOS%-%GOARCH%.zip -r %execName%-%version%-%GOOS%-%GOARCH%
@echo.

@echo Compile Mac 32 bit...
set GOOS=darwin
set GOARCH=386
set GOEXE=
mkdir %execName%-%version%-%GOOS%-%GOARCH%
go build -ldflags="-s -w" -o %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE% %goName%
upx %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE%
::rar a %execName%-%version%-%GOOS%-%GOARCH%.zip -r %execName%-%version%-%GOOS%-%GOARCH%
@echo.

@echo Compile Mac 64 bit...
set GOOS=darwin
set GOARCH=amd64
set GOEXE=
mkdir %execName%-%version%-%GOOS%-%GOARCH%
go build -ldflags="-s -w" -o %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE% %goName%
upx %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE%
::rar a %execName%-%version%-%GOOS%-%GOARCH%.zip -r %execName%-%version%-%GOOS%-%GOARCH%
@echo.

:end