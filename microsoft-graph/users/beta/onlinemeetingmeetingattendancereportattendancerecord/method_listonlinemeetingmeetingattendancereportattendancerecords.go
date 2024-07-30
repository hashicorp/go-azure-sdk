package onlinemeetingmeetingattendancereportattendancerecord

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

type ListOnlineMeetingMeetingAttendanceReportAttendanceRecordsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AttendanceRecord
}

type ListOnlineMeetingMeetingAttendanceReportAttendanceRecordsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AttendanceRecord
}

type ListOnlineMeetingMeetingAttendanceReportAttendanceRecordsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnlineMeetingMeetingAttendanceReportAttendanceRecordsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnlineMeetingMeetingAttendanceReportAttendanceRecords ...
func (c OnlineMeetingMeetingAttendanceReportAttendanceRecordClient) ListOnlineMeetingMeetingAttendanceReportAttendanceRecords(ctx context.Context, id UserIdOnlineMeetingId) (result ListOnlineMeetingMeetingAttendanceReportAttendanceRecordsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListOnlineMeetingMeetingAttendanceReportAttendanceRecordsCustomPager{},
		Path:       fmt.Sprintf("%s/meetingAttendanceReport/attendanceRecords", id.ID()),
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
		Values *[]beta.AttendanceRecord `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOnlineMeetingMeetingAttendanceReportAttendanceRecordsComplete retrieves all the results into a single object
func (c OnlineMeetingMeetingAttendanceReportAttendanceRecordClient) ListOnlineMeetingMeetingAttendanceReportAttendanceRecordsComplete(ctx context.Context, id UserIdOnlineMeetingId) (ListOnlineMeetingMeetingAttendanceReportAttendanceRecordsCompleteResult, error) {
	return c.ListOnlineMeetingMeetingAttendanceReportAttendanceRecordsCompleteMatchingPredicate(ctx, id, AttendanceRecordOperationPredicate{})
}

// ListOnlineMeetingMeetingAttendanceReportAttendanceRecordsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnlineMeetingMeetingAttendanceReportAttendanceRecordClient) ListOnlineMeetingMeetingAttendanceReportAttendanceRecordsCompleteMatchingPredicate(ctx context.Context, id UserIdOnlineMeetingId, predicate AttendanceRecordOperationPredicate) (result ListOnlineMeetingMeetingAttendanceReportAttendanceRecordsCompleteResult, err error) {
	items := make([]beta.AttendanceRecord, 0)

	resp, err := c.ListOnlineMeetingMeetingAttendanceReportAttendanceRecords(ctx, id)
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

	result = ListOnlineMeetingMeetingAttendanceReportAttendanceRecordsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
