package utils

import (
	"os"
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

func setHeaders() http.Header{
	headers := make(http.Header)
	headers.Add("Content-Type", "application/json")
	headers.Add("User-Agent", "Dalvik/2.1.0 (Linux; U; Android 9; Redmi Note 5 MIUI/9.6.27)")
	headers.Add("Host", "efishery.darwinbox.com")
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
		Headers(setHeaders()).
		Host(a.Hostname).
		BodyJSON(&reqCheckIn).ToJSON(&resCheckIn).
		Fetch(context.Background())
	if err != nil {
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
		Headers(setHeaders()).
		Host(a.Hostname).
		BodyJSON(&reqCheckOut).ToJSON(&resCheckOut).
		Fetch(context.Background())
	if err != nil {
		return resCheckOut, err
	}

	return resCheckOut, nil

}

func (a *AppConfig) GetCheckInID() (CheckInIDResponse, error){

	var resCheckID CheckInIDResponse
	t := Token{ Token: a.Token }

	err := requests.URL("/Mobileapi/LastCheckIndeatils").
		Headers(setHeaders()).
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
		Headers(setHeaders()).
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
	defaultAppConfig := AppConfig {
		Token: a.Token,
		Hostname: hostname,
		CheckIn: CheckInTemplate{
			LocationType: 2,
			Message: "",
			LatLong: "",
		},
		CheckOut: CheckOut{
			LocationType: 2,
			Message: "",
			LatLong: "",
		},
		Scheduler: Scheduler{
			CheckIn: "09:00:12",
			CheckOut: "17:30:50",
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

	return nil
}