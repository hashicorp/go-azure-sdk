package groupowner

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

type ListGroupByIdOwnersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListGroupByIdOwnersCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListGroupByIdOwners ...
func (c GroupOwnerClient) ListGroupByIdOwners(ctx context.Context, id GroupId) (result ListGroupByIdOwnersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/owners", id.ID()),
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
		Values *[]models.DirectoryObjectCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdOwnersComplete retrieves all the results into a single object
func (c GroupOwnerClient) ListGroupByIdOwnersComplete(ctx context.Context, id GroupId) (ListGroupByIdOwnersCompleteResult, error) {
	return c.ListGroupByIdOwnersCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListGroupByIdOwnersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupOwnerClient) ListGroupByIdOwnersCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListGroupByIdOwnersCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListGroupByIdOwners(ctx, id)
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

	result = ListGroupByIdOwnersCompleteResult{
		Items: items,
	}
	return
}
