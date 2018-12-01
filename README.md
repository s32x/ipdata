# IPData

[![CircleCI](https://circleci.com/gh/starboy/ipdata.svg?style=svg)](https://circleci.com/gh/starboy/ipdata)
[![GoDoc](https://godoc.org/github.com/starboy/ipdata/ipdata?status.svg)](https://godoc.org/github.com/starboy/ipdata/verifier)

ipdata is a free and open-source ip address lookup system. It combines the results retrieved from multiple maxminddb databases to provide you general geo/isp data for version 4 addresses. The project is available in three forms, the Golang library `ipdata` which can easily be used in your own Go projects, a public API endpoint (more info: https://ipdata.info), and a public Docker image on DockerHub (see: https://hub.docker.com/r/starboy/ipdata/).

## Using the API (public or self-hosted)

Using the API is very simple. All that's needed to lookup an IP is to send a `GET` request using the below URL schema to our origin.
```
https://ipdata.info/lookup/{ip}
```

## Using the library

```go
package main

import (
    "log"

    "s32x.com/ipdata"
)

func main() {
    // Create the ipdata client
    ic, err := ipdata.New()
    if err != nil {
        log.Fatal(err)
    }
    defer ic.Close()

    // Perform the lookup
    log.Println(ic.Lookup("172.217.6.110"))
}
```

## Running with Go

```
go get -u s32x.com/ipdata
ipdata
```

## Running with Docker

```
docker run -p 8080:8080 sdwolfe32/ipdata
```

The BSD 3-clause License
========================

Copyright (c) 2018, Steven Wolfe. All rights reserved.

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