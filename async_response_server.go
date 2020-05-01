package async_response_server

import (
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
)

type ServerInfo struct {
	svr       *http.Server
	Feed_back chan string
	Url       string
}

func (svr *ServerInfo) Close() {
	svr.svr.Close()
}

// feature to be considered
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
		svr:       svr,
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

func mpesaHandlerFunc(path string, feedback_chan chan string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case path:
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("error reading:\n%s\n", err)
			}
			// fmt.Printf("===server message===\ntype:%T\nbody:%s\n===\n", body, body)
			feedback_chan <- string(body)
			// Feature to be considered
			// 1. On the `feedback_chan` / creating another channel for that use
			//    which will also return the header setting
			//    so as to make the response more dynamic
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, <-feedback_chan)
		default:
			http.NotFound(w, r)
			return
		}
	}
}
