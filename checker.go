package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
)

func checkStatus(client *http.Client, userAgent, domain string) (status, error) {
	parsedURL, err := url.Parse(domain)
	if err != nil {
		return status{}, err
	}
	if parsedURL.Scheme == "" {
		parsedURL.Scheme = "http"
	}
	req, err := http.NewRequest(http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return status{}, err
	}
	req.Header.Add("user-agent", userAgent)
	rsp, err := client.Do(req)
	if err != nil {
		if err, ok := err.(*url.Error); ok {
			actualErr := err.Unwrap()
			if _, ok := actualErr.(*net.OpError); ok {
				return status{domain: parsedURL.String(), status: doesNotExistStatus}, nil
			}
		}
		fmt.Printf("%#v\n", err)
		return status{}, err
	}
	return status{
		domain: parsedURL.String(),
		status: rsp.StatusCode,
	}, nil
}
