package applicationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByNamespaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApplicationGroup
}

type ListByNamespaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ApplicationGroup
}

// ListByNamespace ...
func (c ApplicationGroupClient) ListByNamespace(ctx context.Context, id NamespaceId) (result ListByNamespaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/applicationGroups", id.ID()),
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
		Values *[]ApplicationGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByNamespaceComplete retrieves all the results into a single object
func (c ApplicationGroupClient) ListByNamespaceComplete(ctx context.Context, id NamespaceId) (ListByNamespaceCompleteResult, error) {
	return c.ListByNamespaceCompleteMatchingPredicate(ctx, id, ApplicationGroupOperationPredicate{})
}

// ListByNamespaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationGroupClient) ListByNamespaceCompleteMatchingPredicate(ctx context.Context, id NamespaceId, predicate ApplicationGroupOperationPredicate) (result ListByNamespaceCompleteResult, err error) {
	items := make([]ApplicationGroup, 0)

	resp, err := c.ListByNamespace(ctx, id)
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

	result = ListByNamespaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
