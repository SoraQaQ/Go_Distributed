/*
 * @Author: Aquamarine
 * @Date: 2023-11-04 19:10:55
 * @LastEditTime: 2023-11-04 20:30:24
 * @LastEditors: your name
 * @Description:
 * @FilePath: /Distributed/apiServer/locate/locate.go
 */

package locate

import (
	"distributed/apiServer/rabbitmq"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

/**
 * @description: 如果是get调用方法将object_name作为Locate函数的参数进行定位
 * @param {http.ResponseWriter} w
 * @param {*http.Request} r
 * @return {*}
 */
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	info := Locate(strings.Split(r.URL.EscapedPath(), "/")[2])
	if len(info) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, _ := json.Marshal(info)
	w.Write(b)
}

/*
*
  - @description: 接受一个定位的name，创建一个新的消息队列，并向dataServers exchange群发这个对象的定位信息
    并且使用go设置一个超时机制防止无尽的等待
  - @param {string} name
  - @return {*}
*/
func Locate(name string) string {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	q.Publish("dataServers", name)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	return s
}

/**
 * @description: 判断数据是否存在
 * @param {string} name
 * @return {*}
 */
func Exist(name string) bool {
	return Locate(name) != " "
}
