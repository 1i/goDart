# GODART

Alexa skill to have DART times for the requested station.  
_Alexa open Next Dart for Pearse Station_

## build and deploy

Build the go file into an executable.  
`GOOS=linux GOARCH=amd64 go build -o alexa alexa.go`

Zip up the go executable.  
`zip alexa.zip alexa`

Update the lambda zip.  
`aws lambda update-function-code --function-name godart --zip fileb://./alexa.zip`

Invoke the lambda and log the results.  
`aws lambda invoke --cli-binary-format raw-in-base64-out --function-name godart --payload '{"request": {"locale": "en-US","requestId": "amzn1.echo-api.request.a1d80dd9-0538-4f62-ba27-decfd8ade0d8","shouldLinkResultBeReturned": false,"timestamp": "2022-02-12T16:30:15Z","type": "LaunchRequest"},"version": "1.0"}' --log-type Tail outfile | grep "LogResult"| awk -F'"' '{print $4}' | base64 --decode`

## TO DO
Allow saving of a default station if invoked without a station in the request.  
Save the Alexa account Id to DynamoDB and query if there is a default station set for that user.  