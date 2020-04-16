package hook

import (
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
)

// 1. func to for making hook server
//    make sure it's https for safety
func CreateHookServerAsync(port, path string) *http.Server {
	hook_svr := &http.Server{
		Addr:    port,
		Handler: http.HandlerFunc(mpesaHandlerFunc(path)),
	}

	go createServer(hook_svr)
	return hook_svr
}
func createServer(svr *http.Server) {
	if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server exited on the error:\n", err, "\n")
		svr.Close()
	}
}

// 2. func handler for setting up hook
func mpesaHandlerFunc(path string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case path:
			mpesaSucc(w, r)
		// case "/timeout" //for timeout callback
		default:
			http.NotFound(w, r)
			return
		}
	}
}
func mpesaSucc(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading:\n%s\n", err)
	}
	log.Printf("type:%T\nbody:%s\n", body, body)
	succ_mess := `{
		"ResponseCode": "00000000",
		"ResponseDesc": "success"
		}`
	fmt.Fprintf(w, succ_mess)
}
