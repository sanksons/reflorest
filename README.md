# `florest`

Intallation Instructions:

First install the CLI using:
go get -u github.com/sanksons/reflorest/reflorest

Install Govendor Using:
github.com/kardianos/govendor



Make Project Directory under GOPATH, something like:
$GOPATH/github.com/<user>/<your-application>

Move to the folder you want your application to reside:
cd  $GOPATH/github.com/<user>/<your-application>

Once, inside the folder, run following command:
reflorest bootstrap
This will create the default directory structure for your application.

Now run:
govendor fetch +o
This will install all the required dependencies.

Modify the configuration files:
- config/logger/logger.json
- config/newapp/config.json

No just execute:
reflorest deploy
This will install the binary to GOBIN. 


