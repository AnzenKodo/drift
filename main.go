package main

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Proj_Info struct {
    name string
    desc string
    version string
    license string
    repo string
    theme_color rl.Color
    bg_color rl.Color
    fg_color rl.Color
}

var defaultInfo = Proj_Info{
    name: "Drift",
    version: "0.1",
    license: "MIT",
    repo: "https://github.com/AnzenKodo/drift",
}

var engineInfo = Proj_Info{
    name: defaultInfo.name + "Engine",
    desc: "Heart of " + defaultInfo.name,
    repo: defaultInfo.repo,
    version: defaultInfo.version,
    license: defaultInfo.license,
}

var viewInfo = Proj_Info{
    name: defaultInfo.name + "View",
    // Theme: Base2Tone Lavender
    theme_color: rl.NewColor(147, 117, 245, 225),
    bg_color: rl.NewColor(32, 29, 42, 225),
    fg_color: rl.NewColor(239, 235, 255, 225),
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
        pLog(LEng, LInfo, ip, r.Method, r.URL, ua)

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
    // view_start()
    erun, eport, vrun, vport, vpath := cliHandler()
    _ = eport
    _ = vrun
    // engineMux := http.NewServeMux()
    // engineMux.HandleFunc("/", engineHandler)
    // engineMux.HandleFunc("/favicon.ico", showFavicon)

    viewMux := http.NewServeMux()
    viewMux.Handle("/", http.FileServer(http.Dir(vpath)))

    // go func() {
    //     if !vrun {
    //         pLog(LEng, LInfo, "Serving " + engineInfo.name + " at http://localhost:" + eport)
    //         err := http.ListenAndServe(":" + eport, engineMux)
    //         pLog(LEng, LWarn, "", err)
    //     }
    // }()

    if !erun {
        pLog(LView, LInfo, "Serving " + viewInfo.name +" from `" + vpath + "` at http://localhost:" + vport)
        err := http.ListenAndServe(":" + vport, viewMux)
        pLog(LView, LWarn, "", err)
    }
}
