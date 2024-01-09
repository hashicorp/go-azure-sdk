package hypervrunasaccounts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAllRunAsAccountsInSiteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]HyperVRunAsAccount
}

type GetAllRunAsAccountsInSiteCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []HyperVRunAsAccount
}

// GetAllRunAsAccountsInSite ...
func (c HyperVRunAsAccountsClient) GetAllRunAsAccountsInSite(ctx context.Context, id HyperVSiteId) (result GetAllRunAsAccountsInSiteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/runAsAccounts", id.ID()),
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
		Values *[]HyperVRunAsAccount `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetAllRunAsAccountsInSiteComplete retrieves all the results into a single object
func (c HyperVRunAsAccountsClient) GetAllRunAsAccountsInSiteComplete(ctx context.Context, id HyperVSiteId) (GetAllRunAsAccountsInSiteCompleteResult, error) {
	return c.GetAllRunAsAccountsInSiteCompleteMatchingPredicate(ctx, id, HyperVRunAsAccountOperationPredicate{})
}

// GetAllRunAsAccountsInSiteCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HyperVRunAsAccountsClient) GetAllRunAsAccountsInSiteCompleteMatchingPredicate(ctx context.Context, id HyperVSiteId, predicate HyperVRunAsAccountOperationPredicate) (result GetAllRunAsAccountsInSiteCompleteResult, err error) {
	items := make([]HyperVRunAsAccount, 0)

	resp, err := c.GetAllRunAsAccountsInSite(ctx, id)
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

	result = GetAllRunAsAccountsInSiteCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
