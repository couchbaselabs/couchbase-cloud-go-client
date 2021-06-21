package couchbasecloud

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListClusters(t *testing.T) {
	client := NewClient("a", "b")

	var (
		expectedURLParams url.Values
		body              []byte
		statusCode        int
	)

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/v2/clusters", r.URL.Path)
			require.Equal(t, expectedURLParams, r.URL.Query())

			w.WriteHeader(statusCode)
			if body != nil {
				_, _ = w.Write(body)
			}
		}))
	defer server.Close()

	client.BaseURL = server.URL

	type testCase struct {
		name              string
		statusCode        int
		outBody           []byte
		opts              *ListClustersOptions
		expectedURLParams url.Values
		expectedClusters  *ClustersList
		valid             bool
	}

	testCases := []testCase{
		{
			name:              "500",
			expectedURLParams: url.Values{},
			statusCode:        http.StatusInternalServerError,
		},
		{
			name:              "no-pagination-values",
			statusCode:        http.StatusOK,
			expectedURLParams: url.Values{},
			outBody:           []byte(`{"data":[{"id":"cluster-1"},{"id":"cluster-2"}]}`),
			valid:             true,
			expectedClusters: &ClustersList{
				Cursor: Cursor{},
				Data: []Cluster{
					{
						Id: "cluster-1",
					},
					{
						Id: "cluster-2",
					},
				},
			},
		},
		{
			name:       "with-pagination-values",
			statusCode: http.StatusOK,
			expectedURLParams: url.Values{
				"page":    []string{"10"},
				"perPage": []string{"5"},
				"sortBy":  []string{"id"},
			},
			opts: &ListClustersOptions{
				Page:    10,
				PerPage: 5,
				SortBy:  pointerString("id"),
			},
			outBody: []byte(`{"data":[],"cursor":{"pages":{"last":4,"page":10}}}`),
			valid:   true,
			expectedClusters: &ClustersList{
				Cursor: Cursor{Pages: struct {
					Last       int `json:"last"`
					Page       int `json:"page"`
					PerPage    int `json:"perPage"`
					TotalItems int `json:"totalItems"`
				}{
					Last: 4,
					Page: 10,
				},
				},
				Data: []Cluster{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body = tc.outBody
			statusCode = tc.statusCode
			expectedURLParams = tc.expectedURLParams

			clusters, err := client.ListClusters(tc.opts)
			if !tc.valid {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.expectedClusters, clusters)
		})
	}
}

func TestClusterStatus(t *testing.T) {
	client := NewClient("a", "b")

	var (
		clusterID  string
		status     string
		statusCode int
	)

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, fmt.Sprintf("/v2/clusters/%s/status", clusterID), r.URL.Path)

			w.WriteHeader(statusCode)
			if status != "" {
				_, _ = w.Write([]byte(fmt.Sprintf(`{"status":"%s"}`, status)))
			}
		}))
	defer server.Close()

	client.BaseURL = server.URL

	type testCase struct {
		name       string
		statusCode int
		status     string
		valid      bool
		id         string
	}

	testCases := []testCase{
		{
			name:       "500",
			statusCode: http.StatusInternalServerError,
			id:         "a",
		},
		{
			name:       "valid",
			id:         "a",
			statusCode: http.StatusOK,
			status:     "ready",
			valid:      true,
		},
		{
			name:       "no-id",
			statusCode: http.StatusOK,
			status:     "ready",
		},
		{
			name:       "id-needs-escaping",
			id:         "%%1%%",
			statusCode: http.StatusOK,
			status:     "upgrading",
			valid:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status = tc.status
			clusterID = tc.id
			statusCode = tc.statusCode

			clusterStatus, err := client.GetClusterStatus(tc.id)
			if !tc.valid {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.status, clusterStatus.Status)
		})
	}
}

func TestClusterHealth(t *testing.T) {
	client := NewClient("a", "b")

	var (
		clusterID  string
		body       []byte
		statusCode int
	)

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, fmt.Sprintf("/v2/clusters/%s/health", clusterID), r.URL.Path)

			w.WriteHeader(statusCode)
			if body != nil {
				_, _ = w.Write(body)
			}
		}))
	defer server.Close()

	client.BaseURL = server.URL

	type testCase struct {
		name       string
		statusCode int
		body       []byte
		expected   *ClusterHealth
		valid      bool
		id         string
	}

	testCases := []testCase{
		{
			name:       "500",
			statusCode: http.StatusInternalServerError,
			id:         "a",
		},
		{
			name:       "valid",
			id:         "a",
			statusCode: http.StatusOK,
			expected: &ClusterHealth{
				Status: "ready",
				Health: "unhealthy",
				BucketStats: &BucketStats{
					Status:         "ready",
					TotalCount:     1,
					UnhealthyCount: 1,
				},
				NodeStats: &NodeStats{
					Status:          "ready",
					TotalCount:      5,
					FailedOverCount: 1,
					UnhealthyCount:  1,
					ServiceStats: []ServiceStats{
						{
							Services: []string{"kv", "fts"},
							NodeName: "node1",
							Status:   "ready",
						},
					},
				},
			},
			body: []byte(`{"status":"ready","health":"unhealthy","bucketStats":{"status":"ready","totalCount":1,
"unhealthyCount":1},"nodeStats":{"status":"ready","totalCount":5,"failedOverCount":1,"unhealthyCount":1,
"serviceStats":[{"services":["kv","fts"],"nodeName":"node1","status":"ready"}]}}`),
			valid: true,
		},
		{
			name:       "no-id",
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			clusterID = tc.id
			body = tc.body
			statusCode = tc.statusCode

			clusterHealth, err := client.GetClusterHealth(tc.id)
			if !tc.valid {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.expected, clusterHealth)
		})
	}
}

func pointerString(s string) *string {
	return &s
}
