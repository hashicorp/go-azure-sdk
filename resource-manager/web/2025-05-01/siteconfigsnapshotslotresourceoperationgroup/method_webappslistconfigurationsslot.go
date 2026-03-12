package siteconfigsnapshotslotresourceoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListConfigurationsSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SiteConfigResource
}

type WebAppsListConfigurationsSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SiteConfigResource
}

type WebAppsListConfigurationsSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListConfigurationsSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListConfigurationsSlot ...
func (c SiteConfigSnapshotSlotResourceOperationGroupClient) WebAppsListConfigurationsSlot(ctx context.Context, id SlotId) (result WebAppsListConfigurationsSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListConfigurationsSlotCustomPager{},
		Path:       fmt.Sprintf("%s/config", id.ID()),
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
		Values *[]SiteConfigResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListConfigurationsSlotComplete retrieves all the results into a single object
func (c SiteConfigSnapshotSlotResourceOperationGroupClient) WebAppsListConfigurationsSlotComplete(ctx context.Context, id SlotId) (WebAppsListConfigurationsSlotCompleteResult, error) {
	return c.WebAppsListConfigurationsSlotCompleteMatchingPredicate(ctx, id, SiteConfigResourceOperationPredicate{})
}

// WebAppsListConfigurationsSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteConfigSnapshotSlotResourceOperationGroupClient) WebAppsListConfigurationsSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate SiteConfigResourceOperationPredicate) (result WebAppsListConfigurationsSlotCompleteResult, err error) {
	items := make([]SiteConfigResource, 0)

	resp, err := c.WebAppsListConfigurationsSlot(ctx, id)
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

	result = WebAppsListConfigurationsSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
