package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
)

func Hello() string {
	glog.Info("Calling Hello()")
	return "Hello User whoo!!"
}

func main() {
	flag.Parse()
	fmt.Println(Hello())
}
