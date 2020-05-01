package async_response_server

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var succ_mess string = `{
	"ResponseCode": "00000000",
	"ResponseDesc": "success"
	}`

func TestHookWorking(t *testing.T) {
	port := ":8420"
	// success_svr_resp := ""
	ts := CreateHookServerAsync(port)
	t.Logf("The callback URL is '%s'n", ts.Url)
	defer ts.Close()
	go simulateExternalPOSTRequest(t, ts.Url)
	t.Log("waiting on info")
	feedback := <-ts.Feed_back
	ts.Feed_back <- succ_mess
	if feedback != succ_mess {
		t.Errorf("got:%s\n expected:%s\n", feedback, succ_mess)
	}
}
func simulateExternalPOSTRequest(t *testing.T, url string) {
	resp, err := http.Post(url, "application/json", strings.NewReader(succ_mess))
	if err != nil {
		t.Errorf("error posting:\n%s\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("error reading:\n%s\n", err)
	}
	t.Logf("body received from server:\n%s", string(body))
	if succ_mess != string(body) {
		t.Errorf("got:%s\n expected:%s\n", string(body), succ_mess)
	}
}
