package httpclient

import dghttp "dghire.com/libs/go-httpclient"

var httpClient = initHttpClient()

func initHttpClient() *dghttp.DgHttpClient {
	hc := dghttp.GlobalHttpClient
	return hc
}

func Ins() *dghttp.DgHttpClient {
	return httpClient
}
