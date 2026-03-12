package deploymentoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListDeploymentsSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Deployment
}

type WebAppsListDeploymentsSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Deployment
}

type WebAppsListDeploymentsSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListDeploymentsSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListDeploymentsSlot ...
func (c DeploymentOperationGroupClient) WebAppsListDeploymentsSlot(ctx context.Context, id SlotId) (result WebAppsListDeploymentsSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListDeploymentsSlotCustomPager{},
		Path:       fmt.Sprintf("%s/deployments", id.ID()),
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
		Values *[]Deployment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListDeploymentsSlotComplete retrieves all the results into a single object
func (c DeploymentOperationGroupClient) WebAppsListDeploymentsSlotComplete(ctx context.Context, id SlotId) (WebAppsListDeploymentsSlotCompleteResult, error) {
	return c.WebAppsListDeploymentsSlotCompleteMatchingPredicate(ctx, id, DeploymentOperationPredicate{})
}

// WebAppsListDeploymentsSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeploymentOperationGroupClient) WebAppsListDeploymentsSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate DeploymentOperationPredicate) (result WebAppsListDeploymentsSlotCompleteResult, err error) {
	items := make([]Deployment, 0)

	resp, err := c.WebAppsListDeploymentsSlot(ctx, id)
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

	result = WebAppsListDeploymentsSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
