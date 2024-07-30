package onlinemeetingattendancereportattendancerecord

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

type ListOnlineMeetingAttendanceReportAttendanceRecordsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AttendanceRecord
}

type ListOnlineMeetingAttendanceReportAttendanceRecordsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AttendanceRecord
}

type ListOnlineMeetingAttendanceReportAttendanceRecordsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnlineMeetingAttendanceReportAttendanceRecordsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnlineMeetingAttendanceReportAttendanceRecords ...
func (c OnlineMeetingAttendanceReportAttendanceRecordClient) ListOnlineMeetingAttendanceReportAttendanceRecords(ctx context.Context, id MeOnlineMeetingIdAttendanceReportId) (result ListOnlineMeetingAttendanceReportAttendanceRecordsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListOnlineMeetingAttendanceReportAttendanceRecordsCustomPager{},
		Path:       fmt.Sprintf("%s/attendanceRecords", id.ID()),
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
		Values *[]stable.AttendanceRecord `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOnlineMeetingAttendanceReportAttendanceRecordsComplete retrieves all the results into a single object
func (c OnlineMeetingAttendanceReportAttendanceRecordClient) ListOnlineMeetingAttendanceReportAttendanceRecordsComplete(ctx context.Context, id MeOnlineMeetingIdAttendanceReportId) (ListOnlineMeetingAttendanceReportAttendanceRecordsCompleteResult, error) {
	return c.ListOnlineMeetingAttendanceReportAttendanceRecordsCompleteMatchingPredicate(ctx, id, AttendanceRecordOperationPredicate{})
}

// ListOnlineMeetingAttendanceReportAttendanceRecordsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnlineMeetingAttendanceReportAttendanceRecordClient) ListOnlineMeetingAttendanceReportAttendanceRecordsCompleteMatchingPredicate(ctx context.Context, id MeOnlineMeetingIdAttendanceReportId, predicate AttendanceRecordOperationPredicate) (result ListOnlineMeetingAttendanceReportAttendanceRecordsCompleteResult, err error) {
	items := make([]stable.AttendanceRecord, 0)

	resp, err := c.ListOnlineMeetingAttendanceReportAttendanceRecords(ctx, id)
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

	result = ListOnlineMeetingAttendanceReportAttendanceRecordsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
