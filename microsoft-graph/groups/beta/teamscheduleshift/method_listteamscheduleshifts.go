package teamscheduleshift

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

type ListTeamScheduleShiftsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Shift
}

type ListTeamScheduleShiftsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Shift
}

type ListTeamScheduleShiftsOperationOptions struct {
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

func DefaultListTeamScheduleShiftsOperationOptions() ListTeamScheduleShiftsOperationOptions {
	return ListTeamScheduleShiftsOperationOptions{}
}

func (o ListTeamScheduleShiftsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTeamScheduleShiftsOperationOptions) ToOData() *odata.Query {
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

func (o ListTeamScheduleShiftsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTeamScheduleShiftsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamScheduleShiftsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamScheduleShifts - Get shifts from groups. The shifts in the schedule.
func (c TeamScheduleShiftClient) ListTeamScheduleShifts(ctx context.Context, id beta.GroupId, options ListTeamScheduleShiftsOperationOptions) (result ListTeamScheduleShiftsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTeamScheduleShiftsCustomPager{},
		Path:          fmt.Sprintf("%s/team/schedule/shifts", id.ID()),
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
		Values *[]beta.Shift `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamScheduleShiftsComplete retrieves all the results into a single object
func (c TeamScheduleShiftClient) ListTeamScheduleShiftsComplete(ctx context.Context, id beta.GroupId, options ListTeamScheduleShiftsOperationOptions) (ListTeamScheduleShiftsCompleteResult, error) {
	return c.ListTeamScheduleShiftsCompleteMatchingPredicate(ctx, id, options, ShiftOperationPredicate{})
}

// ListTeamScheduleShiftsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamScheduleShiftClient) ListTeamScheduleShiftsCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, options ListTeamScheduleShiftsOperationOptions, predicate ShiftOperationPredicate) (result ListTeamScheduleShiftsCompleteResult, err error) {
	items := make([]beta.Shift, 0)

	resp, err := c.ListTeamScheduleShifts(ctx, id, options)
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

	result = ListTeamScheduleShiftsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
