package main

import (
	"os"
	"fmt"

	"github.com/k1m0ch1/maaingwe/utils"
)

func main(){

	var darwin utils.AppConfig
	
	if len(os.Args) > 0 {
		switch os.Args[1]{
		case "login":
			QRCode := os.Args[3]
			Hostname := os.Args[2]
			resAuth, err := darwin.SetTokenQR(QRCode)
			if err != nil{
				fmt.Println(err)
			}

			if resAuth.ErrorCode == 1 {
				fmt.Printf("Can't Login with message `%s`", resAuth.Message)
				os.Exit(1)
			}else{
				err = darwin.GenerateConfig(Hostname)
				if err != nil {
					fmt.Println(err)
				}
			}
		case "checkin":
			darwin, err := darwin.Load()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			checkIn, err := darwin.DoCheckIn()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(checkIn.Message)
		case "checkout":
			darwin, err := darwin.Load()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			getCheckIn, err := darwin.GetCheckInID()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			checkIn, err := darwin.DoCheckOut(getCheckIn.Message.ID)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(checkIn.Message)
		}
		
	}
}