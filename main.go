/*
MIT License

Copyright (c) 2023 Pierre Santana

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

var port = flag.Int("p", 8000, "Port to listen on")

func main() {
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		buf := make([]byte, r.ContentLength)
		log.Printf("New %s request from %s", r.Method, r.RemoteAddr)
		if r.Method == "POST" {
			defer r.Body.Close()
			if _, err := io.ReadFull(r.Body, buf); err != nil {
				log.Printf("Error reading body: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			log.Printf("Message from %s: %s", r.RemoteAddr, string(buf))
		}
		w.Write(buf)
	})
	log.Printf("Starting http echo server on port %d", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", *port), nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
