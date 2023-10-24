package utils

import (
	"os"
	"io"
	"context"
	"net/http"
	"io/ioutil"
	b64 "encoding/base64"

	requests "github.com/carlmjohnson/requests"
	yaml "github.com/goccy/go-yaml"
)

func (t *Token) SetToken(token string){
	t.Token = token
}

func (a *AppConfig) setHeaders() http.Header{
	headers := make(http.Header)
	headers.Add("Content-Type", "application/json")
	headers.Add("User-Agent", "Dalvik/2.1.0 (Linux; U; Android 9; Redmi Note 5 MIUI/9.6.27)")
	headers.Add("Host", a.Hostname)
	headers.Add("Connection", "Keep-Alive")

	return headers
}

func (a *AppConfig) Load() (*AppConfig, error){
	yamlFile, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, a)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *AppConfig) DoCheckIn() (CheckInResponse, error){

	var resCheckIn CheckInResponse

	latLong := b64.StdEncoding.EncodeToString([]byte(a.CheckIn.LatLong))

	reqCheckIn := CheckIn{
		Location: latLong,
		LatLong: a.CheckIn.LatLong,
		Message: a.CheckIn.Message,
		LocationType: a.CheckIn.LocationType,
		InOut: 1,
		UDID: "",
		Purpose: "",
		Token: a.Token,
	}

	err := requests.URL("/Mobileapi/CheckInPost").
		Headers(a.setHeaders()).
		Host(a.Hostname).
		BodyJSON(&reqCheckIn).ToJSON(&resCheckIn).
		CheckStatus(200).
		Fetch(context.Background())
	if err != nil {
		return resCheckIn, err
	}

	if requests.HasStatusErr(err, 401) {
		return resCheckIn, err
	}

	return resCheckIn, nil

}

func (a *AppConfig) DoCheckOut(checkInID string) (CheckInResponse, error){

	var resCheckOut CheckInResponse

	latLong := b64.StdEncoding.EncodeToString([]byte(a.CheckOut.LatLong))

	reqCheckOut := CheckIn{
		ID: checkInID,
		Location: latLong,
		LatLong: a.CheckOut.LatLong,
		Message: a.CheckOut.Message,
		LocationType: a.CheckOut.LocationType,
		InOut: 2,
		UDID: "",
		Purpose: "",
		Token: a.Token,
	}

	err := requests.URL("/Mobileapi/CheckInPost").
		Headers(a.setHeaders()).
		Host(a.Hostname).
		BodyJSON(&reqCheckOut).ToJSON(&resCheckOut).
		CheckStatus(200).
		Fetch(context.Background())
	if err != nil {
		return resCheckOut, err
	}

	if requests.HasStatusErr(err, 401) {
		return resCheckOut, err
	}

	return resCheckOut, nil

}

func (a *AppConfig) GetCheckInID() (CheckInIDResponse, error){

	var resCheckID CheckInIDResponse
	t := Token{ Token: a.Token }

	err := requests.URL("/Mobileapi/LastCheckIndeatils").
		Headers(a.setHeaders()).
		Host("efishery.darwinbox.com").
		BodyJSON(&t).ToJSON(&resCheckID).
		Fetch(context.Background())
	if err != nil {
		return resCheckID, err
	}

	return resCheckID, nil
}

func (a *AppConfig) SetTokenQR(qrcode string) (AuthResponse, error){
	
	var resAuth AuthResponse

	req := Auth {
		QRCode: qrcode,
		UDID: "",
	}

	err := requests.URL("/Mobileapi/auth").
		Headers(a.setHeaders()).
		Host("efishery.darwinbox.com").
		BodyJSON(&req).ToJSON(&resAuth).
		Fetch(context.Background())
	if err != nil {
		return resAuth, err
	}

	a.Token = resAuth.Token

	return resAuth, nil

}

func (a *AppConfig) GenerateConfig(hostname string) error {
	CheckInTime := a.Scheduler.CheckIn
	CheckOutTime := a.Scheduler.CheckOut
	if len(a.Scheduler.CheckIn) == 0{
		CheckInTime = "09:00:12"
	}
	if len(a.Scheduler.CheckOut) == 0{
		CheckOutTime = "17:30:50"
	}
	
	defaultAppConfig := AppConfig {
		Token: a.Token,
		Hostname: hostname,
		CheckIn: CheckInTemplate{
			LocationType: a.CheckIn.LocationType,
			Message: a.CheckIn.Message,
			LatLong: a.CheckIn.LatLong,
		},
		CheckOut: CheckOut{
			LocationType: a.CheckIn.LocationType,
			Message: a.CheckIn.Message,
			LatLong: a.CheckIn.LatLong,
		},
		Scheduler: Scheduler{
			CheckIn: CheckInTime,
			CheckOut: CheckOutTime,
			// CheckInRandom: []string{
			// 	"09:13", "09:00", "09:09", "09:27", "09:07",
			// },
			// CheckOutRandom: []string{
			// 	"17:30", "17:35", "17:32", "17:29", "17:37",
			// },
		},
	}

	bytes, err := yaml.Marshal(defaultAppConfig)
	if err != nil {
		return err
	}

	err = os.WriteFile("./config.yml", bytes, 0755)
	if err != nil {
		return err
	}

	fileUrl := "https://raw.githubusercontent.com/guangrei/APIHariLibur_V2/main/holidays.json"

	err = DownloadFile("holidays.json", fileUrl)
	if err != nil {
		return err
	}

	return nil
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}