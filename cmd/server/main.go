package main

import (
	"flag"
	"log/slog"
	"net"
	"net/http"
	"os"
)

func main() {
	host := flag.String("host", "0.0.0.0", "Host interface to bind to")
	port := flag.String("port", "8080", "Port to listen on")

	flag.Parse()

	addr := net.JoinHostPort(*host, *port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		if _, err := w.Write([]byte("Hello World!")); err != nil {
			slog.Error("error writing HTTP response body", "err", err, "remote_addr", r.RemoteAddr)
		}
	})

	slog.Info("HTTP server listening", "addr", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		slog.Error("HTTP server exited with error", "err", err, "addr", addr)
		os.Exit(1)
	}
}
