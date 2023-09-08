package useronenotesectionpage

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

type ListUserByIdOnenoteSectionByIdPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenotePageCollectionResponse
}

type ListUserByIdOnenoteSectionByIdPagesCompleteResult struct {
	Items []models.OnenotePageCollectionResponse
}

// ListUserByIdOnenoteSectionByIdPages ...
func (c UserOnenoteSectionPageClient) ListUserByIdOnenoteSectionByIdPages(ctx context.Context, id UserOnenoteSectionId) (result ListUserByIdOnenoteSectionByIdPagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/pages", id.ID()),
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

// ListUserByIdOnenoteSectionByIdPagesComplete retrieves all the results into a single object
func (c UserOnenoteSectionPageClient) ListUserByIdOnenoteSectionByIdPagesComplete(ctx context.Context, id UserOnenoteSectionId) (ListUserByIdOnenoteSectionByIdPagesCompleteResult, error) {
	return c.ListUserByIdOnenoteSectionByIdPagesCompleteMatchingPredicate(ctx, id, models.OnenotePageCollectionResponseOperationPredicate{})
}

// ListUserByIdOnenoteSectionByIdPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOnenoteSectionPageClient) ListUserByIdOnenoteSectionByIdPagesCompleteMatchingPredicate(ctx context.Context, id UserOnenoteSectionId, predicate models.OnenotePageCollectionResponseOperationPredicate) (result ListUserByIdOnenoteSectionByIdPagesCompleteResult, err error) {
	items := make([]models.OnenotePageCollectionResponse, 0)

	resp, err := c.ListUserByIdOnenoteSectionByIdPages(ctx, id)
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

	result = ListUserByIdOnenoteSectionByIdPagesCompleteResult{
		Items: items,
	}
	return
}
