<p align="center">
<img src="static/assets/logo.png" width="310" height="110" border="0" alt="ipdata">
<br>
<a href="https://circleci.com/gh/s32x/ipdata/tree/master"><img src="https://circleci.com/gh/s32x/ipdata/tree/master.svg?style=svg" alt="CircleCI"></a>
<a href="https://goreportcard.com/report/s32x.com/ipdata"><img src="https://goreportcard.com/badge/s32x.com/ipdata" alt="Go Report Card"></a>
<a href="https://godoc.org/s32x.com/ipdata/ipdata"><img src="https://godoc.org/s32x.com/ipdata/ipdata?status.svg" alt="GoDoc"></a>
</p>

<p align="center">
<img src="static/assets/graphic.png" width="650px" height="418px" alt="ipdata curl">
</p>

`ipdata` is a free and open-source ip address lookup system. It combines the results retrieved from multiple maxminddb databases to provide you general geo/isp data for version 4 addresses. The project is currently available as a [publicly consumable API](#public-api-usage), an easily importable [Golang package called `ipdata`](#full-go-example) for use in your own Go projects, and a public [Docker image on DockerHub](https://hub.docker.com/r/s32x/ipdata/).

*Please note: The hosted version of ipdata.info does not store ANY ip data that is requested. If this is a real security concern of yours I recommend either using the self hosted Binary/Docker image or importing and utilizing the package yourself.*

## Features

The lookup system is extremely easy to understand and includes only one output format (JSON). What is returned to you is outlined in the below sample response. If you decide to use the package yourself the response struct can be found [here](https://github.com/s32x/ipdata/blob/master/ipdata/lookup.go#L10-L25).

```json
{
    "ip_address": "123.456.789.012",
    "hostname": "example.com",
    "isp": "ISP LLC Communications",
    "country_code": "US",
    "country_name": "United States",
    "region_code": "CA",
    "region_name": "California",
    "city": "San Francisco",
    "zip_code": "94016",
    "time_zone": "America/California",
    "country_code": "US",
    "country_name": "United States",
    "geohash": "9ydqy",
    "latitude": 37.751,
    "longitude": -97.822,
    "metro_code": 123,
}
```

## Getting Started

### Public API Usage
Using the API is very simple. All that's needed to lookup an IP is to send a `GET` request using the below URL schema to our origin.
```
https://ipdata.info/lookup/{ip}
```

### Installing
To start using IPData on your local system, install Go and run `go get`:
```
$ go get s32x.com/ipdata
```
This will install the ipdata binary on your machine.

### Running with Docker
To start using IPData via Docker, install Docker and run `docker run`:
```
$ docker run -p 8080:8080 s32x/ipdata
```
This will retrieve the remote DockerHub image and start the service on port 8080.

## Full Go Example

```go
package main

import (
    "log"

    "s32x.com/ipdata/ipdata"
)

func main() {
    // Create the ipdata client
    ic, err := ipdata.NewClient()
    if err != nil {
        log.Fatal(err)
    }
    defer ic.Close()

    // Perform the lookup
    log.Println(ic.Lookup("172.217.6.110"))
}
```

The BSD 3-clause License
========================

Copyright (c) 2019, Steven Wolfe. All rights reserved.

Redistribution and use in source and binary forms, with or without modification,
are permitted provided that the following conditions are met:

 - Redistributions of source code must retain the above copyright notice,
   this list of conditions and the following disclaimer.

 - Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

 - Neither the name of ipdata nor the names of its contributors may
   be used to endorse or promote products derived from this software without
   specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.