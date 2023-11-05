/*
 * @Author: Aquamarine
 * @Date: 2023-11-04 19:20:11
 * @LastEditTime: 2023-11-04 19:49:16
 * @LastEditors: your name
 * @Description:
 * @FilePath: /Distributed/apiServer/objects/put.go
 */
package objects

import (
	"distributed/apiServer/heartbeat"
	"distributed/apiServer/objectstream"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

/**
 * @description: put函数
 * @param {http.ResponseWriter} w
 * @param {*http.Request} r
 * @return {*}
 */
func put(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	c, e := storeObject(r.Body, object)
	if e != nil {
		log.Println(e)
	}
	w.WriteHeader(c)
}

/**
 * @description: 调用putStream生成stream
 * @param {io.Reader} r
 * @param {string} object
 * @return {*}
 */
func storeObject(r io.Reader, object string) (int, error) {
	stream, e := putStream(object)
	if e != nil {
		return http.StatusServiceUnavailable, e
	}
	io.Copy(stream, r)
	e = stream.Close()
	if e != nil {
		return http.StatusInternalServerError, e
	}
	return http.StatusOK, nil
}

/**
 * @description: 获取一个随机数据服务节点的地址server
 * @param {string} object
 * @return {*}
 */
func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}
	return objectstream.NewPutStream(server, object), nil
}
