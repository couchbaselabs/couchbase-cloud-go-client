package tests

import (
	"context"
	"fmt"
	"github.com/couchbaselabs/couchbase-cloud-go-client"
	"os"
	"testing"
)

func TestStatus(t *testing.T) {
	ctx := context.WithValue(
		context.Background(),
		couchbasecapella.ContextAPIKeys,
		map[string]couchbasecapella.APIKey{
			"accessKey": {
				Key: os.Getenv("CBC_ACCESS_KEY"),
			},
			"secretKey": {
				Key: os.Getenv("CBC_SECRET_KEY"),
			},
		},
	)
	configuration := couchbasecapella.NewConfiguration()
	apiClient := couchbasecapella.NewAPIClient(configuration)
	resp, _, err := apiClient.StatusApi.StatusShow(ctx).Execute()
	if err != nil {
		t.Fatalf("Error when calling `StatusApi.StatusShow(ctx)`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `StatusApi.StatusShow(ctx)`: %v\n", resp)
}
