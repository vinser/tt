package fb

import (
	"net/http"

	"github.com/golang/glog"
)

type FB struct{}

func (f *FB) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l := r.FormValue("lang")
	switch l {
	case "ru":
		glog.Info("lang=ru")
	case "en":
		glog.Warning("lang=en")
	default:
		glog.Errorf("lang=%s", l)
	}
}
