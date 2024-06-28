package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAvailableStacksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApplicationStackResource
}

type GetAvailableStacksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ApplicationStackResource
}

type GetAvailableStacksOperationOptions struct {
	OsTypeSelected *ProviderOsTypeSelected
}

func DefaultGetAvailableStacksOperationOptions() GetAvailableStacksOperationOptions {
	return GetAvailableStacksOperationOptions{}
}

func (o GetAvailableStacksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAvailableStacksOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o GetAvailableStacksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.OsTypeSelected != nil {
		out.Append("osTypeSelected", fmt.Sprintf("%v", *o.OsTypeSelected))
	}
	return &out
}

type GetAvailableStacksCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GetAvailableStacksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetAvailableStacks ...
func (c ProviderClient) GetAvailableStacks(ctx context.Context, options GetAvailableStacksOperationOptions) (result GetAvailableStacksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Pager:         &GetAvailableStacksCustomPager{},
		Path:          "/providers/Microsoft.Web/availableStacks",
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
		Values *[]ApplicationStackResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetAvailableStacksComplete retrieves all the results into a single object
func (c ProviderClient) GetAvailableStacksComplete(ctx context.Context, options GetAvailableStacksOperationOptions) (GetAvailableStacksCompleteResult, error) {
	return c.GetAvailableStacksCompleteMatchingPredicate(ctx, options, ApplicationStackResourceOperationPredicate{})
}

// GetAvailableStacksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProviderClient) GetAvailableStacksCompleteMatchingPredicate(ctx context.Context, options GetAvailableStacksOperationOptions, predicate ApplicationStackResourceOperationPredicate) (result GetAvailableStacksCompleteResult, err error) {
	items := make([]ApplicationStackResource, 0)

	resp, err := c.GetAvailableStacks(ctx, options)
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

	result = GetAvailableStacksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
