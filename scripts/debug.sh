#! /bin/bash
# if you need to see what is in the requests that transit makes, look out below
# This stop code is associated with the N line in SF
wget -O nstops.json "http://api.511.org/transit/StopMonitoring?api_key=$TRANSIT_DATA_API_KEY&agency=SF&stopCode=17252&format=json"

# This one is to look at all stops for all lines in SF
wget -O allstops.json "http://api.511.org/transit/lines?api_key=$TRANSIT_DATA_API_KEY&operator_id=SF&format=json"
