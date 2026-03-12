package csmdeploymentstatuses

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

type WebAppsListProductionSiteDeploymentStatusesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CsmDeploymentStatus
}

type WebAppsListProductionSiteDeploymentStatusesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []CsmDeploymentStatus
}

type WebAppsListProductionSiteDeploymentStatusesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListProductionSiteDeploymentStatusesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListProductionSiteDeploymentStatuses ...
func (c CsmDeploymentStatusesClient) WebAppsListProductionSiteDeploymentStatuses(ctx context.Context, id commonids.AppServiceId) (result WebAppsListProductionSiteDeploymentStatusesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListProductionSiteDeploymentStatusesCustomPager{},
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

// WebAppsListProductionSiteDeploymentStatusesComplete retrieves all the results into a single object
func (c CsmDeploymentStatusesClient) WebAppsListProductionSiteDeploymentStatusesComplete(ctx context.Context, id commonids.AppServiceId) (WebAppsListProductionSiteDeploymentStatusesCompleteResult, error) {
	return c.WebAppsListProductionSiteDeploymentStatusesCompleteMatchingPredicate(ctx, id, CsmDeploymentStatusOperationPredicate{})
}

// WebAppsListProductionSiteDeploymentStatusesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CsmDeploymentStatusesClient) WebAppsListProductionSiteDeploymentStatusesCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate CsmDeploymentStatusOperationPredicate) (result WebAppsListProductionSiteDeploymentStatusesCompleteResult, err error) {
	items := make([]CsmDeploymentStatus, 0)

	resp, err := c.WebAppsListProductionSiteDeploymentStatuses(ctx, id)
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

	result = WebAppsListProductionSiteDeploymentStatusesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
