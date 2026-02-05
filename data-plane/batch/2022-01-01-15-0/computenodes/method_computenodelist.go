package computenodes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeNodeListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ComputeNode
}

type ComputeNodeListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ComputeNode
}

type ComputeNodeListOperationOptions struct {
	ClientRequestId       *string
	Filter                *string
	Maxresults            *int64
	OcpDate               *string
	ReturnClientRequestId *bool
	Select                *string
	Timeout               *int64
}

func DefaultComputeNodeListOperationOptions() ComputeNodeListOperationOptions {
	return ComputeNodeListOperationOptions{}
}

func (o ComputeNodeListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.ClientRequestId != nil {
		out.Append("client-request-id", fmt.Sprintf("%v", *o.ClientRequestId))
	}
	if o.OcpDate != nil {
		out.Append("ocp-date", fmt.Sprintf("%v", *o.OcpDate))
	}
	if o.ReturnClientRequestId != nil {
		out.Append("return-client-request-id", fmt.Sprintf("%v", *o.ReturnClientRequestId))
	}
	return &out
}

func (o ComputeNodeListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ComputeNodeListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Maxresults != nil {
		out.Append("maxresults", fmt.Sprintf("%v", *o.Maxresults))
	}
	if o.Select != nil {
		out.Append("$select", fmt.Sprintf("%v", *o.Select))
	}
	if o.Timeout != nil {
		out.Append("timeout", fmt.Sprintf("%v", *o.Timeout))
	}
	return &out
}

type ComputeNodeListCustomPager struct {
	NextLink *odata.Link `json:"odata.nextLink"`
}

func (p *ComputeNodeListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ComputeNodeList ...
func (c ComputeNodesClient) ComputeNodeList(ctx context.Context, id PoolId, options ComputeNodeListOperationOptions) (result ComputeNodeListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ComputeNodeListCustomPager{},
		Path:          fmt.Sprintf("%s/nodes", id.Path()),
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
		Values *[]ComputeNode `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ComputeNodeListComplete retrieves all the results into a single object
func (c ComputeNodesClient) ComputeNodeListComplete(ctx context.Context, id PoolId, options ComputeNodeListOperationOptions) (ComputeNodeListCompleteResult, error) {
	return c.ComputeNodeListCompleteMatchingPredicate(ctx, id, options, ComputeNodeOperationPredicate{})
}

// ComputeNodeListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComputeNodesClient) ComputeNodeListCompleteMatchingPredicate(ctx context.Context, id PoolId, options ComputeNodeListOperationOptions, predicate ComputeNodeOperationPredicate) (result ComputeNodeListCompleteResult, err error) {
	items := make([]ComputeNode, 0)

	resp, err := c.ComputeNodeList(ctx, id, options)
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

	result = ComputeNodeListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
