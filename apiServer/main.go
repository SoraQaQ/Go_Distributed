/*
 * @Author: Aquamarine
 * @Date: 2023-11-04 22:04:04
 * @LastEditTime: 2023-11-05 10:25:04
 * @LastEditors: your name
 * @Description:
 * @FilePath: /Distributed/apiServer/main.go
 */
package main

import (
	"distributed/apiServer/heartbeat"
	"distributed/apiServer/locate"
	"distributed/apiServer/objects"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
