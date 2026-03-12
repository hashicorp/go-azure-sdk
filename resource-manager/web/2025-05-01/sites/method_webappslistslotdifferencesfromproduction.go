package sites

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

type WebAppsListSlotDifferencesFromProductionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SlotDifference
}

type WebAppsListSlotDifferencesFromProductionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SlotDifference
}

type WebAppsListSlotDifferencesFromProductionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListSlotDifferencesFromProductionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListSlotDifferencesFromProduction ...
func (c SitesClient) WebAppsListSlotDifferencesFromProduction(ctx context.Context, id commonids.AppServiceId, input CsmSlotEntity) (result WebAppsListSlotDifferencesFromProductionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &WebAppsListSlotDifferencesFromProductionCustomPager{},
		Path:       fmt.Sprintf("%s/slotsdiffs", id.ID()),
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
		Values *[]SlotDifference `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListSlotDifferencesFromProductionComplete retrieves all the results into a single object
func (c SitesClient) WebAppsListSlotDifferencesFromProductionComplete(ctx context.Context, id commonids.AppServiceId, input CsmSlotEntity) (WebAppsListSlotDifferencesFromProductionCompleteResult, error) {
	return c.WebAppsListSlotDifferencesFromProductionCompleteMatchingPredicate(ctx, id, input, SlotDifferenceOperationPredicate{})
}

// WebAppsListSlotDifferencesFromProductionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SitesClient) WebAppsListSlotDifferencesFromProductionCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, input CsmSlotEntity, predicate SlotDifferenceOperationPredicate) (result WebAppsListSlotDifferencesFromProductionCompleteResult, err error) {
	items := make([]SlotDifference, 0)

	resp, err := c.WebAppsListSlotDifferencesFromProduction(ctx, id, input)
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

	result = WebAppsListSlotDifferencesFromProductionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
