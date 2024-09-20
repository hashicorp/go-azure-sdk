package calendarviewinstanceexceptionoccurrence

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

type DeclineCalendarViewInstanceExceptionOccurrenceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeclineCalendarViewInstanceExceptionOccurrenceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultDeclineCalendarViewInstanceExceptionOccurrenceOperationOptions() DeclineCalendarViewInstanceExceptionOccurrenceOperationOptions {
	return DeclineCalendarViewInstanceExceptionOccurrenceOperationOptions{}
}

func (o DeclineCalendarViewInstanceExceptionOccurrenceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DeclineCalendarViewInstanceExceptionOccurrenceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeclineCalendarViewInstanceExceptionOccurrenceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeclineCalendarViewInstanceExceptionOccurrence - Invoke action decline. Decline invitation to the specified event in
// a user calendar. If the event allows proposals for new times, on declining the event, an invitee can choose to
// suggest an alternative time by including the proposedNewTime parameter. For more information on how to propose a
// time, and how to receive and accept a new time proposal, see Propose new meeting times.
func (c CalendarViewInstanceExceptionOccurrenceClient) DeclineCalendarViewInstanceExceptionOccurrence(ctx context.Context, id beta.GroupIdCalendarViewIdInstanceIdExceptionOccurrenceId, input DeclineCalendarViewInstanceExceptionOccurrenceRequest, options DeclineCalendarViewInstanceExceptionOccurrenceOperationOptions) (result DeclineCalendarViewInstanceExceptionOccurrenceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/decline", id.ID()),
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
