package main

import (
	"os"
	"fmt"
	"time"
	"bufio"
	"strings"
	"io/ioutil"
	"encoding/json"

	"github.com/k1m0ch1/maaingwe/utils"
	"github.com/jasonlvhit/gocron"
)

var darwin utils.AppConfig

func doCheckIn() utils.CheckInResponse{
	checkIn, err := darwin.DoCheckIn()
	if err != nil {
		fmt.Println(" |  ├─[-] Error when trying to checkin with Message")
		fmt.Println(" |  └─[-]", err)
		fmt.Printf(" └─[-] Exit the Program\n")
		os.Exit(1)
	}

	return checkIn
}

func doCheckOut() utils.CheckInResponse{

	getCheckIn, err := darwin.GetCheckInID()
	if err != nil {
		fmt.Println(" |  ├─[-] Error when trying to get the CheckIn ID")
		fmt.Println(" |  └─[-]", err)
		fmt.Printf(" └─[-] Exit the Program\n")
		os.Exit(1)
	}

	checkOut, err := darwin.DoCheckOut(getCheckIn.Message.ID)
	if err != nil {
		fmt.Println(" |  ├─[-] Error when trying to checkin with Message")
		fmt.Println(" |  └─[-]", err)
		fmt.Printf(" └─[-] Exit the Program\n")
		os.Exit(1)
	}

	return checkOut
}

type Summary struct {
	Summary string
}

func checking(status string){
	var proceed bool
	proceed = true

	currentTime := time.Now() 
	date := currentTime.Format("2006-01-02")

	fmt.Printf(" ├─[+] Running %s at %s\n", status, darwin.Scheduler.CheckIn)
	fmt.Printf(" |  ├─[+] Check date for today %s\n", date)

	var liburan map[string]Summary
	data, err := ioutil.ReadFile("./holidays.json")
	if err != nil {
		fmt.Printf(" |  └─[-] %v", err)
		os.Exit(1)
	}

	err = json.Unmarshal(data, &liburan)
	if err != nil {
		fmt.Printf(" |  └─[-] %v", err)
		os.Exit(1)
	}

	if currentTime.Weekday() == time.Saturday || currentTime.Weekday() == time.Sunday {
		fmt.Printf(" |  ├─[-] Today (%s) is %s ", date,currentTime.Weekday())
		fmt.Printf(" |  ├─[-] use the argument `%s` if you want to %s today\n",strings.ToLower(status), strings.ToLower(status))
		fmt.Printf(" |  └─[-] Exit the Program\n")
		proceed = false
	}

	for k, v := range liburan {
		if date == k {
			fmt.Printf(" |  ├─[-] Today (%s) is holidays %s \n", k,v.Summary)
			fmt.Printf(" |  ├─[-] Skip the %s process\n", status)
			fmt.Printf(" |  └─[-] use the argument `checkin` if you want to %s today\n", strings.ToLower(status))
			fmt.Printf(" └─[-] Exit the Program\n")
			proceed = false
			break
		}
	}
	
	if proceed{
		fmt.Printf(" |  ├─[+] Today is not holidays or weekend \n")
		fmt.Printf(" |  ├─[+] Proceed %s from config.yml \n", status)
		if status == "CheckIn"{
			resp := doCheckIn()
			fmt.Printf(" |  └─[?] %s\n", resp.Message)
			fmt.Printf(" └─[+] %s Done %s\n", status)
		}else if status == "CheckOut"{
			resp := doCheckOut()
			fmt.Printf(" |  └─[?] %s\n", resp.Message)
			fmt.Printf(" └─[+] %s Done %s\n", status)
		}else{
			fmt.Printf(" |  └─[?] Dude what are you doing here\n")
		}
	}
}

