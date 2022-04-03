package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SentRequest(request string) error {
	fmt.Println("-------------- send start -------")
	req, err := http.NewRequest("POST", "http://localhost:8080/v1/searchtopword", strings.NewReader(request))
	if err != nil {
		fmt.Println("Error --->", err)
		return err
	}
	req.Header.Add("Content-Type", "text/plain")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(" cilent Error --->", err)
		return err
	}
	//fmt.Println(resp.Status)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println()
	fmt.Println("response - ", string(body))
	fmt.Println()
	return nil
}
