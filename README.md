# jsonquery

[![Build Status](https://travis-ci.org/mhemmings/jsonquery.svg?branch=master)](https://travis-ci.org/mhemmings/jsonquery)
[![GoDoc](https://godoc.org/github.com/mhemmings/jsonquery?status.svg)](https://godoc.org/github.com/mhemmings/jsonquery)

Hacked together to get around the AWS API Gateway issue documented [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-known-issues.html) (Duplicated query string parameters are not supported).

All queries for a given prefix are expected to be a JSON array of strings. These are unmarshalled as individual strings, so they are accessible via `url.Query()`.

For example, a `url.URL` of:
	
	https://example.com?json_foo=["a","b","c"]

Becomes available from the `url.URL`:
	
	HandleURL(url, "foo_")
	foos := url.Query()["foo"]
	fmt.Printf("%#v", foos) // []string{"a", "b", "c"}
	
You'll probably want to put this in `http.Handler` middleware.
