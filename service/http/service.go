package http

import (
	log "github.com/liuzheng712/golog"
	"net/http"
	"fmt"
	"github.com/liuzheng712/lzdb/service/conf"
)

func init() {
}
func router() {
	log.Debug("service.http.router", "%v", "load the Routers")
	http.HandleFunc("/query", queryHandler)
}
func queryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Info("http", "%v\t%v\t%v", r.RemoteAddr,r.Method,r.RequestURI)
		log.Debug("http", "%v ", r.Header)
		fmt.Fprint(w, "xxx")
	}
}

func Run() {
	router()
	log.Info("http", "Start http server at %v:%v", conf.Config.Http.Host, conf.Config.Http.Port)
	log.Fatal("http", "%v", http.ListenAndServe(conf.Config.Http.Host + ":" + fmt.Sprint(conf.Config.Http.Port), nil))
}
