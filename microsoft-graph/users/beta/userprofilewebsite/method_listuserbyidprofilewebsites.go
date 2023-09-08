package userprofilewebsite

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

type ListUserByIdProfileWebsitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PersonWebsiteCollectionResponse
}

type ListUserByIdProfileWebsitesCompleteResult struct {
	Items []models.PersonWebsiteCollectionResponse
}

// ListUserByIdProfileWebsites ...
func (c UserProfileWebsiteClient) ListUserByIdProfileWebsites(ctx context.Context, id UserId) (result ListUserByIdProfileWebsitesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/profile/websites", id.ID()),
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
		Values *[]models.PersonWebsiteCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdProfileWebsitesComplete retrieves all the results into a single object
func (c UserProfileWebsiteClient) ListUserByIdProfileWebsitesComplete(ctx context.Context, id UserId) (ListUserByIdProfileWebsitesCompleteResult, error) {
	return c.ListUserByIdProfileWebsitesCompleteMatchingPredicate(ctx, id, models.PersonWebsiteCollectionResponseOperationPredicate{})
}

// ListUserByIdProfileWebsitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserProfileWebsiteClient) ListUserByIdProfileWebsitesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.PersonWebsiteCollectionResponseOperationPredicate) (result ListUserByIdProfileWebsitesCompleteResult, err error) {
	items := make([]models.PersonWebsiteCollectionResponse, 0)

	resp, err := c.ListUserByIdProfileWebsites(ctx, id)
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

	result = ListUserByIdProfileWebsitesCompleteResult{
		Items: items,
	}
	return
}
