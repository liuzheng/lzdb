package http

import (
	log "github.com/liuzheng712/golog"
	"github.com/liuzheng712/lzdb/service/conf"
	"net/http"
	"fmt"
)

func init() {
}

func router() {
	log.Debug("service.http.router", "%v", "load the Routers")
	http.HandleFunc("/test", test)
	http.HandleFunc("/query", queryHandler)
	http.HandleFunc("/status", getstatus)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("http", "%v\t%v\t%v", r.RemoteAddr, r.Method, r.RequestURI)
	w.Header().Set("Server", "LZDB Server")
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	httpHandler(w, r)
	log.Debug("http", "Header: %v ", r.Header)
	log.Debug("http", "Cookies: %v ", r.Cookies())
	log.Debug("http", "Body: %v ", r.Body)

	switch r.Method {
	case "GET":
		w.Write([]byte("xxx"))
	case "HEAD":
	case "POST":
		log.Debug("http", "Form: %v ", r.Form)
		log.Debug("http", "PostForm: %v ", r.PostForm)
	case "OPTIONS":
	case "PUT":
	case "DELETE":
	case "TRACE":
	case "CONNECT":
	default:
		w.WriteHeader(HTTPCodeForbidden)
	}
}
func test(w http.ResponseWriter, r *http.Request) {
	httpHandler(w, r)

	switch r.Method {
	case "GET":
		w.WriteHeader(HTTPCodeForbidden)
		w.Write([]byte(r.Method))
	case "HEAD":
		w.WriteHeader(HTTPCodeForbidden)
		w.Write([]byte(r.Method))
	case "POST":
		log.Debug("http", "Form: %v ", r.Form)
		log.Debug("http", "PostForm: %v ", r.PostForm)
	case "OPTIONS":
		w.Write([]byte(r.Method))
	case "PUT":
		w.Write([]byte(r.Method))
	case "DELETE":
		w.Write([]byte(r.Method))
	case "TRACE":
		w.Write([]byte(r.Method))
	case "CONNECT":
		w.Write([]byte(r.Method))
	default:
		w.WriteHeader(HTTPCodeForbidden)
	}
}

func Run() {
	router()
	log.Info("http", "Start http server at %v:%v", conf.Config.Http.Host, conf.Config.Http.Port)
	log.Fatal("http", "%v", http.ListenAndServe(conf.Config.Http.Host + ":" + fmt.Sprint(conf.Config.Http.Port), nil))
}
