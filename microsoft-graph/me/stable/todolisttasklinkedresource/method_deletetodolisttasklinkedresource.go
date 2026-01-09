package todolisttasklinkedresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteTodoListTaskLinkedResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteTodoListTaskLinkedResourceOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteTodoListTaskLinkedResourceOperationOptions() DeleteTodoListTaskLinkedResourceOperationOptions {
	return DeleteTodoListTaskLinkedResourceOperationOptions{}
}

func (o DeleteTodoListTaskLinkedResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteTodoListTaskLinkedResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteTodoListTaskLinkedResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteTodoListTaskLinkedResource - Delete linkedResource. Delete a linkedResource object.
func (c TodoListTaskLinkedResourceClient) DeleteTodoListTaskLinkedResource(ctx context.Context, id stable.MeTodoListIdTaskIdLinkedResourceId, options DeleteTodoListTaskLinkedResourceOperationOptions) (result DeleteTodoListTaskLinkedResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
