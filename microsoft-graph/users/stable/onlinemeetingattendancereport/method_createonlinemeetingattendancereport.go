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

type CreateOnlineMeetingAttendanceReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.MeetingAttendanceReport
}

type CreateOnlineMeetingAttendanceReportOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateOnlineMeetingAttendanceReportOperationOptions() CreateOnlineMeetingAttendanceReportOperationOptions {
	return CreateOnlineMeetingAttendanceReportOperationOptions{}
}

func (o CreateOnlineMeetingAttendanceReportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateOnlineMeetingAttendanceReportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateOnlineMeetingAttendanceReportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateOnlineMeetingAttendanceReport - Create new navigation property to attendanceReports for users
func (c OnlineMeetingAttendanceReportClient) CreateOnlineMeetingAttendanceReport(ctx context.Context, id stable.UserIdOnlineMeetingId, input stable.MeetingAttendanceReport, options CreateOnlineMeetingAttendanceReportOperationOptions) (result CreateOnlineMeetingAttendanceReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/attendanceReports", id.ID()),
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

	var model stable.MeetingAttendanceReport
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
