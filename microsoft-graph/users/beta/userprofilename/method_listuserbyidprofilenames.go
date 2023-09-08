package userprofilename

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

type ListUserByIdProfileNamesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PersonNameCollectionResponse
}

type ListUserByIdProfileNamesCompleteResult struct {
	Items []models.PersonNameCollectionResponse
}

// ListUserByIdProfileNames ...
func (c UserProfileNameClient) ListUserByIdProfileNames(ctx context.Context, id UserId) (result ListUserByIdProfileNamesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/profile/names", id.ID()),
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
		Values *[]models.PersonNameCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdProfileNamesComplete retrieves all the results into a single object
func (c UserProfileNameClient) ListUserByIdProfileNamesComplete(ctx context.Context, id UserId) (ListUserByIdProfileNamesCompleteResult, error) {
	return c.ListUserByIdProfileNamesCompleteMatchingPredicate(ctx, id, models.PersonNameCollectionResponseOperationPredicate{})
}

// ListUserByIdProfileNamesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserProfileNameClient) ListUserByIdProfileNamesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.PersonNameCollectionResponseOperationPredicate) (result ListUserByIdProfileNamesCompleteResult, err error) {
	items := make([]models.PersonNameCollectionResponse, 0)

	resp, err := c.ListUserByIdProfileNames(ctx, id)
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

	result = ListUserByIdProfileNamesCompleteResult{
		Items: items,
	}
	return
}
