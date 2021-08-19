package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"net/http"
)

func main() {
	dataUrl := "https://github.com/kunsonx/go-reproduce-data/blob/master/server.data?raw=true"
	rsp, err := http.Get(dataUrl)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewReader(data)
	zr, err := gzip.NewReader(buf)
	if err != nil {
		panic(err)
	}

	defer zr.Close()

	actualData, err := ioutil.ReadAll(zr)
	if err != nil {
		// Then panic here.
		// gzip: invalid header
		panic(err)
	}

	println("actual data size : ", len(actualData))
}
