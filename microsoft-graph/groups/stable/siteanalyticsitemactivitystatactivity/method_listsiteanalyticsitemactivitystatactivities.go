package siteanalyticsitemactivitystatactivity

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSiteAnalyticsItemActivityStatActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ItemActivity
}

type ListSiteAnalyticsItemActivityStatActivitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ItemActivity
}

type ListSiteAnalyticsItemActivityStatActivitiesOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListSiteAnalyticsItemActivityStatActivitiesOperationOptions() ListSiteAnalyticsItemActivityStatActivitiesOperationOptions {
	return ListSiteAnalyticsItemActivityStatActivitiesOperationOptions{}
}

func (o ListSiteAnalyticsItemActivityStatActivitiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteAnalyticsItemActivityStatActivitiesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListSiteAnalyticsItemActivityStatActivitiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteAnalyticsItemActivityStatActivitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteAnalyticsItemActivityStatActivitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteAnalyticsItemActivityStatActivities - Get activities from groups. Exposes the itemActivities represented in
// this itemActivityStat resource.
func (c SiteAnalyticsItemActivityStatActivityClient) ListSiteAnalyticsItemActivityStatActivities(ctx context.Context, id stable.GroupIdSiteIdAnalyticsItemActivityStatId, options ListSiteAnalyticsItemActivityStatActivitiesOperationOptions) (result ListSiteAnalyticsItemActivityStatActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteAnalyticsItemActivityStatActivitiesCustomPager{},
		Path:          fmt.Sprintf("%s/activities", id.ID()),
		RetryFunc:     options.RetryFunc,
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
		Values *[]stable.ItemActivity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteAnalyticsItemActivityStatActivitiesComplete retrieves all the results into a single object
func (c SiteAnalyticsItemActivityStatActivityClient) ListSiteAnalyticsItemActivityStatActivitiesComplete(ctx context.Context, id stable.GroupIdSiteIdAnalyticsItemActivityStatId, options ListSiteAnalyticsItemActivityStatActivitiesOperationOptions) (ListSiteAnalyticsItemActivityStatActivitiesCompleteResult, error) {
	return c.ListSiteAnalyticsItemActivityStatActivitiesCompleteMatchingPredicate(ctx, id, options, ItemActivityOperationPredicate{})
}

// ListSiteAnalyticsItemActivityStatActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteAnalyticsItemActivityStatActivityClient) ListSiteAnalyticsItemActivityStatActivitiesCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdSiteIdAnalyticsItemActivityStatId, options ListSiteAnalyticsItemActivityStatActivitiesOperationOptions, predicate ItemActivityOperationPredicate) (result ListSiteAnalyticsItemActivityStatActivitiesCompleteResult, err error) {
	items := make([]stable.ItemActivity, 0)

	resp, err := c.ListSiteAnalyticsItemActivityStatActivities(ctx, id, options)
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

	result = ListSiteAnalyticsItemActivityStatActivitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
