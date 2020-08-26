package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "syscall"

    "golang.org/x/crypto/ssh/terminal"
)

type Client struct {
    ipAddr  string
    sId     string
    passKey string
}

type Torrent struct {
    Hash              string  `json:"hash"`
    Tracker           string  `json:"tracker"`
}

type Trackers struct {
    Msg           string      `json:"msg"`
    NumDownloaded int         `json:"num_downloaded"`
    NumLeeches    int         `json:"num_leeches"`
    NumPeers      int         `json:"num_peers"`
    NumSeeds      int         `json:"num_seeds"`
    Status        int         `json:"status"`
    URL           string      `json:"url"`
}

func inputText(msg string) string {
    fmt.Print(msg)
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        return scanner.Text()
    }
    return ""
}

func inputPassword() string {
    fmt.Print("Password: ")
    bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
    if err != nil {
        log.Fatal(err)
    }
    return string(bytePassword)
}
