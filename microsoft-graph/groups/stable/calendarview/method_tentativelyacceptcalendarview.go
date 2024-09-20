package calendarview

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

type TentativelyAcceptCalendarViewOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type TentativelyAcceptCalendarViewOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultTentativelyAcceptCalendarViewOperationOptions() TentativelyAcceptCalendarViewOperationOptions {
	return TentativelyAcceptCalendarViewOperationOptions{}
}

func (o TentativelyAcceptCalendarViewOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o TentativelyAcceptCalendarViewOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o TentativelyAcceptCalendarViewOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// TentativelyAcceptCalendarView - Invoke action tentativelyAccept. Tentatively accept the specified event in a user
// calendar. If the event allows proposals for new times, on responding tentative to the event, an invitee can choose to
// suggest an alternative time by including the proposedNewTime parameter. For more information on how to propose a
// time, and how to receive and accept a new time proposal, see Propose new meeting times.
func (c CalendarViewClient) TentativelyAcceptCalendarView(ctx context.Context, id stable.GroupIdCalendarViewId, input TentativelyAcceptCalendarViewRequest, options TentativelyAcceptCalendarViewOperationOptions) (result TentativelyAcceptCalendarViewOperationResponse, err error) {
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
