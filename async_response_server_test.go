package async_response_server

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestHookWorking(t *testing.T) {
	// ts := httptest.NewServer(http.HandlerFunc(mpesaHandler))
	// url := ts.URL
	port := ":8420"
	// success_svr_resp := ""
	ts := CreateHookServerAsync(port)
	defer ts.Close()
	//call post web hook for tesing
	succ_mess := `{
		"ResponseCode": "00000000",
		"ResponseDesc": "success"
		}`
	t.Logf("The callback URL is '%s'n", ts.Url)
	resp, err := http.Post(ts.Url, "application/json", strings.NewReader(succ_mess))
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
	t.Log("waiting on info")
	feedback := <-ts.Feed_back
	if feedback != succ_mess {
		t.Errorf("got:%s\n expected:%s\n", feedback, succ_mess)
	}

}
