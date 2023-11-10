/*
 * @Author: Aquamarine
 * @Date: 2023-11-04 14:31:02
 * @LastEditTime: 2023-11-04 19:15:25
 * @LastEditors: your name
 * @Description:
 * @FilePath: /Distributed/dataServer/main.go
 */

package main

import (
	"distributed/dataServer/heartbeat"
	"distributed/dataServer/locate"
	"distributed/demo/objects"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
