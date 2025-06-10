package mobile

import (
	"log"
	"sync"
)

var callback = new(logCallBack)

type logCallBack struct {
}

func (l *logCallBack) OnLog(level int16, time int64, message string) {
	log.Println(level, time, message)
}

var eventEntity = new(EventEntity)

type EventEntity struct {
}

func (e EventEntity) OnStartError(t string, err string) {
	//TODO implement me
	log.Println(t, err)
}

func (e EventEntity) OnShutdown(t string) {
	log.Println(t)
}

func (e EventEntity) OnProcessExit(code int) {
	log.Println(code)
}

func Run() {
	log.Println("Start service")
	//SetConfigData("/data/user/0/cloud.iothub.apps.alistweb.alistweb/app_flutter")
	//err := Init(eventEntity, callback)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//Start()
	RunAPIServer()
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
