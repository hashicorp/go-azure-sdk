package onlinemeeting

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

type SendOnlineMeetingVirtualAppointmentSmsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// SendOnlineMeetingVirtualAppointmentSms - Invoke action sendVirtualAppointmentSms. Send an SMS notification to
// external attendees when a Teams Virtual Appointment is confirmed, rescheduled, or canceled. This feature requires
// Teams Premium. Attendees must have a valid United States phone number to receive these SMS notifications.
func (c OnlineMeetingClient) SendOnlineMeetingVirtualAppointmentSms(ctx context.Context, id beta.MeOnlineMeetingId, input SendOnlineMeetingVirtualAppointmentSmsRequest) (result SendOnlineMeetingVirtualAppointmentSmsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/sendVirtualAppointmentSms", id.ID()),
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
