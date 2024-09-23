package calendareventexceptionoccurrenceextension

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateCalendarEventExceptionOccurrenceExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateCalendarEventExceptionOccurrenceExtensionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateCalendarEventExceptionOccurrenceExtensionOperationOptions() UpdateCalendarEventExceptionOccurrenceExtensionOperationOptions {
	return UpdateCalendarEventExceptionOccurrenceExtensionOperationOptions{}
}

func (o UpdateCalendarEventExceptionOccurrenceExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateCalendarEventExceptionOccurrenceExtensionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateCalendarEventExceptionOccurrenceExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateCalendarEventExceptionOccurrenceExtension - Update the navigation property extensions in groups
func (c CalendarEventExceptionOccurrenceExtensionClient) UpdateCalendarEventExceptionOccurrenceExtension(ctx context.Context, id beta.GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId, input beta.Extension, options UpdateCalendarEventExceptionOccurrenceExtensionOperationOptions) (result UpdateCalendarEventExceptionOccurrenceExtensionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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
