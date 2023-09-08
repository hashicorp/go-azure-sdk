package meonlinemeetingattendancereport

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMeOnlineMeetingByIdAttendanceReportsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MeetingAttendanceReportCollectionResponse
}

type ListMeOnlineMeetingByIdAttendanceReportsCompleteResult struct {
	Items []models.MeetingAttendanceReportCollectionResponse
}

// ListMeOnlineMeetingByIdAttendanceReports ...
func (c MeOnlineMeetingAttendanceReportClient) ListMeOnlineMeetingByIdAttendanceReports(ctx context.Context, id MeOnlineMeetingId) (result ListMeOnlineMeetingByIdAttendanceReportsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/attendanceReports", id.ID()),
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
		Values *[]models.MeetingAttendanceReportCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeOnlineMeetingByIdAttendanceReportsComplete retrieves all the results into a single object
func (c MeOnlineMeetingAttendanceReportClient) ListMeOnlineMeetingByIdAttendanceReportsComplete(ctx context.Context, id MeOnlineMeetingId) (ListMeOnlineMeetingByIdAttendanceReportsCompleteResult, error) {
	return c.ListMeOnlineMeetingByIdAttendanceReportsCompleteMatchingPredicate(ctx, id, models.MeetingAttendanceReportCollectionResponseOperationPredicate{})
}

// ListMeOnlineMeetingByIdAttendanceReportsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOnlineMeetingAttendanceReportClient) ListMeOnlineMeetingByIdAttendanceReportsCompleteMatchingPredicate(ctx context.Context, id MeOnlineMeetingId, predicate models.MeetingAttendanceReportCollectionResponseOperationPredicate) (result ListMeOnlineMeetingByIdAttendanceReportsCompleteResult, err error) {
	items := make([]models.MeetingAttendanceReportCollectionResponse, 0)

	resp, err := c.ListMeOnlineMeetingByIdAttendanceReports(ctx, id)
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

	result = ListMeOnlineMeetingByIdAttendanceReportsCompleteResult{
		Items: items,
	}
	return
}
