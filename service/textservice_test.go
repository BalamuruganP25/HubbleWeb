package service

import (
	"errors"
	"fmt"
	"testing"
)

func TestFindToptenwords(t *testing.T) {
	text_input := "bala murugan bala"
	resp := FindToptenwords(text_input)

	if resp.Topwords[0].Word == "bala" && resp.Topwords[0].Count == 2 {
		fmt.Println("testing success")
	} else {
		fmt.Errorf("%s", errors.New("testing error"))
	}
}
