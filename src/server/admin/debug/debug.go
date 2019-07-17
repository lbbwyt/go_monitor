package debug
import (
	"net/http"
	"net/http/pprof"
	log "github.com/astaxie/beego/logs"

)

func InitDebug(url string) {
	mux := http.NewServeMux()
	handlePprof(mux)

	log.Info("url:" + url)

	go func() {
		err := http.ListenAndServe(url, mux)
		if err != nil {
			log.Error(err)
		}
	}()
}

func handlePprof(mux *http.ServeMux) {
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
}
