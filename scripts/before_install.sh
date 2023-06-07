#!/bin/bash

sudo yum update
sudo yum install golang

DIR="/home/ec2-user/url-app"

if [-d "$DIR" ] then 
    echo "$DIR exists"
else
    mkdir ${DIR}
fi 

