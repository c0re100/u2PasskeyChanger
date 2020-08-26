package main

import (
    "fmt"
    "net/url"
    "strings"
)

func (qB *Client) changeKey(hash, origUrl string) {
    nUrl := "https://daydream.dmhy.best/announce?passkey="+qB.passKey
    if nUrl == origUrl {
        return
    }

    data := url.Values{}
    data.Add("hash", hash)
    data.Add("origUrl", origUrl)
    data.Add("newUrl", nUrl)

    editUrl := qB.ipAddr + "/api/v2/torrents/editTracker"
    res, _ := qB.createReq("POST", editUrl, strings.NewReader(data.Encode()))

    if res.StatusCode == 200 {
        fmt.Printf("Hash '%v' tracker changed.\n", hash)
    }
}
