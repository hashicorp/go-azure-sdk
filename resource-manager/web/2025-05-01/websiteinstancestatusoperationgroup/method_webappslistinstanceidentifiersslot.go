package websiteinstancestatusoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListInstanceIdentifiersSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WebSiteInstanceStatus
}

type WebAppsListInstanceIdentifiersSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WebSiteInstanceStatus
}

type WebAppsListInstanceIdentifiersSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListInstanceIdentifiersSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListInstanceIdentifiersSlot ...
func (c WebSiteInstanceStatusOperationGroupClient) WebAppsListInstanceIdentifiersSlot(ctx context.Context, id SlotId) (result WebAppsListInstanceIdentifiersSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListInstanceIdentifiersSlotCustomPager{},
		Path:       fmt.Sprintf("%s/instances", id.ID()),
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
		Values *[]WebSiteInstanceStatus `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListInstanceIdentifiersSlotComplete retrieves all the results into a single object
func (c WebSiteInstanceStatusOperationGroupClient) WebAppsListInstanceIdentifiersSlotComplete(ctx context.Context, id SlotId) (WebAppsListInstanceIdentifiersSlotCompleteResult, error) {
	return c.WebAppsListInstanceIdentifiersSlotCompleteMatchingPredicate(ctx, id, WebSiteInstanceStatusOperationPredicate{})
}

// WebAppsListInstanceIdentifiersSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WebSiteInstanceStatusOperationGroupClient) WebAppsListInstanceIdentifiersSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate WebSiteInstanceStatusOperationPredicate) (result WebAppsListInstanceIdentifiersSlotCompleteResult, err error) {
	items := make([]WebSiteInstanceStatus, 0)

	resp, err := c.WebAppsListInstanceIdentifiersSlot(ctx, id)
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

	result = WebAppsListInstanceIdentifiersSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
