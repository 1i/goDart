# GODART

Alexa skill to have DART times for the requested station.

## build and deploy

build the go file into an executable.  
`GOOS=linux GOARCH=amd64 go build -o alexa alexa.go`

zip up the go executable.  
`zip alexa.zip alexa`

update the lambda zip.  
`aws lambda update-function-code --function-name godart --zip fileb://./alexa.zip`

invoke the lambda and log the results.  
`aws lambda invoke --cli-binary-format raw-in-base64-out --function-name godart --payload '{"request": {"locale": "en-US","requestId": "amzn1.echo-api.request.a1d80dd9-0538-4f62-ba27-decfd8ade0d8","shouldLinkResultBeReturned": false,"timestamp": "2022-02-12T16:30:15Z","type": "LaunchRequest"},"version": "1.0"}' --log-type Tail outfile | grep "LogResult"| awk -F'"' '{print $4}' | base64 --decode`
