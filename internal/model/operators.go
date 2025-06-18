package model

/*
{
    "Id": "SS",
    "Name": "City of South San Francisco",
    "ShortName": "South City Shuttle",
    "SiriOperatorRef": null,
    "TimeZone": "America/Vancouver",
    "DefaultLanguage": "en",
    "ContactTelephoneNumber": null,
    "WebSite": "http://www.ssf.net/SCS",
    "PrimaryMode": "bus",
    "PrivateCode": "SS",
    "Monitored": false,
    "OtherModes": ""
  },
*/

type Operator struct {
	OperatorID string `json:"Id"`
	Name       string `json:"Name"`
}
