package client

import (
	"bufio"
	"fmt"
	"golculator/internal/numerical"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type ApiClient struct {
	Port string
}

func (ac ApiClient) RunClient() {
	// get numbers and operator from user for calculation
	// send request to server
	// get result from server
	// print result
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter numbers and operator separated by space")
	fmt.Println("Example: 1 + 2")
	for scanner.Scan() {
		var input string
		var num1, num2, op string
		input = scanner.Text()
		if input == "exit" {
			break
		}

		// split input by space and store into array of strings
		inputArr := strings.Fields(input)
		num1 = inputArr[0]
		op = inputArr[1]
		num2 = inputArr[2]

		params := url.Values{}
		params.Add("number1", num1)
		params.Add("number2", num2)
		params.Add("operator", numerical.OperatorToString(op))
		body := strings.NewReader(params.Encode())

		req, err := http.NewRequest("POST", "http://localhost:"+ac.Port+"/", body)
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Print(err)
		}
		// read body
		resBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Print(err)
		}
		fmt.Println(string(resBody))
		defer resp.Body.Close()
	}

}
