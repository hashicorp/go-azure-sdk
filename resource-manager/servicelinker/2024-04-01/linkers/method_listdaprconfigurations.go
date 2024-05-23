package linkers

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

type ListDaprConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DaprConfigurationResource
}

type ListDaprConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DaprConfigurationResource
}

// ListDaprConfigurations ...
func (c LinkersClient) ListDaprConfigurations(ctx context.Context, id commonids.ScopeId) (result ListDaprConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.ServiceLinker/daprConfigurations", id.ID()),
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
		Values *[]DaprConfigurationResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDaprConfigurationsComplete retrieves all the results into a single object
func (c LinkersClient) ListDaprConfigurationsComplete(ctx context.Context, id commonids.ScopeId) (ListDaprConfigurationsCompleteResult, error) {
	return c.ListDaprConfigurationsCompleteMatchingPredicate(ctx, id, DaprConfigurationResourceOperationPredicate{})
}

// ListDaprConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LinkersClient) ListDaprConfigurationsCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, predicate DaprConfigurationResourceOperationPredicate) (result ListDaprConfigurationsCompleteResult, err error) {
	items := make([]DaprConfigurationResource, 0)

	resp, err := c.ListDaprConfigurations(ctx, id)
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

	result = ListDaprConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
