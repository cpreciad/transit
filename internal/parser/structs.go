// TODO: depreciate this file
package parser

type StopJSON struct {
	Contents Contents `json:"Contents"`
}

type Contents struct {
	DataObjects DataObjects `json:"dataObjects"`
}

type DataObjects struct {
	ScheduledStopPoint []Stops `json:"ScheduledStopPoint"`
}

type Stops struct {
	Id       string   `json:"id"`
	Name     string   `json:"Name"`
	Location Location `json:"Location"`
	// Url string `json:"Url"`
	// StopType string `json:"StopType"`
}

// stop monitoring data

type StopMonitoringJSON struct {
	ServiceDelivery ServiceDelivery `json:"ServiceDelivery"`
}

type ServiceDelivery struct {
	StopMonitoringDelivery StopMonitoringDelivery `json:"StopMonitoringDelivery"`
}

type StopMonitoringDelivery struct {
	MonitoredStopVisits []MonitoredStopVisit `json:"MonitoredStopVisit"`
}

type MonitoredStopVisit struct {
	MonitoredVehicleJourney MonitoredVehicleJourney `json:"MonitoredVehicleJourney"`
	OperatorRef             string                  `json:"OperatorRef"`
}

type MonitoredVehicleJourney struct {
	LineRef         string        `json:"LineRef"`
	DirectionRef    string        `json:"DirectionRef"`
	VehicleLocation Location      `json:"VehicleLocation"`
	MonitoredCall   MonitoredCall `json:"MonitoredCall"`
}

type MonitoredCall struct {
	StopPointName       string `json:"StopPointName"`
	ExpectedArrivalTime string `json:"ExpectedArrivalTime"`
}

// for both
type Location struct {
	Longitude string
	Latitude  string
}
