<h1 align="center">
Stiletto 
</h1>

<div align="center">
  Configurable lookup if a certain feature is matched.
  <br />
  <br />
</div>

## About

Stiletto is a fully customizable and extensible feature store
built for looking up certain queries in production. One my pains
has been allowing certain actions to happen for users currently
under investigation. I could have built a one of table to do lookups,
but a more generalized solution could come in handy for the rest of
my team and hopefully company.

Stiletto works of off the assumption that sometimes you want a fast lookup for simple things such as user ids, device
ids and such. Stiletto provides various different stores for different needs, some will allow refreshing of data on each
query, others cache data to allow maximum speed and highest throughput. These can be combined, such as an.

```go
package main

func main() {

	sClient := NewStilettoClient().
		SetFeatureStore(
			"userIDs",
			NewEagerFeatureStore,
			NewCacheableFeatureStore,
			NewRemoteFeatureStore("https://some-url-to-do-lookup/users/"))

	featureActive, err := sClient.GetFeature(
		context.Background(),
		"userIDs",
		"some-user-id")
	if featureActive {
		log.Debugf("user: %s has transferred: %d to %s", userID, amount, otherUserID)
		dumpStore.DumpIntermediaryState( ... )
	}
}
```

This should provide some rationale behind the design. Stiletto is meant to run on both the client (server) as well as on
the server. However, it isn't a requirement. Though running it on the server provides the best ease of use and
usability.

Stiletto on the server provides a robust toolkit for integrating with third-parties like Jira, Trello, git and so on, to
fill it's feature stores.

## Built With

- Go

## Getting Started

### Prerequisites

- Go >=1.18

### Usage

```bash
go get github.com/kjuulh/stiletto
```

#### Client

```go
package main

import (
	"context"
	"fmt"
	"github.com/kjuulh/stiletto/pkg/client"
	"github.com/kjuulh/stiletto/pkg/featurestores"
)

func main() {
	sClient := client.NewStilettoClient().
		SetFeatureStore(
			"userIDs",
			featurestores.NewInMemoryFeatureStore([]string{"some-user-id"}))

	featureActive, err := sClient.GetFeature(
		context.Background(),
		"userIDs",
		"some-user-id")
	if err != nil {
		panic(err)
	}

	if featureActive {
		fmt.Println("Feature is active, yay!")
	}
}
```

## Misc

- [Readme Template](https://github.com/dec0dOS/amazing-github-template/blob/main/README.md)
