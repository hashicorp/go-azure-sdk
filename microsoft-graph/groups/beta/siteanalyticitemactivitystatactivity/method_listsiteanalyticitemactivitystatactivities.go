package siteanalyticitemactivitystatactivity

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSiteAnalyticItemActivityStatActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ItemActivity
}

type ListSiteAnalyticItemActivityStatActivitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ItemActivity
}

type ListSiteAnalyticItemActivityStatActivitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteAnalyticItemActivityStatActivitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteAnalyticItemActivityStatActivities ...
func (c SiteAnalyticItemActivityStatActivityClient) ListSiteAnalyticItemActivityStatActivities(ctx context.Context, id GroupIdSiteIdAnalyticItemActivityStatId) (result ListSiteAnalyticItemActivityStatActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteAnalyticItemActivityStatActivitiesCustomPager{},
		Path:       fmt.Sprintf("%s/activities", id.ID()),
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
		Values *[]beta.ItemActivity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteAnalyticItemActivityStatActivitiesComplete retrieves all the results into a single object
func (c SiteAnalyticItemActivityStatActivityClient) ListSiteAnalyticItemActivityStatActivitiesComplete(ctx context.Context, id GroupIdSiteIdAnalyticItemActivityStatId) (ListSiteAnalyticItemActivityStatActivitiesCompleteResult, error) {
	return c.ListSiteAnalyticItemActivityStatActivitiesCompleteMatchingPredicate(ctx, id, ItemActivityOperationPredicate{})
}

// ListSiteAnalyticItemActivityStatActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteAnalyticItemActivityStatActivityClient) ListSiteAnalyticItemActivityStatActivitiesCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteIdAnalyticItemActivityStatId, predicate ItemActivityOperationPredicate) (result ListSiteAnalyticItemActivityStatActivitiesCompleteResult, err error) {
	items := make([]beta.ItemActivity, 0)

	resp, err := c.ListSiteAnalyticItemActivityStatActivities(ctx, id)
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

	result = ListSiteAnalyticItemActivityStatActivitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
