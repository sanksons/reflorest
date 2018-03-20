# Reflorest
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/sanksons/reflorest/master/LICENSE)
[![Build](https://api.travis-ci.org/sanksons/reflorest.svg?branch=master)](https://travis-ci.org/sanksons/reflorest)




Reflorest is a reincarnation of a REST API Framework `florest`. To know more about florest 
[Click Here](https://github.com/jabong/florest-core)

### Why A reincarnation?

Despite of exposing a beautiful workflow based approach for writing REST API's, florest suffered from several limitations. This reincarnation takes the florest to next level by removing those limitations and making florest project
more user friendly. 

### Florest limitations fixed in Reflorest
- All Platform Support. Supports Windows, Mac and Linux.
- New generation Dependency Management using Vendor folder.
- No relative package referencing.
- Easy debugging with delve.
- QueryString support.

### Pending Items:
- Use jsonitor instead of encoding/json
- Modify Gzip handler so that it gzips only when the content size is higher than a minimum threshold.
- Rest Friendly URL's

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
$ reflorest bootstrap github.com/<user>/<your-application>
```
>This will create the default directory structure for your application.

**6.** Now execute:
```
$ govendor init
$ govendor fetch +o
```
> This will install all the required dependencies.

**7.** No need to modify any configuration file. Just for Info, files are kept at:
```
- conf/logger.json
- conf/config.json
```

**8.** Now just execute:
```
$ reflorest deploy
```
This will install the binary to $GOBIN. 

### Debugging

This illustrates how you can use delve debugger to debug incoming requests to your application server.

1. Install delve debugger:
```
$ go get -u github.com/derekparker/delve/cmd/dlv
```
2. Move inside your project folder and run:
```
$ dlv debug
```
3. Run following commands to setup important breakpoints. 
```
dlv> break reflorest/src/core/service.(*Webserver).ServiceHandler
dlv> break reflorest/src/core/common/orchestrator.run
dlv> break reflorest/src/core/common/orchestrator.execExecuteNode
dlv> break reflorest/src/core/common/orchestrator.execDecisionNode
```
4. Now run ```continue``` command, this will boot up the server.

5. Once the server is up and running simply hit your application url:
```
curl http://localhost:8080/restful/v1/hello
``` 



