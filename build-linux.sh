#!/bin/sh

APP=powercord-installer
APPDIR=${APP}-linux

cd ./src

go build -v -o ../$APPDIR/$APP
