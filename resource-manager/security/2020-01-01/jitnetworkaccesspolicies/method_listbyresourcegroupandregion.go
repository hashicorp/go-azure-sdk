package jitnetworkaccesspolicies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByResourceGroupAndRegionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]JitNetworkAccessPolicy
}

type ListByResourceGroupAndRegionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []JitNetworkAccessPolicy
}

// ListByResourceGroupAndRegion ...
func (c JitNetworkAccessPoliciesClient) ListByResourceGroupAndRegion(ctx context.Context, id ProviderLocationId) (result ListByResourceGroupAndRegionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/jitNetworkAccessPolicies", id.ID()),
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
		Values *[]JitNetworkAccessPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByResourceGroupAndRegionComplete retrieves all the results into a single object
func (c JitNetworkAccessPoliciesClient) ListByResourceGroupAndRegionComplete(ctx context.Context, id ProviderLocationId) (ListByResourceGroupAndRegionCompleteResult, error) {
	return c.ListByResourceGroupAndRegionCompleteMatchingPredicate(ctx, id, JitNetworkAccessPolicyOperationPredicate{})
}

// ListByResourceGroupAndRegionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JitNetworkAccessPoliciesClient) ListByResourceGroupAndRegionCompleteMatchingPredicate(ctx context.Context, id ProviderLocationId, predicate JitNetworkAccessPolicyOperationPredicate) (result ListByResourceGroupAndRegionCompleteResult, err error) {
	items := make([]JitNetworkAccessPolicy, 0)

	resp, err := c.ListByResourceGroupAndRegion(ctx, id)
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

	result = ListByResourceGroupAndRegionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
