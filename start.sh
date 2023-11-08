#!/bin/sh
###
 # @Author: Aquamarine
 # @Date: 2023-11-07 21:34:10
 # @LastEditTime: 2023-11-07 22:04:32
 # @LastEditors: your name
 # @Description: 
 # @FilePath: /Distributed/start.sh
### 
export RABBITMQ_SERVER=amqp://localhost:5672
LISTEN_ADDRESS=10.29.1.1:12345 STORAGE_ROOT=/tmp/1 go run dataServer/main.go &
LISTEN_ADDRESS=10.29.1.2:12345 STORAGE_ROOT=/tmp/2 go run dataServer/main.go &
LISTEN_ADDRESS=10.29.1.3:12345 STORAGE_ROOT=/tmp/3 go run dataServer/main.go &
LISTEN_ADDRESS=10.29.1.4:12345 STORAGE_ROOT=/tmp/4 go run dataServer/main.go &
LISTEN_ADDRESS=10.29.1.5:12345 STORAGE_ROOT=/tmp/5 go run dataServer/main.go &
LISTEN_ADDRESS=10.29.1.6:12345 STORAGE_ROOT=/tmp/6 go run dataServer/main.go &
LISTEN_ADDRESS=10.29.2.1:12345 go run apiServer/main.go &
LISTEN_ADDRESS=10.29.2.2:12345 go run apiServer/main.go &
echo "Start"
