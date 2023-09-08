package meonenotenotebooksectionpage

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

type ListMeOnenoteNotebookByIdSectionByIdPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenotePageCollectionResponse
}

type ListMeOnenoteNotebookByIdSectionByIdPagesCompleteResult struct {
	Items []models.OnenotePageCollectionResponse
}

// ListMeOnenoteNotebookByIdSectionByIdPages ...
func (c MeOnenoteNotebookSectionPageClient) ListMeOnenoteNotebookByIdSectionByIdPages(ctx context.Context, id MeOnenoteNotebookSectionId) (result ListMeOnenoteNotebookByIdSectionByIdPagesOperationResponse, err error) {
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

// ListMeOnenoteNotebookByIdSectionByIdPagesComplete retrieves all the results into a single object
func (c MeOnenoteNotebookSectionPageClient) ListMeOnenoteNotebookByIdSectionByIdPagesComplete(ctx context.Context, id MeOnenoteNotebookSectionId) (ListMeOnenoteNotebookByIdSectionByIdPagesCompleteResult, error) {
	return c.ListMeOnenoteNotebookByIdSectionByIdPagesCompleteMatchingPredicate(ctx, id, models.OnenotePageCollectionResponseOperationPredicate{})
}

// ListMeOnenoteNotebookByIdSectionByIdPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOnenoteNotebookSectionPageClient) ListMeOnenoteNotebookByIdSectionByIdPagesCompleteMatchingPredicate(ctx context.Context, id MeOnenoteNotebookSectionId, predicate models.OnenotePageCollectionResponseOperationPredicate) (result ListMeOnenoteNotebookByIdSectionByIdPagesCompleteResult, err error) {
	items := make([]models.OnenotePageCollectionResponse, 0)

	resp, err := c.ListMeOnenoteNotebookByIdSectionByIdPages(ctx, id)
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

	result = ListMeOnenoteNotebookByIdSectionByIdPagesCompleteResult{
		Items: items,
	}
	return
}
