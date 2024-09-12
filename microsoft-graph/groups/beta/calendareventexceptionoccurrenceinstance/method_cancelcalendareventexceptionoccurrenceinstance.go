package calendareventexceptionoccurrenceinstance

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

type CancelCalendarEventExceptionOccurrenceInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// CancelCalendarEventExceptionOccurrenceInstance - Invoke action cancel. This action allows the organizer of a meeting
// to send a cancellation message and cancel the event. The action moves the event to the Deleted Items folder. The
// organizer can also cancel an occurrence of a recurring meeting by providing the occurrence event ID. An attendee
// calling this action gets an error (HTTP 400 Bad Request), with the following error message: 'Your request can't be
// completed. You need to be an organizer to cancel a meeting.' This action differs from Delete in that Cancel is
// available to only the organizer, and lets the organizer send a custom message to the attendees about the
// cancellation.
func (c CalendarEventExceptionOccurrenceInstanceClient) CancelCalendarEventExceptionOccurrenceInstance(ctx context.Context, id beta.GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId, input CancelCalendarEventExceptionOccurrenceInstanceRequest) (result CancelCalendarEventExceptionOccurrenceInstanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/cancel", id.ID()),
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
