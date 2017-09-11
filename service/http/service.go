package http

import (
	log "github.com/liuzheng712/golog"
	"net/http"
)

func init() {

}
func Router() {
	log.Debug("service.http.router", "%v", )
	http.Handle("/query", queryHandler)

}
func queryHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("http", "%v", r.Form)

	switch r.Method {
	case "GET":
		log.Info("service.http.queryHandler", "%v", r.Form)

	}
}

