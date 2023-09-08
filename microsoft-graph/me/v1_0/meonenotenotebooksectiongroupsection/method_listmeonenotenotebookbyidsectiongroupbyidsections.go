package meonenotenotebooksectiongroupsection

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

type ListMeOnenoteNotebookByIdSectionGroupByIdSectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenoteSectionCollectionResponse
}

type ListMeOnenoteNotebookByIdSectionGroupByIdSectionsCompleteResult struct {
	Items []models.OnenoteSectionCollectionResponse
}

// ListMeOnenoteNotebookByIdSectionGroupByIdSections ...
func (c MeOnenoteNotebookSectionGroupSectionClient) ListMeOnenoteNotebookByIdSectionGroupByIdSections(ctx context.Context, id MeOnenoteNotebookSectionGroupId) (result ListMeOnenoteNotebookByIdSectionGroupByIdSectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/sections", id.ID()),
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

// ListMeOnenoteNotebookByIdSectionGroupByIdSectionsComplete retrieves all the results into a single object
func (c MeOnenoteNotebookSectionGroupSectionClient) ListMeOnenoteNotebookByIdSectionGroupByIdSectionsComplete(ctx context.Context, id MeOnenoteNotebookSectionGroupId) (ListMeOnenoteNotebookByIdSectionGroupByIdSectionsCompleteResult, error) {
	return c.ListMeOnenoteNotebookByIdSectionGroupByIdSectionsCompleteMatchingPredicate(ctx, id, models.OnenoteSectionCollectionResponseOperationPredicate{})
}

// ListMeOnenoteNotebookByIdSectionGroupByIdSectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOnenoteNotebookSectionGroupSectionClient) ListMeOnenoteNotebookByIdSectionGroupByIdSectionsCompleteMatchingPredicate(ctx context.Context, id MeOnenoteNotebookSectionGroupId, predicate models.OnenoteSectionCollectionResponseOperationPredicate) (result ListMeOnenoteNotebookByIdSectionGroupByIdSectionsCompleteResult, err error) {
	items := make([]models.OnenoteSectionCollectionResponse, 0)

	resp, err := c.ListMeOnenoteNotebookByIdSectionGroupByIdSections(ctx, id)
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

	result = ListMeOnenoteNotebookByIdSectionGroupByIdSectionsCompleteResult{
		Items: items,
	}
	return
}
