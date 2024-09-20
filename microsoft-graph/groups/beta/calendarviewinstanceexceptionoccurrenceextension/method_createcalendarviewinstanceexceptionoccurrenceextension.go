package calendarviewinstanceexceptionoccurrenceextension

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateCalendarViewInstanceExceptionOccurrenceExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.Extension
}

type CreateCalendarViewInstanceExceptionOccurrenceExtensionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateCalendarViewInstanceExceptionOccurrenceExtensionOperationOptions() CreateCalendarViewInstanceExceptionOccurrenceExtensionOperationOptions {
	return CreateCalendarViewInstanceExceptionOccurrenceExtensionOperationOptions{}
}

func (o CreateCalendarViewInstanceExceptionOccurrenceExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateCalendarViewInstanceExceptionOccurrenceExtensionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateCalendarViewInstanceExceptionOccurrenceExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateCalendarViewInstanceExceptionOccurrenceExtension - Create new navigation property to extensions for groups
func (c CalendarViewInstanceExceptionOccurrenceExtensionClient) CreateCalendarViewInstanceExceptionOccurrenceExtension(ctx context.Context, id beta.GroupIdCalendarViewIdInstanceIdExceptionOccurrenceId, input beta.Extension, options CreateCalendarViewInstanceExceptionOccurrenceExtensionOperationOptions) (result CreateCalendarViewInstanceExceptionOccurrenceExtensionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/extensions", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalExtensionImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
