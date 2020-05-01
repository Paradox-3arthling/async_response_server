package hook

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
	path := "/mpesa_hook"
	ts, feedback_c := CreateHookServerAsync(port, path)
	url := "http://127.0.0.1" + port + path
	defer ts.Close()
	//call post web hook for tesing
	succ_mess := `{
		"ResponseCode": "00000000",
		"ResponseDesc": "success"
		}`
	t.Log("url for safaricom callback: ", url)
	resp, err := http.Post(url, "application/json", strings.NewReader(succ_mess))
	if err != nil {
		t.Errorf("error posting:\n%s\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("error reading:\n%s\n", err)
	}
	// t.Logf("body is:\n%s", string(body))
	if succ_mess != string(body) {
		t.Logf("got:%s\n expected:%s\n", string(body), succ_mess)
		t.Fail()
	}
	t.Log("waiting on info")
	feedback := <-feedback_c
	if feedback != succ_mess {
		t.Logf("got:%s\n expected:%s\n", feedback, succ_mess)
		t.Fail()
	}

}
