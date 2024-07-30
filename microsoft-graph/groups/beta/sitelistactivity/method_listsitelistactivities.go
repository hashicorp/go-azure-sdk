package sitelistactivity

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

type ListSiteListActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ItemActivityOLD
}

type ListSiteListActivitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ItemActivityOLD
}

type ListSiteListActivitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteListActivitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteListActivities ...
func (c SiteListActivityClient) ListSiteListActivities(ctx context.Context, id GroupIdSiteIdListId) (result ListSiteListActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteListActivitiesCustomPager{},
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

// ListSiteListActivitiesComplete retrieves all the results into a single object
func (c SiteListActivityClient) ListSiteListActivitiesComplete(ctx context.Context, id GroupIdSiteIdListId) (ListSiteListActivitiesCompleteResult, error) {
	return c.ListSiteListActivitiesCompleteMatchingPredicate(ctx, id, ItemActivityOLDOperationPredicate{})
}

// ListSiteListActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteListActivityClient) ListSiteListActivitiesCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteIdListId, predicate ItemActivityOLDOperationPredicate) (result ListSiteListActivitiesCompleteResult, err error) {
	items := make([]beta.ItemActivityOLD, 0)

	resp, err := c.ListSiteListActivities(ctx, id)
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

	result = ListSiteListActivitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
