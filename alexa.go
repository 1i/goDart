package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
)

func IntentDispatcher(ctx context.Context, request alexa.Request) (alexa.Response, error) {
	fmt.Print("start IntentDispatcher \n")
	spew.Dump(request)
	var response alexa.Response

	switch request.Body.Intent.Name {
	case "LaunchRequest":
		response = alexa.NewSimpleResponse("Unknown Request", makeRequest("Bray"))
	case "stationIntent":
		station := request.Body.Intent.Slots["station"].Value
		if len(station) == 0 {
			fmt.Println("Unable to get station")

			response = alexa.NewSimpleResponse("Unknown Request", "Unable to get station")
		}
		fmt.Printf("station %s . \n" + station)

		response = alexa.NewSimpleResponse("Unknown Request", makeRequest(station))
	default:
		response = alexa.NewSimpleResponse("Unknown Request", "The intent was unrecognized")
	}
	fmt.Print("end IntentDispatcher \n")

	return response, nil
}

func main() {
	lambda.Start(IntentDispatcher)
}

type Result struct {
	XMLName  xml.Name   `xml:"ArrayOfObjStationData"`
	Stations []Stations `xml:"objStationData"`
}

type Stations struct {
	XMLName         xml.Name `xml:"objStationData"`
	TrainType       string   `xml:"Traintype"`
	StationFullName string   `xml:"Stationfullname"`
	Destination     string   `xml:"Destination"`
	Direction       string   `xml:"Direction"`
	DueIn           string   `xml:"Duein"`
}

func makeRequest(station string) string {
	startionUrl := fmt.Sprintf("http://api.irishrail.ie/realtime/realtime.asmx/getStationDataByNameXML?StationDesc=%s", station)
	xmlFile, err := http.Get(startionUrl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfull request")
	defer xmlFile.Body.Close()
	byteValue, _ := io.ReadAll(xmlFile.Body)

	var results Result
	xml.Unmarshal(byteValue, &results)

	//fmt.Printf("%+v\n", results)
	fmt.Println("Station : " + results.Stations[0].StationFullName)
	northBoundresponse := "Northbound Trains due in "
	southBoundresponse := "Southbound Trains due in "
	for i := 0; i < len(results.Stations); i++ {

		dueInTime, err := strconv.Atoi(results.Stations[i].DueIn)
		if err != nil {
			fmt.Println(err)
		}

		if results.Stations[i].Direction == "Northbound" && dueInTime < 60 {
			northBoundresponse = northBoundresponse + ", " + results.Stations[i].DueIn
		}
		if results.Stations[i].Direction == "Southbound" && dueInTime < 60 {
			southBoundresponse = southBoundresponse + ", " + results.Stations[i].DueIn
		}

		fmt.Printf("Direction: %s ", results.Stations[i].Direction)
		fmt.Printf("Destination: %s ", results.Stations[i].Destination)
		fmt.Printf("DueIn : %s \n", results.Stations[i].DueIn)

	}
	response := "For " + results.Stations[0].StationFullName + " " + northBoundresponse + " minutes. " + southBoundresponse + " minutes. "
	fmt.Println("response " + response)
	return response

}
