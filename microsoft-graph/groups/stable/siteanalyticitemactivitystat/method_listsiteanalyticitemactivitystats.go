package siteanalyticitemactivitystat

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSiteAnalyticItemActivityStatsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ItemActivityStat
}

type ListSiteAnalyticItemActivityStatsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ItemActivityStat
}

type ListSiteAnalyticItemActivityStatsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteAnalyticItemActivityStatsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteAnalyticItemActivityStats ...
func (c SiteAnalyticItemActivityStatClient) ListSiteAnalyticItemActivityStats(ctx context.Context, id GroupIdSiteId) (result ListSiteAnalyticItemActivityStatsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteAnalyticItemActivityStatsCustomPager{},
		Path:       fmt.Sprintf("%s/analytics/itemActivityStats", id.ID()),
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
		Values *[]stable.ItemActivityStat `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteAnalyticItemActivityStatsComplete retrieves all the results into a single object
func (c SiteAnalyticItemActivityStatClient) ListSiteAnalyticItemActivityStatsComplete(ctx context.Context, id GroupIdSiteId) (ListSiteAnalyticItemActivityStatsCompleteResult, error) {
	return c.ListSiteAnalyticItemActivityStatsCompleteMatchingPredicate(ctx, id, ItemActivityStatOperationPredicate{})
}

// ListSiteAnalyticItemActivityStatsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteAnalyticItemActivityStatClient) ListSiteAnalyticItemActivityStatsCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteId, predicate ItemActivityStatOperationPredicate) (result ListSiteAnalyticItemActivityStatsCompleteResult, err error) {
	items := make([]stable.ItemActivityStat, 0)

	resp, err := c.ListSiteAnalyticItemActivityStats(ctx, id)
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

	result = ListSiteAnalyticItemActivityStatsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
