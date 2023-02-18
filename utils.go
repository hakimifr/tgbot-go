package main

import "runtime"

func check(err error, message string) {
    if err != nil {
        panic(message + ": " + err.Error())
    }
}

func getFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

