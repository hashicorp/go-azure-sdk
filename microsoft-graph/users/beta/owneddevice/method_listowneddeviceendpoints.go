package owneddevice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOwnedDeviceEndpointsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Endpoint
}

type ListOwnedDeviceEndpointsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Endpoint
}

type ListOwnedDeviceEndpointsOperationOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	Count            *bool
	Expand           *odata.Expand
	Filter           *string
	Metadata         *odata.Metadata
	OrderBy          *odata.OrderBy
	RetryFunc        client.RequestRetryFunc
	Search           *string
	Select           *[]string
	Skip             *int64
	Top              *int64
}

func DefaultListOwnedDeviceEndpointsOperationOptions() ListOwnedDeviceEndpointsOperationOptions {
	return ListOwnedDeviceEndpointsOperationOptions{}
}

func (o ListOwnedDeviceEndpointsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOwnedDeviceEndpointsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.ConsistencyLevel != nil {
		out.ConsistencyLevel = *o.ConsistencyLevel
	}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListOwnedDeviceEndpointsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOwnedDeviceEndpointsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOwnedDeviceEndpointsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOwnedDeviceEndpoints - Get the items of type microsoft.graph.endpoint in the microsoft.graph.directoryObject
// collection
func (c OwnedDeviceClient) ListOwnedDeviceEndpoints(ctx context.Context, id beta.UserId, options ListOwnedDeviceEndpointsOperationOptions) (result ListOwnedDeviceEndpointsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOwnedDeviceEndpointsCustomPager{},
		Path:          fmt.Sprintf("%s/ownedDevices/endpoint", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]beta.Endpoint `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOwnedDeviceEndpointsComplete retrieves all the results into a single object
func (c OwnedDeviceClient) ListOwnedDeviceEndpointsComplete(ctx context.Context, id beta.UserId, options ListOwnedDeviceEndpointsOperationOptions) (ListOwnedDeviceEndpointsCompleteResult, error) {
	return c.ListOwnedDeviceEndpointsCompleteMatchingPredicate(ctx, id, options, EndpointOperationPredicate{})
}

// ListOwnedDeviceEndpointsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OwnedDeviceClient) ListOwnedDeviceEndpointsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListOwnedDeviceEndpointsOperationOptions, predicate EndpointOperationPredicate) (result ListOwnedDeviceEndpointsCompleteResult, err error) {
	items := make([]beta.Endpoint, 0)

	resp, err := c.ListOwnedDeviceEndpoints(ctx, id, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListOwnedDeviceEndpointsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
