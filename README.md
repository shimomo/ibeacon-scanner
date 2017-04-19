# iBeacon Scanner

[![MIT License](http://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](LICENSE)

## Overview
ibeacon-scanner is a program written in Go.

Scan iBeacon on Raspberry Pi.

## Installation
```
$ go get github.com/shimomo/ibeacon-scanner
```

## Usage
```
$ cd $GOPATH/src/github.com/shimomo/ibeacon-scanner/examples
$ GOARCH=arm GOARM=6 GOOS=linux go build discover.go
$ sudo ./discover
```

## License
ibeacon-scanner is open-sourced software licensed under the [MIT license](LICENSE).
