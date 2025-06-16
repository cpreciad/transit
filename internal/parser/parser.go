// TODO: depreciate this file
package parser

import (
	"encoding/json"
	"fmt"
)

type ConciseStopInfo struct {
	Operator     string
	Line         string
	StopName     string
	Direction    string
	ExpectedTime string
	Location     Location
	Next         *ConciseStopInfo
}

// ParseStopID - takes a byte slice and a string, parses the stop ID for a given stop name
func ParseStopID(data []byte, stops map[string][]string) error {
	var stopsJson StopJSON

	err := json.Unmarshal(data, &stopsJson)

	if err != nil {
		return fmt.Errorf("parser: json unmarchal error: %w", err)
	}

	// pass the stops map in, and get stop IDs mapped to each stop
	parseFindStopID(stopsJson, stops)

	return nil
}

// ParseNextArrival - takes a byte slice and parses the next predicted times for the line
func ParseNextArrival(data []byte, stopId string) (*ConciseStopInfo, error) {
	var stopMonitoringJson StopMonitoringJSON

	err := json.Unmarshal(data, &stopMonitoringJson)
	if err != nil {
		return nil, err
	}
	stopInfo := parseRestructureTimes(stopMonitoringJson)
	if stopInfo == nil {
		return nil, fmt.Errorf("parser: no arrivals for the stop ID %s", stopId)
	}
	return stopInfo, nil
}

// see if there are stop Ids to parse from the stopsJson given some target stops
// and append them to the particular stop
func parseFindStopID(stopsJson StopJSON, targetStops map[string][]string) {

	stopDataList := stopsJson.Contents.DataObjects.ScheduledStopPoint
	for _, stopData := range stopDataList {
		if _, ok := targetStops[stopData.Name]; ok {
			targetStops[stopData.Name] = append(targetStops[stopData.Name], stopData.Id)
		}
	}
}

func parseRestructureTimes(stopMonitoringJson StopMonitoringJSON) *ConciseStopInfo {
	// need to check for case if there are no entries in MonitoredStopVisit list
	// seems like there is a max of 3

	// the only reason this is a linked list is because I wanted to practice
	// building one...😱
	dummy := &ConciseStopInfo{}
	builder := dummy

	MSV := stopMonitoringJson.ServiceDelivery.StopMonitoringDelivery.MonitoredStopVisits

	for _, object := range MSV {
		builder.Next = &ConciseStopInfo{
			Operator:     object.OperatorRef,
			Line:         object.MonitoredVehicleJourney.LineRef,
			StopName:     object.MonitoredVehicleJourney.MonitoredCall.StopPointName,
			Direction:    object.MonitoredVehicleJourney.DirectionRef,
			ExpectedTime: object.MonitoredVehicleJourney.MonitoredCall.ExpectedArrivalTime,
			Location:     object.MonitoredVehicleJourney.VehicleLocation,
			Next:         nil,
		}

		builder = builder.Next
	}

	return dummy.Next
}
