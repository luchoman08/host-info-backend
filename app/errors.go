package app

import "errors"
var (
	ErrDomainNotFound    = errors.New("Domain is not found.")
	ErrWhoIsInvalidData = errors.New("Domain whois data invalid.")
	ErrUrlMalformed = errors.New("The given URL is malformed.")
	ErrWebInfoError = errors.New("Can not extract the web page info (icon and title)")
)
