package teamprimarychannelplannerplantask

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

type ListTeamPrimaryChannelPlannerPlanTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PlannerTask
}

type ListTeamPrimaryChannelPlannerPlanTasksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PlannerTask
}

type ListTeamPrimaryChannelPlannerPlanTasksOperationOptions struct {
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

func DefaultListTeamPrimaryChannelPlannerPlanTasksOperationOptions() ListTeamPrimaryChannelPlannerPlanTasksOperationOptions {
	return ListTeamPrimaryChannelPlannerPlanTasksOperationOptions{}
}

func (o ListTeamPrimaryChannelPlannerPlanTasksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTeamPrimaryChannelPlannerPlanTasksOperationOptions) ToOData() *odata.Query {
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

func (o ListTeamPrimaryChannelPlannerPlanTasksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTeamPrimaryChannelPlannerPlanTasksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamPrimaryChannelPlannerPlanTasksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamPrimaryChannelPlannerPlanTasks - Get tasks from groups. Collection of tasks in the plan. Read-only. Nullable.
func (c TeamPrimaryChannelPlannerPlanTaskClient) ListTeamPrimaryChannelPlannerPlanTasks(ctx context.Context, id beta.GroupIdTeamPrimaryChannelPlannerPlanId, options ListTeamPrimaryChannelPlannerPlanTasksOperationOptions) (result ListTeamPrimaryChannelPlannerPlanTasksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTeamPrimaryChannelPlannerPlanTasksCustomPager{},
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

// ListTeamPrimaryChannelPlannerPlanTasksComplete retrieves all the results into a single object
func (c TeamPrimaryChannelPlannerPlanTaskClient) ListTeamPrimaryChannelPlannerPlanTasksComplete(ctx context.Context, id beta.GroupIdTeamPrimaryChannelPlannerPlanId, options ListTeamPrimaryChannelPlannerPlanTasksOperationOptions) (ListTeamPrimaryChannelPlannerPlanTasksCompleteResult, error) {
	return c.ListTeamPrimaryChannelPlannerPlanTasksCompleteMatchingPredicate(ctx, id, options, PlannerTaskOperationPredicate{})
}

// ListTeamPrimaryChannelPlannerPlanTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamPrimaryChannelPlannerPlanTaskClient) ListTeamPrimaryChannelPlannerPlanTasksCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdTeamPrimaryChannelPlannerPlanId, options ListTeamPrimaryChannelPlannerPlanTasksOperationOptions, predicate PlannerTaskOperationPredicate) (result ListTeamPrimaryChannelPlannerPlanTasksCompleteResult, err error) {
	items := make([]beta.PlannerTask, 0)

	resp, err := c.ListTeamPrimaryChannelPlannerPlanTasks(ctx, id, options)
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

	result = ListTeamPrimaryChannelPlannerPlanTasksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
