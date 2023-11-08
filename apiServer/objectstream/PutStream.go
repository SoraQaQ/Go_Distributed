/*
 * @Author: Aquamarine
 * @Date: 2023-11-04 19:24:14
 * @LastEditTime: 2023-11-08 10:58:55
 * @LastEditors: your name
 * @Description:
 * @FilePath: /Distributed/apiServer/objectstream/PutStream.go
 */
package objectstream

import (
	"fmt"
	"io"
	"net/http"
)

type PutStream struct {
	writer *io.PipeWriter
	c      chan error
}

/**
 * @description: 新建
 * @param {*} server
 * @param {string} object
 * @return {*}
 */
func NewPutStream(server, object string) *PutStream {
	reader, writer := io.Pipe()
	c := make(chan error)
	go func() {
		request, _ := http.NewRequest("PUT", "http://"+server+"/objects/"+object, reader)
		client := http.Client{}
		r, e := client.Do(request)
		if e == nil && r.StatusCode != http.StatusOK {
			e = fmt.Errorf("dataServer return http code %d", r.StatusCode)
		}
		c <- e
	}()
	return &PutStream{writer, c}
}

/**
 * @description: 实现io.Writer接口
 * @param {[]byte} p
 * @return {*}
 */
func (w *PutStream) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

/**
 * @description: 关闭
 * @return {*}
 */
func (w *PutStream) Close() error {
	w.writer.Close()
	return <-w.c
}
