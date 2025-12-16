package monitorresources

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

type MonitorsListBySubscriptionIdOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]MonitorResource
}

type MonitorsListBySubscriptionIdCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []MonitorResource
}

type MonitorsListBySubscriptionIdCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *MonitorsListBySubscriptionIdCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// MonitorsListBySubscriptionId ...
func (c MonitorResourcesClient) MonitorsListBySubscriptionId(ctx context.Context, id commonids.SubscriptionId) (result MonitorsListBySubscriptionIdOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &MonitorsListBySubscriptionIdCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Dynatrace.Observability/monitors", id.ID()),
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
		Values *[]MonitorResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// MonitorsListBySubscriptionIdComplete retrieves all the results into a single object
func (c MonitorResourcesClient) MonitorsListBySubscriptionIdComplete(ctx context.Context, id commonids.SubscriptionId) (MonitorsListBySubscriptionIdCompleteResult, error) {
	return c.MonitorsListBySubscriptionIdCompleteMatchingPredicate(ctx, id, MonitorResourceOperationPredicate{})
}

// MonitorsListBySubscriptionIdCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MonitorResourcesClient) MonitorsListBySubscriptionIdCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate MonitorResourceOperationPredicate) (result MonitorsListBySubscriptionIdCompleteResult, err error) {
	items := make([]MonitorResource, 0)

	resp, err := c.MonitorsListBySubscriptionId(ctx, id)
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

	result = MonitorsListBySubscriptionIdCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
