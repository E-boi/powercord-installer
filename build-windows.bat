@echo off
set APP=powercord-installer
set APPDIR=%APP%_1.0.0

cd src/
go generate
go build -ldflags "-H windowsgui" -o ../%APPDIR%/%APP%.exe
cd ../
