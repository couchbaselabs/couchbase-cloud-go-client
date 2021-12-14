package tests

import (
	"context"
	"github.com/couchbaselabs/couchbase-cloud-go-client"
	"math"
	"os"
	"testing"
)

func TestGetClouds(t *testing.T) {
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

	var page int32 = 1
	var lastPage int32 = math.MaxInt16


	for ok := true; ok; ok = page <= lastPage {
		listCloudsResponse, _, err := apiClient.CloudsApi.CloudsList(ctx).Page(page).PerPage(10).Execute()

		if err != nil {
			t.Fatalf("Error when calling `CloudsApi.CloudsList(ctx)`: %v\n", err)
		}

		lastPage = *listCloudsResponse.Cursor.Pages.Last
		page++
	}

}
