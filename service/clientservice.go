package service

import (
	client "HubbleWeb/client"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UserInteraction() {
	for {
		fmt.Println("Enter the your options")
		fmt.Println("1-Default request")
		fmt.Println("2- i want enter my request body")
		fmt.Print("enter options -")
		var options string
		fmt.Scanf("%s", &options)
		if options == "1" {
			err := client.SentRequest("Life is a characteristic that distinguishes physical entities that have biological processes")
			if err != nil {
				fmt.Println("Http Request error -", err)
			}
		} else if options == "2" {

			fmt.Print("enter string =")
			in := bufio.NewReader(os.Stdin)
			line, err := in.ReadString('\n') //Read the console input
			if err != nil {
				fmt.Println("Input reading error -", err)
				continue
			}
			if line != "" {
				line = strings.ReplaceAll(line, "\n", "")
				err := client.SentRequest(line)
				if err != nil {
					fmt.Println("Http Request error -", err)
				}
			} else {
				fmt.Println("enter valid options")
			}

		} else {
			fmt.Println("Enter valid options")
			continue
		}
	}
}
