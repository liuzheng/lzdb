package http

import (
	"net/http"
	//json "github.com/pquerna/ffjson/ffjson"
	"github.com/liuzheng712/lzdb/service/status"
)

func getstatus(w http.ResponseWriter, r *http.Request) {
	httpHandler(w, r)

	status.Mem()
	status.CPU()
	status.Load()
	status.Disk()
	status.Net()

	switch r.Method {
	case "GET":
		w.WriteHeader(HTTPCodeForbidden)
	default:
		w.WriteHeader(HTTPCodeForbidden)
	}
}