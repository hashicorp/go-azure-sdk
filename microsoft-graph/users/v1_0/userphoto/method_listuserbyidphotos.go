package userphoto

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

type ListUserByIdPhotosOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ProfilePhotoCollectionResponse
}

type ListUserByIdPhotosCompleteResult struct {
	Items []models.ProfilePhotoCollectionResponse
}

// ListUserByIdPhotos ...
func (c UserPhotoClient) ListUserByIdPhotos(ctx context.Context, id UserId) (result ListUserByIdPhotosOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/photos", id.ID()),
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
		Values *[]models.ProfilePhotoCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdPhotosComplete retrieves all the results into a single object
func (c UserPhotoClient) ListUserByIdPhotosComplete(ctx context.Context, id UserId) (ListUserByIdPhotosCompleteResult, error) {
	return c.ListUserByIdPhotosCompleteMatchingPredicate(ctx, id, models.ProfilePhotoCollectionResponseOperationPredicate{})
}

// ListUserByIdPhotosCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPhotoClient) ListUserByIdPhotosCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.ProfilePhotoCollectionResponseOperationPredicate) (result ListUserByIdPhotosCompleteResult, err error) {
	items := make([]models.ProfilePhotoCollectionResponse, 0)

	resp, err := c.ListUserByIdPhotos(ctx, id)
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

	result = ListUserByIdPhotosCompleteResult{
		Items: items,
	}
	return
}
