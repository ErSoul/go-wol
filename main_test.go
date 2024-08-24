package main

import (
	"testing"
	"regexp"
)

func TestParseMACAddress(t *testing.T){
	address := "fffffffffff"

	unParsedAddress := ParseMACAddress(address)

	want := regexp.MustCompilePOSIX("^[a-fA-F0-9]{12}$")

	if ! want.MatchString(unParsedAddress) {
		t.Errorf("orale")
	}
}
