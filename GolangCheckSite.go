package main

import (
    "net/http"
    "net/url"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
    "time"
    "strconv"
)

var timer int = 1
var proxy string = ""

func testSite(u string) bool {
    if u == "" {
        return false
    }
    client := http.Client{
        Timeout: time.Duration(5 * time.Second),
    }
    if proxy != "" {
        proxyUrl, _ := url.Parse(proxy)
        client = http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
    }
    req, _ := http.NewRequest("GET", u, nil)
    req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT x.y; Win64; x64; rv:10.0) Gecko/20100101 Firefox/10.0")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("error: ", err)
        return false
    } 
    _, err = ioutil.ReadAll(resp.Body)
    if err == nil {
        if resp.StatusCode == 200 {
            fmt.Println(u + ": " + "ok")
        } else {
            fmt.Println("Site ", u, " returned error code: ", resp.StatusCode)   
        }
    }
    defer resp.Body.Close()
    return true
}

func getSitesFromFile(namefile string) []string {
    data, err := ioutil.ReadFile(namefile)
    if err != nil {
        fmt.Println("error: ", err)
        os.Exit(2)
    }
    result := string(data)
    
    test := strings.Split(result, "\r\n")
    return test
}

func getParamsFromArgs() {
    if len(os.Args) > 1 {
        for i, _ := range os.Args {
            if os.Args[i] == "proxy" {
                i++
                proxy = os.Args[i]
            }
            if os.Args[i] == "timer" {
                i++
                timer, _ = strconv.Atoi(os.Args[i])
            }
        }
    }
}

func main() {
    sites := getSitesFromFile("sites.txt")
    getParamsFromArgs()
    if timer > 1 {
        fmt.Println("use timer: ", timer)
    }
    if proxy != "" {
        fmt.Println("use proxy: ", proxy)
    }
    ch := make(chan byte, 1)
    for _, site := range sites {
        go func (site string) {
            for {
                testSite(site)
                time.Sleep(time.Duration(timer) * time.Second)
            }
            ch <- 1
        }(site)
    }
    <-ch
}