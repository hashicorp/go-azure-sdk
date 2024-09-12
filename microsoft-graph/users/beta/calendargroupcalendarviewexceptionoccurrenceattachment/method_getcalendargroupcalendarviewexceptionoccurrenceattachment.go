package calendargroupcalendarviewexceptionoccurrenceattachment

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetCalendarGroupCalendarViewExceptionOccurrenceAttachmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.Attachment
}

type GetCalendarGroupCalendarViewExceptionOccurrenceAttachmentOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetCalendarGroupCalendarViewExceptionOccurrenceAttachmentOperationOptions() GetCalendarGroupCalendarViewExceptionOccurrenceAttachmentOperationOptions {
	return GetCalendarGroupCalendarViewExceptionOccurrenceAttachmentOperationOptions{}
}

func (o GetCalendarGroupCalendarViewExceptionOccurrenceAttachmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCalendarGroupCalendarViewExceptionOccurrenceAttachmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetCalendarGroupCalendarViewExceptionOccurrenceAttachmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetCalendarGroupCalendarViewExceptionOccurrenceAttachment - Get attachments from users. The collection of
// FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only.
// Nullable.
func (c CalendarGroupCalendarViewExceptionOccurrenceAttachmentClient) GetCalendarGroupCalendarViewExceptionOccurrenceAttachment(ctx context.Context, id beta.UserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdAttachmentId, options GetCalendarGroupCalendarViewExceptionOccurrenceAttachmentOperationOptions) (result GetCalendarGroupCalendarViewExceptionOccurrenceAttachmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalAttachmentImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
