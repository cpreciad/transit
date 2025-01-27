package main

import (
    "log"
    "fmt"
    "time"
    "github.com/cpreciad/transit/internal/request"
    "github.com/cpreciad/transit/internal/parser"
    "github.com/cpreciad/transit/internal/helpers"
)

// 3 is to take into account for the consistent time it takes for the 
// N to travel through the tunnel between Cole valley to Duboce Park
const TunnelTime = time.Duration(3) * time.Minute


func main(){
    body, err := request.RequestStops()
    if err != nil{
        log.Fatal(err)
    }

    stopIds, err := parser.ParseStopID(body)
    if err != nil{
        log.Fatal(err)
    }

    var inbound, outbound *parser.ConciseStopInfo
    for _, stopId := range stopIds{
        body, err := request.RequestNextArrivals(stopId)
        if err != nil{
            log.Fatal(err)
        }
        stopInfo, err := parser.ParseNextArrival(body, stopId)
        if err != nil{
            log.Fatal(err)
        }
        if stopInfo.Direction == "IB"{
            inbound = stopInfo
        } else if stopInfo.Name == "Duboce St/Noe St/Duboce Park"{
            outbound = stopInfo
        }
    }
    display(inbound, outbound)
}

func display(inboundStopInfo *parser.ConciseStopInfo, outboundStopInfo *parser.ConciseStopInfo){
    // display inbound times
    fmt.Printf("Inbound N line train times for Duboce Stop:\n")
    for stopInfo := inboundStopInfo; stopInfo != nil ; stopInfo = stopInfo.Next{
        t, err := helpers.UTCtoPST(stopInfo.ExpectedTime)
        formattedTime := convertTime(stopInfo, t)
        if err != nil{
            log.Println(err)
        }
        if stopInfo.Next == nil{
            fmt.Printf("%s\n\n", formattedTime)
        }else{
            fmt.Printf("%s <- ", formattedTime)
        }
    }

    // display outbound times
    fmt.Printf("Outbound N line train times for Duboce Stop:\n")
    for stopInfo := outboundStopInfo; stopInfo != nil; stopInfo = stopInfo.Next{
        t, err := helpers.UTCtoPST(stopInfo.ExpectedTime)
        formattedTime := convertTime(stopInfo, t)
        if err != nil{
            log.Println(err)
        }
        if stopInfo.Next == nil{
            fmt.Printf("%s\n\n", formattedTime)
        }else{
            fmt.Printf("%s <- ", formattedTime)
        }
    }
}

func convertTime(stopInfo *parser.ConciseStopInfo, t time.Time) string{
    if stopInfo.Direction == "IB"{
        t = t.Add(TunnelTime)
    }
    return t.Format(time.Kitchen)
}
