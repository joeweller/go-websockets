#!/bin/sh

if [ ! -d "bin" ] 
then
    mkdir bin
else
    rm -R bin/*
fi

go build -o=bin main.go
cp -r assets bin/