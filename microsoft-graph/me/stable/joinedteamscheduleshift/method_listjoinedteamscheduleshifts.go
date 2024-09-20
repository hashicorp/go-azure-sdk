package joinedteamscheduleshift

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

type ListJoinedTeamScheduleShiftsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Shift
}

type ListJoinedTeamScheduleShiftsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Shift
}

type ListJoinedTeamScheduleShiftsOperationOptions struct {
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

func DefaultListJoinedTeamScheduleShiftsOperationOptions() ListJoinedTeamScheduleShiftsOperationOptions {
	return ListJoinedTeamScheduleShiftsOperationOptions{}
}

func (o ListJoinedTeamScheduleShiftsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListJoinedTeamScheduleShiftsOperationOptions) ToOData() *odata.Query {
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

func (o ListJoinedTeamScheduleShiftsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListJoinedTeamScheduleShiftsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamScheduleShiftsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamScheduleShifts - Get shifts from me. The shifts in the schedule.
func (c JoinedTeamScheduleShiftClient) ListJoinedTeamScheduleShifts(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamScheduleShiftsOperationOptions) (result ListJoinedTeamScheduleShiftsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListJoinedTeamScheduleShiftsCustomPager{},
		Path:          fmt.Sprintf("%s/schedule/shifts", id.ID()),
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
		Values *[]stable.Shift `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamScheduleShiftsComplete retrieves all the results into a single object
func (c JoinedTeamScheduleShiftClient) ListJoinedTeamScheduleShiftsComplete(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamScheduleShiftsOperationOptions) (ListJoinedTeamScheduleShiftsCompleteResult, error) {
	return c.ListJoinedTeamScheduleShiftsCompleteMatchingPredicate(ctx, id, options, ShiftOperationPredicate{})
}

// ListJoinedTeamScheduleShiftsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamScheduleShiftClient) ListJoinedTeamScheduleShiftsCompleteMatchingPredicate(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamScheduleShiftsOperationOptions, predicate ShiftOperationPredicate) (result ListJoinedTeamScheduleShiftsCompleteResult, err error) {
	items := make([]stable.Shift, 0)

	resp, err := c.ListJoinedTeamScheduleShifts(ctx, id, options)
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

	result = ListJoinedTeamScheduleShiftsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
