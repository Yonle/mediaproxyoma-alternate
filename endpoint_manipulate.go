package main

import (
	"net/url"
	"os"
	"strings"
)

var oldMedia_Host = os.Getenv("OLD_MEDIA_HOST")
var oldMedia_PathPrefix = os.Getenv("OLD_MEDIA_PATHPREFIX")

var newMedia_Host = os.Getenv("NEW_MEDIA_HOST")
var newMedia_Scheme = os.Getenv("NEW_MEDIA_SCHEME")
var newMedia_PathPrefix = os.Getenv("NEW_MEDIA_PATHPREFIX")

func checkAndReplace(u string) (string, error) {
	URL, err := url.Parse(u)

	if err != nil {
		return "", err
	}

	isOurHost := URL.Host == oldMedia_Host && strings.HasPrefix(URL.Path, oldMedia_PathPrefix)

	if isOurHost {
		URL.Host = newMedia_Host
		URL.Scheme = newMedia_Scheme
		URL.Path = newMedia_PathPrefix + URL.Path[len(oldMedia_PathPrefix):]
	}

	return URL.String(), nil
}
