package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	couchbasecloud "github.com/couchbaselabs/couchbase-cloud-go-client"
)

func TestStatus(t *testing.T) {
	ctx := context.WithValue(
		context.Background(),
		couchbasecloud.ContextAPIKeys,
		map[string]couchbasecloud.APIKey{
			"accessKey": {
				Key: os.Getenv("CBC_ACCESS_KEY"),
			},
			"secretKey": {
				Key: os.Getenv("CBC_SECRET_KEY"),
			},
		},
	)
	configuration := couchbasecloud.NewConfiguration()
	api_client := couchbasecloud.NewAPIClient(configuration)
	resp, r, err := api_client.StatusApi.StatusShow(ctx).Execute()
	if err != nil {
		t.Fatalf("Error when calling `StatusApi.StatusShow(ctx)`: %v\n", err)
		t.Fatalf("Full HTTP response: %v\n", r)

	}
	fmt.Fprintf(os.Stdout, "Response from `StatusApi.StatusShow(ctx)`: %v\n", resp)
}
