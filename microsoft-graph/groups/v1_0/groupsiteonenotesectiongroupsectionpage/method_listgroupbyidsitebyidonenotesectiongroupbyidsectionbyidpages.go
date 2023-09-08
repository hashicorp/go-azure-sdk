package groupsiteonenotesectiongroupsectionpage

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

type ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenotePageCollectionResponse
}

type ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPagesCompleteResult struct {
	Items []models.OnenotePageCollectionResponse
}

// ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPages ...
func (c GroupSiteOnenoteSectionGroupSectionPageClient) ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPages(ctx context.Context, id GroupSiteOnenoteSectionGroupSectionId) (result ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPagesOperationResponse, err error) {
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

// ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPagesComplete retrieves all the results into a single object
func (c GroupSiteOnenoteSectionGroupSectionPageClient) ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPagesComplete(ctx context.Context, id GroupSiteOnenoteSectionGroupSectionId) (ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPagesCompleteResult, error) {
	return c.ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPagesCompleteMatchingPredicate(ctx, id, models.OnenotePageCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteOnenoteSectionGroupSectionPageClient) ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPagesCompleteMatchingPredicate(ctx context.Context, id GroupSiteOnenoteSectionGroupSectionId, predicate models.OnenotePageCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPagesCompleteResult, err error) {
	items := make([]models.OnenotePageCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPages(ctx, id)
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

	result = ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPagesCompleteResult{
		Items: items,
	}
	return
}
