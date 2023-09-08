package meonenotesection

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

type ListMeOnenoteSectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenoteSectionCollectionResponse
}

type ListMeOnenoteSectionsCompleteResult struct {
	Items []models.OnenoteSectionCollectionResponse
}

// ListMeOnenoteSections ...
func (c MeOnenoteSectionClient) ListMeOnenoteSections(ctx context.Context) (result ListMeOnenoteSectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/onenote/sections",
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

// ListMeOnenoteSectionsComplete retrieves all the results into a single object
func (c MeOnenoteSectionClient) ListMeOnenoteSectionsComplete(ctx context.Context) (ListMeOnenoteSectionsCompleteResult, error) {
	return c.ListMeOnenoteSectionsCompleteMatchingPredicate(ctx, models.OnenoteSectionCollectionResponseOperationPredicate{})
}

// ListMeOnenoteSectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOnenoteSectionClient) ListMeOnenoteSectionsCompleteMatchingPredicate(ctx context.Context, predicate models.OnenoteSectionCollectionResponseOperationPredicate) (result ListMeOnenoteSectionsCompleteResult, err error) {
	items := make([]models.OnenoteSectionCollectionResponse, 0)

	resp, err := c.ListMeOnenoteSections(ctx)
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

	result = ListMeOnenoteSectionsCompleteResult{
		Items: items,
	}
	return
}
