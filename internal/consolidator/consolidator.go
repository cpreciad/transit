package consolidator

import (
    "fmt"

	"github.com/cpreciad/transit/internal/parser"
	"github.com/cpreciad/transit/internal/request"
)

type Info struct {
	StopName string
	Direction *direction
}

type direction struct {
	Inbound  *parser.ConciseStopInfo
	Outbound *parser.ConciseStopInfo
}

// GetStopInfo makes two requests to the 511 api. keep that in mind when calling this on an interval
func GetStopInfo(operatorId, lineId string, stops map[string][]string) ([]*Info, error) {
	body, err := request.RequestStops(operatorId, lineId)
	if err != nil {
		return nil, fmt.Errorf("consolidator: request for fresh stop data for %s-%s failed: %v\n", operatorId, lineId, err)
	}

	err = parser.ParseStopID(body, stops)
	if err != nil {
        return nil, fmt.Errorf("consolidator: failed to parse stop ids from input: %v", err)
	}
	info := make([]*Info, 0)
	for stop, stopIds := range stops {
		i := &Info{
			StopName: stop,
		}
		var inbound, outbound *parser.ConciseStopInfo = nil, nil
		for _, stopId := range stopIds {
			body, err := request.RequestUpcomingArrivals(operatorId, stopId)

			if err != nil {
				return nil, fmt.Errorf("consolidator: request for fresh arrival data for %s-%s-%s failed: %v", operatorId, lineId, stopId, err)
			}

			stopInfo, err := parser.ParseUpcomingArrivals(body, stopId)
			if err != nil {
                return nil, fmt.Errorf("consolidator could not parse arrival data: %v:", err)
			}
			if stopInfo.Direction == "IB" {
				inbound = stopInfo
			} else {
				outbound = stopInfo
			}
		}

		i.Direction = &direction{
				Inbound:  inbound,
				Outbound: outbound,
		}

		info = append(info, i)

	}
	return info, nil
}

