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
`aws lambda invoke --function-name godart --log-type Tail outfile | grep "LogResult"| awk -F'"' '{print $4}' | base64 --decode`