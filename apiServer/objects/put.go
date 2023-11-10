/*
 * @Author: Aquamarine
 * @Date: 2023-11-04 19:20:11
 * @LastEditTime: 2023-11-09 11:00:44
 * @LastEditors: your name
 * @Description:
 * @FilePath: /Distributed/apiServer/objects/put.go
 */
package objects

import (
	"distributed/apiServer/es"
	"distributed/apiServer/heartbeat"
	"distributed/apiServer/objectstream"
	"distributed/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

/**
 * @description: put函数
 * @param {http.ResponseWriter} w
 * @param {*http.Request} r
 * @return {*}
 */
func put(w http.ResponseWriter, r *http.Request) {
	hash := utils.GetHashFromnHeader(r.Header)
	if hash == "" {
		log.Println("missing object hash in digest header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c, e := storeObject(r.Body, url.PathEscape(hash))
	if e != nil {
		log.Println(e)
		w.WriteHeader(c)
		return
	}
	if c != http.StatusOK {
		w.WriteHeader(c)
		return
	}

	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	size := utils.GetSizeFromHeader(r.Header)
	e = es.AddVersion(name, hash, size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}
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
