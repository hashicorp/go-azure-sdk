package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VaultsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TrackedResource
}

type VaultsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []TrackedResource
}

type VaultsListOperationOptions struct {
	Filter *Filter
	Top    *int64
}

func DefaultVaultsListOperationOptions() VaultsListOperationOptions {
	return VaultsListOperationOptions{}
}

func (o VaultsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o VaultsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o VaultsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type VaultsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *VaultsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// VaultsList ...
func (c OpenapisClient) VaultsList(ctx context.Context, id commonids.SubscriptionId, options VaultsListOperationOptions) (result VaultsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &VaultsListCustomPager{},
		Path:          fmt.Sprintf("%s/resources", id.ID()),
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
		Values *[]TrackedResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// VaultsListComplete retrieves all the results into a single object
func (c OpenapisClient) VaultsListComplete(ctx context.Context, id commonids.SubscriptionId, options VaultsListOperationOptions) (VaultsListCompleteResult, error) {
	return c.VaultsListCompleteMatchingPredicate(ctx, id, options, TrackedResourceOperationPredicate{})
}

// VaultsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) VaultsListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options VaultsListOperationOptions, predicate TrackedResourceOperationPredicate) (result VaultsListCompleteResult, err error) {
	items := make([]TrackedResource, 0)

	resp, err := c.VaultsList(ctx, id, options)
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

	result = VaultsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
