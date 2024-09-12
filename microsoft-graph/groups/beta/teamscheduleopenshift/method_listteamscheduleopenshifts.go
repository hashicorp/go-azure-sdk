package teamscheduleopenshift

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

type ListTeamScheduleOpenShiftsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OpenShift
}

type ListTeamScheduleOpenShiftsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OpenShift
}

type ListTeamScheduleOpenShiftsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListTeamScheduleOpenShiftsOperationOptions() ListTeamScheduleOpenShiftsOperationOptions {
	return ListTeamScheduleOpenShiftsOperationOptions{}
}

func (o ListTeamScheduleOpenShiftsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTeamScheduleOpenShiftsOperationOptions) ToOData() *odata.Query {
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

func (o ListTeamScheduleOpenShiftsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTeamScheduleOpenShiftsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamScheduleOpenShiftsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamScheduleOpenShifts - Get openShifts from groups. The set of open shifts in a scheduling group in the
// schedule.
func (c TeamScheduleOpenShiftClient) ListTeamScheduleOpenShifts(ctx context.Context, id beta.GroupId, options ListTeamScheduleOpenShiftsOperationOptions) (result ListTeamScheduleOpenShiftsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTeamScheduleOpenShiftsCustomPager{},
		Path:          fmt.Sprintf("%s/team/schedule/openShifts", id.ID()),
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
		Values *[]beta.OpenShift `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamScheduleOpenShiftsComplete retrieves all the results into a single object
func (c TeamScheduleOpenShiftClient) ListTeamScheduleOpenShiftsComplete(ctx context.Context, id beta.GroupId, options ListTeamScheduleOpenShiftsOperationOptions) (ListTeamScheduleOpenShiftsCompleteResult, error) {
	return c.ListTeamScheduleOpenShiftsCompleteMatchingPredicate(ctx, id, options, OpenShiftOperationPredicate{})
}

// ListTeamScheduleOpenShiftsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamScheduleOpenShiftClient) ListTeamScheduleOpenShiftsCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, options ListTeamScheduleOpenShiftsOperationOptions, predicate OpenShiftOperationPredicate) (result ListTeamScheduleOpenShiftsCompleteResult, err error) {
	items := make([]beta.OpenShift, 0)

	resp, err := c.ListTeamScheduleOpenShifts(ctx, id, options)
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

	result = ListTeamScheduleOpenShiftsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
