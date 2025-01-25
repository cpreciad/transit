package parser

import (
    "fmt"
    "encoding/json"
)

const (
    //targetStopName = "Duboce St/Noe St/Duboce Park"
    targetStopName = "Duboce Ave & Church St"
)

type JSON struct{
    Contents Contents
}

type Contents struct{
    ResponseTimestamp string
    DataObjects DataObjects
}

type DataObjects struct{
    Id string
    ScheduledStopPoint []Stops
}

type Stops struct{
    Id string
    Extensions map[string]interface{}
    Name string
    Location Location
    Url string
    StopType string
}

type Location struct{
    Longitude string
    Latitude string
}

// ParseStopID - takes a byte slice and a string, parses the stop ID for a given stop name
func ParseStopID(data []byte) (string, error){
    var stopsJson JSON

    err := json.Unmarshal(data, &stopsJson)

    if err != nil{
        return "", fmt.Errorf("Parser: json unmarchal error: %w",err)
    }

    stopId := parseFindStopID(stopsJson, targetStopName)

    if stopId == ""{
        return "", fmt.Errorf("Parser: could not find stopID for %s", targetStopName)
    }

    return stopId, nil
}


// ParseNextArrival - takes a byte slice and parses the next predicted times for the line
func ParseNextArrival() error{
    return nil
}

func parseFindStopID(stopsJson JSON, stopName string) string{
    // walk the struct to get to the stop data list of objects
    stopDataList := stopsJson.Contents.DataObjects.ScheduledStopPoint
    for _, stopData := range stopDataList{
        if stopData.Name == stopName{
            fmt.Println(stopData)
            //return stopData.Id
        }
    }
    return ""
}
