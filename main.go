package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	version = "dev"
)

type Config struct {
	Files []struct {
		Filename string `yaml:"filename"`
		URL      string `yaml:"url"`
	} `yaml:"files"`
}

func DownloadFile(filepath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	fmt.Printf("Version: %s\n", version)
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", "files", "the directory of static file to host")
	config := flag.String("c", "config.yaml", "Configuration for files to load")
	flag.Parse()

	var conf Config

	b, err := ioutil.ReadFile(*config)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(b, &conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	for _, fileRequest := range conf.Files {
		log.Printf("Fetching %s from %s\n", fileRequest.Filename, fileRequest.URL)
		if err := DownloadFile(fmt.Sprintf("%s/%s", *directory, fileRequest.Filename), fileRequest.URL); err != nil {
			log.Fatal(err)
		}
	}

	files, err := ioutil.ReadDir(*directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Printf("file found -> %s\n", f.Name())
	}

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
