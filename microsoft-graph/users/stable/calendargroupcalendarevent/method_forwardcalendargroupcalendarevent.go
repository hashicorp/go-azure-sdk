package calendargroupcalendarevent

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

type ForwardCalendarGroupCalendarEventOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// ForwardCalendarGroupCalendarEvent - Invoke action forward. This action allows the organizer or attendee of a meeting
// event to forward the meeting request to a new recipient. If the meeting event is forwarded from an attendee's
// Microsoft 365 mailbox to another recipient, this action also sends a message to notify the organizer of the
// forwarding, and adds the recipient to the organizer's copy of the meeting event. This convenience is not available
// when forwarding from an Outlook.com account.
func (c CalendarGroupCalendarEventClient) ForwardCalendarGroupCalendarEvent(ctx context.Context, id stable.UserIdCalendarGroupIdCalendarIdEventId, input ForwardCalendarGroupCalendarEventRequest) (result ForwardCalendarGroupCalendarEventOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/forward", id.ID()),
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
