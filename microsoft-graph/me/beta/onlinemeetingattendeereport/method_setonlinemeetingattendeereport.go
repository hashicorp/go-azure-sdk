package onlinemeetingattendeereport

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

type SetOnlineMeetingAttendeeReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetOnlineMeetingAttendeeReportOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetOnlineMeetingAttendeeReportOperationOptions() SetOnlineMeetingAttendeeReportOperationOptions {
	return SetOnlineMeetingAttendeeReportOperationOptions{}
}

func (o SetOnlineMeetingAttendeeReportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetOnlineMeetingAttendeeReportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetOnlineMeetingAttendeeReportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetOnlineMeetingAttendeeReport - Update attendeeReport for the navigation property onlineMeetings in me. The content
// stream of the attendee report of a Teams live event. Read-only.
func (c OnlineMeetingAttendeeReportClient) SetOnlineMeetingAttendeeReport(ctx context.Context, id beta.MeOnlineMeetingId, input []byte, options SetOnlineMeetingAttendeeReportOperationOptions) (result SetOnlineMeetingAttendeeReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/attendeeReport", id.ID()),
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
