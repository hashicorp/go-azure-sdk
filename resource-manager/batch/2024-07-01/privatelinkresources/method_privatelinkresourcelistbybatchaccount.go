package privatelinkresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkResourceListByBatchAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateLinkResource
}

type PrivateLinkResourceListByBatchAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateLinkResource
}

type PrivateLinkResourceListByBatchAccountOperationOptions struct {
	Maxresults *int64
}

func DefaultPrivateLinkResourceListByBatchAccountOperationOptions() PrivateLinkResourceListByBatchAccountOperationOptions {
	return PrivateLinkResourceListByBatchAccountOperationOptions{}
}

func (o PrivateLinkResourceListByBatchAccountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o PrivateLinkResourceListByBatchAccountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o PrivateLinkResourceListByBatchAccountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Maxresults != nil {
		out.Append("maxresults", fmt.Sprintf("%v", *o.Maxresults))
	}
	return &out
}

type PrivateLinkResourceListByBatchAccountCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *PrivateLinkResourceListByBatchAccountCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// PrivateLinkResourceListByBatchAccount ...
func (c PrivateLinkResourcesClient) PrivateLinkResourceListByBatchAccount(ctx context.Context, id BatchAccountId, options PrivateLinkResourceListByBatchAccountOperationOptions) (result PrivateLinkResourceListByBatchAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &PrivateLinkResourceListByBatchAccountCustomPager{},
		Path:          fmt.Sprintf("%s/privateLinkResources", id.ID()),
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

// PrivateLinkResourceListByBatchAccountComplete retrieves all the results into a single object
func (c PrivateLinkResourcesClient) PrivateLinkResourceListByBatchAccountComplete(ctx context.Context, id BatchAccountId, options PrivateLinkResourceListByBatchAccountOperationOptions) (PrivateLinkResourceListByBatchAccountCompleteResult, error) {
	return c.PrivateLinkResourceListByBatchAccountCompleteMatchingPredicate(ctx, id, options, PrivateLinkResourceOperationPredicate{})
}

// PrivateLinkResourceListByBatchAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateLinkResourcesClient) PrivateLinkResourceListByBatchAccountCompleteMatchingPredicate(ctx context.Context, id BatchAccountId, options PrivateLinkResourceListByBatchAccountOperationOptions, predicate PrivateLinkResourceOperationPredicate) (result PrivateLinkResourceListByBatchAccountCompleteResult, err error) {
	items := make([]PrivateLinkResource, 0)

	resp, err := c.PrivateLinkResourceListByBatchAccount(ctx, id, options)
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

	result = PrivateLinkResourceListByBatchAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
