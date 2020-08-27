package internal

import "errors"

var NotFound = errors.New("not found")
var InvalidPassword = errors.New("invalid password")
var NoCredentialsProvided = errors.New("no credentials provided")
var JSONSerialization = errors.New("json serialization")

var UrlParams = errors.New("url params not as expected")