package groupsiteonenotesectiongroupsectiongroup

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

type ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SectionGroupCollectionResponse
}

type ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionGroupsCompleteResult struct {
	Items []models.SectionGroupCollectionResponse
}

// ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionGroups ...
func (c GroupSiteOnenoteSectionGroupSectionGroupClient) ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionGroups(ctx context.Context, id GroupSiteOnenoteSectionGroupId) (result ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/sectionGroups", id.ID()),
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

// ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionGroupsComplete retrieves all the results into a single object
func (c GroupSiteOnenoteSectionGroupSectionGroupClient) ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionGroupsComplete(ctx context.Context, id GroupSiteOnenoteSectionGroupId) (ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionGroupsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionGroupsCompleteMatchingPredicate(ctx, id, models.SectionGroupCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteOnenoteSectionGroupSectionGroupClient) ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionGroupsCompleteMatchingPredicate(ctx context.Context, id GroupSiteOnenoteSectionGroupId, predicate models.SectionGroupCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionGroupsCompleteResult, err error) {
	items := make([]models.SectionGroupCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionGroups(ctx, id)
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

	result = ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionGroupsCompleteResult{
		Items: items,
	}
	return
}
