package meonenotesectiongroup

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

type ListMeOnenoteSectionGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SectionGroupCollectionResponse
}

type ListMeOnenoteSectionGroupsCompleteResult struct {
	Items []models.SectionGroupCollectionResponse
}

// ListMeOnenoteSectionGroups ...
func (c MeOnenoteSectionGroupClient) ListMeOnenoteSectionGroups(ctx context.Context) (result ListMeOnenoteSectionGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/onenote/sectionGroups",
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

// ListMeOnenoteSectionGroupsComplete retrieves all the results into a single object
func (c MeOnenoteSectionGroupClient) ListMeOnenoteSectionGroupsComplete(ctx context.Context) (ListMeOnenoteSectionGroupsCompleteResult, error) {
	return c.ListMeOnenoteSectionGroupsCompleteMatchingPredicate(ctx, models.SectionGroupCollectionResponseOperationPredicate{})
}

// ListMeOnenoteSectionGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOnenoteSectionGroupClient) ListMeOnenoteSectionGroupsCompleteMatchingPredicate(ctx context.Context, predicate models.SectionGroupCollectionResponseOperationPredicate) (result ListMeOnenoteSectionGroupsCompleteResult, err error) {
	items := make([]models.SectionGroupCollectionResponse, 0)

	resp, err := c.ListMeOnenoteSectionGroups(ctx)
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

	result = ListMeOnenoteSectionGroupsCompleteResult{
		Items: items,
	}
	return
}
