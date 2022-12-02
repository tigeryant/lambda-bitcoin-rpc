package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	testRequest := `{"jsonrpc":"1.0","id":"curltext","method":"getblockchaininfo","params":[]}`
	url := fmt.Sprintf("http://%s:%s@192.168.1.154:8332/", os.Getenv("RPCUSER"), os.Getenv("RPCPASSWORD"))
	req, _ := http.NewRequest("POST", url, strings.NewReader(testRequest))
	req.Header.Add("content-type", "text/plain;")

	res, e := http.DefaultClient.Do(req)
	output := ""
	if e != nil {
		output = fmt.Sprint(e)
		fmt.Println(e)
	} else {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		output = string(body)
		fmt.Println(string(body))
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       output,
	}

	return response, nil
}
