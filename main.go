package main

import (
	"log"
	"net"
	"net/http"
	"strings"
	"github.com/coder/websocket"
	"fmt"
)

type Proj_Info struct {
    name string
    desc string
    version string
    license string
    repo string
}

var defaultInfo = Proj_Info{
    name: "drift",
    version: "0.1",
    license: "MIT",
    repo: "https://github.com/AnzenKodo/drift",
}

var engineInfo = Proj_Info{
    name: defaultInfo.name + "/engine",
    desc: "Heart of " + defaultInfo.name,
    repo: defaultInfo.repo,
    version: defaultInfo.version,
    license: defaultInfo.license,
}

var viewInfo = Proj_Info{
    name: defaultInfo.name + "/view",
}

var cliInfo = Proj_Info{
    desc: "cli for managing drift/engine & drift/view",
    repo: defaultInfo.repo,
    version: defaultInfo.version,
    license: defaultInfo.license,
}

func engineHandler(w http.ResponseWriter, r *http.Request) {
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
        log.Println("[Engine WS]", ip, " connected ", ua)
	} else {
		log.Println("[Engine Site]", ip, r.Method, r.URL, ua)

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
}

func main() {
    erun, eport, vrun, vport, vpath := cliHandler()

    engineMux := http.NewServeMux()
    engineMux.HandleFunc("/", engineHandler)
    engineMux.HandleFunc("/favicon.ico", showFavicon)

    viewMux := http.NewServeMux()
    viewMux.Handle("/", http.FileServer(http.Dir(vpath)))

    go func() {
        if !vrun {
            fmt.Println("Serving " + engineInfo.name + " at https://localhost:" + eport)
            err := http.ListenAndServe(":" + eport, engineMux)
            log.Println("[Engine]", err)
        }
    }()

    if !erun {
        fmt.Println("Serving " + viewInfo.name +" from `" + vpath + "` at https://localhost:" + vport)
        err := http.ListenAndServe(":" + vport, viewMux)
        log.Println("[View]", err)
    }
}
