package consolidator

import (
	"fmt"
	"log"

	"github.com/cpreciad/transit/internal/consolidator/backup"
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
		log.Printf("consolidator: request for fresh stop data for %s-%s failed: %v\n", operatorId, lineId, err)
		log.Println("consolidator: attempting to load backup data")
	}

	fileName := fmt.Sprintf("%s_%s.json", operatorId, lineId)
	if body, err = handleBackup(fileName, body); err != nil {
		log.Fatalf("consolidator: failed to load backup data: %v\n", err)
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
				log.Printf("consolidator: request for fresh arrival data for %s-%s-%s failed: %v \n", operatorId, lineId, stopId, err)
				log.Println("attempting to load backup data")
			}
			fileName := fmt.Sprintf("%s_%s.json", operatorId, stopId)
			if body, err = handleBackup(fileName, body); err != nil {
				log.Fatalf("consolidator: failed to load backup data: %v\n", err)
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

func handleBackup(fileName string, body []byte) ([]byte, error) {
	if body != nil {
		// success was seen, we'll write this to disc
		err := backup.StoreBackup(fileName, body)
		if err != nil {
			// we won't panic on error here, just log the fail to backup
			log.Printf("consolidator: failed to store backup for %s to disc: %v\n", fileName, err)
		}
	} else {
		backup, err := backup.LoadBackup(fileName)
		if err != nil {
			return nil, err
		}
		log.Printf("consolidator: backup data for %s loaded successfully\n", fileName)
		body = backup
	}
	return body, nil
}
