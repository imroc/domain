package domain

import (
	"regexp"
	"strings"

	"github.com/goware/urlx"
	"github.com/pkg/errors"
)

type URL struct {
	Subdomain    string
	Domain       string
	PublicSuffix string
}

var regLocalhost = regexp.MustCompile(`(\A|\.)localhost\z`)
var regIP = regexp.MustCompile(`\A\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\z`)

func Parse(uri string) (url *URL, err error) {
	url = new(URL)
	u, err := urlx.Parse(uri)
	if err != nil {
		return nil, errors.Wrap(err, "invalid url")
	}
	host := u.Host
	if host == "" {
		return
	}
	if i := strings.IndexByte(host, ':'); i > 0 { // remove port
		host = host[:i]
	}
	// parse
	if regLocalhost.MatchString(host) { // loclahost
		url.Domain = "localhost"
		//TODO localhost subdomain
		return
	}
	if regIP.MatchString(host) { // ip
		url.Domain = host
		return
	}
	parts := strings.Split(host, ".")
	getPart := func(i int) string {
		if i < 0 {
			return ""
		}
		return parts[i]
	}
	var suffix []string
	addSuffix := func(s string) {
		suffix = append(suffix, s)
	}
	subHash := publicSuffix
	var ok bool
	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]
		subHash, ok = subHash[part].(hash)
		if !ok {
			url.PublicSuffix = reverseJoin(suffix, 0, len(suffix), ".")
			url.Domain = part
			url.Subdomain = join(parts, 0, i, ".")
			return
		}
		addSuffix(part)
		if _, ok = subHash["*"]; ok {
			addSuffix(getPart(i - 1))
			url.PublicSuffix = reverseJoin(suffix, 0, len(suffix), ".")
			url.Domain = getPart(i - 2)
			url.Subdomain = join(parts, 0, i-2, ".")
			return
		}
	}
	return
}

func join(a []string, i, j int, sep string) (s string) {
	if i > j {
		return
	}
	a = a[i:j]
	if len(a) == 0 {
		return
	}
	return strings.Join(a, ".")
}

func reverseJoin(a []string, i, j int, sep string) (s string) {
	if i > j {
		return
	}
	a = a[i:j]
	if len(a) == 0 {
		return
	}
	for i := len(a) - 1; i >= 0; i-- {
		s += a[i]
		if i != 0 {
			s += sep
		}
	}
	return
}
