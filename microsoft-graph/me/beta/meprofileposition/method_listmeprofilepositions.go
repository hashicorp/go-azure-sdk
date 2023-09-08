package meprofileposition

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

type ListMeProfilePositionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.WorkPositionCollectionResponse
}

type ListMeProfilePositionsCompleteResult struct {
	Items []models.WorkPositionCollectionResponse
}

// ListMeProfilePositions ...
func (c MeProfilePositionClient) ListMeProfilePositions(ctx context.Context) (result ListMeProfilePositionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/profile/positions",
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

// ListMeProfilePositionsComplete retrieves all the results into a single object
func (c MeProfilePositionClient) ListMeProfilePositionsComplete(ctx context.Context) (ListMeProfilePositionsCompleteResult, error) {
	return c.ListMeProfilePositionsCompleteMatchingPredicate(ctx, models.WorkPositionCollectionResponseOperationPredicate{})
}

// ListMeProfilePositionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeProfilePositionClient) ListMeProfilePositionsCompleteMatchingPredicate(ctx context.Context, predicate models.WorkPositionCollectionResponseOperationPredicate) (result ListMeProfilePositionsCompleteResult, err error) {
	items := make([]models.WorkPositionCollectionResponse, 0)

	resp, err := c.ListMeProfilePositions(ctx)
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

	result = ListMeProfilePositionsCompleteResult{
		Items: items,
	}
	return
}
