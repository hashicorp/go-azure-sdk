package plannerplanbucket

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPlannerPlanBucketsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PlannerBucket
}

type ListPlannerPlanBucketsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PlannerBucket
}

type ListPlannerPlanBucketsOperationOptions struct {
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

func DefaultListPlannerPlanBucketsOperationOptions() ListPlannerPlanBucketsOperationOptions {
	return ListPlannerPlanBucketsOperationOptions{}
}

func (o ListPlannerPlanBucketsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPlannerPlanBucketsOperationOptions) ToOData() *odata.Query {
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

func (o ListPlannerPlanBucketsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPlannerPlanBucketsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPlannerPlanBucketsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPlannerPlanBuckets - Get buckets from me. Collection of buckets in the plan. Read-only. Nullable.
func (c PlannerPlanBucketClient) ListPlannerPlanBuckets(ctx context.Context, id beta.MePlannerPlanId, options ListPlannerPlanBucketsOperationOptions) (result ListPlannerPlanBucketsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPlannerPlanBucketsCustomPager{},
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

// ListPlannerPlanBucketsComplete retrieves all the results into a single object
func (c PlannerPlanBucketClient) ListPlannerPlanBucketsComplete(ctx context.Context, id beta.MePlannerPlanId, options ListPlannerPlanBucketsOperationOptions) (ListPlannerPlanBucketsCompleteResult, error) {
	return c.ListPlannerPlanBucketsCompleteMatchingPredicate(ctx, id, options, PlannerBucketOperationPredicate{})
}

// ListPlannerPlanBucketsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PlannerPlanBucketClient) ListPlannerPlanBucketsCompleteMatchingPredicate(ctx context.Context, id beta.MePlannerPlanId, options ListPlannerPlanBucketsOperationOptions, predicate PlannerBucketOperationPredicate) (result ListPlannerPlanBucketsCompleteResult, err error) {
	items := make([]beta.PlannerBucket, 0)

	resp, err := c.ListPlannerPlanBuckets(ctx, id, options)
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

	result = ListPlannerPlanBucketsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
