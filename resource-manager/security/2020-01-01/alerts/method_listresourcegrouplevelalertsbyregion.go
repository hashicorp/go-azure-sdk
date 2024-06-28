package alerts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListResourceGroupLevelAlertsByRegionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Alert
}

type ListResourceGroupLevelAlertsByRegionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Alert
}

type ListResourceGroupLevelAlertsByRegionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListResourceGroupLevelAlertsByRegionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListResourceGroupLevelAlertsByRegion ...
func (c AlertsClient) ListResourceGroupLevelAlertsByRegion(ctx context.Context, id ProviderLocationId) (result ListResourceGroupLevelAlertsByRegionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListResourceGroupLevelAlertsByRegionCustomPager{},
		Path:       fmt.Sprintf("%s/alerts", id.ID()),
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
		Values *[]Alert `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListResourceGroupLevelAlertsByRegionComplete retrieves all the results into a single object
func (c AlertsClient) ListResourceGroupLevelAlertsByRegionComplete(ctx context.Context, id ProviderLocationId) (ListResourceGroupLevelAlertsByRegionCompleteResult, error) {
	return c.ListResourceGroupLevelAlertsByRegionCompleteMatchingPredicate(ctx, id, AlertOperationPredicate{})
}

// ListResourceGroupLevelAlertsByRegionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AlertsClient) ListResourceGroupLevelAlertsByRegionCompleteMatchingPredicate(ctx context.Context, id ProviderLocationId, predicate AlertOperationPredicate) (result ListResourceGroupLevelAlertsByRegionCompleteResult, err error) {
	items := make([]Alert, 0)

	resp, err := c.ListResourceGroupLevelAlertsByRegion(ctx, id)
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

	result = ListResourceGroupLevelAlertsByRegionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
