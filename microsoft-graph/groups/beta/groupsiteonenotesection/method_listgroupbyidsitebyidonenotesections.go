package groupsiteonenotesection

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

type ListGroupByIdSiteByIdOnenoteSectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenoteSectionCollectionResponse
}

type ListGroupByIdSiteByIdOnenoteSectionsCompleteResult struct {
	Items []models.OnenoteSectionCollectionResponse
}

// ListGroupByIdSiteByIdOnenoteSections ...
func (c GroupSiteOnenoteSectionClient) ListGroupByIdSiteByIdOnenoteSections(ctx context.Context, id GroupSiteId) (result ListGroupByIdSiteByIdOnenoteSectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/onenote/sections", id.ID()),
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
		Values *[]models.OnenoteSectionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdOnenoteSectionsComplete retrieves all the results into a single object
func (c GroupSiteOnenoteSectionClient) ListGroupByIdSiteByIdOnenoteSectionsComplete(ctx context.Context, id GroupSiteId) (ListGroupByIdSiteByIdOnenoteSectionsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdOnenoteSectionsCompleteMatchingPredicate(ctx, id, models.OnenoteSectionCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdOnenoteSectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteOnenoteSectionClient) ListGroupByIdSiteByIdOnenoteSectionsCompleteMatchingPredicate(ctx context.Context, id GroupSiteId, predicate models.OnenoteSectionCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdOnenoteSectionsCompleteResult, err error) {
	items := make([]models.OnenoteSectionCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdOnenoteSections(ctx, id)
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

	result = ListGroupByIdSiteByIdOnenoteSectionsCompleteResult{
		Items: items,
	}
	return
}
