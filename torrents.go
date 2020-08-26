package main

import (
    "encoding/json"
    "log"
    "regexp"
)

func (qB *Client) getTorrentsTracker(torrentList []string) {
    for _, hash := range torrentList {
        hashUrl := qB.ipAddr + "/api/v2/torrents/trackers?hash=" + hash
        _, body := qB.createReq("GET", hashUrl, nil)

        var trackers []Trackers
        err := json.Unmarshal(body.Bytes(), &trackers)
        if err != nil {
            log.Fatal(err)
        }

        for _, tracker := range trackers {
            tUrl := tracker.URL
            b, _ := regexp.MatchString(`(?i)(daydream.dmhy.best|tracker.dmhy.org)/announce\?(secure|passkey)`, tUrl)
            if b {
                qB.changeKey(hash, tUrl)
            }
        }
    }
}
