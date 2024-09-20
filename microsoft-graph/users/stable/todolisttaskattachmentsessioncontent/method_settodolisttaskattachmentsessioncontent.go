package todolisttaskattachmentsessioncontent

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetTodoListTaskAttachmentSessionContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetTodoListTaskAttachmentSessionContentOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetTodoListTaskAttachmentSessionContentOperationOptions() SetTodoListTaskAttachmentSessionContentOperationOptions {
	return SetTodoListTaskAttachmentSessionContentOperationOptions{}
}

func (o SetTodoListTaskAttachmentSessionContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetTodoListTaskAttachmentSessionContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetTodoListTaskAttachmentSessionContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetTodoListTaskAttachmentSessionContent - Update content for the navigation property attachmentSessions in users. The
// content streams that are uploaded.
func (c TodoListTaskAttachmentSessionContentClient) SetTodoListTaskAttachmentSessionContent(ctx context.Context, id stable.UserIdTodoListIdTaskIdAttachmentSessionId, input []byte, options SetTodoListTaskAttachmentSessionContentOperationOptions) (result SetTodoListTaskAttachmentSessionContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/content", id.ID()),
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
