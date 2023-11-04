package heartbeat

import (
	"distributed/apiServer/rabbitmq"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var dataServers = make(map[string]time.Time)
var mutex sync.Mutex

/**
 * @description: 创建消息队列，绑定apiServers exchange监听心跳消息存入dataServers
 * @return {*}
 */
func ListenHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("apiServer")
	c := q.Consume()
	go removeExpiredDataServer()
	for msg := range c {
		dataServer, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		mutex.Lock()
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}
}

/**
 * @description: 在goroutine中运行， 每隔5s扫描一遍dataServers，并清除其中超过10s没收心跳消息的数据服务节点
 * @return {*}
 */
func removeExpiredDataServer() {
	for {
		time.Sleep(5 * time.Second)
		mutex.Lock()
		for s, t := range dataServers {
			if t.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers, s)
			}
		}
		mutex.Unlock()
	}
}

/**
 * @description: 获取当前所有服务节点
 * @return {*}
 */
func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()
	ds := make([]string, 0)
	for s, _ := range dataServers {
		ds = append(ds, s)
	}
	return ds
}

/**
 * @description: 会在当前中随机选择一个节点返回，如果当前数据服务节点为空，则返回空字符串
 * @return {*}
 */
func ChooseRandomDataServer() string {
	ds := GetDataServers()
	n := len(ds)
	if n == 0 {
		return ""
	}
	return ds[rand.Intn(n)]
}
