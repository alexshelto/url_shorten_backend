#!/bin/bash

yum update
yum install golang

DIR="/home/ec2-user/url-app"

if [-d "$DIR" ] then 
    echo "$DIR exists"
else
    mkdir ${DIR}
fi 

