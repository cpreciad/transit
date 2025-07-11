package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cpreciad/transit/cmd/transit/duboce/consolidator"
	"github.com/cpreciad/transit/cmd/transit/duboce/dubocehelpers"
	"github.com/cpreciad/transit/cmd/transit/duboce/parser"
)

const (
	operatorId = "SF"
	lineId     = "N"
)

var stops map[string][]string

// 3 is to take into account for the consistent time it takes for the
// N to travel through the tunnel between Cole valley to Duboce Park
const TunnelTime = time.Duration(3) * time.Minute

func init() {

	stops = make(map[string][]string)
	stops["Duboce St/Noe St/Duboce Park"] = make([]string, 0)
	stops["Carl St & Cole St"] = make([]string, 0)

}

func main() {
	info := consolidator.GetStopInfo(operatorId, lineId, stops)
	display(info)
}

func display(infos []*consolidator.Info) {

	localDisplay := func(i *parser.ConciseStopInfo) {
		if i != nil {
			fmt.Printf("%s line (%s) train times for %s:\n", i.Line, i.Direction, i.StopName)
		}
		for stopInfo := i; stopInfo != nil; stopInfo = stopInfo.Next {
			t, err := dubocehelpers.UTCtoPST(stopInfo.ExpectedTime)
			if err != nil {
				log.Println(err)
			}
			formattedTime := convertTime(t, false)
			if stopInfo.Next == nil {
				fmt.Printf("%s\n\n", formattedTime)
			} else {
				fmt.Printf("%s <- ", formattedTime)
			}
		}
	}
	specialDisplay := func(i *parser.ConciseStopInfo) {
		special := false
		if i != nil {
			if i.StopName == "Carl St & Cole St" {
				i.StopName = "Duboce St/Noe St/Duboce Park"
				special = true
			}
			fmt.Printf("%s line (%s) train times for %s:\n", i.Line, i.Direction, i.StopName)
		}

		for stopInfo := i; stopInfo != nil; stopInfo = stopInfo.Next {
			t, err := dubocehelpers.UTCtoPST(stopInfo.ExpectedTime)
			if err != nil {
				log.Println(err)
			}
			formattedTime := convertTime(t, special)
			if stopInfo.Next == nil {
				fmt.Printf("%s\n\n", formattedTime)
			} else {
				fmt.Printf("%s <- ", formattedTime)
			}
		}
	}

	for _, info := range infos {
		inboundStopInfo := info.Direction.Inbound
		outboundStopInfo := info.Direction.Outbound
		// the only reason this is using a scoped function is because I felt like it
		specialDisplay(inboundStopInfo)
		localDisplay(outboundStopInfo)
	}

}

func convertTime(t time.Time, specialCase bool) string {
	if specialCase {
		t = t.Add(TunnelTime)
	}
	return t.Format(time.Kitchen)
}
