package userinsightused

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdInsightUsedsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UsedInsightCollectionResponse
}

type ListUserByIdInsightUsedsCompleteResult struct {
	Items []models.UsedInsightCollectionResponse
}

// ListUserByIdInsightUseds ...
func (c UserInsightUsedClient) ListUserByIdInsightUseds(ctx context.Context, id UserId) (result ListUserByIdInsightUsedsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/insights/used", id.ID()),
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
		Values *[]models.UsedInsightCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdInsightUsedsComplete retrieves all the results into a single object
func (c UserInsightUsedClient) ListUserByIdInsightUsedsComplete(ctx context.Context, id UserId) (ListUserByIdInsightUsedsCompleteResult, error) {
	return c.ListUserByIdInsightUsedsCompleteMatchingPredicate(ctx, id, models.UsedInsightCollectionResponseOperationPredicate{})
}

// ListUserByIdInsightUsedsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightUsedClient) ListUserByIdInsightUsedsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.UsedInsightCollectionResponseOperationPredicate) (result ListUserByIdInsightUsedsCompleteResult, err error) {
	items := make([]models.UsedInsightCollectionResponse, 0)

	resp, err := c.ListUserByIdInsightUseds(ctx, id)
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

	result = ListUserByIdInsightUsedsCompleteResult{
		Items: items,
	}
	return
}
