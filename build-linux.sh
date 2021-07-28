#!/bin/sh

APP=powercord-installer
APPDIR=${APP}_1.1.0

cd ./src

go build -v -o ../$APPDIR/$APP
