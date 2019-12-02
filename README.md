# New Relic Api

## Description

Package for accessing the new relic API in golang.  
Examples in the example folder

Implemented functions:

* alert channels

## Usage

Create client:
```go
client := new_relic_api.New("your api key here")
```

## Development

Prepare environment: 

* Setup golang [golang.org](https://golang.org)
* Install dep [Instructions](https://github.com/golang/dep)
* Clone this repo into `$GOPATH/src/stash.massiveinteractive.com/to`
* Get dependencies `dep ensure`

Guidelines:

* Create a branch for development
* Add your feature
* Create a PR