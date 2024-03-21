package main

import (
    "flag"
    "log"
    "os"
    "path/filepath"
    "time"
)

var (
    fileExt = flag.String("e", "", "File Extension")
    mainDir = flag.String("m", "", "Main Directory")
    tempDir = flag.String("t", "", "Temp Directory")
)

func main() {
    flag.Parse()
    if *mainDir == "" || *tempDir == "" || *fileExt == "" {
        flag.Usage()
        log.Fatalf("[ERRO] %v", "Invalid Flag(s)")
    }
    for {
        moveTempFiles(*mainDir, *tempDir, *fileExt)
        time.Sleep(1 * time.Hour)
    }
}

func moveTempFiles(directory, tempDirectory, fileExt string) {
    if err := os.MkdirAll(tempDirectory, os.ModePerm); err != nil {
        log.Fatalf("[ERRO] %v", err)
    }
    err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            log.Printf("[WARN] %v", err)
            return nil
        }
        if info.IsDir() {
            return nil
        }
        if filepath.Ext(info.Name()) == fileExt {
            oldPath := path
            newPath := filepath.Join(tempDirectory, info.Name())
            if err := os.Rename(oldPath, newPath); err != nil {
                log.Printf("[WARN] %v", err)
            }
        }
        return nil
    })
    if err != nil {
        log.Fatalf("[ERRO] %v", err)
    }
}
