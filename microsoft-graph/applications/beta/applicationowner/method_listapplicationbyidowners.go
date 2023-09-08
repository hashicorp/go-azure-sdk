package applicationowner

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

type ListApplicationByIdOwnersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListApplicationByIdOwnersCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListApplicationByIdOwners ...
func (c ApplicationOwnerClient) ListApplicationByIdOwners(ctx context.Context, id ApplicationId) (result ListApplicationByIdOwnersOperationResponse, err error) {
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

// ListApplicationByIdOwnersComplete retrieves all the results into a single object
func (c ApplicationOwnerClient) ListApplicationByIdOwnersComplete(ctx context.Context, id ApplicationId) (ListApplicationByIdOwnersCompleteResult, error) {
	return c.ListApplicationByIdOwnersCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListApplicationByIdOwnersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationOwnerClient) ListApplicationByIdOwnersCompleteMatchingPredicate(ctx context.Context, id ApplicationId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListApplicationByIdOwnersCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListApplicationByIdOwners(ctx, id)
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

	result = ListApplicationByIdOwnersCompleteResult{
		Items: items,
	}
	return
}
