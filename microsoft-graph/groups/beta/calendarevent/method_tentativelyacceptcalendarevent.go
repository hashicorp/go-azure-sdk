package calendarevent

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

type TentativelyAcceptCalendarEventOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type TentativelyAcceptCalendarEventOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultTentativelyAcceptCalendarEventOperationOptions() TentativelyAcceptCalendarEventOperationOptions {
	return TentativelyAcceptCalendarEventOperationOptions{}
}

func (o TentativelyAcceptCalendarEventOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o TentativelyAcceptCalendarEventOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o TentativelyAcceptCalendarEventOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// TentativelyAcceptCalendarEvent - Invoke action tentativelyAccept. Tentatively accept the specified event in a user
// calendar. If the event allows proposals for new times, on responding tentative to the event, an invitee can choose to
// suggest an alternative time by including the proposedNewTime parameter. For more information on how to propose a
// time, and how to receive and accept a new time proposal, see Propose new meeting times.
func (c CalendarEventClient) TentativelyAcceptCalendarEvent(ctx context.Context, id beta.GroupIdCalendarEventId, input TentativelyAcceptCalendarEventRequest, options TentativelyAcceptCalendarEventOperationOptions) (result TentativelyAcceptCalendarEventOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/tentativelyAccept", id.ID()),
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
