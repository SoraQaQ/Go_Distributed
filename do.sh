#!/bin/sh
###
 # @Author: Aquamarine
 # @Date: 2023-11-07 22:05:23
 # @LastEditTime: 2023-11-08 10:42:09
 # @LastEditors: your name
 # @Description: 
 # @FilePath: /Distributed/do.sh
### 
 for i in `seq 1 6`; do mkdir -p /tmp/$i/objects; done

 sudo chmod -R 777 /tmp/1 /tmp/2 /tmp/3 /tmp/4 /tmp/5 /tmp/6

ifconfig eth0:1 10.29.1.1 netmask 255.255.255.0 up 
ifconfig eth0:2 10.29.1.2 netmask 255.255.255.0 up 
ifconfig eth0:3 10.29.1.3 netmask 255.255.255.0 up 
ifconfig eth0:4 10.29.1.4 netmask 255.255.255.0 up 
ifconfig eth0:5 10.29.1.5 netmask 255.255.255.0 up 
ifconfig eth0:6 10.29.1.6 netmask 255.255.255.0 up 
ifconfig eth0:7 10.29.2.1 netmask 255.255.255.0 up 
ifconfig eth0:8 10.29.2.2 netmask 255.255.255.0 up 