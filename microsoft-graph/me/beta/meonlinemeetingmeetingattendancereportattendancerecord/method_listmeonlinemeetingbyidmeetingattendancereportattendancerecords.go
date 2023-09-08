package meonlinemeetingmeetingattendancereportattendancerecord

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMeOnlineMeetingByIdMeetingAttendanceReportAttendanceRecordsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttendanceRecordCollectionResponse
}

type ListMeOnlineMeetingByIdMeetingAttendanceReportAttendanceRecordsCompleteResult struct {
	Items []models.AttendanceRecordCollectionResponse
}

// ListMeOnlineMeetingByIdMeetingAttendanceReportAttendanceRecords ...
func (c MeOnlineMeetingMeetingAttendanceReportAttendanceRecordClient) ListMeOnlineMeetingByIdMeetingAttendanceReportAttendanceRecords(ctx context.Context, id MeOnlineMeetingId) (result ListMeOnlineMeetingByIdMeetingAttendanceReportAttendanceRecordsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.AttendanceRecordCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeOnlineMeetingByIdMeetingAttendanceReportAttendanceRecordsComplete retrieves all the results into a single object
func (c MeOnlineMeetingMeetingAttendanceReportAttendanceRecordClient) ListMeOnlineMeetingByIdMeetingAttendanceReportAttendanceRecordsComplete(ctx context.Context, id MeOnlineMeetingId) (ListMeOnlineMeetingByIdMeetingAttendanceReportAttendanceRecordsCompleteResult, error) {
	return c.ListMeOnlineMeetingByIdMeetingAttendanceReportAttendanceRecordsCompleteMatchingPredicate(ctx, id, models.AttendanceRecordCollectionResponseOperationPredicate{})
}

// ListMeOnlineMeetingByIdMeetingAttendanceReportAttendanceRecordsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOnlineMeetingMeetingAttendanceReportAttendanceRecordClient) ListMeOnlineMeetingByIdMeetingAttendanceReportAttendanceRecordsCompleteMatchingPredicate(ctx context.Context, id MeOnlineMeetingId, predicate models.AttendanceRecordCollectionResponseOperationPredicate) (result ListMeOnlineMeetingByIdMeetingAttendanceReportAttendanceRecordsCompleteResult, err error) {
	items := make([]models.AttendanceRecordCollectionResponse, 0)

	resp, err := c.ListMeOnlineMeetingByIdMeetingAttendanceReportAttendanceRecords(ctx, id)
	if err != nil {
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

	result = ListMeOnlineMeetingByIdMeetingAttendanceReportAttendanceRecordsCompleteResult{
		Items: items,
	}
	return
}
