package useronenotenotebooksectionpage

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

type ListUserByIdOnenoteNotebookByIdSectionByIdPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenotePageCollectionResponse
}

type ListUserByIdOnenoteNotebookByIdSectionByIdPagesCompleteResult struct {
	Items []models.OnenotePageCollectionResponse
}

// ListUserByIdOnenoteNotebookByIdSectionByIdPages ...
func (c UserOnenoteNotebookSectionPageClient) ListUserByIdOnenoteNotebookByIdSectionByIdPages(ctx context.Context, id UserOnenoteNotebookSectionId) (result ListUserByIdOnenoteNotebookByIdSectionByIdPagesOperationResponse, err error) {
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

// ListUserByIdOnenoteNotebookByIdSectionByIdPagesComplete retrieves all the results into a single object
func (c UserOnenoteNotebookSectionPageClient) ListUserByIdOnenoteNotebookByIdSectionByIdPagesComplete(ctx context.Context, id UserOnenoteNotebookSectionId) (ListUserByIdOnenoteNotebookByIdSectionByIdPagesCompleteResult, error) {
	return c.ListUserByIdOnenoteNotebookByIdSectionByIdPagesCompleteMatchingPredicate(ctx, id, models.OnenotePageCollectionResponseOperationPredicate{})
}

// ListUserByIdOnenoteNotebookByIdSectionByIdPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOnenoteNotebookSectionPageClient) ListUserByIdOnenoteNotebookByIdSectionByIdPagesCompleteMatchingPredicate(ctx context.Context, id UserOnenoteNotebookSectionId, predicate models.OnenotePageCollectionResponseOperationPredicate) (result ListUserByIdOnenoteNotebookByIdSectionByIdPagesCompleteResult, err error) {
	items := make([]models.OnenotePageCollectionResponse, 0)

	resp, err := c.ListUserByIdOnenoteNotebookByIdSectionByIdPages(ctx, id)
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

	result = ListUserByIdOnenoteNotebookByIdSectionByIdPagesCompleteResult{
		Items: items,
	}
	return
}
