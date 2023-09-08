package meprofileaward

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

type ListMeProfileAwardsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PersonAwardCollectionResponse
}

type ListMeProfileAwardsCompleteResult struct {
	Items []models.PersonAwardCollectionResponse
}

// ListMeProfileAwards ...
func (c MeProfileAwardClient) ListMeProfileAwards(ctx context.Context) (result ListMeProfileAwardsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/profile/awards",
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

// ListMeProfileAwardsComplete retrieves all the results into a single object
func (c MeProfileAwardClient) ListMeProfileAwardsComplete(ctx context.Context) (ListMeProfileAwardsCompleteResult, error) {
	return c.ListMeProfileAwardsCompleteMatchingPredicate(ctx, models.PersonAwardCollectionResponseOperationPredicate{})
}

// ListMeProfileAwardsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeProfileAwardClient) ListMeProfileAwardsCompleteMatchingPredicate(ctx context.Context, predicate models.PersonAwardCollectionResponseOperationPredicate) (result ListMeProfileAwardsCompleteResult, err error) {
	items := make([]models.PersonAwardCollectionResponse, 0)

	resp, err := c.ListMeProfileAwards(ctx)
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

	result = ListMeProfileAwardsCompleteResult{
		Items: items,
	}
	return
}
