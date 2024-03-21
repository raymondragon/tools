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
    scanDir = flag.String("s", "", "Scan Directory")
    tempDir = flag.String("t", "", "Temp Directory")
)

func main() {
    flag.Parse()
    if *fileExt == "" || *scanDir == "" || *tempDir == "" {
        flag.Usage()
        log.Fatalf("[ERRO] %v", "Invalid Flag(s)")
    }
    log.Printf("[INFO] %v <-> [*%v] --> %v", *scanDir, *fileExt, *tempDir)
    for {
        moveTempFiles(*fileExt, *scanDir, *tempDir)
        time.Sleep(1 * time.Hour)
    }
}

func moveTempFiles(fileExt, scanDir, tempDir string) {
    if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
        log.Fatalf("[ERRO] %v", err)
    }
    if walkErr := filepath.Walk(scanDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            return nil
        }
        if filepath.Ext(info.Name()) == fileExt {
            oldPath := path
            newPath := filepath.Join(tempDir, info.Name())
            if err := os.Rename(oldPath, newPath); err != nil {
                return err
            }
        }
        return nil
    }); walkErr != nil {
        log.Fatalf("[ERRO] %v", walkErr)
    }
}
