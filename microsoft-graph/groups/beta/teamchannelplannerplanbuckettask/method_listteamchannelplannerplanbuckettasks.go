package teamchannelplannerplanbuckettask

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListTeamChannelPlannerPlanBucketTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PlannerTask
}

type ListTeamChannelPlannerPlanBucketTasksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PlannerTask
}

type ListTeamChannelPlannerPlanBucketTasksOperationOptions struct {
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

func DefaultListTeamChannelPlannerPlanBucketTasksOperationOptions() ListTeamChannelPlannerPlanBucketTasksOperationOptions {
	return ListTeamChannelPlannerPlanBucketTasksOperationOptions{}
}

func (o ListTeamChannelPlannerPlanBucketTasksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTeamChannelPlannerPlanBucketTasksOperationOptions) ToOData() *odata.Query {
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

func (o ListTeamChannelPlannerPlanBucketTasksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTeamChannelPlannerPlanBucketTasksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamChannelPlannerPlanBucketTasksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamChannelPlannerPlanBucketTasks - Get tasks from groups. Read-only. Nullable. The collection of tasks in the
// bucket.
func (c TeamChannelPlannerPlanBucketTaskClient) ListTeamChannelPlannerPlanBucketTasks(ctx context.Context, id beta.GroupIdTeamChannelIdPlannerPlanIdBucketId, options ListTeamChannelPlannerPlanBucketTasksOperationOptions) (result ListTeamChannelPlannerPlanBucketTasksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTeamChannelPlannerPlanBucketTasksCustomPager{},
		Path:          fmt.Sprintf("%s/tasks", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.PlannerTask, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalPlannerTaskImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.PlannerTask (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListTeamChannelPlannerPlanBucketTasksComplete retrieves all the results into a single object
func (c TeamChannelPlannerPlanBucketTaskClient) ListTeamChannelPlannerPlanBucketTasksComplete(ctx context.Context, id beta.GroupIdTeamChannelIdPlannerPlanIdBucketId, options ListTeamChannelPlannerPlanBucketTasksOperationOptions) (ListTeamChannelPlannerPlanBucketTasksCompleteResult, error) {
	return c.ListTeamChannelPlannerPlanBucketTasksCompleteMatchingPredicate(ctx, id, options, PlannerTaskOperationPredicate{})
}

// ListTeamChannelPlannerPlanBucketTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamChannelPlannerPlanBucketTaskClient) ListTeamChannelPlannerPlanBucketTasksCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdTeamChannelIdPlannerPlanIdBucketId, options ListTeamChannelPlannerPlanBucketTasksOperationOptions, predicate PlannerTaskOperationPredicate) (result ListTeamChannelPlannerPlanBucketTasksCompleteResult, err error) {
	items := make([]beta.PlannerTask, 0)

	resp, err := c.ListTeamChannelPlannerPlanBucketTasks(ctx, id, options)
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

	result = ListTeamChannelPlannerPlanBucketTasksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
