package parser

import (
    "fmt"
    "log"
    "encoding/json"
)

const (
    duboceStopName = "Duboce St/Noe St/Duboce Park"
    coleStopName = "Carl St & Cole St"
)

// stop data
type StopJSON struct{
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


// stop monitoring data

type StopMonitoringJSON struct{
    ServiceDelivery ServiceDelivery "json:ServiceDelivery"
}

type ServiceDelivery struct{
    StopMonitoringDelivery StopMonitoringDelivery "json:StopMonitoringDelivery"
}

type StopMonitoringDelivery struct {
    MonitoredStopVisits []MonitoredStopVisit "json:MonitoredStopVisit"
}

type MonitoredStopVisit struct {
    MonitoredVehicleJourney MonitoredVehicleJourney "json:MonitoredVehicleJourney"
}

type MonitoredVehicleJourney struct{
    DirectionRef string "json:DirectionRef"
    VehicleLocation Location "json:VehicleLocation"
    MonitoredCall MonitoredCall "json:MonitoredCall"
}

type MonitoredCall struct{
    StopPointName string "json:StopPointName"
    ExpectedArrivalTime string "json:ExpectedArrivalTime"
}

type ConciseStopInfo struct{
    Name string
    Direction string
    ExpectedTime string
    Location Location
    Next *ConciseStopInfo
}

// for both
type Location struct{
    Longitude string
    Latitude string
}

// ParseStopID - takes a byte slice and a string, parses the stop ID for a given stop name
func ParseStopID(data []byte) ([]string, error){
    var stopsJson StopJSON

    err := json.Unmarshal(data, &stopsJson)

    if err != nil{
        return nil, fmt.Errorf("Parser: json unmarchal error: %w",err)
    }

    stopIds := parseFindStopID(stopsJson)

    if len(stopIds) < 2{
        return nil, fmt.Errorf("Parser: an insufficient number of stop ids were obtained")
    } else if len(stopIds) == 4{
        log.Println("Parser: 4 stops were returned, check number of duboce stops")
    }

    return stopIds, nil
}


// ParseNextArrival - takes a byte slice and parses the next predicted times for the line
func ParseNextArrival(data []byte) (ConciseStopInfo, error){
    var stopMonitoringJson StopMonitoringJSON

    err := json.Unmarshal(data, &stopMonitoringJson)
    if err != nil{
        return ConciseStopInfo{}, err
    }
    stopInfo := parseRestructureTimes(stopMonitoringJson)
    return stopInfo, nil
}

func parseFindStopID(stopsJson StopJSON) []string{
    // walk the struct to get to the stop data list of objects

    // two annoyances here: for one, there isn't any info at this level for 
    // whether the stop is for inbound or outbound transit, so for a given 
    // stop, grab both and deal with it later
    // second, "Duboce St/Noe St/Duboce Park" only has data for outbound
    // trains (smfh) so in order to compensate, we'll have to grab the stop
    // data for Carl & Cole and approximate the arrival time off that

    // as of now, predicting we'll get three stops
    // if for some reason 4 get returned, and data is added for duboce, 
    // we'll have the capacity for it
    stopIds := make([]string, 0, 4)
    stopDataList := stopsJson.Contents.DataObjects.ScheduledStopPoint
    for _, stopData := range stopDataList{
        if stopData.Name == duboceStopName || stopData.Name == coleStopName{
            stopIds = append(stopIds, stopData.Id)
        }
    }
    return stopIds
}

func parseRestructureTimes(stopMonitoringJson StopMonitoringJSON) ConciseStopInfo{
    // implement restructuring linked list, return head 
    return ConciseStopInfo{}
}
