package siteextensioninfos

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

type WebAppsListSiteExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SiteExtensionInfo
}

type WebAppsListSiteExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SiteExtensionInfo
}

type WebAppsListSiteExtensionsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListSiteExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListSiteExtensions ...
func (c SiteExtensionInfosClient) WebAppsListSiteExtensions(ctx context.Context, id commonids.AppServiceId) (result WebAppsListSiteExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListSiteExtensionsCustomPager{},
		Path:       fmt.Sprintf("%s/siteExtensions", id.ID()),
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
		Values *[]SiteExtensionInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListSiteExtensionsComplete retrieves all the results into a single object
func (c SiteExtensionInfosClient) WebAppsListSiteExtensionsComplete(ctx context.Context, id commonids.AppServiceId) (WebAppsListSiteExtensionsCompleteResult, error) {
	return c.WebAppsListSiteExtensionsCompleteMatchingPredicate(ctx, id, SiteExtensionInfoOperationPredicate{})
}

// WebAppsListSiteExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteExtensionInfosClient) WebAppsListSiteExtensionsCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate SiteExtensionInfoOperationPredicate) (result WebAppsListSiteExtensionsCompleteResult, err error) {
	items := make([]SiteExtensionInfo, 0)

	resp, err := c.WebAppsListSiteExtensions(ctx, id)
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

	result = WebAppsListSiteExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
