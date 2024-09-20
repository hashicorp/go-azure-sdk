package driverootanalyticsitemactivitystat

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

type ListDriveRootAnalyticsItemActivityStatsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ItemActivityStat
}

type ListDriveRootAnalyticsItemActivityStatsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ItemActivityStat
}

type ListDriveRootAnalyticsItemActivityStatsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListDriveRootAnalyticsItemActivityStatsOperationOptions() ListDriveRootAnalyticsItemActivityStatsOperationOptions {
	return ListDriveRootAnalyticsItemActivityStatsOperationOptions{}
}

func (o ListDriveRootAnalyticsItemActivityStatsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveRootAnalyticsItemActivityStatsOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveRootAnalyticsItemActivityStatsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveRootAnalyticsItemActivityStatsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveRootAnalyticsItemActivityStatsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveRootAnalyticsItemActivityStats - Get itemActivityStats from groups
func (c DriveRootAnalyticsItemActivityStatClient) ListDriveRootAnalyticsItemActivityStats(ctx context.Context, id stable.GroupIdDriveId, options ListDriveRootAnalyticsItemActivityStatsOperationOptions) (result ListDriveRootAnalyticsItemActivityStatsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveRootAnalyticsItemActivityStatsCustomPager{},
		Path:          fmt.Sprintf("%s/root/analytics/itemActivityStats", id.ID()),
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

// ListDriveRootAnalyticsItemActivityStatsComplete retrieves all the results into a single object
func (c DriveRootAnalyticsItemActivityStatClient) ListDriveRootAnalyticsItemActivityStatsComplete(ctx context.Context, id stable.GroupIdDriveId, options ListDriveRootAnalyticsItemActivityStatsOperationOptions) (ListDriveRootAnalyticsItemActivityStatsCompleteResult, error) {
	return c.ListDriveRootAnalyticsItemActivityStatsCompleteMatchingPredicate(ctx, id, options, ItemActivityStatOperationPredicate{})
}

// ListDriveRootAnalyticsItemActivityStatsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveRootAnalyticsItemActivityStatClient) ListDriveRootAnalyticsItemActivityStatsCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdDriveId, options ListDriveRootAnalyticsItemActivityStatsOperationOptions, predicate ItemActivityStatOperationPredicate) (result ListDriveRootAnalyticsItemActivityStatsCompleteResult, err error) {
	items := make([]stable.ItemActivityStat, 0)

	resp, err := c.ListDriveRootAnalyticsItemActivityStats(ctx, id, options)
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

	result = ListDriveRootAnalyticsItemActivityStatsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
