package trustedidproviders

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TrustedIdProvider
}

type ListByAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []TrustedIdProvider
}

// ListByAccount ...
func (c TrustedIdProvidersClient) ListByAccount(ctx context.Context, id AccountId) (result ListByAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/trustedIdProviders", id.ID()),
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
		Values *[]TrustedIdProvider `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByAccountComplete retrieves all the results into a single object
func (c TrustedIdProvidersClient) ListByAccountComplete(ctx context.Context, id AccountId) (ListByAccountCompleteResult, error) {
	return c.ListByAccountCompleteMatchingPredicate(ctx, id, TrustedIdProviderOperationPredicate{})
}

// ListByAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TrustedIdProvidersClient) ListByAccountCompleteMatchingPredicate(ctx context.Context, id AccountId, predicate TrustedIdProviderOperationPredicate) (result ListByAccountCompleteResult, err error) {
	items := make([]TrustedIdProvider, 0)

	resp, err := c.ListByAccount(ctx, id)
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

	result = ListByAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
