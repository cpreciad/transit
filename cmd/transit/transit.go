package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cpreciad/transit/internal/consolidator"
	"github.com/cpreciad/transit/internal/helpers"
	"github.com/cpreciad/transit/internal/parser"
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
			t, err := helpers.UTCtoPST(stopInfo.ExpectedTime)
			if err != nil {
				log.Println(err)
			}
			formattedTime := convertTime(t)
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
		localDisplay(inboundStopInfo)
		localDisplay(outboundStopInfo)
	}

}

func convertTime(t time.Time) string {
	/*
		if stopInfo.Direction == "IB" {
			t = t.Add(TunnelTime)
		}
	*/
	return t.Format(time.Kitchen)
}
