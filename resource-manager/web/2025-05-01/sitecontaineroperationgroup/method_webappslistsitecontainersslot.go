package sitecontaineroperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListSiteContainersSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SiteContainer
}

type WebAppsListSiteContainersSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SiteContainer
}

type WebAppsListSiteContainersSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListSiteContainersSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListSiteContainersSlot ...
func (c SiteContainerOperationGroupClient) WebAppsListSiteContainersSlot(ctx context.Context, id SlotId) (result WebAppsListSiteContainersSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListSiteContainersSlotCustomPager{},
		Path:       fmt.Sprintf("%s/sitecontainers", id.ID()),
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
		Values *[]SiteContainer `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListSiteContainersSlotComplete retrieves all the results into a single object
func (c SiteContainerOperationGroupClient) WebAppsListSiteContainersSlotComplete(ctx context.Context, id SlotId) (WebAppsListSiteContainersSlotCompleteResult, error) {
	return c.WebAppsListSiteContainersSlotCompleteMatchingPredicate(ctx, id, SiteContainerOperationPredicate{})
}

// WebAppsListSiteContainersSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteContainerOperationGroupClient) WebAppsListSiteContainersSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate SiteContainerOperationPredicate) (result WebAppsListSiteContainersSlotCompleteResult, err error) {
	items := make([]SiteContainer, 0)

	resp, err := c.WebAppsListSiteContainersSlot(ctx, id)
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

	result = WebAppsListSiteContainersSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
