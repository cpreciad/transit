package main

import (
    "log"
    "os"
    "github.com/cpreciad/transit/internal/request"
    "github.com/cpreciad/transit/internal/parser"
)

func main(){
    body, err := request.RequestStops()
    if err != nil{
        log.Fatal(err)
    }

    stopId, err := parser.ParseStopID(body)

    if err != nil{
        log.Fatal(err)
    }
    log.Println(stopId)
}

func writeData(body []byte) error{
    var filePath string = "data/stops_data.txt"
    fd, err := os.Create(filePath)
    defer fd.Close()
    if err != nil{
        log.Fatal(err)
    }

    n, err := fd.Write(body)
    if err != nil{
        log.Fatal(err)
    }

    log.Printf("Main: wrote %d bytes to %s", n, filePath)
    return nil
}
