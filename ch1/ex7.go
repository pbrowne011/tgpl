// Exercise 1.7
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// Use io.Copy() for unbufferd io
		_, err = io.Copy(os.Stdout, resp.Body)

		// Note: this line is essential. When working as a client, the package
		// manages "Transport" type connections, reusing connections when
		// efficient. This consumes file descriptors and network sockets in the
		// OS. It is essential to close these resources after they are used to
		// prevent OS errors from having too many conenctions open at once.
		//
		// Response.Body.Close() closes this network connection, freeing up
		// network resources. I'm not sure how it does this (how does it keep
		// track of network connections? why not resp.Close()?), but will dig
		// into later.
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
