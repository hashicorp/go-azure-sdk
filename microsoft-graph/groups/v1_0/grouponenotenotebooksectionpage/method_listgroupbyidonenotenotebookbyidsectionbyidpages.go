package grouponenotenotebooksectionpage

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

type ListGroupByIdOnenoteNotebookByIdSectionByIdPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenotePageCollectionResponse
}

type ListGroupByIdOnenoteNotebookByIdSectionByIdPagesCompleteResult struct {
	Items []models.OnenotePageCollectionResponse
}

// ListGroupByIdOnenoteNotebookByIdSectionByIdPages ...
func (c GroupOnenoteNotebookSectionPageClient) ListGroupByIdOnenoteNotebookByIdSectionByIdPages(ctx context.Context, id GroupOnenoteNotebookSectionId) (result ListGroupByIdOnenoteNotebookByIdSectionByIdPagesOperationResponse, err error) {
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

// ListGroupByIdOnenoteNotebookByIdSectionByIdPagesComplete retrieves all the results into a single object
func (c GroupOnenoteNotebookSectionPageClient) ListGroupByIdOnenoteNotebookByIdSectionByIdPagesComplete(ctx context.Context, id GroupOnenoteNotebookSectionId) (ListGroupByIdOnenoteNotebookByIdSectionByIdPagesCompleteResult, error) {
	return c.ListGroupByIdOnenoteNotebookByIdSectionByIdPagesCompleteMatchingPredicate(ctx, id, models.OnenotePageCollectionResponseOperationPredicate{})
}

// ListGroupByIdOnenoteNotebookByIdSectionByIdPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupOnenoteNotebookSectionPageClient) ListGroupByIdOnenoteNotebookByIdSectionByIdPagesCompleteMatchingPredicate(ctx context.Context, id GroupOnenoteNotebookSectionId, predicate models.OnenotePageCollectionResponseOperationPredicate) (result ListGroupByIdOnenoteNotebookByIdSectionByIdPagesCompleteResult, err error) {
	items := make([]models.OnenotePageCollectionResponse, 0)

	resp, err := c.ListGroupByIdOnenoteNotebookByIdSectionByIdPages(ctx, id)
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

	result = ListGroupByIdOnenoteNotebookByIdSectionByIdPagesCompleteResult{
		Items: items,
	}
	return
}
