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

type Settings struct {
    timer int
    proxy string
    timeout int
    file string
    useragent string
}

func testSite(u string, s Settings) bool {
    if u == "" {
        return false
    }
    client := http.Client{
        Timeout: time.Duration(s.timeout) * time.Second,
    }
    if s.proxy != "" {
        proxyUrl, _ := url.Parse(s.proxy)
        client = http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
    }
    req, _ := http.NewRequest("GET", u, nil)
    req.Header.Add("User-Agent", s.useragent)
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

func getParamsFromArgs(s Settings) Settings {
    if len(os.Args) > 1 {
        for i, _ := range os.Args {
            if os.Args[i] == "proxy" {
                i++
                s.proxy = os.Args[i]
                fmt.Println("use proxy: ", os.Args[i])
            }
            if os.Args[i] == "timer" {
                i++
                s.timer, _ = strconv.Atoi(os.Args[i])
                fmt.Println("use timer: ", os.Args[i])
            }
            if os.Args[i] == "timeout" {
                i++
                s.timeout, _ = strconv.Atoi(os.Args[i])
                fmt.Println("use timeout: ", os.Args[i])
            }
            if os.Args[i] == "file" {
                i++
                s.file = os.Args[i]
                fmt.Println("use file: ", os.Args[i])
            }
            if os.Args[i] == "useragent" {
                i++
                s.useragent = os.Args[i]
                fmt.Println("use useragent: ", os.Args[i])
            }
        }
    }
    return s
}

func initSettings() Settings {
    s := Settings{
        timer: 1,
        proxy: "",
        timeout: 5,
        file: "sites.txt",
        useragent: "Mozilla/5.0 (Windows NT x.y; Win64; x64; rv:10.0) Gecko/20100101 Firefox/10.0",
    }
    return s
}

func main() {
    s := initSettings()
    s = getParamsFromArgs(s)
    sites := getSitesFromFile(s.file)
    ch := make(chan byte, 1)
    for _, site := range sites {
        go func (site string) {
            for {
                testSite(site, s)
                time.Sleep(time.Duration(s.timer) * time.Second)
            }
            ch <- 1
        }(site)
    }
    <-ch
}