package groupsiteonenotenotebooksectionpage

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

type ListGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenotePageCollectionResponse
}

type ListGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPagesCompleteResult struct {
	Items []models.OnenotePageCollectionResponse
}

// ListGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPages ...
func (c GroupSiteOnenoteNotebookSectionPageClient) ListGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPages(ctx context.Context, id GroupSiteOnenoteNotebookSectionId) (result ListGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPagesOperationResponse, err error) {
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

// ListGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPagesComplete retrieves all the results into a single object
func (c GroupSiteOnenoteNotebookSectionPageClient) ListGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPagesComplete(ctx context.Context, id GroupSiteOnenoteNotebookSectionId) (ListGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPagesCompleteResult, error) {
	return c.ListGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPagesCompleteMatchingPredicate(ctx, id, models.OnenotePageCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteOnenoteNotebookSectionPageClient) ListGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPagesCompleteMatchingPredicate(ctx context.Context, id GroupSiteOnenoteNotebookSectionId, predicate models.OnenotePageCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPagesCompleteResult, err error) {
	items := make([]models.OnenotePageCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPages(ctx, id)
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

	result = ListGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPagesCompleteResult{
		Items: items,
	}
	return
}
