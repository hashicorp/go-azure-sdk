package provider

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

type GetAvailableStacksOnPremOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApplicationStackResource
}

type GetAvailableStacksOnPremCompleteResult struct {
	Items []ApplicationStackResource
}

type GetAvailableStacksOnPremOperationOptions struct {
	OsTypeSelected *ProviderOsTypeSelected
}

func DefaultGetAvailableStacksOnPremOperationOptions() GetAvailableStacksOnPremOperationOptions {
	return GetAvailableStacksOnPremOperationOptions{}
}

func (o GetAvailableStacksOnPremOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAvailableStacksOnPremOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o GetAvailableStacksOnPremOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.OsTypeSelected != nil {
		out.Append("osTypeSelected", fmt.Sprintf("%v", *o.OsTypeSelected))
	}
	return &out
}

// GetAvailableStacksOnPrem ...
func (c ProviderClient) GetAvailableStacksOnPrem(ctx context.Context, id commonids.SubscriptionId, options GetAvailableStacksOnPremOperationOptions) (result GetAvailableStacksOnPremOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.Web/availableStacks", id.ID()),
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

// GetAvailableStacksOnPremComplete retrieves all the results into a single object
func (c ProviderClient) GetAvailableStacksOnPremComplete(ctx context.Context, id commonids.SubscriptionId, options GetAvailableStacksOnPremOperationOptions) (GetAvailableStacksOnPremCompleteResult, error) {
	return c.GetAvailableStacksOnPremCompleteMatchingPredicate(ctx, id, options, ApplicationStackResourceOperationPredicate{})
}

// GetAvailableStacksOnPremCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProviderClient) GetAvailableStacksOnPremCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options GetAvailableStacksOnPremOperationOptions, predicate ApplicationStackResourceOperationPredicate) (result GetAvailableStacksOnPremCompleteResult, err error) {
	items := make([]ApplicationStackResource, 0)

	resp, err := c.GetAvailableStacksOnPrem(ctx, id, options)
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

	result = GetAvailableStacksOnPremCompleteResult{
		Items: items,
	}
	return
}
