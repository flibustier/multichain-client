package multichain

import (
	"fmt"
	"testing"
)

func TestListPermissions(t *testing.T) {

	x, err := client.ListPermissions([]string{"receive", "send"}, []string{}, false)
	if err != nil {
		t.Fail()
	}

	fmt.Println(x)
}
