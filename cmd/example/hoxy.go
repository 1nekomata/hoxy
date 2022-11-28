package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/kyoukaya/hoxy/proxy"
	"github.com/kyoukaya/hoxy/proxy/core/userauth"
	"github.com/kyoukaya/hoxy/utils"

	_ "github.com/kyoukaya/hoxy/mods/constructioninfo"
	_ "github.com/kyoukaya/hoxy/mods/packetlogger"
)

func main() {
	// accept command line argument -c
	var floc string
	flag.StringVar(&floc, "c", "", "Defines optional Blocklist Path")
	flag.Parse()

	list := readFileIntoList(floc)
	fmt.Println(list)

	filters := proxy.Filters{
		HTTPSFilter:     proxy.DefaultHTTPSFilter,
		TelemetryFilter: proxy.DefaultTelemetryFilter,
		LogFilter:       proxy.DefaultLogFilter,
	}

	hoxy := proxy.NewHoxy(proxy.GlobalGameBaseURL, userauth.AuthHandler, filters)
	hoxy.LogGamePackets(true)
	hoxy.Start()
}

// read a file from specified location
func readFileIntoList(fileloc string) []string {
	// open file from string and check for errors
	fr, err := os.Open(fileloc)
	utils.Check(err)

	// create a new buffered reader
	bufr := bufio.NewReader(fr)

	// create the finished []string
	var list []string
	// loop counter
	for {
		// initialize buffers and read a line from a file in a loop
		dat, err := bufr.ReadString('\n')
		if err == io.EOF {
			break
		}
		utils.Check(err)
		list = append(list, string(dat))
	}

	return list
}
