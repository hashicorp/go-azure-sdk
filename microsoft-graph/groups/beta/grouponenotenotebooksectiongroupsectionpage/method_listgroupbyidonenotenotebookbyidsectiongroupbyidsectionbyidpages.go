package grouponenotenotebooksectiongroupsectionpage

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

type ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenotePageCollectionResponse
}

type ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesCompleteResult struct {
	Items []models.OnenotePageCollectionResponse
}

// ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPages ...
func (c GroupOnenoteNotebookSectionGroupSectionPageClient) ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPages(ctx context.Context, id GroupOnenoteNotebookSectionGroupSectionId) (result ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesOperationResponse, err error) {
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

// ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesComplete retrieves all the results into a single object
func (c GroupOnenoteNotebookSectionGroupSectionPageClient) ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesComplete(ctx context.Context, id GroupOnenoteNotebookSectionGroupSectionId) (ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesCompleteResult, error) {
	return c.ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesCompleteMatchingPredicate(ctx, id, models.OnenotePageCollectionResponseOperationPredicate{})
}

// ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupOnenoteNotebookSectionGroupSectionPageClient) ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesCompleteMatchingPredicate(ctx context.Context, id GroupOnenoteNotebookSectionGroupSectionId, predicate models.OnenotePageCollectionResponseOperationPredicate) (result ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesCompleteResult, err error) {
	items := make([]models.OnenotePageCollectionResponse, 0)

	resp, err := c.ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPages(ctx, id)
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

	result = ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesCompleteResult{
		Items: items,
	}
	return
}
