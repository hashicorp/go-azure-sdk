package privatelink

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionsListByHostPoolOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateEndpointConnection
}

type PrivateEndpointConnectionsListByHostPoolCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateEndpointConnection
}

type PrivateEndpointConnectionsListByHostPoolOperationOptions struct {
	InitialSkip  *int64
	IsDescending *bool
	PageSize     *int64
}

func DefaultPrivateEndpointConnectionsListByHostPoolOperationOptions() PrivateEndpointConnectionsListByHostPoolOperationOptions {
	return PrivateEndpointConnectionsListByHostPoolOperationOptions{}
}

func (o PrivateEndpointConnectionsListByHostPoolOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o PrivateEndpointConnectionsListByHostPoolOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o PrivateEndpointConnectionsListByHostPoolOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.InitialSkip != nil {
		out.Append("initialSkip", fmt.Sprintf("%v", *o.InitialSkip))
	}
	if o.IsDescending != nil {
		out.Append("isDescending", fmt.Sprintf("%v", *o.IsDescending))
	}
	if o.PageSize != nil {
		out.Append("pageSize", fmt.Sprintf("%v", *o.PageSize))
	}
	return &out
}

type PrivateEndpointConnectionsListByHostPoolCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *PrivateEndpointConnectionsListByHostPoolCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// PrivateEndpointConnectionsListByHostPool ...
func (c PrivateLinkClient) PrivateEndpointConnectionsListByHostPool(ctx context.Context, id HostPoolId, options PrivateEndpointConnectionsListByHostPoolOperationOptions) (result PrivateEndpointConnectionsListByHostPoolOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &PrivateEndpointConnectionsListByHostPoolCustomPager{},
		Path:          fmt.Sprintf("%s/privateEndpointConnections", id.ID()),
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
		Values *[]PrivateEndpointConnection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PrivateEndpointConnectionsListByHostPoolComplete retrieves all the results into a single object
func (c PrivateLinkClient) PrivateEndpointConnectionsListByHostPoolComplete(ctx context.Context, id HostPoolId, options PrivateEndpointConnectionsListByHostPoolOperationOptions) (PrivateEndpointConnectionsListByHostPoolCompleteResult, error) {
	return c.PrivateEndpointConnectionsListByHostPoolCompleteMatchingPredicate(ctx, id, options, PrivateEndpointConnectionOperationPredicate{})
}

// PrivateEndpointConnectionsListByHostPoolCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateLinkClient) PrivateEndpointConnectionsListByHostPoolCompleteMatchingPredicate(ctx context.Context, id HostPoolId, options PrivateEndpointConnectionsListByHostPoolOperationOptions, predicate PrivateEndpointConnectionOperationPredicate) (result PrivateEndpointConnectionsListByHostPoolCompleteResult, err error) {
	items := make([]PrivateEndpointConnection, 0)

	resp, err := c.PrivateEndpointConnectionsListByHostPool(ctx, id, options)
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

	result = PrivateEndpointConnectionsListByHostPoolCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
