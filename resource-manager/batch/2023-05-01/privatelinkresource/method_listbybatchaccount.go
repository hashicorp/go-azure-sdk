package privatelinkresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByBatchAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateLinkResource
}

type ListByBatchAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateLinkResource
}

type ListByBatchAccountOperationOptions struct {
	Maxresults *int64
}

func DefaultListByBatchAccountOperationOptions() ListByBatchAccountOperationOptions {
	return ListByBatchAccountOperationOptions{}
}

func (o ListByBatchAccountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByBatchAccountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByBatchAccountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Maxresults != nil {
		out.Append("maxresults", fmt.Sprintf("%v", *o.Maxresults))
	}
	return &out
}

// ListByBatchAccount ...
func (c PrivateLinkResourceClient) ListByBatchAccount(ctx context.Context, id BatchAccountId, options ListByBatchAccountOperationOptions) (result ListByBatchAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/privateLinkResources", id.ID()),
		OptionsObject: options,
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
		Values *[]PrivateLinkResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByBatchAccountComplete retrieves all the results into a single object
func (c PrivateLinkResourceClient) ListByBatchAccountComplete(ctx context.Context, id BatchAccountId, options ListByBatchAccountOperationOptions) (ListByBatchAccountCompleteResult, error) {
	return c.ListByBatchAccountCompleteMatchingPredicate(ctx, id, options, PrivateLinkResourceOperationPredicate{})
}

// ListByBatchAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateLinkResourceClient) ListByBatchAccountCompleteMatchingPredicate(ctx context.Context, id BatchAccountId, options ListByBatchAccountOperationOptions, predicate PrivateLinkResourceOperationPredicate) (result ListByBatchAccountCompleteResult, err error) {
	items := make([]PrivateLinkResource, 0)

	resp, err := c.ListByBatchAccount(ctx, id, options)
	if err != nil {
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

	result = ListByBatchAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
