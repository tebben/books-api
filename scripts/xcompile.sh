#!/bin/sh

if [ -d "..\bin" ];
then
   mkdir ..\bin\linux64
   mkdir ..\bin\darwin64
   echo "Build directories created"
else
   echo "Build directories already exist"
fi

SET GOOS=linux
SET GOARCH=amd64
go build -o ..\bin\linux64\gost ../src/main.go
echo "Built application for Linux/amd64"

SET GOOS=darwin
SET GOARCH=amd64
go build -o ..\bin\darwin64\gost ../src/main.go
echo "Built application for Darwin/amd64"
