package messageattachment

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMessageAttachmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.Attachment
}

type GetMessageAttachmentOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetMessageAttachmentOperationOptions() GetMessageAttachmentOperationOptions {
	return GetMessageAttachmentOperationOptions{}
}

func (o GetMessageAttachmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetMessageAttachmentOperationOptions) ToOData() *odata.Query {
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

func (o GetMessageAttachmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetMessageAttachment - Get attachment. Read the properties, relationships, or raw contents of an attachment that is
// attached to a user event, message, or group post. An attachment can be one of the following types: All these types of
// attachments are derived from the attachment resource.
func (c MessageAttachmentClient) GetMessageAttachment(ctx context.Context, id stable.MeMessageIdAttachmentId, options GetMessageAttachmentOperationOptions) (result GetMessageAttachmentOperationResponse, err error) {
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
	model, err := stable.UnmarshalAttachmentImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
