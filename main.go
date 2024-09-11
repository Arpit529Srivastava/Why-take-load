package main

import (
	"fmt"
	"net/http/httputil"
)

type simpleserver struct{
	addr string
	proxy httputil.ReverseProxy
}

func newsimplesever(add string)