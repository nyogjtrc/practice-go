package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var now = time.Now()

var stopCh = make(chan bool)
var reloadCh = make(chan bool)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	io.WriteString(w, fmt.Sprintln("Hello world!", now.Format(time.RFC3339)))
}

func reload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reload")
	io.WriteString(w, "Reloading.")
	reloadCh <- true
}

func stop(w http.ResponseWriter, r *http.Request) {
	fmt.Println("stop")
	io.WriteString(w, "Stopping.")
	stopCh <- true
}

func serverControl(server *http.Server) {
	for {
		select {
		case <-stopCh:
			fmt.Println("stop chan")
			if err := server.Shutdown(nil); err != nil {
				panic(err)
			}
			return
		case <-reloadCh:
			now = time.Now()
		default:
		}
	}
}

func main() {
	fmt.Println("graceful stop")

	http.HandleFunc("/", hello)
	http.HandleFunc("/stop", stop)
	http.HandleFunc("/reload", reload)

	server := &http.Server{Addr: ":8080"}
	go serverControl(server)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
