package joinedteamscheduletimesoff

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

type ListJoinedTeamScheduleTimesOffsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TimeOff
}

type ListJoinedTeamScheduleTimesOffsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TimeOff
}

type ListJoinedTeamScheduleTimesOffsOperationOptions struct {
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

func DefaultListJoinedTeamScheduleTimesOffsOperationOptions() ListJoinedTeamScheduleTimesOffsOperationOptions {
	return ListJoinedTeamScheduleTimesOffsOperationOptions{}
}

func (o ListJoinedTeamScheduleTimesOffsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListJoinedTeamScheduleTimesOffsOperationOptions) ToOData() *odata.Query {
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

func (o ListJoinedTeamScheduleTimesOffsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListJoinedTeamScheduleTimesOffsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamScheduleTimesOffsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamScheduleTimesOffs - Get timesOff from me. The instances of times off in the schedule.
func (c JoinedTeamScheduleTimesOffClient) ListJoinedTeamScheduleTimesOffs(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamScheduleTimesOffsOperationOptions) (result ListJoinedTeamScheduleTimesOffsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListJoinedTeamScheduleTimesOffsCustomPager{},
		Path:          fmt.Sprintf("%s/schedule/timesOff", id.ID()),
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
		Values *[]stable.TimeOff `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamScheduleTimesOffsComplete retrieves all the results into a single object
func (c JoinedTeamScheduleTimesOffClient) ListJoinedTeamScheduleTimesOffsComplete(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamScheduleTimesOffsOperationOptions) (ListJoinedTeamScheduleTimesOffsCompleteResult, error) {
	return c.ListJoinedTeamScheduleTimesOffsCompleteMatchingPredicate(ctx, id, options, TimeOffOperationPredicate{})
}

// ListJoinedTeamScheduleTimesOffsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamScheduleTimesOffClient) ListJoinedTeamScheduleTimesOffsCompleteMatchingPredicate(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamScheduleTimesOffsOperationOptions, predicate TimeOffOperationPredicate) (result ListJoinedTeamScheduleTimesOffsCompleteResult, err error) {
	items := make([]stable.TimeOff, 0)

	resp, err := c.ListJoinedTeamScheduleTimesOffs(ctx, id, options)
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

	result = ListJoinedTeamScheduleTimesOffsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
