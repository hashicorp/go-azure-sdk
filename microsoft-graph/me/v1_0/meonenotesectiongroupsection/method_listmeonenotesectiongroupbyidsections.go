package meonenotesectiongroupsection

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

type ListMeOnenoteSectionGroupByIdSectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenoteSectionCollectionResponse
}

type ListMeOnenoteSectionGroupByIdSectionsCompleteResult struct {
	Items []models.OnenoteSectionCollectionResponse
}

// ListMeOnenoteSectionGroupByIdSections ...
func (c MeOnenoteSectionGroupSectionClient) ListMeOnenoteSectionGroupByIdSections(ctx context.Context, id MeOnenoteSectionGroupId) (result ListMeOnenoteSectionGroupByIdSectionsOperationResponse, err error) {
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

// ListMeOnenoteSectionGroupByIdSectionsComplete retrieves all the results into a single object
func (c MeOnenoteSectionGroupSectionClient) ListMeOnenoteSectionGroupByIdSectionsComplete(ctx context.Context, id MeOnenoteSectionGroupId) (ListMeOnenoteSectionGroupByIdSectionsCompleteResult, error) {
	return c.ListMeOnenoteSectionGroupByIdSectionsCompleteMatchingPredicate(ctx, id, models.OnenoteSectionCollectionResponseOperationPredicate{})
}

// ListMeOnenoteSectionGroupByIdSectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOnenoteSectionGroupSectionClient) ListMeOnenoteSectionGroupByIdSectionsCompleteMatchingPredicate(ctx context.Context, id MeOnenoteSectionGroupId, predicate models.OnenoteSectionCollectionResponseOperationPredicate) (result ListMeOnenoteSectionGroupByIdSectionsCompleteResult, err error) {
	items := make([]models.OnenoteSectionCollectionResponse, 0)

	resp, err := c.ListMeOnenoteSectionGroupByIdSections(ctx, id)
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

	result = ListMeOnenoteSectionGroupByIdSectionsCompleteResult{
		Items: items,
	}
	return
}
