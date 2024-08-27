package main

import (
    "log"
    "net/url"
    "net/http"
    "net/http/httputil"
    "os"
    "strings"

    "github.com/raymondragon/link/pkg/autotls"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatalf("[ERRO] Usage: https://(username)@hostname:port/targetURL")
    }
    rawURL := os.Args[1]
    parsedURL, err := url.Parse(rawURL)
    if err != nil {
        log.Printf("[WARN] %v", err)
    }
    parsedTargetURL, err := url.Parse(strings.TrimPrefix(parsedURL.Path, "/"))
    if err != nil {
        log.Printf("[WARN] %v", err)
    }
    proxyHandler := httputil.NewSingleHostReverseProxy(parsedTargetURL)
    tlsConfig, err := autotls.Setup(parsedURL.User.Username(), parsedURL.Hostname())
    if err != nil {
        log.Printf("[WARN] %v", err)
    }
    proxyServer := &http.Server{
        Addr:      parsedURL.Host,
        Handler:   proxyHandler,
        TLSConfig: tlsConfig,
    }
    log.Printf("[INFO] %v", rawURL)
    if err := proxyServer.ListenAndServeTLS("", ""); err != nil {
        log.Fatalf("[ERRO] %v", err)
    }
}
