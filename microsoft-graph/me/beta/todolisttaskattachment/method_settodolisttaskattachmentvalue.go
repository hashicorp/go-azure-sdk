package todolisttaskattachment

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

type SetTodoListTaskAttachmentValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetTodoListTaskAttachmentValueOperationOptions struct {
	ContentType string
	Metadata    *odata.Metadata
	RetryFunc   client.RequestRetryFunc
}

func DefaultSetTodoListTaskAttachmentValueOperationOptions() SetTodoListTaskAttachmentValueOperationOptions {
	return SetTodoListTaskAttachmentValueOperationOptions{}
}

func (o SetTodoListTaskAttachmentValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetTodoListTaskAttachmentValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetTodoListTaskAttachmentValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetTodoListTaskAttachmentValue - Update media content for the navigation property attachments in me. The unique
// identifier for an entity. Read-only.
func (c TodoListTaskAttachmentClient) SetTodoListTaskAttachmentValue(ctx context.Context, id beta.MeTodoListIdTaskIdAttachmentId, input []byte, options SetTodoListTaskAttachmentValueOperationOptions) (result SetTodoListTaskAttachmentValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: options.ContentType,
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$value", id.ID()),
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
