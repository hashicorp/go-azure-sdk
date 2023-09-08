package userprofileinterest

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

type ListUserByIdProfileInterestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PersonInterestCollectionResponse
}

type ListUserByIdProfileInterestsCompleteResult struct {
	Items []models.PersonInterestCollectionResponse
}

// ListUserByIdProfileInterests ...
func (c UserProfileInterestClient) ListUserByIdProfileInterests(ctx context.Context, id UserId) (result ListUserByIdProfileInterestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/profile/interests", id.ID()),
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
		Values *[]models.PersonInterestCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdProfileInterestsComplete retrieves all the results into a single object
func (c UserProfileInterestClient) ListUserByIdProfileInterestsComplete(ctx context.Context, id UserId) (ListUserByIdProfileInterestsCompleteResult, error) {
	return c.ListUserByIdProfileInterestsCompleteMatchingPredicate(ctx, id, models.PersonInterestCollectionResponseOperationPredicate{})
}

// ListUserByIdProfileInterestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserProfileInterestClient) ListUserByIdProfileInterestsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.PersonInterestCollectionResponseOperationPredicate) (result ListUserByIdProfileInterestsCompleteResult, err error) {
	items := make([]models.PersonInterestCollectionResponse, 0)

	resp, err := c.ListUserByIdProfileInterests(ctx, id)
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

	result = ListUserByIdProfileInterestsCompleteResult{
		Items: items,
	}
	return
}
