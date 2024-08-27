package main

import (
    "flag"
    "io"
    "log"
    "os"
    "path/filepath"
    "time"
)

var (
    sourceDir = flag.String("s", "", "Source Directory")
    targetDir = flag.String("t", "", "Target Directory")
)

func main() {
    flag.Parse()
    if *sourceDir == "" || *targetDir == "" {
        flag.Usage()
        log.Fatalf("[ERRO] %v", "Invalid Flag(s)")
    }
    log.Printf("[INFO] %v --> %v", *sourceDir, *targetDir)
    for {
        backupFiles(*sourceDir, *targetDir)
        time.Sleep(24 * time.Hour)
    }
}

func backupFiles(sourceDir, targetDir string) {
    if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
        log.Fatalf("[ERRO] %v", err)
    }
    if walkErr := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            return nil
        }
        relPath, err := filepath.Rel(sourceDir, path)
        if err != nil {
            return err
        }
        targetPath := filepath.Join(targetDir, relPath)
        if err := os.MkdirAll(filepath.Dir(targetPath), os.ModePerm); err != nil {
            return err
        }
        if err := copyFile(path, targetPath); err != nil {
            return err
        }
        return nil
    }); walkErr != nil {
        log.Fatalf("[ERRO] %v", walkErr)
    }
}

func copyFile(src, dest string) error {
    srcFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    destFile, err := os.Create(dest)
    if err != nil {
        return err
    }
    defer destFile.Close()

    if _, err := io.Copy(destFile, srcFile); err != nil {
        return err
    }

    if err := destFile.Sync(); err != nil {
        return err
    }

    return nil
}
