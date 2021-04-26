package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "time"
    //"net/url"
)

type Response struct {
    Original_url   string
    Status_code    int
    Redirected_url string
}

func main () {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/checklink/desktop", checkLinkDesktop)
    router.HandleFunc("/checklink/mobile", checkLinkMobile)
    router.HandleFunc("/healthcheck", healthCheck)
    log.Fatal(http.ListenAndServe(":8080", router))
}


func checkLinkDesktop(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    decoder := json.NewDecoder(r.Body)
    var link string
    err := decoder.Decode(&link)
    user_agent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36"
    original_url,status_code,redirected_url := checkURL(link,user_agent)
    log.Println("source url:",original_url,"status code:",status_code,"redirected url:",redirected_url)
    resp := Response{original_url,status_code,redirected_url}
    js, err := json.Marshal(resp)
    w.Write(js)
    if err != nil {
        panic(err)
    }
}

func checkLinkMobile(w http.ResponseWriter, r *http.Request) {
    //w.Header().Set("Content-Type", "application/json")
    decoder := json.NewDecoder(r.Body)
    var link string
    err := decoder.Decode(&link)
    user_agent := "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1"
    original_url,status_code,redirected_url := checkURL(link,user_agent)
    log.Println("source url:",original_url,"status code:",status_code,"redirected url:",redirected_url)
    resp := Response{original_url,status_code,redirected_url}
    js, err := json.Marshal(resp)
    w.Write(js)
    if err != nil {
        panic(err)
    }
}

func checkURL(link string, client string) (string, int, string) {
    webclient := &http.Client{Timeout: 3 * time.Second,}
    /*_, valid_err := url.ParseRequestURI(link)
        if valid_err != nil {
        log.Println(valid_err)                      //better to implement validation step on frontend
        return link, 400, "Invalid URL provided"
       }
    */
    req, err := http.NewRequest("GET", link, nil)
    req.Header.Add("User-Agent", client)
    resp, err := webclient.Do(req)
        if err != nil {
            log.Println(err)
            return link, 522, "Can not be reached or request timeout"
        }
    finalURL := resp.Request.URL.String()
    return link, resp.StatusCode, finalURL
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("ALIVE"))
}