package main

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
)

var versionCode = "v0.2"
var proxyBufferSize = 4096
var httpListenPort = 80
var httpsConnectingPort = 443
var allowInsecure = false

func handler(responseToRequest http.ResponseWriter, incomingRequest *http.Request) {

	host := incomingRequest.Host
	url := incomingRequest.URL
	log.Printf("Received request to route to host %s and url %s", host, url)

	// Get the raw request bytes
	requestDump, err := httputil.DumpRequest(incomingRequest, true)
	if err != nil {
		log.Printf("cannot dump %s", err)
		http.Error(responseToRequest, "Cannot dump request", http.StatusBadRequest)
	}

	//ioutil.WriteFile("input.txt", requestDump, 0644)

	// You can uncomment to view the raw http request for debugging
	//log.Printf("Dump:\n%s", string(requestDump))

	conf := &tls.Config{}

	if allowInsecure {
		conf = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	conn, err := tls.Dial("tcp", host+":"+strconv.Itoa(httpsConnectingPort), conf)
	if err != nil {
		log.Printf("Cannot dial host %s", err)
		http.Error(responseToRequest, "Cannot dial host", http.StatusGatewayTimeout)
		return
	}
	defer conn.Close()

	n, err := conn.Write(requestDump)
	if err != nil {
		log.Printf("Cannot write request %d %s\n", n, err)
		http.Error(responseToRequest, "Cannot write request"+err.Error(), http.StatusBadGateway)
		return
	}

	// Prepare the requesting socket for writing. Access raw socket by hijacking
	// Reference: https://stackoverflow.com/questions/29531993/accessing-the-underlying-socket-of-a-net-http-response

	hj, ok := responseToRequest.(http.Hijacker)
	if !ok {
		http.Error(responseToRequest, "webserver doesn't support hijacking", http.StatusInternalServerError)
		return
	}
	returnConn, _, err := hj.Hijack()

	if err != nil {
		http.Error(responseToRequest, err.Error(), http.StatusInternalServerError)
		return
	}

	defer returnConn.Close()

	readBuf := make([]byte, proxyBufferSize)

	for {
		//Read from response socket from external server and pass data back
		bytesRead, err := conn.Read(readBuf)

		if err != nil {

			if err == io.EOF {
				log.Printf("EOF reached")
			} else {
				log.Printf("Error getting bytes from server %d %s", bytesRead, err)
			}

			break
		}

		bytesWritten, err := returnConn.Write(readBuf[:bytesRead])

		//ioutil.WriteFile("output.txt", readBuf[:bytesRead], 0644)

		if err != nil {
			log.Printf("Error writing bytes to requester %d %s", bytesWritten, err)
			break
		}

	}

	log.Println("End of handler")

}

func main() {

	argsWithoutProg := os.Args[1:]

	for _, arg := range os.Args[1:] {
		if arg == "-i"{
			allowInsecure = true
		}
	}

	if len(argsWithoutProg) >= 3 {
		parsedHTTPPort, err := strconv.ParseInt(argsWithoutProg[0], 10, 32)

		if err != nil {
			log.Printf("Cannot parse argument %s", argsWithoutProg[0])
			return
		}

		parsedHTTPSPort, err := strconv.ParseInt(argsWithoutProg[1], 10, 32)

		if err != nil {
			log.Printf("Cannot parse argument %s", argsWithoutProg[1])
			return
		}

		parsedProxyBuffer, err := strconv.ParseInt(argsWithoutProg[2], 10, 32)

		if err != nil {
			log.Printf("Cannot parse argument %s", argsWithoutProg[2])
			return
		}

		httpListenPort = int(parsedHTTPPort)
		httpsConnectingPort = int(parsedHTTPSPort)
		proxyBufferSize = int(parsedProxyBuffer)
	}

	log.Printf("HTTP to HTTPS proxy %s listening to %d, forward to %d with listening buffer %d", versionCode, httpListenPort, httpsConnectingPort, proxyBufferSize)
	
	if allowInsecure {
		log.Printf("Allow insecure TLS certificates")
	}
	
	log.Printf("You can supply the listening port, forward port, buffer size, insecure -i cert as command line args")

	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":"+strconv.Itoa(httpListenPort), nil); err != nil {
		log.Fatal(err)
	}
}
