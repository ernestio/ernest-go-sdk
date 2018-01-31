# ernest-go-sdk
Ernest SDK for golang


## Build status

* Master: [![CircleCI  Master](https://circleci.com/gh/ernestio/ernest-go-sdk/tree/master.svg?style=svg)](https://circleci.com/gh/ernestio/ernest-go-sdk/tree/master)

## Quick start

To install:

```sh
$ go get github.com/ernestio/ernest-go-sdk
```

To Test:

```sh
$ make dev-deps
$ make test
```

## Basic Usage

#### Authentication

To create a new client:

```go
import (
    "github.com/ernestio/ernest-go-sdk/config"
    "github.com/ernestio/ernest-go-sdk/client"
)

func main() {
    cfg := config.New("https://my-ernest-instance").WithCredentials("username", "password")
    ernest := client.New(cfg)

    err := ernest.Authenticate()
}
```

#### Services

To list all services available to you:

```go
services, err := ernest.Services.List()
if err != nil {
    // handle error
}

for _, service := range services {
    fmt.Println(service.Name)
}
```

To get details about a single service:

```go
service, err := ernest.Services.Get("some-service")
if err != nil {
    // handle error
}

fmt.Println(service.Name)
```

To get create a service:

```go
import (
    "github.com/ernestio/ernest-go-sdk/config"
    "github.com/ernestio/ernest-go-sdk/client"
    "github.com/ernestio/ernest-go-sdk/models"
)

func main() {
    cfg := config.New("https://my-ernest-instance").WithCredentials("username", "password")
    ernest := client.New(cfg)

    err := ernest.Authenticate()
    if err != nil {
        // handle error
    }

    service := &models.Service{
        DatacenterID: 1,
        Name: "test-service",
        Type: "aws",
    }

    err := ernest.Services.Create(service)
    if err != nil {
        // handle error
    }

    fmt.Println(service.ID)
}
```

#### Builds

To get create a service build:

```go
import (
    "github.com/ernestio/ernest-go-sdk/config"
    "github.com/ernestio/ernest-go-sdk/client"
    "github.com/ernestio/ernest-go-sdk/models"
)

func main() {
    cfg := config.New("https://my-ernest-instance").WithCredentials("username", "password")
    ernest := client.New(cfg)

    err := ernest.Authenticate()
    if err != nil {
        // handle error
    }

    yml := `
        ---
        name: some-service
        ...
    `

    build, err := ernest.Builds.Create("some-service", []byte(yml))
    if err != nil {
        // handle error
    }

    fmt.Println(build.ID)
}
```

To view events from a service build:

```go
import (
    "github.com/ernestio/ernest-go-sdk/config"
    "github.com/ernestio/ernest-go-sdk/client"
    "github.com/ernestio/ernest-go-sdk/models"
)

func main() {
    cfg := config.New("https://my-ernest-instance").WithCredentials("username", "password")
    ernest := client.New(cfg)

    err := ernest.Authenticate()
    if err != nil {
        // handle error
    }

    stream, err := ernest.Builds.Stream("build-id")
    if err != nil {
        // handle error
    }

    for {
        event, ok := <-stream
        if ok {
            fmt.Println(string(event.Data))
        }
    }
}
```


#### Custom HTTP client

To add additional parameters to the http client, such as disabling ssl verification for self signed certs, you can override the http client or update its options:

```go
func main() {
    cfg := config.New("https://my-ernest-instance").WithCredentials("username", "password")
    ernest := client.New(cfg)

    ernest.Conn.HTTPClient.Transport =  &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}
```

## Contributing

Please read through our
[contributing guidelines](CONTRIBUTING.md).
Included are directions for opening issues, coding standards, and notes on
development.

Moreover, if your pull request contains patches or features, you must include
relevant unit tests.


## Copyright and License

Code and documentation copyright since 2015 ernest.io authors.

Code released under
[the Mozilla Public License Version 2.0](LICENSE).
