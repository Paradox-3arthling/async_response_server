package hook

import (
	"fmt"
	"log"

	// "log"
	"net/http"
)

// 1. func to for making hook server
//    make sure it's https for safety
func CreateHookServerAsync(port, path string) *http.Server {
	hook_svr := &http.Server{
		Addr:    port,
		Handler: http.HandlerFunc(mpesaHandler),
	}

	go createServer(hook_svr)
	return hook_svr
}
func createServer(svr *http.Server) {
	if err := svr.ListenAndServe(); err != nil {
		log.Fatal("Server exited on the error:\n", err, "\n")
		svr.Close()
	}
}

// make handler made up of params to make url to be used

// 2. func handler for setting up hook
func mpesaHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/success":
		mpesaSucc(w)
	// case "/timeout" //for timeout callback
	default:
		http.NotFound(w, r)
		return
	}
}
func mpesaSucc(w http.ResponseWriter) {
	succ_mess := `{
		"ResponseCode": "00000000",
		"ResponseDesc": "success"
		}`
	fmt.Fprintf(w, succ_mess)
}
