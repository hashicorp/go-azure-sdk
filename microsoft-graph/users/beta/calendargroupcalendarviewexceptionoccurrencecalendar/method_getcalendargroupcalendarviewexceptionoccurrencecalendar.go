package calendargroupcalendarviewexceptionoccurrencecalendar

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

type GetCalendarGroupCalendarViewExceptionOccurrenceCalendarOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Calendar
}

type GetCalendarGroupCalendarViewExceptionOccurrenceCalendarOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetCalendarGroupCalendarViewExceptionOccurrenceCalendarOperationOptions() GetCalendarGroupCalendarViewExceptionOccurrenceCalendarOperationOptions {
	return GetCalendarGroupCalendarViewExceptionOccurrenceCalendarOperationOptions{}
}

func (o GetCalendarGroupCalendarViewExceptionOccurrenceCalendarOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCalendarGroupCalendarViewExceptionOccurrenceCalendarOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetCalendarGroupCalendarViewExceptionOccurrenceCalendarOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetCalendarGroupCalendarViewExceptionOccurrenceCalendar - Get calendar from users. The calendar that contains the
// event. Navigation property. Read-only.
func (c CalendarGroupCalendarViewExceptionOccurrenceCalendarClient) GetCalendarGroupCalendarViewExceptionOccurrenceCalendar(ctx context.Context, id beta.UserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId, options GetCalendarGroupCalendarViewExceptionOccurrenceCalendarOperationOptions) (result GetCalendarGroupCalendarViewExceptionOccurrenceCalendarOperationResponse, err error) {
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
