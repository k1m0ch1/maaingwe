package main

import (
	"os"
	"fmt"
	// "time"

	"github.com/k1m0ch1/maaingwe/utils"
	"github.com/jasonlvhit/gocron"
)

var darwin utils.AppConfig

func doCheckIn(){
	checkIn, err := darwin.DoCheckIn()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(checkIn.Message)
}

func doCheckOut(){

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

func checkCI(){
	// currentTime := time.Now() 

	fmt.Printf("[+] Running CheckIn at %s\n", darwin.Scheduler.CheckIn)
	fmt.Printf("[+] Check date for today ")
	doCheckIn()
}

func checkCO(){
	doCheckOut()
}

func main(){
	darwin, err := darwin.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	if len(os.Args) > 0 {
		switch os.Args[0]{
		case "login":
			QRCode := os.Args[2]
			Hostname := os.Args[1]
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
			doCheckIn()
		case "checkout":
			doCheckOut()
		case "scheduler":
			schedule := gocron.NewScheduler()

			schedule.Every(1).Day().At(darwin.Scheduler.CheckIn).Do(checkCI)
			schedule.Every(1).Day().At(darwin.Scheduler.CheckOut).Do(checkCO)

			fmt.Printf("\nScheduler is Running at %s and %s\n",darwin.Scheduler.CheckIn,darwin.Scheduler.CheckOut)
			fmt.Println("Ctrl+C to Cancel")
			<- schedule.Start()

		}
	}
}