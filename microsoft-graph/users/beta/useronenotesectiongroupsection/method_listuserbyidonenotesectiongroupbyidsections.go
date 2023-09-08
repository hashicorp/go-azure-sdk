package useronenotesectiongroupsection

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

type ListUserByIdOnenoteSectionGroupByIdSectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenoteSectionCollectionResponse
}

type ListUserByIdOnenoteSectionGroupByIdSectionsCompleteResult struct {
	Items []models.OnenoteSectionCollectionResponse
}

// ListUserByIdOnenoteSectionGroupByIdSections ...
func (c UserOnenoteSectionGroupSectionClient) ListUserByIdOnenoteSectionGroupByIdSections(ctx context.Context, id UserOnenoteSectionGroupId) (result ListUserByIdOnenoteSectionGroupByIdSectionsOperationResponse, err error) {
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

// ListUserByIdOnenoteSectionGroupByIdSectionsComplete retrieves all the results into a single object
func (c UserOnenoteSectionGroupSectionClient) ListUserByIdOnenoteSectionGroupByIdSectionsComplete(ctx context.Context, id UserOnenoteSectionGroupId) (ListUserByIdOnenoteSectionGroupByIdSectionsCompleteResult, error) {
	return c.ListUserByIdOnenoteSectionGroupByIdSectionsCompleteMatchingPredicate(ctx, id, models.OnenoteSectionCollectionResponseOperationPredicate{})
}

// ListUserByIdOnenoteSectionGroupByIdSectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOnenoteSectionGroupSectionClient) ListUserByIdOnenoteSectionGroupByIdSectionsCompleteMatchingPredicate(ctx context.Context, id UserOnenoteSectionGroupId, predicate models.OnenoteSectionCollectionResponseOperationPredicate) (result ListUserByIdOnenoteSectionGroupByIdSectionsCompleteResult, err error) {
	items := make([]models.OnenoteSectionCollectionResponse, 0)

	resp, err := c.ListUserByIdOnenoteSectionGroupByIdSections(ctx, id)
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

	result = ListUserByIdOnenoteSectionGroupByIdSectionsCompleteResult{
		Items: items,
	}
	return
}
