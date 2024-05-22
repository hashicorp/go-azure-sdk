package adaptivenetworkhardenings

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

type ListByExtendedResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AdaptiveNetworkHardening
}

type ListByExtendedResourceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AdaptiveNetworkHardening
}

// ListByExtendedResource ...
func (c AdaptiveNetworkHardeningsClient) ListByExtendedResource(ctx context.Context, id commonids.ScopeId) (result ListByExtendedResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.Security/adaptiveNetworkHardenings", id.ID()),
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
		Values *[]AdaptiveNetworkHardening `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByExtendedResourceComplete retrieves all the results into a single object
func (c AdaptiveNetworkHardeningsClient) ListByExtendedResourceComplete(ctx context.Context, id commonids.ScopeId) (ListByExtendedResourceCompleteResult, error) {
	return c.ListByExtendedResourceCompleteMatchingPredicate(ctx, id, AdaptiveNetworkHardeningOperationPredicate{})
}

// ListByExtendedResourceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AdaptiveNetworkHardeningsClient) ListByExtendedResourceCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, predicate AdaptiveNetworkHardeningOperationPredicate) (result ListByExtendedResourceCompleteResult, err error) {
	items := make([]AdaptiveNetworkHardening, 0)

	resp, err := c.ListByExtendedResource(ctx, id)
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

	result = ListByExtendedResourceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
