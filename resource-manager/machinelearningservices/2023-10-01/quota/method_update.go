package quota

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]UpdateWorkspaceQuotas
}

type UpdateCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []UpdateWorkspaceQuotas
}

// Update ...
func (c QuotaClient) Update(ctx context.Context, id LocationId, input QuotaUpdateParameters) (result UpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/updateQuotas", id.ID()),
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
		Values *[]UpdateWorkspaceQuotas `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// UpdateComplete retrieves all the results into a single object
func (c QuotaClient) UpdateComplete(ctx context.Context, id LocationId, input QuotaUpdateParameters) (UpdateCompleteResult, error) {
	return c.UpdateCompleteMatchingPredicate(ctx, id, input, UpdateWorkspaceQuotasOperationPredicate{})
}

// UpdateCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c QuotaClient) UpdateCompleteMatchingPredicate(ctx context.Context, id LocationId, input QuotaUpdateParameters, predicate UpdateWorkspaceQuotasOperationPredicate) (result UpdateCompleteResult, err error) {
	items := make([]UpdateWorkspaceQuotas, 0)

	resp, err := c.Update(ctx, id, input)
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

	result = UpdateCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
