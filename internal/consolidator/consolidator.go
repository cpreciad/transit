package consolidator

import (
	"fmt"
	"log"

	"github.com/cpreciad/transit/internal/parser"
	"github.com/cpreciad/transit/internal/request"
)

type Info struct {
	Direction *direction
}

type direction struct {
	Inbound  *parser.ConciseStopInfo
	Outbound *parser.ConciseStopInfo
}

func GetStopInfo(operatorId, lineId string, stops map[string][]string) []*Info {
	body, err := request.RequestStops(operatorId, lineId)
	if err != nil {
		log.Fatal(err)
	}

	err = parser.ParseStopID(body, stops)
	if err != nil {
		log.Fatal(err)
	}
	info := make([]*Info, 0)
	for stop, stopIds := range stops {
		// logs if there are missing or extra stops
		_ = validate(stop, stopIds)

		var inbound, outbound *parser.ConciseStopInfo = nil, nil
		for _, stopId := range stopIds {
			body, err := request.RequestNextArrivals(operatorId, stopId)
			if err != nil {
				log.Fatal(err)
			}
			stopInfo, err := parser.ParseNextArrival(body, stopId)
			if err != nil {
				log.Fatal(err)
			}
			if stopInfo.Direction == "IB" {
				inbound = stopInfo
			} else {
				outbound = stopInfo
			}
		}

		info = append(info, &Info{
			Direction: &direction{
				Inbound:  inbound,
				Outbound: outbound,
			},
		})

	}
	return info
}

func validate(stop string, stopIds []string) error {
	var err error
	if len(stopIds) < 2 {
		err = fmt.Errorf("consolidator: %s has missing stop data, see logs", stop)
	} else if len(stopIds) > 2 {
		err = fmt.Errorf("consolidator: %s has extra stop data, see logs", stop)
	}
	if err != nil {
		log.Println(err)
		log.Println("consolidator: stopIds for ", stop, stopIds)
	}
	return err
}
