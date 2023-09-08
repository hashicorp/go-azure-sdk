package userprofileaward

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

type ListUserByIdProfileAwardsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PersonAwardCollectionResponse
}

type ListUserByIdProfileAwardsCompleteResult struct {
	Items []models.PersonAwardCollectionResponse
}

// ListUserByIdProfileAwards ...
func (c UserProfileAwardClient) ListUserByIdProfileAwards(ctx context.Context, id UserId) (result ListUserByIdProfileAwardsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/profile/awards", id.ID()),
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
		Values *[]models.PersonAwardCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdProfileAwardsComplete retrieves all the results into a single object
func (c UserProfileAwardClient) ListUserByIdProfileAwardsComplete(ctx context.Context, id UserId) (ListUserByIdProfileAwardsCompleteResult, error) {
	return c.ListUserByIdProfileAwardsCompleteMatchingPredicate(ctx, id, models.PersonAwardCollectionResponseOperationPredicate{})
}

// ListUserByIdProfileAwardsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserProfileAwardClient) ListUserByIdProfileAwardsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.PersonAwardCollectionResponseOperationPredicate) (result ListUserByIdProfileAwardsCompleteResult, err error) {
	items := make([]models.PersonAwardCollectionResponse, 0)

	resp, err := c.ListUserByIdProfileAwards(ctx, id)
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

	result = ListUserByIdProfileAwardsCompleteResult{
		Items: items,
	}
	return
}
