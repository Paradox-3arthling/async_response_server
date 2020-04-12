package hook_test

import (
	"testing"

	"github.com/paradox-3arthling/mpesa_api/hook"
)

func TestHookWorking(t *testing.T) {
	port := ":8080"
	url, err := hook.CreateServer(port, "/mpesa_hook")
	if err != nil {
		t.Errorf("err->%d", err)
	}
	//call post web hook for tesing
}
