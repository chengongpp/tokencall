package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/chengongpp/tokencall/pkg/tokencall"
)

var services = make(map[string]tokencall.ApiService)

func init() {
	//services["aliyun"] = &tokencall.AliyunService{}
}

func main() {
	accessKey := flag.String("k", "", "access key")
	accessSecret := flag.String("s", "", "access secret")
	url := flag.String("u", "", "url")
	serviceType := flag.String("t", "aliyun", "service provider type")
	outputFile := flag.String("o", "result", "output folder/file")
	listServices := flag.Bool("l", false, "list services")
	jsonOutput := flag.Bool("json", false, "output in plain json")
	flag.Parse()
	if *listServices {
		printServices()
		os.Exit(0)
	}
	if *accessKey == "" || *accessSecret == "" {
		flag.Usage()
		os.Exit(1)
	}

	apiSvc, ok := services[strings.ToLower(*serviceType)]
	if !ok {
		_, _ = fmt.Fprintln(os.Stderr, "Service", *serviceType, "is not supported yet")
		os.Exit(1)
	}
	leaked, err := apiSvc.Configure(*accessKey, *accessSecret, *url, nil).Leak()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error leaking:", err.Error())
	}
	if leaked == nil {
		_, _ = fmt.Fprintln(os.Stderr, "No leaks exported. Quit")
		os.Exit(1)
	}
	_, _ = fmt.Fprintln(os.Stdout, leaked)
	if *jsonOutput {
		content, err := json.Marshal(leaked)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Error marshalling leaked data")
			os.Exit(1)
		}
		file, err := os.OpenFile(*outputFile, os.O_APPEND, 0755)
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, "Error closing file:", err.Error())
			}
		}(file)
		_, err = file.Write(content)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Error writing to", file)
			os.Exit(1)
		}
		os.Exit(0)
	}
	if _, err := os.Stat(*outputFile); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(*outputFile, 666)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Error making directory", *outputFile, ":", err.Error())
			os.Exit(1)
		}
	}
}

func printServices() {
	println(`
Supported services:
`)
}
