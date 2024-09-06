package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"github.com/coder/websocket"
)

const proj_version = "0.1"
const proj_name = "drift/engine"
const proj_desc = "Heart of DriftNet"
const proj_license = "MIT"
const proj_repo = "https://github.com/AnzenKodo/drift/engine"

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    xff := r.Header.Get("X-Forwarded-For")
	ua := r.Header.Get("User-Agent")
	ip := strings.Split(xff, ",")[0]

	if len(ip) < 1 {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	if r.Method == http.MethodGet && r.Header.Get("Upgrade") == "websocket" {
	    c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			InsecureSkipVerify: true,
		})
	    if err != nil {
			log.Println(err)
			return
		}
		defer c.CloseNow()
        log.Printf("%s connected (%s)", ip, ua)
	} else {
		log.Println(ip, r.Method, r.URL, ua)

		if len(r.Header.Get("Upgrade")) > 0 {
			http.Error(w, "Invalid Upgrade Header", 400)
			return
		}

		accept := r.Header.Get("Accept")
		if strings.Contains(accept, "application/nostr+json") {
			// ShowNIP11(w)
		} else {
			showWebsite(w, r)
		}
	}
 //    c, err := websocket.Accept(w, r, nil)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// defer c.CloseNow()
}

func main() {
    var port = ""
    flag.StringVar(&port,"port", "7447", "Set port number")
    show_ver := flag.Bool("version", false, "Show version")
    show_help := flag.Bool("help", false, "Show help")
    flag.Set("help", "h")
	flag.Set("version", "v")
    flag.Parse()

    if *show_ver {
        fmt.Println(proj_version)
        os.Exit(0)
    }

    if *show_help {
        fmt.Print(proj_name+": ", proj_desc+"\n\n")
        flag.Usage()
        fmt.Println("\nLicense: ["+proj_license+"](https://spdx.org/licenses/" + proj_license + ")")
        os.Exit(0)
    }

    http.HandleFunc("/", handlerFunc)
    http.HandleFunc("/favicon.ico", showFavicon)

    err := http.ListenAndServe(":" + port, nil)
    log.Println(err)
}
