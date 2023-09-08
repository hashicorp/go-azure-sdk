package useronenotenotebooksectiongroupsectionpage

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

type ListUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenotePageCollectionResponse
}

type ListUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesCompleteResult struct {
	Items []models.OnenotePageCollectionResponse
}

// ListUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPages ...
func (c UserOnenoteNotebookSectionGroupSectionPageClient) ListUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPages(ctx context.Context, id UserOnenoteNotebookSectionGroupSectionId) (result ListUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesOperationResponse, err error) {
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

// ListUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesComplete retrieves all the results into a single object
func (c UserOnenoteNotebookSectionGroupSectionPageClient) ListUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesComplete(ctx context.Context, id UserOnenoteNotebookSectionGroupSectionId) (ListUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesCompleteResult, error) {
	return c.ListUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesCompleteMatchingPredicate(ctx, id, models.OnenotePageCollectionResponseOperationPredicate{})
}

// ListUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOnenoteNotebookSectionGroupSectionPageClient) ListUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesCompleteMatchingPredicate(ctx context.Context, id UserOnenoteNotebookSectionGroupSectionId, predicate models.OnenotePageCollectionResponseOperationPredicate) (result ListUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesCompleteResult, err error) {
	items := make([]models.OnenotePageCollectionResponse, 0)

	resp, err := c.ListUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPages(ctx, id)
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

	result = ListUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesCompleteResult{
		Items: items,
	}
	return
}
