package calendargroupcalendarviewexceptionoccurrenceinstanceextension

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

type DeleteCalendarGroupCalendarViewExceptionOccurrenceInstanceExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteCalendarGroupCalendarViewExceptionOccurrenceInstanceExtensionOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteCalendarGroupCalendarViewExceptionOccurrenceInstanceExtensionOperationOptions() DeleteCalendarGroupCalendarViewExceptionOccurrenceInstanceExtensionOperationOptions {
	return DeleteCalendarGroupCalendarViewExceptionOccurrenceInstanceExtensionOperationOptions{}
}

func (o DeleteCalendarGroupCalendarViewExceptionOccurrenceInstanceExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteCalendarGroupCalendarViewExceptionOccurrenceInstanceExtensionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteCalendarGroupCalendarViewExceptionOccurrenceInstanceExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteCalendarGroupCalendarViewExceptionOccurrenceInstanceExtension - Delete navigation property extensions for me
func (c CalendarGroupCalendarViewExceptionOccurrenceInstanceExtensionClient) DeleteCalendarGroupCalendarViewExceptionOccurrenceInstanceExtension(ctx context.Context, id beta.MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId, options DeleteCalendarGroupCalendarViewExceptionOccurrenceInstanceExtensionOperationOptions) (result DeleteCalendarGroupCalendarViewExceptionOccurrenceInstanceExtensionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
