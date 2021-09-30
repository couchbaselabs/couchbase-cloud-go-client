package openapi

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestGetProject(t *testing.T) {
	ctx := context.WithValue(
		context.Background(),
		ContextAPIKeys,
		map[string]APIKey{
			"publicKeyAuth": {
				Key: "1jSX7XQRMN5qu4YZaB4SN5N4J0l497sz",
			},
			"secretKeyAuth": {
				Key: "647Q982Dv589XpMhWAafBM0EGcMWvN04TviBBIDzsPOlCo8qz8pTtlPHCcRkIuVd",
			},
		},
	)
	configuration := NewConfiguration()
	api_client := NewAPIClient(configuration)
	resp, r, err := api_client.ProjectsApi.ProjectsShow(ctx, "723956e2-7043-4fdc-80eb-16881f1a4127").Execute()
	if err != nil {
		t.Fatalf("Error when calling `ProjectsApi.ProjectsShow()`: %v\n", err)
		t.Fatalf("Full HTTP response: %v\n", r)

	}
	fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsShow()`: %v\n", resp)
}
