package userinsightshared

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdInsightSharedsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SharedInsightCollectionResponse
}

type ListUserByIdInsightSharedsCompleteResult struct {
	Items []models.SharedInsightCollectionResponse
}

// ListUserByIdInsightShareds ...
func (c UserInsightSharedClient) ListUserByIdInsightShareds(ctx context.Context, id UserId) (result ListUserByIdInsightSharedsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/insights/shared", id.ID()),
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
		Values *[]models.SharedInsightCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdInsightSharedsComplete retrieves all the results into a single object
func (c UserInsightSharedClient) ListUserByIdInsightSharedsComplete(ctx context.Context, id UserId) (ListUserByIdInsightSharedsCompleteResult, error) {
	return c.ListUserByIdInsightSharedsCompleteMatchingPredicate(ctx, id, models.SharedInsightCollectionResponseOperationPredicate{})
}

// ListUserByIdInsightSharedsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightSharedClient) ListUserByIdInsightSharedsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.SharedInsightCollectionResponseOperationPredicate) (result ListUserByIdInsightSharedsCompleteResult, err error) {
	items := make([]models.SharedInsightCollectionResponse, 0)

	resp, err := c.ListUserByIdInsightShareds(ctx, id)
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

	result = ListUserByIdInsightSharedsCompleteResult{
		Items: items,
	}
	return
}
