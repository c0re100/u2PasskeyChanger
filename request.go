package main

import (
    "bytes"
    "io"
    "log"
    "net/http"
)

func (qB *Client) createReq(method, endpoint string, reader io.Reader) (*http.Response, *bytes.Buffer) {
    req, err := http.NewRequest(method, endpoint, reader)
    if err != nil {
        log.Fatal(err)
    }

    req.Header.Add("user-agent", "dllmch hacker")
    req.Header.Add("content-type", "application/x-www-form-urlencoded")
    req.Header.Add("referer", qB.ipAddr)
    req.Header.Add("origin", qB.ipAddr)
    if qB.sId != "" {
        req.Header.Add("cookie", "SID="+qB.sId)
    }

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }

    body := &bytes.Buffer{}
    _, err = body.ReadFrom(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    return res, body
}
