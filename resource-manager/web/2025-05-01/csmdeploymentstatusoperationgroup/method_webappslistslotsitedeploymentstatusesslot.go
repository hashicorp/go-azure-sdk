package csmdeploymentstatusoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListSlotSiteDeploymentStatusesSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CsmDeploymentStatus
}

type WebAppsListSlotSiteDeploymentStatusesSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []CsmDeploymentStatus
}

type WebAppsListSlotSiteDeploymentStatusesSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListSlotSiteDeploymentStatusesSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListSlotSiteDeploymentStatusesSlot ...
func (c CsmDeploymentStatusOperationGroupClient) WebAppsListSlotSiteDeploymentStatusesSlot(ctx context.Context, id SlotId) (result WebAppsListSlotSiteDeploymentStatusesSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListSlotSiteDeploymentStatusesSlotCustomPager{},
		Path:       fmt.Sprintf("%s/deploymentStatus", id.ID()),
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
		Values *[]CsmDeploymentStatus `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListSlotSiteDeploymentStatusesSlotComplete retrieves all the results into a single object
func (c CsmDeploymentStatusOperationGroupClient) WebAppsListSlotSiteDeploymentStatusesSlotComplete(ctx context.Context, id SlotId) (WebAppsListSlotSiteDeploymentStatusesSlotCompleteResult, error) {
	return c.WebAppsListSlotSiteDeploymentStatusesSlotCompleteMatchingPredicate(ctx, id, CsmDeploymentStatusOperationPredicate{})
}

// WebAppsListSlotSiteDeploymentStatusesSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CsmDeploymentStatusOperationGroupClient) WebAppsListSlotSiteDeploymentStatusesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate CsmDeploymentStatusOperationPredicate) (result WebAppsListSlotSiteDeploymentStatusesSlotCompleteResult, err error) {
	items := make([]CsmDeploymentStatus, 0)

	resp, err := c.WebAppsListSlotSiteDeploymentStatusesSlot(ctx, id)
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

	result = WebAppsListSlotSiteDeploymentStatusesSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
