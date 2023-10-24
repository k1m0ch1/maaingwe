# MaAingWe - Darwin box Automated CheckIn/Checkout

MaAingWe from "Kumaha Aing We" (whatever I like), This tools is dedicated to automted prosess checkin and checkout with scheduler

> !!Disclaimer!! use at your own risk

## How to use

for the process from 1 to 5 you can see from this video
[![Login darwinbox with maaingwe](http://img.youtube.com/vi/xbSWfFoy8fw/0.jpg)](http://www.youtube.com/watch?v=xbSWfFoy8fw "maaingwe | login")

1. Login to use workspace darwinbox usually at <workspace>.darwinbox.com
2. click Profile icon at the right top, and click menu Mobile QR Code
3. now you only have 30 second to do this
4. get the string from the QR Code
5. run `maaingwe login workspace.darwinbox.com` and input the decoded QR Code (you can use this extension https://chrome.google.com/webstore/detail/qr-code-reader-for-google/gmloihcgbhbonllenincdakeijmikcne)
6. `config.yml` will be generated
7. edit the `config.yml` and set the schedule time, lat long, message and location type
8. run `maaingwe scheduler` to run the following schedule at `config.yml`
9. just wait and it will run

## Available command

- `maaingwe login <hostname>`
- `maaingwe scheduler`
- `maaingwe checkin`
- `maaingwe checkout`