func main(){
	darwin, err := darwin.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	if len(os.Args) > 1 {
		switch os.Args[1]{
		case "login":
			if len(os.Args) < 3 {
				fmt.Printf(" ├─[-] Please include the argument\n")
				fmt.Printf(" └─[-] Exit the Program\n\n")
				asDefault()
				os.Exit(1)
			}
			fmt.Println("[+] Run Login")
			fmt.Println(" ├─[-] Type the QR Code String :")
			fmt.Printf(" ├─[input] ")
			reader := bufio.NewReader(os.Stdin)
			line, err := reader.ReadString('\n')
			if err != nil{
				fmt.Printf(" ├─[-] Error when trying to read string input\n")
				fmt.Printf(" |  └─[?] %s\n", err)
				fmt.Printf(" └─[-] Exit the Program\n")
				fmt.Println(err)
			}
			QRCode := line
			Hostname := os.Args[2]
			resAuth, err := darwin.SetTokenQR(QRCode)
			if err != nil{
				fmt.Printf(" ├─[-] Can't Login with message\n")
				fmt.Printf(" |  └─[?] %s\n", resAuth.Message)
				fmt.Printf(" └─[-] Exit the Program\n")
				fmt.Println(err)
			}

			if resAuth.ErrorCode == 1 {
				fmt.Printf(" ├─[-] Can't Login with message\n")
				fmt.Printf(" |  └─[?] %s\n", resAuth.Message)
				fmt.Printf(" └─[-] Exit the Program\n")
				os.Exit(1)
			}else{
				err = darwin.GenerateConfig(Hostname)
				if err != nil {
					fmt.Printf(" ├─[-] Can't Login with message\n")
					fmt.Printf(" |  └─[?] %s\n", resAuth.Message)
					fmt.Printf(" └─[-] Exit the Program\n")
					os.Exit(1)
				}

				usersProfile := resAuth.UserDetails
				fmt.Printf(" ├─[+] Hi! %s (%s) %s\n", usersProfile.Name, usersProfile.Email, usersProfile.EmployeeNo)
				fmt.Printf(" |  └─[+] Role. %s at %s\n", usersProfile.Designation, usersProfile.Department)
				fmt.Printf(" ├─[+] File config.yml is generated\n")
				fmt.Printf(" ├─[+] Now change the configuration to config.yml\n")
				fmt.Printf(" └─[+] Enjoy\n")
				os.Exit(1)
			}
			fmt.Printf(" ├─[-] Can't Process, please check the argument\n")
			fmt.Printf(" └─[-] Exit the Program\n")
			os.Exit(1)
		case "checkin":
			fmt.Println("[+] Run Checkin")
			fmt.Println(" ├─[+] Proceed CheckIn from config.yml")
			resp := doCheckIn()
			fmt.Printf(" |  └─[?] %s\n", resp.Message)
			fmt.Printf(" └─[+] CheckIn Done %s\n")
		case "checkout":
			fmt.Println("[+] Run Checkout")
			fmt.Println(" ├─[+] Proceed CheckOut from config.yml")
			resp := doCheckOut()
			fmt.Printf(" |  └─[?] %s\n", resp.Message)
			fmt.Printf(" └─[+] CheckOut Done %s\n")
		case "scheduler":
			schedule := gocron.NewScheduler()

			schedule.Every(1).Day().At(darwin.Scheduler.CheckIn).Do(checking, "CheckIn")
			schedule.Every(1).Day().At(darwin.Scheduler.CheckOut).Do(checking, "CheckOut")

			fmt.Printf("\nScheduler is Running at %s and %s\n",darwin.Scheduler.CheckIn,darwin.Scheduler.CheckOut)
			fmt.Println("Ctrl+C to Cancel")
			<- schedule.Start()
		default:
			asDefault()
		}
	}else{
		asDefault()
	}
}

func asDefault(){
	fmt.Println("MaAingWe - The Darwinbox Automated CheckIn/CheckOut\n")
	fmt.Println("available command:")
	fmt.Println("`login <Hostname>` to generate token and set to config.yml")
	fmt.Println("`checkin` to checkin following with message and location at config.yml")
	fmt.Println("`checkout` to checkout following with message and location at config.yml")
	fmt.Println("`scheduler` to schedule following with schedule rule at config.yml")
}