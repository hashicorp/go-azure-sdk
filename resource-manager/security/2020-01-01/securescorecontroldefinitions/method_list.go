package securescorecontroldefinitions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SecureScoreControlDefinitionItem
}

type ListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SecureScoreControlDefinitionItem
}

// List ...
func (c SecureScoreControlDefinitionsClient) List(ctx context.Context) (result ListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/providers/Microsoft.Security/secureScoreControlDefinitions",
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
		Values *[]SecureScoreControlDefinitionItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComplete retrieves all the results into a single object
func (c SecureScoreControlDefinitionsClient) ListComplete(ctx context.Context) (ListCompleteResult, error) {
	return c.ListCompleteMatchingPredicate(ctx, SecureScoreControlDefinitionItemOperationPredicate{})
}

// ListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SecureScoreControlDefinitionsClient) ListCompleteMatchingPredicate(ctx context.Context, predicate SecureScoreControlDefinitionItemOperationPredicate) (result ListCompleteResult, err error) {
	items := make([]SecureScoreControlDefinitionItem, 0)

	resp, err := c.List(ctx)
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

	result = ListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
