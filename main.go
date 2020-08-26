package main

import (
    "fmt"
    "log"
)

func main() {
    fmt.Println("A program for qBittorrent edit U2 tracker")

    client := &Client{
        ipAddr: inputText("URL: "),
    }
    for {
        if client.passKey == "" {
            client.passKey = inputText("Passkey: ")
        } else {
            break
        }
    }
    client.sId = client.login()
    if client.sId == "" {
        log.Fatal("Login fail.")
    }
    torrentList := client.getInfo()
    client.getTorrentsTracker(torrentList)
    fmt.Println("All u2 torrent tracker updated.")
}
