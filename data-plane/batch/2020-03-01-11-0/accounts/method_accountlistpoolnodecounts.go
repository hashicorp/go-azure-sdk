package accounts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountListPoolNodeCountsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PoolNodeCounts
}

type AccountListPoolNodeCountsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PoolNodeCounts
}

type AccountListPoolNodeCountsOperationOptions struct {
	ClientRequestId       *string
	Filter                *string
	Maxresults            *int64
	OcpDate               *string
	ReturnClientRequestId *bool
	Timeout               *int64
}

func DefaultAccountListPoolNodeCountsOperationOptions() AccountListPoolNodeCountsOperationOptions {
	return AccountListPoolNodeCountsOperationOptions{}
}

func (o AccountListPoolNodeCountsOperationOptions) ToHeaders() *client.Headers {
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

func (o AccountListPoolNodeCountsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o AccountListPoolNodeCountsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Maxresults != nil {
		out.Append("maxresults", fmt.Sprintf("%v", *o.Maxresults))
	}
	if o.Timeout != nil {
		out.Append("timeout", fmt.Sprintf("%v", *o.Timeout))
	}
	return &out
}

type AccountListPoolNodeCountsCustomPager struct {
	NextLink *odata.Link `json:"odata.nextLink"`
}

func (p *AccountListPoolNodeCountsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AccountListPoolNodeCounts ...
func (c AccountsClient) AccountListPoolNodeCounts(ctx context.Context, options AccountListPoolNodeCountsOperationOptions) (result AccountListPoolNodeCountsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &AccountListPoolNodeCountsCustomPager{},
		Path:          "/nodeCounts",
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
		Values *[]PoolNodeCounts `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AccountListPoolNodeCountsComplete retrieves all the results into a single object
func (c AccountsClient) AccountListPoolNodeCountsComplete(ctx context.Context, options AccountListPoolNodeCountsOperationOptions) (AccountListPoolNodeCountsCompleteResult, error) {
	return c.AccountListPoolNodeCountsCompleteMatchingPredicate(ctx, options, PoolNodeCountsOperationPredicate{})
}

// AccountListPoolNodeCountsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccountsClient) AccountListPoolNodeCountsCompleteMatchingPredicate(ctx context.Context, options AccountListPoolNodeCountsOperationOptions, predicate PoolNodeCountsOperationPredicate) (result AccountListPoolNodeCountsCompleteResult, err error) {
	items := make([]PoolNodeCounts, 0)

	resp, err := c.AccountListPoolNodeCounts(ctx, options)
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

	result = AccountListPoolNodeCountsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
