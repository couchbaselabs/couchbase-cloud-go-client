package couchbasecloud

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListClouds(t *testing.T) {
	client := NewClient("a", "b")

	var (
		expectedURLParams url.Values
		body              []byte
		statusCode        int
	)

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/v2/clouds", r.URL.Path)
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
		opts              *ListCloudsOptions
		expectedURLParams url.Values
		expectedClusters  *CloudsList
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
			outBody:           []byte(`{"data":[{"id":"1"},{"id":"2"}]}`),
			valid:             true,
			expectedClusters: &CloudsList{
				Cursor: Cursor{},
				Data: []Cloud{
					{Id: "1"}, {Id: "2"},
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
			opts: &ListCloudsOptions{
				Page:    10,
				PerPage: 5,
				SortBy:  pointerString("id"),
			},
			outBody: []byte(`{"data":[],"cursor":{"pages":{"last":4,"page":10}}}`),
			valid:   true,
			expectedClusters: &CloudsList{
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
				Data: []Cloud{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body = tc.outBody
			statusCode = tc.statusCode
			expectedURLParams = tc.expectedURLParams

			clusters, err := client.ListClouds(tc.opts)
			if !tc.valid {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.expectedClusters, clusters)
		})
	}
}
