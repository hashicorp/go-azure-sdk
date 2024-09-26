package calendareventattachment

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

type CreateCalendarEventAttachmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.Attachment
}

type CreateCalendarEventAttachmentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateCalendarEventAttachmentOperationOptions() CreateCalendarEventAttachmentOperationOptions {
	return CreateCalendarEventAttachmentOperationOptions{}
}

func (o CreateCalendarEventAttachmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateCalendarEventAttachmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateCalendarEventAttachmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateCalendarEventAttachment - Create new navigation property to attachments for groups
func (c CalendarEventAttachmentClient) CreateCalendarEventAttachment(ctx context.Context, id beta.GroupIdCalendarEventId, input beta.Attachment, options CreateCalendarEventAttachmentOperationOptions) (result CreateCalendarEventAttachmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/attachments", id.ID()),
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
