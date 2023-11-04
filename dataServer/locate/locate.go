/*
 * @Author: your name
 * @Date: 2023-11-04 15:28:29
 * @LastEditTime: 2023-11-04 15:36:49
 * @LastEditors: your name
 * @Description:数据服务locate包
 * @FilePath: /Distributed/dataServer/locate/locate.go
 * 可以输入预定的版权声明、个性签名、空行等
 */
package locate

import (
	"distributed/dataServer/rabbitmq"
	"os"
	"strconv"
)

/**
 * @description: 用于os.Start访问磁盘上对应的文件名
 * @param {string} name
 * @return {*} 定位成功返回true, 否则返回false
 */
func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

/**
 * @description: 启动
 * @param {*}
 * @return {*}
 */
func StartLocate() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		object, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		if Locate(os.Getenv("STORAGE_ROOT") + "/objects/" + object) {
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		}
	}
}
