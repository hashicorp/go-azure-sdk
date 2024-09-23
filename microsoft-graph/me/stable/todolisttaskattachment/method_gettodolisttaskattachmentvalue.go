package todolisttaskattachment

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

type GetTodoListTaskAttachmentValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetTodoListTaskAttachmentValueOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetTodoListTaskAttachmentValueOperationOptions() GetTodoListTaskAttachmentValueOperationOptions {
	return GetTodoListTaskAttachmentValueOperationOptions{}
}

func (o GetTodoListTaskAttachmentValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTodoListTaskAttachmentValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetTodoListTaskAttachmentValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTodoListTaskAttachmentValue - List taskFileAttachments. Get a list of the taskFileAttachment objects and their
// properties. The contentBytes property will not be returned in the response. Use the Get attachment API to view the
// contentBytes.
func (c TodoListTaskAttachmentClient) GetTodoListTaskAttachmentValue(ctx context.Context, id stable.MeTodoListIdTaskIdAttachmentId, options GetTodoListTaskAttachmentValueOperationOptions) (result GetTodoListTaskAttachmentValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$value", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
