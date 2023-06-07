#!/bin/bash

sudo apt install 
sudo apt update
sudo apt install golang-go --fix-missing

DIR="/home/ec2-user/url-app"

if [-d "$DIR" ] then 
    echo "$DIR exists"
else
    mkdir ${DIR}
fi 

