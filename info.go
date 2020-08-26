package main

import (
    "encoding/json"
    "log"
)

func (qB *Client) getInfo() []string {
    dataURL := qB.ipAddr + "/api/v2/torrents/info"
    _, body := qB.createReq("GET", dataURL, nil)

    var tList []Torrent
    err := json.Unmarshal(body.Bytes(), &tList)
    if err != nil {
        log.Fatal(err)
    }

    var hashList []string
    for _, t := range tList {
        hashList = append(hashList, t.Hash)
    }
    return hashList
}
