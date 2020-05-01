package async_response_server

import (
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
)

type ServerInfo struct {
	Svr       *http.Server
	Feed_back chan string
	Url       string
}

func (svr *ServerInfo) Close() {
	svr.Svr.Close()
}

// features to be considered
// 1. On making successful http server we'll create a secure TLS server
func CreateHookServerAsync(port string) *ServerInfo {
	feed_back_chan := make(chan string)
	path := "/response_path"
	svr := &http.Server{
		Addr:    port,
		Handler: http.HandlerFunc(mpesaHandlerFunc(path, feed_back_chan)),
	}
	go createServer(svr)
	return &ServerInfo{
		Svr:       svr,
		Feed_back: feed_back_chan,
		Url:       "http://127.0.0.1" + port + path,
	}

}
func createServer(svr *http.Server) {
	if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server exited on the error:\n", err, "\n")
		svr.Close()
	}
}
func sendBackInfo(feedback_chan chan string, body string) {
	feedback_chan <- body
	close(feedback_chan)
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
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(body))
		default:
			http.NotFound(w, r)
			return
		}
	}
}
