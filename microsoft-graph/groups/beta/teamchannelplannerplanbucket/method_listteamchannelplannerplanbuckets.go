package teamchannelplannerplanbucket

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

type ListTeamChannelPlannerPlanBucketsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PlannerBucket
}

type ListTeamChannelPlannerPlanBucketsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PlannerBucket
}

type ListTeamChannelPlannerPlanBucketsOperationOptions struct {
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

func DefaultListTeamChannelPlannerPlanBucketsOperationOptions() ListTeamChannelPlannerPlanBucketsOperationOptions {
	return ListTeamChannelPlannerPlanBucketsOperationOptions{}
}

func (o ListTeamChannelPlannerPlanBucketsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTeamChannelPlannerPlanBucketsOperationOptions) ToOData() *odata.Query {
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

func (o ListTeamChannelPlannerPlanBucketsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTeamChannelPlannerPlanBucketsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamChannelPlannerPlanBucketsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamChannelPlannerPlanBuckets - Get buckets from groups. Collection of buckets in the plan. Read-only. Nullable.
func (c TeamChannelPlannerPlanBucketClient) ListTeamChannelPlannerPlanBuckets(ctx context.Context, id beta.GroupIdTeamChannelIdPlannerPlanId, options ListTeamChannelPlannerPlanBucketsOperationOptions) (result ListTeamChannelPlannerPlanBucketsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTeamChannelPlannerPlanBucketsCustomPager{},
		Path:          fmt.Sprintf("%s/buckets", id.ID()),
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
		Values *[]beta.PlannerBucket `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamChannelPlannerPlanBucketsComplete retrieves all the results into a single object
func (c TeamChannelPlannerPlanBucketClient) ListTeamChannelPlannerPlanBucketsComplete(ctx context.Context, id beta.GroupIdTeamChannelIdPlannerPlanId, options ListTeamChannelPlannerPlanBucketsOperationOptions) (ListTeamChannelPlannerPlanBucketsCompleteResult, error) {
	return c.ListTeamChannelPlannerPlanBucketsCompleteMatchingPredicate(ctx, id, options, PlannerBucketOperationPredicate{})
}

// ListTeamChannelPlannerPlanBucketsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamChannelPlannerPlanBucketClient) ListTeamChannelPlannerPlanBucketsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdTeamChannelIdPlannerPlanId, options ListTeamChannelPlannerPlanBucketsOperationOptions, predicate PlannerBucketOperationPredicate) (result ListTeamChannelPlannerPlanBucketsCompleteResult, err error) {
	items := make([]beta.PlannerBucket, 0)

	resp, err := c.ListTeamChannelPlannerPlanBuckets(ctx, id, options)
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

	result = ListTeamChannelPlannerPlanBucketsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
