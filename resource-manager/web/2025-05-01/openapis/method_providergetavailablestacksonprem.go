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

type ProviderGetAvailableStacksOnPremOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApplicationStackResource
}

type ProviderGetAvailableStacksOnPremCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ApplicationStackResource
}

type ProviderGetAvailableStacksOnPremOperationOptions struct {
	OsTypeSelected *ProviderOsTypeSelected
}

func DefaultProviderGetAvailableStacksOnPremOperationOptions() ProviderGetAvailableStacksOnPremOperationOptions {
	return ProviderGetAvailableStacksOnPremOperationOptions{}
}

func (o ProviderGetAvailableStacksOnPremOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ProviderGetAvailableStacksOnPremOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ProviderGetAvailableStacksOnPremOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.OsTypeSelected != nil {
		out.Append("osTypeSelected", fmt.Sprintf("%v", *o.OsTypeSelected))
	}
	return &out
}

type ProviderGetAvailableStacksOnPremCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ProviderGetAvailableStacksOnPremCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ProviderGetAvailableStacksOnPrem ...
func (c OpenapisClient) ProviderGetAvailableStacksOnPrem(ctx context.Context, id commonids.SubscriptionId, options ProviderGetAvailableStacksOnPremOperationOptions) (result ProviderGetAvailableStacksOnPremOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ProviderGetAvailableStacksOnPremCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.Web/availableStacks", id.ID()),
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

// ProviderGetAvailableStacksOnPremComplete retrieves all the results into a single object
func (c OpenapisClient) ProviderGetAvailableStacksOnPremComplete(ctx context.Context, id commonids.SubscriptionId, options ProviderGetAvailableStacksOnPremOperationOptions) (ProviderGetAvailableStacksOnPremCompleteResult, error) {
	return c.ProviderGetAvailableStacksOnPremCompleteMatchingPredicate(ctx, id, options, ApplicationStackResourceOperationPredicate{})
}

// ProviderGetAvailableStacksOnPremCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) ProviderGetAvailableStacksOnPremCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options ProviderGetAvailableStacksOnPremOperationOptions, predicate ApplicationStackResourceOperationPredicate) (result ProviderGetAvailableStacksOnPremCompleteResult, err error) {
	items := make([]ApplicationStackResource, 0)

	resp, err := c.ProviderGetAvailableStacksOnPrem(ctx, id, options)
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

	result = ProviderGetAvailableStacksOnPremCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
