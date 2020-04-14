package hook

import (
	"fmt"
	// "log"
	"net/http"
)

// 1. func to for making hook server
//    make sure it's https for safety
func CreateHookServerAsync(port, path string, err chan error) (*http.Server, error) {
	// (string, error) {
	//new_str := port + path
	// http.HandleFunc(path, handler)
	// make server
	// handle := http.HandlerFunc(handler)
	hook_svr := &http.Server{
		Addr:    port,
		Handler: http.HandlerFunc(mpesaHandler),
	}
	// mess := make(chan string)
	// return http.ListenAndServe(port, nil)
	return hook_svr, nil
}
func createServer(svr *http.Server) {

}

// make handler made up of params to make url to be used

// 2. func handler for setting up hook
func mpesaHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/success" {
		mpesaSucc(w)
		// } else if r.URL.Path == "" { //for timeout callback
		//URL
	} else {
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
