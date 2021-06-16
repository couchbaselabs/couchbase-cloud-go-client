# couchbase-cloud-go-client
Work in progress Go wrapper for the Couchbase Cloud REST API

## Usage

`go get github.com/couchbaselabs/cloud-monitoring-tool/couchbasecloudclient`

### Import
```go
import (
  "github.com/couchbaselabs/cloud-monitoring-tool/couchbasecloudclient"
)
```

### Client

```go
client := couchbasecloudclient.NewClient(os.Getenv("CBC_ACCESS_KEY"), os.Getenv("CBC_SECRET_KEY"))
```

### List all clouds
Options can be specified with the `ListCloudOptions` type (some are optional)
```go
page := 1
lastPage := math.MaxInt16

for ok := true; ok; ok = page <= lastPage {
    listCloudsResponse, err := client.ListClouds(&couchbasecloudclient.ListCloudsOptions{Page: page, PerPage: 10})

    if err != nil {
        return nil, err
    }

    for _, cloud := range listCloudsResponse.Data {
        fmt.Printf(cloud.Name)
    }

    lastPage = listCloudsResponse.Cursor.Pages.Last
    page++
}
```

### List all clusters
Options can be specified with the `ListClustersOptions` type (some are optional)
```go
page := 1
lastPage := math.MaxInt16

for ok := true; ok; ok = page <= lastPage {
    listClustersResponse, err := client.ListClusters(&couchbasecloudclient.ListClustersOptions{Page: page, PerPage: 10})

    if err != nil {
        return nil, err
    }

    for _, cluster := range listClustersResponse.Data {
        fmt.Printf(cluster.Name)
    }

    lastPage = listClustersResponse.Cursor.Pages.Last
    page++
}
```

## TODO

- Build better (functional) pagination capabilities into the client
