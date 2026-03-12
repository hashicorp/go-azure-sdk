package sitecontainers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListSiteContainersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SiteContainer
}

type WebAppsListSiteContainersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SiteContainer
}

type WebAppsListSiteContainersCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListSiteContainersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListSiteContainers ...
func (c SiteContainersClient) WebAppsListSiteContainers(ctx context.Context, id commonids.AppServiceId) (result WebAppsListSiteContainersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListSiteContainersCustomPager{},
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

// WebAppsListSiteContainersComplete retrieves all the results into a single object
func (c SiteContainersClient) WebAppsListSiteContainersComplete(ctx context.Context, id commonids.AppServiceId) (WebAppsListSiteContainersCompleteResult, error) {
	return c.WebAppsListSiteContainersCompleteMatchingPredicate(ctx, id, SiteContainerOperationPredicate{})
}

// WebAppsListSiteContainersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteContainersClient) WebAppsListSiteContainersCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate SiteContainerOperationPredicate) (result WebAppsListSiteContainersCompleteResult, err error) {
	items := make([]SiteContainer, 0)

	resp, err := c.WebAppsListSiteContainers(ctx, id)
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

	result = WebAppsListSiteContainersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
