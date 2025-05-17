package parser

import (
	"encoding/json"
	"fmt"
	"log"
)

type ConciseStopInfo struct {
	StopName     string
	Direction    string
	ExpectedTime string
	Location     Location
	Next         *ConciseStopInfo
}

// ParseStopID - takes a byte slice and a string, parses the stop ID for a given stop name
func ParseStopID(data []byte, stops map[string]string) ([]string, error) {
	var stopsJson StopJSON

	err := json.Unmarshal(data, &stopsJson)

	if err != nil {
		return nil, fmt.Errorf("parser: json unmarchal error: %w", err)
	}

	stopIds := parseFindStopID(stopsJson, stops)

	if len(stopIds) < 2 {
		return nil, fmt.Errorf("parser: an insufficient number of stop ids were obtained")
	} else if len(stopIds) == 4 {
		log.Println("Parser: 4 stops were returned, check number of duboce stops")
	}

	return stopIds, nil
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

func parseFindStopID(stopsJson StopJSON, stops map[string]string) []string {
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
	for _, stopData := range stopDataList {
		if _, ok := stops[stopData.Name]; ok {
			stopIds = append(stopIds, stopData.Id)
		}
	}
	return stopIds
}

func parseRestructureTimes(stopMonitoringJson StopMonitoringJSON) *ConciseStopInfo {
	// need to check for case if there are no entries in MonitoredStopVisit list
	// seems like there is a max of 3

	dummy := &ConciseStopInfo{}
	builder := dummy

	MSV := stopMonitoringJson.ServiceDelivery.StopMonitoringDelivery.MonitoredStopVisits

	for _, object := range MSV {
		builder.Next = &ConciseStopInfo{
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
