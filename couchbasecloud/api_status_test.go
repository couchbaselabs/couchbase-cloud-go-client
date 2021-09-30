package openapi

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestStatus(t *testing.T) {
	ctx := context.WithValue(
		context.Background(),
		ContextAPIKeys,
		map[string]APIKey{
			"aKeyAuth": {
				Key: "1jSX7XQRMN5qu4YZaB4SN5N4J0l497sz",
			},
			"secretKeyAuth": {
				Key: "647Q982Dv589XpMhWAafBM0EGcMWvN04TviBBIDzsPOlCo8qz8pTtlPHCcRkIuVd",
			},
		},
	)
	configuration := NewConfiguration()
	api_client := NewAPIClient(configuration)
	resp, r, err := api_client.StatusApi.StatusShow(ctx).Execute()
	if err != nil {
		t.Fatalf("Error when calling `StatusApi.StatusShow(ctx)`: %v\n", err)
		t.Fatalf("Full HTTP response: %v\n", r)

	}
	fmt.Fprintf(os.Stdout, "Response from `StatusApi.StatusShow(ctx)`: %v\n", resp)
}
