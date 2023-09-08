package useronlinemeetingattendancereportattendancerecord

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

type ListUserByIdOnlineMeetingByIdAttendanceReportByIdAttendanceRecordsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttendanceRecordCollectionResponse
}

type ListUserByIdOnlineMeetingByIdAttendanceReportByIdAttendanceRecordsCompleteResult struct {
	Items []models.AttendanceRecordCollectionResponse
}

// ListUserByIdOnlineMeetingByIdAttendanceReportByIdAttendanceRecords ...
func (c UserOnlineMeetingAttendanceReportAttendanceRecordClient) ListUserByIdOnlineMeetingByIdAttendanceReportByIdAttendanceRecords(ctx context.Context, id UserOnlineMeetingAttendanceReportId) (result ListUserByIdOnlineMeetingByIdAttendanceReportByIdAttendanceRecordsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.AttendanceRecordCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdOnlineMeetingByIdAttendanceReportByIdAttendanceRecordsComplete retrieves all the results into a single object
func (c UserOnlineMeetingAttendanceReportAttendanceRecordClient) ListUserByIdOnlineMeetingByIdAttendanceReportByIdAttendanceRecordsComplete(ctx context.Context, id UserOnlineMeetingAttendanceReportId) (ListUserByIdOnlineMeetingByIdAttendanceReportByIdAttendanceRecordsCompleteResult, error) {
	return c.ListUserByIdOnlineMeetingByIdAttendanceReportByIdAttendanceRecordsCompleteMatchingPredicate(ctx, id, models.AttendanceRecordCollectionResponseOperationPredicate{})
}

// ListUserByIdOnlineMeetingByIdAttendanceReportByIdAttendanceRecordsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOnlineMeetingAttendanceReportAttendanceRecordClient) ListUserByIdOnlineMeetingByIdAttendanceReportByIdAttendanceRecordsCompleteMatchingPredicate(ctx context.Context, id UserOnlineMeetingAttendanceReportId, predicate models.AttendanceRecordCollectionResponseOperationPredicate) (result ListUserByIdOnlineMeetingByIdAttendanceReportByIdAttendanceRecordsCompleteResult, err error) {
	items := make([]models.AttendanceRecordCollectionResponse, 0)

	resp, err := c.ListUserByIdOnlineMeetingByIdAttendanceReportByIdAttendanceRecords(ctx, id)
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

	result = ListUserByIdOnlineMeetingByIdAttendanceReportByIdAttendanceRecordsCompleteResult{
		Items: items,
	}
	return
}
