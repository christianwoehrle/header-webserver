// Command alertmanager-sns-forwarder provides a Prometheus Alertmanager Webhook Receiver for forwarding alerts to AWS SNS.
package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	log = logrus.New()

	headerSizeString = kingpin.Flag("size", "Size of Header").Default(":4098").Envar("HEADER_SIZE").String()
	headerSize       = 0
)

func main() {
	kingpin.Parse()

	http.HandleFunc("/", HelloHeader)
	http.ListenAndServe(":8082", nil)
}

func HelloHeader(w http.ResponseWriter, r *http.Request) {
	size, ok := r.URL.Query()["size"]
	s := 5000
	if ok {
		s, _ = strconv.Atoi(string(size[0]))
	}

	fmt.Printf("Write %d bytes\n", s)

	status, ok := r.URL.Query()["status"]
	sc := 200
	if ok {
		sc, _ = strconv.Atoi(string(status[0]))
	}
	fmt.Printf("Write Status Code %d\n", sc)

	header := ""
	for i := 1; i < s; i++ {
		header = header + "x"
	}

	w.Header().Set("h", header)
	w.WriteHeader(sc)
}
