#!/bin/bash

sudo chmod -R 777 /home/ec2-user/url-app

cd /home/ec2-user/url-app
ls

go mod download
go build -o main . 
./main
