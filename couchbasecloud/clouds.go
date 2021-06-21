package couchbasecloud

import (
	"net/http"
	"net/url"
	"strconv"
)

type Clouds []Cloud

type CloudsList struct {
	Cursor Cursor  `json:"cursor"`
	Data   []Cloud `json:"data"`
}

type Cloud struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	Provider           string `json:"provider"`
	Region             string `json:"region"`
	Status             string `json:"status"`
	VirtualNetworkCIDR string `json:"virtualNetworkCidr"`
	VirtualNetworkID   string `json:"virtualNetworkId"`
}

type ListCloudsOptions struct {
	Page    int     `json:"page"`
	PerPage int     `json:"perPage"`
	SortBy  *string `json:"sortBy"`
}

// ListClouds returns a page of clouds results. The page returned can be controlled via the options argument.
func (client *CouchbaseCloudClient) ListClouds(options *ListCloudsOptions) (*CloudsList, error) {
	var clouds *CloudsList
	err := client.do(&request{
		endpoint:        listCloudsEndpoint.format(),
		method:          http.MethodGet,
		contentType:     contentTypeJSON,
		queryParameters: getCloudsPagingOption(options),
	}, &clouds)

	return clouds, err
}

// ListCloudPages allows iterating over all the clouds. For every page of cloud items it will call the callback and pass
// the page worth of clouds as well as a boolean that indicates whether is is the last page or not.
// The function iterates over all the pages either until the callback returns false, the REST endpoint returns an error
// or it runs out of pages.
func (client *CouchbaseCloudClient) ListCloudPages(options *ListCloudsOptions, fn func(Clouds, bool) bool) error {
	var localOpts ListCloudsOptions
	if options != nil {
		localOpts = *options
	}

	for {
		clouds, err := client.ListClouds(&localOpts)
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

func getCloudsPagingOption(opts *ListCloudsOptions) url.Values {
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
