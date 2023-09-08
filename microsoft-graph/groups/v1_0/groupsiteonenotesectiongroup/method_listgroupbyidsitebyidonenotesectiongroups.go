package groupsiteonenotesectiongroup

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

type ListGroupByIdSiteByIdOnenoteSectionGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SectionGroupCollectionResponse
}

type ListGroupByIdSiteByIdOnenoteSectionGroupsCompleteResult struct {
	Items []models.SectionGroupCollectionResponse
}

// ListGroupByIdSiteByIdOnenoteSectionGroups ...
func (c GroupSiteOnenoteSectionGroupClient) ListGroupByIdSiteByIdOnenoteSectionGroups(ctx context.Context, id GroupSiteId) (result ListGroupByIdSiteByIdOnenoteSectionGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/onenote/sectionGroups", id.ID()),
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
		Values *[]models.SectionGroupCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdOnenoteSectionGroupsComplete retrieves all the results into a single object
func (c GroupSiteOnenoteSectionGroupClient) ListGroupByIdSiteByIdOnenoteSectionGroupsComplete(ctx context.Context, id GroupSiteId) (ListGroupByIdSiteByIdOnenoteSectionGroupsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdOnenoteSectionGroupsCompleteMatchingPredicate(ctx, id, models.SectionGroupCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdOnenoteSectionGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteOnenoteSectionGroupClient) ListGroupByIdSiteByIdOnenoteSectionGroupsCompleteMatchingPredicate(ctx context.Context, id GroupSiteId, predicate models.SectionGroupCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdOnenoteSectionGroupsCompleteResult, err error) {
	items := make([]models.SectionGroupCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdOnenoteSectionGroups(ctx, id)
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

	result = ListGroupByIdSiteByIdOnenoteSectionGroupsCompleteResult{
		Items: items,
	}
	return
}
