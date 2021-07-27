#!/bin/sh

APP=powercord-installer
APPDIR=${APP}_1.0.0

cd ./src

go build -v -o ../$APPDIR/$APP
