package todolisttaskattachment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTodoListTaskAttachmentsUploadSessionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.UploadSession
}

type CreateTodoListTaskAttachmentsUploadSessionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTodoListTaskAttachmentsUploadSessionOperationOptions() CreateTodoListTaskAttachmentsUploadSessionOperationOptions {
	return CreateTodoListTaskAttachmentsUploadSessionOperationOptions{}
}

func (o CreateTodoListTaskAttachmentsUploadSessionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTodoListTaskAttachmentsUploadSessionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTodoListTaskAttachmentsUploadSessionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTodoListTaskAttachmentsUploadSession - Invoke action createUploadSession. Create an upload session to
// iteratively upload ranges of a file as an attachment to a todoTask. As part of the response, this action returns an
// upload URL that you can use in subsequent sequential PUT queries. The request headers for each PUT operation let you
// specify the exact range of bytes to be uploaded. This allows the transfer to be resumed, in case the network
// connection is dropped during the upload. The following are the steps to attach a file to a Microsoft To Do task using
// an upload session: For an example that describes the end-to-end attachment process, see attach files to a To Do task.
func (c TodoListTaskAttachmentClient) CreateTodoListTaskAttachmentsUploadSession(ctx context.Context, id beta.MeTodoListIdTaskId, input CreateTodoListTaskAttachmentsUploadSessionRequest, options CreateTodoListTaskAttachmentsUploadSessionOperationOptions) (result CreateTodoListTaskAttachmentsUploadSessionOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/attachments/createUploadSession", id.ID()),
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
	model, err := beta.UnmarshalUploadSessionImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
