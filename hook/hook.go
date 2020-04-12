package hook

import (
	"fmt"
	"log"
	"net/http"
)

// 1. func to for making hook server
//    make sure it's https for safety
func CreateServer(port, path string) {
	// (string, error) {
	//new_str := port + path
	// make server
	http.HandleFunc(path, handler)
	// mess := make(chan string)
	log.Fatal(http.ListenAndServe(port, nil))
	// return new_str, nil
}

// make handler made up of params to make url to be used

// 2. func handler for setting up hook
func handler(w http.ResponseWriter, r *http.Request) {
	succ_mess := `{
	"ResponseCode": "00000000",
	"ResponseDesc": "success"
	}`
	fmt.Fprintf(w, succ_mess)
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
