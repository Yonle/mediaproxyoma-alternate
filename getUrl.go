package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

var errorBad = fmt.Errorf("bad")

func getUrl(rP string) (url string, err error) {
	sp := strings.SplitN(rP, "/", 3)

	if len(sp) < 3 {
		return "", errorBad
	}

	decb, err := base64.RawURLEncoding.DecodeString(sp[1])
	if err != nil {
		return "", errorBad
	}

	return checkAndReplace(string(decb))
}
