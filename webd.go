package main

import (
    "log"
    "net/url"
    "net/http"
    "os"

    "github.com/raymondragon/golib"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatalf("[ERRO] Usage: http://hostname:port/path#directory")
    }
    rawURL := os.Args[1]
    parsedURL, err := url.Parse(rawURL)
    if err != nil {
        log.Printf("[WARN] %v", err)
    }
    webdavHandler := golib.WebdavHandler(parsedURL.Fragment, parsedURL.Path)
    log.Printf("[INFO] %v", rawURL)
    if err := http.ListenAndServe(parsedURL.Host, webdavHandler); err != nil {
        log.Fatalf("[ERRO] %v", err)
    }
}