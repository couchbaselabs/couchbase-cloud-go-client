package couchbasecloud

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Clusters []Cluster

type ClustersList struct {
	Cursor Cursor    `json:"cursor"`
	Data   []Cluster `json:"data"`
}

type ClusterStatus struct {
	Status string `json:"status"`
}

type Cluster struct {
	Id        string   `json:"id"`
	CloudId   string   `json:"cloudId"`
	Name      string   `json:"name"`
	Nodes     int      `json:"nodes"`
	ProjectId string   `json:"projectId"`
	Services  []string `json:"services"`
	TenantId  string   `json:"tenantId"`
}

type ListClustersOptions struct {
	Page      int     `json:"page"`
	PerPage   int     `json:"perPage"`
	SortBy    *string `json:"sortBy"`
	CloudId   *string `json:"cloudId"`
	ProjectId *string `json:"projectId"`
}

// ListClusters returns a page of clusters results. The page returned can be controlled with the 'options' argument.
func (client *CouchbaseCloudClient) ListClusters(options *ListClustersOptions) (*ClustersList, error) {
	var clusters *ClustersList
	err := client.do(&request{
		endpoint:        listClustersEndpoint.format(),
		method:          http.MethodGet,
		contentType:     contentTypeJSON,
		queryParameters: getClustersPagingOption(options),
	}, &clusters)

	return clusters, err
}

// ListClusterPages allows iterating over all the clusters. For every page of cluster items it will call the callback
// and pass the page worth of clouds as well as a boolean that indicates whether is is the last page or not. The
// function iterates over all the pages either until the callback returns false, the REST endpoint returns an error
// or it runs out of pages.
func (client *CouchbaseCloudClient) ListClusterPages(opts *ListClustersOptions, fn func(Clusters, bool) bool) error {
	var localOpts ListClustersOptions
	if opts != nil {
		localOpts = *opts
	}

	for {
		clouds, err := client.ListClusters(&localOpts)
		if err != nil {
			return err
		}

		if len(clouds.Data) == 0 {
			return nil
		}

		cont := fn(clouds.Data, clouds.Cursor.Pages.Last >= clouds.Cursor.Pages.Page)
		if !cont {
			return nil
		}

		localOpts.Page++
	}
}

// GetClusterStatus returns the status of the cluster identified by 'id'.
func (client *CouchbaseCloudClient) GetClusterStatus(id string) (*ClusterStatus, error) {
	if id == "" {
		return nil, fmt.Errorf("cluster id is required")
	}

	var status *ClusterStatus
	err := client.do(&request{
		endpoint:    clusterStatusEndpoint.format(id),
		method:      http.MethodGet,
		contentType: contentTypeJSON,
	}, &status)

	return status, err
}

// GetClusterHealth retrieves cluster health and status. Health stats of the cluster are returned only for clusters
// with ready status.
func (client *CouchbaseCloudClient) GetClusterHealth(id string) (*ClusterHealth, error) {
	if id == "" {
		return nil, fmt.Errorf("cluster id is required")
	}

	var health *ClusterHealth
	err := client.do(&request{
		endpoint:    clusterHealthEndpoint.format(id),
		method:      http.MethodGet,
		contentType: contentTypeJSON,
	}, &health)

	return health, err
}

func getClustersPagingOption(opts *ListClustersOptions) url.Values {
	params := url.Values{}

	if opts == nil {
		return params
	}

	if opts.SortBy != nil {
		params.Add("sortBy", *opts.SortBy)
	}

	if opts.Page != 0 {
		params.Add("page", strconv.Itoa(opts.Page))
	}

	if opts.PerPage != 0 {
		params.Add("perPage", strconv.Itoa(opts.PerPage))
	}

	return params
}
