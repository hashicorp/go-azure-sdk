package joinedteamscheduleopenshiftchangerequest

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

type ListJoinedTeamScheduleOpenShiftChangeRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.OpenShiftChangeRequest
}

type ListJoinedTeamScheduleOpenShiftChangeRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.OpenShiftChangeRequest
}

type ListJoinedTeamScheduleOpenShiftChangeRequestsOperationOptions struct {
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

func DefaultListJoinedTeamScheduleOpenShiftChangeRequestsOperationOptions() ListJoinedTeamScheduleOpenShiftChangeRequestsOperationOptions {
	return ListJoinedTeamScheduleOpenShiftChangeRequestsOperationOptions{}
}

func (o ListJoinedTeamScheduleOpenShiftChangeRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListJoinedTeamScheduleOpenShiftChangeRequestsOperationOptions) ToOData() *odata.Query {
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

func (o ListJoinedTeamScheduleOpenShiftChangeRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListJoinedTeamScheduleOpenShiftChangeRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamScheduleOpenShiftChangeRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamScheduleOpenShiftChangeRequests - Get openShiftChangeRequests from me. The open shift requests in the
// schedule.
func (c JoinedTeamScheduleOpenShiftChangeRequestClient) ListJoinedTeamScheduleOpenShiftChangeRequests(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamScheduleOpenShiftChangeRequestsOperationOptions) (result ListJoinedTeamScheduleOpenShiftChangeRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListJoinedTeamScheduleOpenShiftChangeRequestsCustomPager{},
		Path:          fmt.Sprintf("%s/schedule/openShiftChangeRequests", id.ID()),
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
		Values *[]stable.OpenShiftChangeRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamScheduleOpenShiftChangeRequestsComplete retrieves all the results into a single object
func (c JoinedTeamScheduleOpenShiftChangeRequestClient) ListJoinedTeamScheduleOpenShiftChangeRequestsComplete(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamScheduleOpenShiftChangeRequestsOperationOptions) (ListJoinedTeamScheduleOpenShiftChangeRequestsCompleteResult, error) {
	return c.ListJoinedTeamScheduleOpenShiftChangeRequestsCompleteMatchingPredicate(ctx, id, options, OpenShiftChangeRequestOperationPredicate{})
}

// ListJoinedTeamScheduleOpenShiftChangeRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamScheduleOpenShiftChangeRequestClient) ListJoinedTeamScheduleOpenShiftChangeRequestsCompleteMatchingPredicate(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamScheduleOpenShiftChangeRequestsOperationOptions, predicate OpenShiftChangeRequestOperationPredicate) (result ListJoinedTeamScheduleOpenShiftChangeRequestsCompleteResult, err error) {
	items := make([]stable.OpenShiftChangeRequest, 0)

	resp, err := c.ListJoinedTeamScheduleOpenShiftChangeRequests(ctx, id, options)
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

	result = ListJoinedTeamScheduleOpenShiftChangeRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
