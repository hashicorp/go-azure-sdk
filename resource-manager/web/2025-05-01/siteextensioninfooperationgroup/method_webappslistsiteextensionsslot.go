package siteextensioninfooperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListSiteExtensionsSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SiteExtensionInfo
}

type WebAppsListSiteExtensionsSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SiteExtensionInfo
}

type WebAppsListSiteExtensionsSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListSiteExtensionsSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListSiteExtensionsSlot ...
func (c SiteExtensionInfoOperationGroupClient) WebAppsListSiteExtensionsSlot(ctx context.Context, id SlotId) (result WebAppsListSiteExtensionsSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListSiteExtensionsSlotCustomPager{},
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

// WebAppsListSiteExtensionsSlotComplete retrieves all the results into a single object
func (c SiteExtensionInfoOperationGroupClient) WebAppsListSiteExtensionsSlotComplete(ctx context.Context, id SlotId) (WebAppsListSiteExtensionsSlotCompleteResult, error) {
	return c.WebAppsListSiteExtensionsSlotCompleteMatchingPredicate(ctx, id, SiteExtensionInfoOperationPredicate{})
}

// WebAppsListSiteExtensionsSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteExtensionInfoOperationGroupClient) WebAppsListSiteExtensionsSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate SiteExtensionInfoOperationPredicate) (result WebAppsListSiteExtensionsSlotCompleteResult, err error) {
	items := make([]SiteExtensionInfo, 0)

	resp, err := c.WebAppsListSiteExtensionsSlot(ctx, id)
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

	result = WebAppsListSiteExtensionsSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
