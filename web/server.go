package web

import (
	"fmt"
	"net/http"

	"github.com/pdxjohnny/telem/log"
)

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func Start(mux *http.ServeMux, address, port, cert, key string) {
	mux.HandleFunc("/", handler)
  listen := fmt.Sprintf("%s:%s", address, port)
	fmt.Println(address, port, cert, key)
	if cert == "" || key == "" {
		fmt.Printf("About to listen on http://%s/\n", listen)
		err := http.ListenAndServe(listen, mux)
		log.PrintError("ERROR serving frontend", err)
	} else {
		fmt.Printf("About to listen on https://%s/\n", listen)
		err := http.ListenAndServeTLS(listen, cert, key, mux)
		log.PrintError("ERROR serving frontend", err)
	}
}
