package main

import (
    "log"
    "fmt"
)

type LogComp int

const (
    LEng LogComp = iota
    LView LogComp = iota
    LCli LogComp = iota
)

type LogLevel int

const (
	LDebug LogLevel = iota
	LInfo  LogLevel = iota
	LWarn  LogLevel = iota
	LError LogLevel = iota
)

func pLog(comp LogComp, level LogLevel, v ...any) {
    var str string

    switch level {
        case LDebug:
            str = "DEBUG "
        case LInfo:
            str = "INFO  "
        case LWarn:
            str = "WARN  "
        case LError:
            str = "ERROR "
        default:
            log.Panic("Wrong Log Level: ", level)
    }

    switch comp {
        case LEng:
            str += "[Engine] "
        case LView:
            str += "[View]   "
        case LCli:
            str += "[Cli]    "
        default:
            log.Panic("Wrong Componet Level: ", comp)
    }

    log.Println(str + fmt.Sprint(v...))
}
