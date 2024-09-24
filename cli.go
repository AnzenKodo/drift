package main

import (
    "flag"
    "os"
    "fmt"
)

func printHelp() {
    fmt.Print(os.Args[0]+": ", cliInfo.desc+"\n\n")
    fmt.Println("Usage of ", os.Args[0])
    flag.PrintDefaults()
    fmt.Println("\nVersion: " + cliInfo.version)
    fmt.Println("Source: " + cliInfo.repo)
    fmt.Println("License: ["+cliInfo.license+"](https://spdx.org/licenses/" + cliInfo.license + ")")
    os.Exit(0)
}

func cliHandler() (bool, string, bool, string, string) {
    var eport = ""
    var vport = ""
    var vpath = ""

    erun := flag.Bool("erun", false, "Only run " + engineInfo.name)
    flag.StringVar(&eport,"eport", "7447", "Set " + engineInfo.name + " port number")

    vrun := flag.Bool("vrun", false, "Only run " + engineInfo.name)
    flag.StringVar(&vport,"vport", "8080", "Set view port number")
    flag.StringVar(&vpath, "vpath", "./view", "Path of the " + viewInfo.name)

    show_ver := flag.Bool("version", false, "Show version")
	flag.Set("version", "v")
    show_help := flag.Bool("help", false, "Show help")
    flag.Set("help", "h")

    flag.Usage = func() {
        fmt.Println("")
        printHelp()
    }
    flag.Parse()

    if len(os.Args) > 1 {
        if os.Args[1][0] != '-' {
            pLog(LCli, LWarn, "flag provided but not defined: ", os.Args[1])
            printHelp()
        }
    }

    if *show_ver {
        fmt.Println(cliInfo.version)
        os.Exit(0)
    }

    if *show_help {
        printHelp()
    }

    return *erun, eport, *vrun, vport, vpath
}
