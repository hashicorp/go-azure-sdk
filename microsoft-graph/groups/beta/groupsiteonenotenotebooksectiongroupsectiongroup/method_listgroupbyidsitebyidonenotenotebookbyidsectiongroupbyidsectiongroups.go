package groupsiteonenotenotebooksectiongroupsectiongroup

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

type ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SectionGroupCollectionResponse
}

type ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsCompleteResult struct {
	Items []models.SectionGroupCollectionResponse
}

// ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionGroups ...
func (c GroupSiteOnenoteNotebookSectionGroupSectionGroupClient) ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionGroups(ctx context.Context, id GroupSiteOnenoteNotebookSectionGroupId) (result ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsOperationResponse, err error) {
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

// ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsComplete retrieves all the results into a single object
func (c GroupSiteOnenoteNotebookSectionGroupSectionGroupClient) ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsComplete(ctx context.Context, id GroupSiteOnenoteNotebookSectionGroupId) (ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsCompleteMatchingPredicate(ctx, id, models.SectionGroupCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteOnenoteNotebookSectionGroupSectionGroupClient) ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsCompleteMatchingPredicate(ctx context.Context, id GroupSiteOnenoteNotebookSectionGroupId, predicate models.SectionGroupCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsCompleteResult, err error) {
	items := make([]models.SectionGroupCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionGroups(ctx, id)
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

	result = ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsCompleteResult{
		Items: items,
	}
	return
}
