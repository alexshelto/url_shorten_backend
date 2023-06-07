#!/bin/bash

sudo chmod -R 777 /home/ec2-user/url-app
go mod download
go build -o main . 
./main
