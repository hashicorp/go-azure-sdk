package privateendpointconnections

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionListByBatchAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateEndpointConnection
}

type PrivateEndpointConnectionListByBatchAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateEndpointConnection
}

type PrivateEndpointConnectionListByBatchAccountOperationOptions struct {
	Maxresults *int64
}

func DefaultPrivateEndpointConnectionListByBatchAccountOperationOptions() PrivateEndpointConnectionListByBatchAccountOperationOptions {
	return PrivateEndpointConnectionListByBatchAccountOperationOptions{}
}

func (o PrivateEndpointConnectionListByBatchAccountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o PrivateEndpointConnectionListByBatchAccountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o PrivateEndpointConnectionListByBatchAccountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Maxresults != nil {
		out.Append("maxresults", fmt.Sprintf("%v", *o.Maxresults))
	}
	return &out
}

type PrivateEndpointConnectionListByBatchAccountCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *PrivateEndpointConnectionListByBatchAccountCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// PrivateEndpointConnectionListByBatchAccount ...
func (c PrivateEndpointConnectionsClient) PrivateEndpointConnectionListByBatchAccount(ctx context.Context, id BatchAccountId, options PrivateEndpointConnectionListByBatchAccountOperationOptions) (result PrivateEndpointConnectionListByBatchAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &PrivateEndpointConnectionListByBatchAccountCustomPager{},
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

// PrivateEndpointConnectionListByBatchAccountComplete retrieves all the results into a single object
func (c PrivateEndpointConnectionsClient) PrivateEndpointConnectionListByBatchAccountComplete(ctx context.Context, id BatchAccountId, options PrivateEndpointConnectionListByBatchAccountOperationOptions) (PrivateEndpointConnectionListByBatchAccountCompleteResult, error) {
	return c.PrivateEndpointConnectionListByBatchAccountCompleteMatchingPredicate(ctx, id, options, PrivateEndpointConnectionOperationPredicate{})
}

// PrivateEndpointConnectionListByBatchAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateEndpointConnectionsClient) PrivateEndpointConnectionListByBatchAccountCompleteMatchingPredicate(ctx context.Context, id BatchAccountId, options PrivateEndpointConnectionListByBatchAccountOperationOptions, predicate PrivateEndpointConnectionOperationPredicate) (result PrivateEndpointConnectionListByBatchAccountCompleteResult, err error) {
	items := make([]PrivateEndpointConnection, 0)

	resp, err := c.PrivateEndpointConnectionListByBatchAccount(ctx, id, options)
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

	result = PrivateEndpointConnectionListByBatchAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
