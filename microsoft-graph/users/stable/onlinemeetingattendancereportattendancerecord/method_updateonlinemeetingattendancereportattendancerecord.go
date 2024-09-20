package onlinemeetingattendancereportattendancerecord

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateOnlineMeetingAttendanceReportAttendanceRecordOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateOnlineMeetingAttendanceReportAttendanceRecordOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateOnlineMeetingAttendanceReportAttendanceRecordOperationOptions() UpdateOnlineMeetingAttendanceReportAttendanceRecordOperationOptions {
	return UpdateOnlineMeetingAttendanceReportAttendanceRecordOperationOptions{}
}

func (o UpdateOnlineMeetingAttendanceReportAttendanceRecordOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateOnlineMeetingAttendanceReportAttendanceRecordOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateOnlineMeetingAttendanceReportAttendanceRecordOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateOnlineMeetingAttendanceReportAttendanceRecord - Update the navigation property attendanceRecords in users
func (c OnlineMeetingAttendanceReportAttendanceRecordClient) UpdateOnlineMeetingAttendanceReportAttendanceRecord(ctx context.Context, id stable.UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId, input stable.AttendanceRecord, options UpdateOnlineMeetingAttendanceReportAttendanceRecordOperationOptions) (result UpdateOnlineMeetingAttendanceReportAttendanceRecordOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
