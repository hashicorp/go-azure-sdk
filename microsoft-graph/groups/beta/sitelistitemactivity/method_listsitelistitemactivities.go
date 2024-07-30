package sitelistitemactivity

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

type ListSiteListItemActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ItemActivityOLD
}

type ListSiteListItemActivitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ItemActivityOLD
}

type ListSiteListItemActivitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteListItemActivitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteListItemActivities ...
func (c SiteListItemActivityClient) ListSiteListItemActivities(ctx context.Context, id GroupIdSiteIdListIdItemId) (result ListSiteListItemActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteListItemActivitiesCustomPager{},
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
		Values *[]beta.ItemActivityOLD `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteListItemActivitiesComplete retrieves all the results into a single object
func (c SiteListItemActivityClient) ListSiteListItemActivitiesComplete(ctx context.Context, id GroupIdSiteIdListIdItemId) (ListSiteListItemActivitiesCompleteResult, error) {
	return c.ListSiteListItemActivitiesCompleteMatchingPredicate(ctx, id, ItemActivityOLDOperationPredicate{})
}

// ListSiteListItemActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteListItemActivityClient) ListSiteListItemActivitiesCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteIdListIdItemId, predicate ItemActivityOLDOperationPredicate) (result ListSiteListItemActivitiesCompleteResult, err error) {
	items := make([]beta.ItemActivityOLD, 0)

	resp, err := c.ListSiteListItemActivities(ctx, id)
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

	result = ListSiteListItemActivitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
