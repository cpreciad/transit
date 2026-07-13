package fetch

import (
	"fmt"
	"time"
	"log"
	"github.com/cpreciad/transit/internal/consolidator"
	"github.com/cpreciad/transit/internal/parser"

)

// 3 is to take into account for the consistent time it takes for the
// N to travel through the tunnel between Cole valley to Duboce Park
const (
	TunnelTime = time.Duration(3) * time.Minute
)

func DisplayDuboceIoT() string{
	info := fetchDuboce()
	return getIOTString(info.Direction.Inbound)
}

func DisplayDubocePST(){
	info := fetchDuboce()
	displayPst(info.Direction.Inbound)
	displayPst(info.Direction.Outbound)
}

// getIOTString - returns a single string with newlines with the following format: 
//              - : N, to Ball Park
//                  10, 12, 29
// input is a conciseStopInfo. I hate that I made a ConciseStopInfo as a function of direction fml
func getIOTString(i *parser.ConciseStopInfo) string{
	if i == nil {
		return fmt.Sprintf("no more stops for the day")
	}
	var destination string

	switch i.Direction {
		case "IB":
			destination = "Ball Park"
		case "OB":
			destination = "Ocean Beach"
		default:
			return fmt.Sprintf("unknown direction")
	}

	iotString := fmt.Sprintf("%s, to %s\n", i.Line, destination)
	for stopInfo := i; stopInfo != nil; stopInfo = stopInfo.Next {
		t := stopInfo.ExpectedTime
		minutesTil := time.Until(t).Minutes()
		formattedTime := fmt.Sprintf("%.0f", minutesTil)

		if stopInfo.Next == nil {
			iotString = iotString + fmt.Sprintf("%s", formattedTime)
		} else {
			iotString = iotString + fmt.Sprintf("%s, ", formattedTime)
		}
	}
	return iotString
}

func displayPst(i *parser.ConciseStopInfo){
	if i == nil {
		fmt.Println("no more stops for the day")
		return
	}
	var arrow string
	var destination string

	switch i.Direction {
		case "IB":
			arrow = "<=="
			destination = "the Ball Park"
		case "OB":
			arrow = "==>"
			destination = "Ocean Beach"
		default:
			fmt.Println("unknown direction")
			return 
	}
	fmt.Printf("%s, towards %s: arrival times at Duboce Park\n", i.Line, destination) 
	for stopInfo := i; stopInfo != nil; stopInfo = stopInfo.Next {
		t := stopInfo.ExpectedTime
		kitchenTime := t.Format(time.Kitchen)
		minutesTil := time.Until(t).Minutes()
		formattedTime := fmt.Sprintf("%s (in %.0f minutes)", kitchenTime, minutesTil)

		if stopInfo.Next == nil {
			fmt.Printf("%s\n\n", formattedTime)
		} else {
			fmt.Printf("%s %s ", formattedTime, arrow)
		}
	}
}

func fetchDuboce() *consolidator.Info{
	operatorId  := "SF"
	lineId      := "N"
	// turns out that for the same stop location, there are different IDs and names 
	// associated with it. I figured this out by going online to munis website,
	// seeing that the two different ids correlated to the same stop, and found 
	// the correct names to use
	outboundStopName  := "Duboce St/Noe St/Duboce Park"
	inboundStopName := "Sunset Tunnel East Portal"

	stops := make(map[string][]string)
	stops[outboundStopName] = make([]string, 0)
	stops[inboundStopName ] = make([]string, 0)

	allInfo := fetch(operatorId, lineId, stops) 
	
	duboceInfo := allInfo[outboundStopName]

	if duboceInfo.Direction.Inbound == nil{	
		duboceInfo.Direction.Inbound = allInfo[inboundStopName].Direction.Inbound
	}

	return duboceInfo
}

func fetch (operatorId, lineId string, stops map[string][]string) map[string]*consolidator.Info{
	infoMap := make(map[string]*consolidator.Info)
	for _, info := range consolidator.GetStopInfo(operatorId, lineId, stops){
		stopName := info.StopName
		if stopName == "" {
			log.Printf("fetch: could not derrive stop name from info")
		}
		infoMap[stopName] = info
	} 
	return infoMap
}

