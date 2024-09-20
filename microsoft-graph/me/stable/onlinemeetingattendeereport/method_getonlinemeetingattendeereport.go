package onlinemeetingattendeereport

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

type GetOnlineMeetingAttendeeReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetOnlineMeetingAttendeeReportOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultGetOnlineMeetingAttendeeReportOperationOptions() GetOnlineMeetingAttendeeReportOperationOptions {
	return GetOnlineMeetingAttendeeReportOperationOptions{}
}

func (o GetOnlineMeetingAttendeeReportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetOnlineMeetingAttendeeReportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetOnlineMeetingAttendeeReportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetOnlineMeetingAttendeeReport - Get onlineMeeting. Retrieve the properties and relationships of an onlineMeeting
// object. For example, you can: Teams live event attendee report (deprecated) is an online meeting artifact. For
// details, see Online meeting artifacts and permissions.
func (c OnlineMeetingAttendeeReportClient) GetOnlineMeetingAttendeeReport(ctx context.Context, id stable.MeOnlineMeetingId, options GetOnlineMeetingAttendeeReportOperationOptions) (result GetOnlineMeetingAttendeeReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/attendeeReport", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
