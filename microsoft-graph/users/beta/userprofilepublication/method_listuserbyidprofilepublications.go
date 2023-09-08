package userprofilepublication

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

type ListUserByIdProfilePublicationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ItemPublicationCollectionResponse
}

type ListUserByIdProfilePublicationsCompleteResult struct {
	Items []models.ItemPublicationCollectionResponse
}

// ListUserByIdProfilePublications ...
func (c UserProfilePublicationClient) ListUserByIdProfilePublications(ctx context.Context, id UserId) (result ListUserByIdProfilePublicationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/profile/publications", id.ID()),
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
		Values *[]models.ItemPublicationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdProfilePublicationsComplete retrieves all the results into a single object
func (c UserProfilePublicationClient) ListUserByIdProfilePublicationsComplete(ctx context.Context, id UserId) (ListUserByIdProfilePublicationsCompleteResult, error) {
	return c.ListUserByIdProfilePublicationsCompleteMatchingPredicate(ctx, id, models.ItemPublicationCollectionResponseOperationPredicate{})
}

// ListUserByIdProfilePublicationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserProfilePublicationClient) ListUserByIdProfilePublicationsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.ItemPublicationCollectionResponseOperationPredicate) (result ListUserByIdProfilePublicationsCompleteResult, err error) {
	items := make([]models.ItemPublicationCollectionResponse, 0)

	resp, err := c.ListUserByIdProfilePublications(ctx, id)
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

	result = ListUserByIdProfilePublicationsCompleteResult{
		Items: items,
	}
	return
}
