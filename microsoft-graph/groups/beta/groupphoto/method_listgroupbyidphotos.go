package groupphoto

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

type ListGroupByIdPhotosOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ProfilePhotoCollectionResponse
}

type ListGroupByIdPhotosCompleteResult struct {
	Items []models.ProfilePhotoCollectionResponse
}

// ListGroupByIdPhotos ...
func (c GroupPhotoClient) ListGroupByIdPhotos(ctx context.Context, id GroupId) (result ListGroupByIdPhotosOperationResponse, err error) {
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

// ListGroupByIdPhotosComplete retrieves all the results into a single object
func (c GroupPhotoClient) ListGroupByIdPhotosComplete(ctx context.Context, id GroupId) (ListGroupByIdPhotosCompleteResult, error) {
	return c.ListGroupByIdPhotosCompleteMatchingPredicate(ctx, id, models.ProfilePhotoCollectionResponseOperationPredicate{})
}

// ListGroupByIdPhotosCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPhotoClient) ListGroupByIdPhotosCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.ProfilePhotoCollectionResponseOperationPredicate) (result ListGroupByIdPhotosCompleteResult, err error) {
	items := make([]models.ProfilePhotoCollectionResponse, 0)

	resp, err := c.ListGroupByIdPhotos(ctx, id)
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

	result = ListGroupByIdPhotosCompleteResult{
		Items: items,
	}
	return
}
