package util

import (
	"fmt"
	"testing"
)

func CheckForFailure(err error) {
	if err != nil {
		fmt.Println("FAILED: ", err)
	}
}

func CheckForFailureTest(err error, t *testing.T) {
	if err != nil {
		t.Fatal("FAILED: ", err)
	}
}
