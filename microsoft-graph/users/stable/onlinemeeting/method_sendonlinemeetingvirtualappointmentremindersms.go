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

type SendOnlineMeetingVirtualAppointmentReminderSmsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SendOnlineMeetingVirtualAppointmentReminderSmsOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSendOnlineMeetingVirtualAppointmentReminderSmsOperationOptions() SendOnlineMeetingVirtualAppointmentReminderSmsOperationOptions {
	return SendOnlineMeetingVirtualAppointmentReminderSmsOperationOptions{}
}

func (o SendOnlineMeetingVirtualAppointmentReminderSmsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SendOnlineMeetingVirtualAppointmentReminderSmsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SendOnlineMeetingVirtualAppointmentReminderSmsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SendOnlineMeetingVirtualAppointmentReminderSms - Invoke action sendVirtualAppointmentReminderSms. Send an SMS
// reminder to external attendees for a Teams virtual appointment. This feature requires Teams premium and attendees
// must have a valid United States phone number to receive SMS notifications.
func (c OnlineMeetingClient) SendOnlineMeetingVirtualAppointmentReminderSms(ctx context.Context, id stable.UserIdOnlineMeetingId, input SendOnlineMeetingVirtualAppointmentReminderSmsRequest, options SendOnlineMeetingVirtualAppointmentReminderSmsOperationOptions) (result SendOnlineMeetingVirtualAppointmentReminderSmsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/sendVirtualAppointmentReminderSms", id.ID()),
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
