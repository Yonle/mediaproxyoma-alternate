package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

var errorBad = fmt.Errorf("bad")
var errorBadSig = fmt.Errorf("bad signature")

func getUrl(rP string) (url string, err error) {
	sp := strings.SplitN(rP, "/", 3)

	if len(sp) < 3 {
		return "", errorBad
	}

	sig, err := base64.RawURLEncoding.DecodeString(sp[0])
	if err != nil {
		return "", errorBad
	}

	decb, err := base64.RawURLEncoding.DecodeString(sp[1])
	if err != nil {
		return "", errorBad
	}

	if !verifySig64(sp[1], sig) {
		return "", errorBadSig
	}

	return checkAndReplace(string(decb))
}
