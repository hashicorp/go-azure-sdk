package meonenotesectiongroupsectionpage

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

type ListMeOnenoteSectionGroupByIdSectionByIdPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenotePageCollectionResponse
}

type ListMeOnenoteSectionGroupByIdSectionByIdPagesCompleteResult struct {
	Items []models.OnenotePageCollectionResponse
}

// ListMeOnenoteSectionGroupByIdSectionByIdPages ...
func (c MeOnenoteSectionGroupSectionPageClient) ListMeOnenoteSectionGroupByIdSectionByIdPages(ctx context.Context, id MeOnenoteSectionGroupSectionId) (result ListMeOnenoteSectionGroupByIdSectionByIdPagesOperationResponse, err error) {
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

// ListMeOnenoteSectionGroupByIdSectionByIdPagesComplete retrieves all the results into a single object
func (c MeOnenoteSectionGroupSectionPageClient) ListMeOnenoteSectionGroupByIdSectionByIdPagesComplete(ctx context.Context, id MeOnenoteSectionGroupSectionId) (ListMeOnenoteSectionGroupByIdSectionByIdPagesCompleteResult, error) {
	return c.ListMeOnenoteSectionGroupByIdSectionByIdPagesCompleteMatchingPredicate(ctx, id, models.OnenotePageCollectionResponseOperationPredicate{})
}

// ListMeOnenoteSectionGroupByIdSectionByIdPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOnenoteSectionGroupSectionPageClient) ListMeOnenoteSectionGroupByIdSectionByIdPagesCompleteMatchingPredicate(ctx context.Context, id MeOnenoteSectionGroupSectionId, predicate models.OnenotePageCollectionResponseOperationPredicate) (result ListMeOnenoteSectionGroupByIdSectionByIdPagesCompleteResult, err error) {
	items := make([]models.OnenotePageCollectionResponse, 0)

	resp, err := c.ListMeOnenoteSectionGroupByIdSectionByIdPages(ctx, id)
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

	result = ListMeOnenoteSectionGroupByIdSectionByIdPagesCompleteResult{
		Items: items,
	}
	return
}
