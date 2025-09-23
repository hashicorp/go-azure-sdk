package groupquotasentities

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupQuotaUsagesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ResourceUsages
}

type GroupQuotaUsagesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ResourceUsages
}

type GroupQuotaUsagesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GroupQuotaUsagesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GroupQuotaUsagesList ...
func (c GroupQuotasEntitiesClient) GroupQuotaUsagesList(ctx context.Context, id LocationUsageId) (result GroupQuotaUsagesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &GroupQuotaUsagesListCustomPager{},
		Path:       id.ID(),
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
		Values *[]ResourceUsages `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GroupQuotaUsagesListComplete retrieves all the results into a single object
func (c GroupQuotasEntitiesClient) GroupQuotaUsagesListComplete(ctx context.Context, id LocationUsageId) (GroupQuotaUsagesListCompleteResult, error) {
	return c.GroupQuotaUsagesListCompleteMatchingPredicate(ctx, id, ResourceUsagesOperationPredicate{})
}

// GroupQuotaUsagesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupQuotasEntitiesClient) GroupQuotaUsagesListCompleteMatchingPredicate(ctx context.Context, id LocationUsageId, predicate ResourceUsagesOperationPredicate) (result GroupQuotaUsagesListCompleteResult, err error) {
	items := make([]ResourceUsages, 0)

	resp, err := c.GroupQuotaUsagesList(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = GroupQuotaUsagesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
