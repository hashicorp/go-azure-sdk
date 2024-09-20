package eventinstanceexceptionoccurrencecalendar

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

type GetEventInstanceExceptionOccurrenceCalendarOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Calendar
}

type GetEventInstanceExceptionOccurrenceCalendarOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetEventInstanceExceptionOccurrenceCalendarOperationOptions() GetEventInstanceExceptionOccurrenceCalendarOperationOptions {
	return GetEventInstanceExceptionOccurrenceCalendarOperationOptions{}
}

func (o GetEventInstanceExceptionOccurrenceCalendarOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEventInstanceExceptionOccurrenceCalendarOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEventInstanceExceptionOccurrenceCalendarOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEventInstanceExceptionOccurrenceCalendar - Get calendar from users. The calendar that contains the event.
// Navigation property. Read-only.
func (c EventInstanceExceptionOccurrenceCalendarClient) GetEventInstanceExceptionOccurrenceCalendar(ctx context.Context, id beta.UserIdEventIdInstanceIdExceptionOccurrenceId, options GetEventInstanceExceptionOccurrenceCalendarOperationOptions) (result GetEventInstanceExceptionOccurrenceCalendarOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/calendar", id.ID()),
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

	var model beta.Calendar
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
