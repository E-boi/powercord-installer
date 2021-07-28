@echo off
set APP=powercord-installer
set APPDIR=%APP%-windows

cd src/
go generate
go build -ldflags "-H windowsgui" -o ../%APPDIR%/%APP%.exe
cd ../
