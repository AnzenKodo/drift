package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
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
	    wsHandler(w, r, ip, ua)
	} else {
        log.Println("[Engine Site]", ip, r.Method, r.URL, ua)

        if len(r.Header.Get("Upgrade")) > 0 {
           	http.Error(w, "Invalid Upgrade Header", 400)
           	return
        }

        accept := r.Header.Get("Accept")

        if strings.Contains(accept, "application/nostr+json") || strings.Contains(accept, "application/json") {
            w.Header().Set("Content-Type", "application/nostr+json")
            w.Header().Set("Access-Control-Allow-Origin", "*")

            d, err := json.Marshal(config)
            if err != nil {
                fmt.Fprint(w, "{}")
                return
            }
            w.Write(d)
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
            fmt.Print("Serving " + engineInfo.name + " at http://localhost:" + eport, "\n\n")
            err := http.ListenAndServe(":" + eport, engineMux)
            log.Println("[Engine]", err)
        }
    }()

    if !erun {
        fmt.Println("Serving " + viewInfo.name +" from `" + vpath + "` at http://localhost:" + vport)
        err := http.ListenAndServe(":" + vport, viewMux)
        log.Println("[View]", err)
    }
}
