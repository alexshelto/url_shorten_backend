#!/bin/bash

sudo yum -y update
sudo yum -y install golang

DIR="/home/ec2-user/url-app"

if [-d "$DIR" ] then 
    echo "$DIR exists"
else
    mkdir ${DIR}
fi 

