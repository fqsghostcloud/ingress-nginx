package utils

import (
	"net/url"
	"regexp"
	"strings"
)

// BaseEndpoint will return a URL without the /vX.Y
// portion of the URL.
func BaseEndpoint(endpoint string) (string, error) {
	var base string

	u, err := url.Parse(endpoint)
	if err != nil {
		return base, err
	}

	u.RawQuery, u.Fragment = "", ""

	base = u.String()
	versionRe := regexp.MustCompile("v[0-9.]+/?")

	if version := versionRe.FindString(base); version != "" {
		versionIndex := strings.Index(base, version)
		base = base[:versionIndex]
	}

	return base, nil
}
