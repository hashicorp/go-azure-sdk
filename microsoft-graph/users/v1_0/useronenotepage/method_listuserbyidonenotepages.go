package useronenotepage

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

type ListUserByIdOnenotePagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenotePageCollectionResponse
}

type ListUserByIdOnenotePagesCompleteResult struct {
	Items []models.OnenotePageCollectionResponse
}

// ListUserByIdOnenotePages ...
func (c UserOnenotePageClient) ListUserByIdOnenotePages(ctx context.Context, id UserId) (result ListUserByIdOnenotePagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/onenote/pages", id.ID()),
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
		Values *[]models.OnenotePageCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdOnenotePagesComplete retrieves all the results into a single object
func (c UserOnenotePageClient) ListUserByIdOnenotePagesComplete(ctx context.Context, id UserId) (ListUserByIdOnenotePagesCompleteResult, error) {
	return c.ListUserByIdOnenotePagesCompleteMatchingPredicate(ctx, id, models.OnenotePageCollectionResponseOperationPredicate{})
}

// ListUserByIdOnenotePagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOnenotePageClient) ListUserByIdOnenotePagesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.OnenotePageCollectionResponseOperationPredicate) (result ListUserByIdOnenotePagesCompleteResult, err error) {
	items := make([]models.OnenotePageCollectionResponse, 0)

	resp, err := c.ListUserByIdOnenotePages(ctx, id)
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

	result = ListUserByIdOnenotePagesCompleteResult{
		Items: items,
	}
	return
}
