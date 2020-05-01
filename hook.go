package hook

import (
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
)

func CreateHookServerSync(port, path string) {
	feedback_c := make(chan string)
	hook_svr := &http.Server{
		Addr:    port,
		Handler: http.HandlerFunc(mpesaHandlerFunc(path, feedback_c)),
	}
	createServer(hook_svr)
}

// 1. func to for making hook server
//    make sure it's https for safety
func CreateHookServerAsync(port, path string) (*http.Server, chan string) {
	feedback_c := make(chan string)
	hook_svr := &http.Server{
		Addr:    port,
		Handler: http.HandlerFunc(mpesaHandlerFunc(path, feedback_c)),
	}

	go createServer(hook_svr)
	return hook_svr, feedback_c
}
func createServer(svr *http.Server) {
	if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server exited on the error:\n", err, "\n")
		svr.Close()
	}
}
func sendBackInfo(feedback_chan chan string, body string) {
	feedback_chan <- body
}

// 2. func handler for setting up hook
func mpesaHandlerFunc(path string, feedback_chan chan string) func(http.ResponseWriter, *http.Request) {
	// succ_mess := `{
	// "ResponseCode": "00000000",
	// "ResponseDesc": "success"
	// }`
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case path:
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("error reading:\n%s\n", err)
			}
			// fmt.Printf("===server message===\ntype:%T\nbody:%s\n===\n", body, body)
			go sendBackInfo(feedback_chan, string(body))
			// defer close(feedback_chan)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(body))
		default:
			http.NotFound(w, r)
			return
		}
	}
}
