package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := http.Get("https://distopia.savi2w.workers.dev/")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for k, v := range resp.Header {
		io.WriteString(os.Stdout, fmt.Sprintf("%s: %s\n", k, v))
	}

	fmt.Println("Response status: ", resp.Status)
}
