package onlinemeetingattendancereport

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

type ListOnlineMeetingAttendanceReportsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.MeetingAttendanceReport
}

type ListOnlineMeetingAttendanceReportsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.MeetingAttendanceReport
}

type ListOnlineMeetingAttendanceReportsOperationOptions struct {
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

func DefaultListOnlineMeetingAttendanceReportsOperationOptions() ListOnlineMeetingAttendanceReportsOperationOptions {
	return ListOnlineMeetingAttendanceReportsOperationOptions{}
}

func (o ListOnlineMeetingAttendanceReportsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOnlineMeetingAttendanceReportsOperationOptions) ToOData() *odata.Query {
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

func (o ListOnlineMeetingAttendanceReportsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOnlineMeetingAttendanceReportsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnlineMeetingAttendanceReportsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnlineMeetingAttendanceReports - List meetingAttendanceReports. Get a list of meetingAttendanceReport objects for
// an onlineMeeting or a virtualEvent. Each time an online meeting or a virtual event ends, an attendance report is
// generated for that session.
func (c OnlineMeetingAttendanceReportClient) ListOnlineMeetingAttendanceReports(ctx context.Context, id stable.MeOnlineMeetingId, options ListOnlineMeetingAttendanceReportsOperationOptions) (result ListOnlineMeetingAttendanceReportsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOnlineMeetingAttendanceReportsCustomPager{},
		Path:          fmt.Sprintf("%s/attendanceReports", id.ID()),
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
		Values *[]stable.MeetingAttendanceReport `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOnlineMeetingAttendanceReportsComplete retrieves all the results into a single object
func (c OnlineMeetingAttendanceReportClient) ListOnlineMeetingAttendanceReportsComplete(ctx context.Context, id stable.MeOnlineMeetingId, options ListOnlineMeetingAttendanceReportsOperationOptions) (ListOnlineMeetingAttendanceReportsCompleteResult, error) {
	return c.ListOnlineMeetingAttendanceReportsCompleteMatchingPredicate(ctx, id, options, MeetingAttendanceReportOperationPredicate{})
}

// ListOnlineMeetingAttendanceReportsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnlineMeetingAttendanceReportClient) ListOnlineMeetingAttendanceReportsCompleteMatchingPredicate(ctx context.Context, id stable.MeOnlineMeetingId, options ListOnlineMeetingAttendanceReportsOperationOptions, predicate MeetingAttendanceReportOperationPredicate) (result ListOnlineMeetingAttendanceReportsCompleteResult, err error) {
	items := make([]stable.MeetingAttendanceReport, 0)

	resp, err := c.ListOnlineMeetingAttendanceReports(ctx, id, options)
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

	result = ListOnlineMeetingAttendanceReportsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
