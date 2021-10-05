package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	couchbasecloud "github.com/couchbaselabs/couchbase-cloud-go-client"
)

func TestGetProject(t *testing.T) {
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

	projectId := "5efaa2fc-3917-4da7-8e84-5321c56bb40f"
	configuration := couchbasecloud.NewConfiguration()
	api_client := couchbasecloud.NewAPIClient(configuration)
	fmt.Print("Hello")
	fmt.Print(os.Getenv("CBC_ACCESS_KEY"))
	fmt.Print(os.Getenv("CBC_SECRET_KEY"))
	project, r, err := api_client.ProjectsApi.ProjectsShow(ctx, projectId).Execute()
	if err != nil {
		t.Fatalf("Error when calling `ProjectsApi.ProjectsShow(ctx)`: %v\n", err)
		t.Fatalf("Full HTTP response: %v\n", r)
	}

	fmt.Println("Project name is " + project.Name)
	fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsShow(ctx)`: %v\n", project)

}
