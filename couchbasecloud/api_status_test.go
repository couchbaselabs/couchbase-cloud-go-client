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
			"accessKey": {
				Key: "",
			},
			"secretKey": {
				Key: "",
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

func TestProjects(t *testing.T) {
	ctx := context.WithValue(
		context.Background(),
		ContextAPIKeys,
		map[string]APIKey{
			"accessKey": {
				Key: "1jSX7XQRMN5qu4YZaB4SN5N4J0l497sz",
			},
			"secretKey": {
				Key: "647Q982Dv589XpMhWAafBM0EGcMWvN04TviBBIDzsPOlCo8qz8pTtlPHCcRkIuVd",
			},
		},
	)

	projectId := "5efaa2fc-3917-4da7-8e84-5321c56bb40f"
	configuration := NewConfiguration()
	api_client := NewAPIClient(configuration)
	project, r, err := api_client.ProjectsApi.ProjectsShow(ctx, projectId).Execute()
	if err != nil {
		t.Fatalf("Error when calling `ProjectsApi.ProjectsShow(ctx)`: %v\n", err)
		t.Fatalf("Full HTTP response: %v\n", r)
	}

	fmt.Println("Project name is " + project.Name)
	fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsShow(ctx)`: %v\n", project)
}