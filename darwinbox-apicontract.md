# API Contract Darwinbox to automated checkin/checkout

## Common Header

| Key | Value |
|-----|-----|
| Content-Type | application/json |
| User-Agent | Dalvik/2.1.0 (Linux; U; Android 9; Redmi Note 5 MIUI/9.6.27) |
| Host | workspace.darwinbox.com |
| Connection | Keep-Alive

## Login

`POST <workspace>/Mobileapi/auth`

JSON Body Request
`qrcode` string
`uuid` string

Response

```
	"token": string,
	"user_id": string,
	"tenant_id": string,
	"expires": timestamp,
	"is_manager": boolean,
	"status": int,
	"message": string,
	"user_details": {
		"name": string,
		"email": string,
		"user_id": string,
		"tenant_id": string,
		"mongo_id": string,
		"designation": string,
		"department": string,
		"business_unit": string,
		"mobile": string,
		"office": string,
		"office_address": string,
		"dob": string,
		"doj": string,
		"employee_no": string,
		"manager_name": string,
		"pic48": string,
		"pic320": string,
		"pic25": string
	}
```

## CheckIn

`POST <workspace>/Mobileapi/CheckInPost`

JSON Body Request
`location` string (base64 from latlng)
`message` string
`latlng` string (lat long, you can get from map google)
`location_type` int (1 for Office, 2 For Home, 3 [Default] for Field Duty )
`in_out` int (1 for checkin, 2 for checkout)
`purpose` string
`token` string (use current Token get from Login)

Response

```
{
	"status": int,
	"message": string
}
```

## Get CheckIn ID

`POST <workspace>/Mobileapi/LastCheckIndeatils` (I know its deatils not details, I'm not typo)

JSON Body Request
`token` string (use current Token get from Login)

Response

```
{
	"status": int,
	"message": {
		"id": string (uuid),
		"date": string (date YYYY-MM-DD),
		"last_action": int
	},
	"error": string
}
```

## CheckOut

`POST <workspace>/Mobileapi/CheckInPost`

JSON Body Request
`location` string (base64 from latlng)
`message` string
`latlng` string (lat long, you can get from map google)
`checkin_id` string (get from Get CheckIn ID)
`location_type` int (1 for Office, 2 For Home, 3 [Default] for Field Duty )
`in_out` int (1 for checkin, 2 for checkout)
`purpose` string
`token` string (use current Token get from Login)

Response

```
{
	"status": int,
	"message": string
}
```