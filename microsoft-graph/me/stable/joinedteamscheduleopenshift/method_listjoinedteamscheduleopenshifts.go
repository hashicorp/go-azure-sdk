package joinedteamscheduleopenshift

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

type ListJoinedTeamScheduleOpenShiftsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.OpenShift
}

type ListJoinedTeamScheduleOpenShiftsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.OpenShift
}

type ListJoinedTeamScheduleOpenShiftsOperationOptions struct {
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

func DefaultListJoinedTeamScheduleOpenShiftsOperationOptions() ListJoinedTeamScheduleOpenShiftsOperationOptions {
	return ListJoinedTeamScheduleOpenShiftsOperationOptions{}
}

func (o ListJoinedTeamScheduleOpenShiftsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListJoinedTeamScheduleOpenShiftsOperationOptions) ToOData() *odata.Query {
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

func (o ListJoinedTeamScheduleOpenShiftsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListJoinedTeamScheduleOpenShiftsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamScheduleOpenShiftsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamScheduleOpenShifts - Get openShifts from me. The set of open shifts in a scheduling group in the
// schedule.
func (c JoinedTeamScheduleOpenShiftClient) ListJoinedTeamScheduleOpenShifts(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamScheduleOpenShiftsOperationOptions) (result ListJoinedTeamScheduleOpenShiftsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListJoinedTeamScheduleOpenShiftsCustomPager{},
		Path:          fmt.Sprintf("%s/schedule/openShifts", id.ID()),
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
		Values *[]stable.OpenShift `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamScheduleOpenShiftsComplete retrieves all the results into a single object
func (c JoinedTeamScheduleOpenShiftClient) ListJoinedTeamScheduleOpenShiftsComplete(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamScheduleOpenShiftsOperationOptions) (ListJoinedTeamScheduleOpenShiftsCompleteResult, error) {
	return c.ListJoinedTeamScheduleOpenShiftsCompleteMatchingPredicate(ctx, id, options, OpenShiftOperationPredicate{})
}

// ListJoinedTeamScheduleOpenShiftsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamScheduleOpenShiftClient) ListJoinedTeamScheduleOpenShiftsCompleteMatchingPredicate(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamScheduleOpenShiftsOperationOptions, predicate OpenShiftOperationPredicate) (result ListJoinedTeamScheduleOpenShiftsCompleteResult, err error) {
	items := make([]stable.OpenShift, 0)

	resp, err := c.ListJoinedTeamScheduleOpenShifts(ctx, id, options)
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

	result = ListJoinedTeamScheduleOpenShiftsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
