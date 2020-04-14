package hook

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHookWorking(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mpesaHandler))
	defer ts.Close()
	//call post web hook for tesing
	test_inp := `{test: "abc"}`
	t.Log("url for safaricom callback: ", ts.URL+"/success")
	resp, err := http.Post(ts.URL+"/success", "application/json", strings.NewReader(test_inp))
	if err != nil {
		t.Errorf("error posting:\n%s\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("error reading:\n%s\n", err)
	}
	// t.Logf("body is:\n%s", string(body))
	succ_mess := `{
		"ResponseCode": "00000000",
		"ResponseDesc": "success"
		}`
	if succ_mess != string(body) {
		t.Logf("got:%s\n expected:%s\n", string(body), succ_mess)
		t.Fail()
	}

}
