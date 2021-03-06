package tests

import (
	"context"
	"fmt"
	"github.com/couchbaselabs/couchbase-cloud-go-client"
	"os"
	"testing"
)

func TestGetProject(t *testing.T) {
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

	projectId := "5efaa2fc-3917-4da7-8e84-5321c56bb40f"
	configuration := couchbasecapella.NewConfiguration()
	apiClient := couchbasecapella.NewAPIClient(configuration)
	project, _, err := apiClient.ProjectsApi.ProjectsShow(ctx, projectId).Execute()
	if err != nil {
		t.Fatalf("Error when calling `ProjectsApi.ProjectsShow(ctx)`: %v\n", err)
	}

	fmt.Println("Project name is " + project.Name)
	fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsShow(ctx)`: %v\n", project)
}
