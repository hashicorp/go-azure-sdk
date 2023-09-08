package grouponenotenotebooksectiongroupsectiongroup

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

type ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SectionGroupCollectionResponse
}

type ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsCompleteResult struct {
	Items []models.SectionGroupCollectionResponse
}

// ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionGroups ...
func (c GroupOnenoteNotebookSectionGroupSectionGroupClient) ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionGroups(ctx context.Context, id GroupOnenoteNotebookSectionGroupId) (result ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsOperationResponse, err error) {
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

// ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsComplete retrieves all the results into a single object
func (c GroupOnenoteNotebookSectionGroupSectionGroupClient) ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsComplete(ctx context.Context, id GroupOnenoteNotebookSectionGroupId) (ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsCompleteResult, error) {
	return c.ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsCompleteMatchingPredicate(ctx, id, models.SectionGroupCollectionResponseOperationPredicate{})
}

// ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupOnenoteNotebookSectionGroupSectionGroupClient) ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsCompleteMatchingPredicate(ctx context.Context, id GroupOnenoteNotebookSectionGroupId, predicate models.SectionGroupCollectionResponseOperationPredicate) (result ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsCompleteResult, err error) {
	items := make([]models.SectionGroupCollectionResponse, 0)

	resp, err := c.ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionGroups(ctx, id)
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

	result = ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionGroupsCompleteResult{
		Items: items,
	}
	return
}
