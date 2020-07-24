package main

import (
	"fmt"
	"testing"
)

func TestDecrypt(t *testing.T) {
	s := `2pVbjH2CBpXCLJVpqWkUvvYRRDNUZVn1MxDlX1ONMWcRmECdmnL+NXwlSuCDQCNxrxvafr5zZFBQtUT5iRf8H/3MrUMAAxeBZC2bBz7QJtBd2LI0hnPJYK3sOY9UnXlaGM1MzGQcKhc4uilX1D9XAw==`
	res, err := decrypt(s)
	if err != nil {
		t.Errorf("decrypt failed. %v", err)
		return
	}

	fmt.Println(res)
}
