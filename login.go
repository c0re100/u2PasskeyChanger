package main

import (
    "fmt"
    "net/url"
    "strings"
)

func (qB *Client) login() string {
    data := url.Values{}
    data.Add("username", inputText("Username: "))
    data.Add("password", inputPassword())

    loginUrl := qB.ipAddr + "/api/v2/auth/login"
    res, _ := qB.createReq("POST",loginUrl, strings.NewReader(data.Encode()))
    if res.StatusCode == 200 {
        fmt.Println()
        fmt.Println("Login success.")
        return res.Cookies()[0].Value
    }
    return ""
}
