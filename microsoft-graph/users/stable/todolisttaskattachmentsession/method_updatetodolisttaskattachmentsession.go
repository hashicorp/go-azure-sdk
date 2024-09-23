package todolisttaskattachmentsession

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateTodoListTaskAttachmentSessionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateTodoListTaskAttachmentSessionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateTodoListTaskAttachmentSessionOperationOptions() UpdateTodoListTaskAttachmentSessionOperationOptions {
	return UpdateTodoListTaskAttachmentSessionOperationOptions{}
}

func (o UpdateTodoListTaskAttachmentSessionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateTodoListTaskAttachmentSessionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateTodoListTaskAttachmentSessionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateTodoListTaskAttachmentSession - Update the navigation property attachmentSessions in users
func (c TodoListTaskAttachmentSessionClient) UpdateTodoListTaskAttachmentSession(ctx context.Context, id stable.UserIdTodoListIdTaskIdAttachmentSessionId, input stable.AttachmentSession, options UpdateTodoListTaskAttachmentSessionOperationOptions) (result UpdateTodoListTaskAttachmentSessionOperationResponse, err error) {
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
