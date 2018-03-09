# Reflorest

Reflorest is a reincarnation of a REST API Framework `florest`. To know more about florest 
[Click Here](https://github.com/jabong/florest-core)

### Why A reincarnation?

Despite of exposing a beautiful workflow based approach for writing REST API's, florest suffered from several limitations. This reincarnation takes the florest to next level by removing those limitations and making florest project
more user friendly. 

### Florest limitations fixed in Reflorest
- All Platform Support. Supports Windows, Mac and Linux.
- New generation Dependency Management using Vendor folder.
- No relative package referencing.

### Intallation Instructions:

**1.** Install the Reflorest CLI, using command:
```
$ go get -u github.com/sanksons/reflorest/reflorest
```
**2.** Install Govendor (dependency manager), using command
```
$ github.com/kardianos/govendor
```
**3.** Make Your Application Project Directory under GOPATH, something like:

 $GOPATH/github.com/```user```/```your-application```

**4.** Move inside the application folder created in previous step. like:
```
$ cd  $GOPATH/github.com/<user>/<your-application>
```
**5.** Once, inside the folder, run following command:
```
$ reflorest bootstrap github.com/```user```/```your-application```
```
>This will create the default directory structure for your application.

**6.** Now execute:
```
$ govendor fetch +o
```
> This will install all the required dependencies.

**7.** Modify the configuration files:
```
- config/logger/logger.json
- config/newapp/config.json
```

**8.** Now just execute:
```
$ reflorest deploy
```
This will install the binary to $GOBIN. 


