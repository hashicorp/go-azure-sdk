package userprofileposition

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

type ListUserByIdProfilePositionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.WorkPositionCollectionResponse
}

type ListUserByIdProfilePositionsCompleteResult struct {
	Items []models.WorkPositionCollectionResponse
}

// ListUserByIdProfilePositions ...
func (c UserProfilePositionClient) ListUserByIdProfilePositions(ctx context.Context, id UserId) (result ListUserByIdProfilePositionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/profile/positions", id.ID()),
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
		Values *[]models.WorkPositionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdProfilePositionsComplete retrieves all the results into a single object
func (c UserProfilePositionClient) ListUserByIdProfilePositionsComplete(ctx context.Context, id UserId) (ListUserByIdProfilePositionsCompleteResult, error) {
	return c.ListUserByIdProfilePositionsCompleteMatchingPredicate(ctx, id, models.WorkPositionCollectionResponseOperationPredicate{})
}

// ListUserByIdProfilePositionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserProfilePositionClient) ListUserByIdProfilePositionsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.WorkPositionCollectionResponseOperationPredicate) (result ListUserByIdProfilePositionsCompleteResult, err error) {
	items := make([]models.WorkPositionCollectionResponse, 0)

	resp, err := c.ListUserByIdProfilePositions(ctx, id)
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

	result = ListUserByIdProfilePositionsCompleteResult{
		Items: items,
	}
	return
}
