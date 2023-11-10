/*
 * @Author: Aquamarine
 * @Date: 2023-11-09 10:46:41
 * @LastEditTime: 2023-11-09 10:49:29
 * @LastEditors: your name
 * @Description:
 * @FilePath: /Distributed/apiServer/objects/del.go
 */
package objects

import (
	"distributed/apiServer/es"
	"log"
	"net/http"
	"strings"
)

func del(w http.ResponseWriter, r *http.Request) {
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	version, e := es.SerachLatesVersion(name)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	e = es.PutMetadata(name, version.Version+1, 0, "")
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
