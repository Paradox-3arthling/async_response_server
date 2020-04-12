package hook

import "testing"

func TestHookWorking(t *testing.T) {
	port := ":8080"
	path := "/mpesa_hook"
	url := "http://localhost" + port + path
	url_got, err := CreateServer(port, path)
	if err != nil {
		t.Errorf("Got an error:%s\n", err)
	}
	if url_got != url {
		t.Errorf("expected:%s, got:%s \n", url, url_got)
	}
	//call post web hook for tesing
}
