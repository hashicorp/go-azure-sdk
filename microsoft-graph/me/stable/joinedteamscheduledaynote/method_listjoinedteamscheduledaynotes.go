package joinedteamscheduledaynote

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

type ListJoinedTeamScheduleDayNotesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DayNote
}

type ListJoinedTeamScheduleDayNotesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DayNote
}

type ListJoinedTeamScheduleDayNotesOperationOptions struct {
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

func DefaultListJoinedTeamScheduleDayNotesOperationOptions() ListJoinedTeamScheduleDayNotesOperationOptions {
	return ListJoinedTeamScheduleDayNotesOperationOptions{}
}

func (o ListJoinedTeamScheduleDayNotesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListJoinedTeamScheduleDayNotesOperationOptions) ToOData() *odata.Query {
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

func (o ListJoinedTeamScheduleDayNotesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListJoinedTeamScheduleDayNotesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamScheduleDayNotesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamScheduleDayNotes - Get dayNotes from me. The day notes in the schedule.
func (c JoinedTeamScheduleDayNoteClient) ListJoinedTeamScheduleDayNotes(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamScheduleDayNotesOperationOptions) (result ListJoinedTeamScheduleDayNotesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListJoinedTeamScheduleDayNotesCustomPager{},
		Path:          fmt.Sprintf("%s/schedule/dayNotes", id.ID()),
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
		Values *[]stable.DayNote `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamScheduleDayNotesComplete retrieves all the results into a single object
func (c JoinedTeamScheduleDayNoteClient) ListJoinedTeamScheduleDayNotesComplete(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamScheduleDayNotesOperationOptions) (ListJoinedTeamScheduleDayNotesCompleteResult, error) {
	return c.ListJoinedTeamScheduleDayNotesCompleteMatchingPredicate(ctx, id, options, DayNoteOperationPredicate{})
}

// ListJoinedTeamScheduleDayNotesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamScheduleDayNoteClient) ListJoinedTeamScheduleDayNotesCompleteMatchingPredicate(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamScheduleDayNotesOperationOptions, predicate DayNoteOperationPredicate) (result ListJoinedTeamScheduleDayNotesCompleteResult, err error) {
	items := make([]stable.DayNote, 0)

	resp, err := c.ListJoinedTeamScheduleDayNotes(ctx, id, options)
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

	result = ListJoinedTeamScheduleDayNotesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
