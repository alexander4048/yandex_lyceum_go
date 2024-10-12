package main

import (
	"errors"
	"testing"
)

func TestGetUTFLength(t *testing.T) {

	var bytes = []byte{33, 28, 64}
	var bytes2 = []byte{255, 43, 51}

	got, gotError := GetUTFLength(bytes)
	expect := 3
	var expectError error = nil
	if got != expect && gotError != expectError {
		t.Fatalf("Ошибка")
	}

	got2, gotError2 := GetUTFLength(bytes2)
	expect2 := 0
	var expectError2 error = errors.New("invalid utf8")
	if got2 != expect2 && gotError2 != expectError2 {
		t.Fatalf("Ошибка")
	}

}
