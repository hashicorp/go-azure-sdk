package onlinemeeting

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

type SendOnlineMeetingVirtualAppointmentSmsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SendOnlineMeetingVirtualAppointmentSmsOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSendOnlineMeetingVirtualAppointmentSmsOperationOptions() SendOnlineMeetingVirtualAppointmentSmsOperationOptions {
	return SendOnlineMeetingVirtualAppointmentSmsOperationOptions{}
}

func (o SendOnlineMeetingVirtualAppointmentSmsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SendOnlineMeetingVirtualAppointmentSmsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SendOnlineMeetingVirtualAppointmentSmsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SendOnlineMeetingVirtualAppointmentSms - Invoke action sendVirtualAppointmentSms. Send an SMS notification to
// external attendees when a Teams virtual appointment is confirmed, rescheduled, or canceled. This feature requires
// Teams premium. Attendees must have a valid United States phone number to receive these SMS notifications.
func (c OnlineMeetingClient) SendOnlineMeetingVirtualAppointmentSms(ctx context.Context, id stable.UserIdOnlineMeetingId, input SendOnlineMeetingVirtualAppointmentSmsRequest, options SendOnlineMeetingVirtualAppointmentSmsOperationOptions) (result SendOnlineMeetingVirtualAppointmentSmsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/sendVirtualAppointmentSms", id.ID()),
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
