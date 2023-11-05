/*
 * @Author: Aquamarine
 * @Date: 2023-11-04 23:47:16
 * @LastEditTime: 2023-11-04 23:48:43
 * @LastEditors: your name
 * @Description:
 * @FilePath: /Distributed/apiServer/objects/objects.go
 */
package objects

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		put(w, r)
		return
	}
	if m == http.MethodGet {
		get(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
