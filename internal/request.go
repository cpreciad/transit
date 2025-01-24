package request

import (
    "net/http"
    "fmt"
    "io"
)

const(
    url = "http://example.com/"
)

func Request() {
    resp, err := http.Get(url)
    if err != nil{
        fmt.Printf("Request: failed to get response from url %s\n", url)
    }
    body, err := io.ReadAll(resp.Body)
    resp.Body.Close()

    fmt.Printf("%s\n", body)

}
