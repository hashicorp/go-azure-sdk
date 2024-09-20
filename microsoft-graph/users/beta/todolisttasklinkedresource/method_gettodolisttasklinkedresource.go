package todolisttasklinkedresource

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTodoListTaskLinkedResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.LinkedResource
}

type GetTodoListTaskLinkedResourceOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetTodoListTaskLinkedResourceOperationOptions() GetTodoListTaskLinkedResourceOperationOptions {
	return GetTodoListTaskLinkedResourceOperationOptions{}
}

func (o GetTodoListTaskLinkedResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTodoListTaskLinkedResourceOperationOptions) ToOData() *odata.Query {
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

func (o GetTodoListTaskLinkedResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTodoListTaskLinkedResource - Get linkedResources from users. A collection of resources linked to the task.
func (c TodoListTaskLinkedResourceClient) GetTodoListTaskLinkedResource(ctx context.Context, id beta.UserIdTodoListIdTaskIdLinkedResourceId, options GetTodoListTaskLinkedResourceOperationOptions) (result GetTodoListTaskLinkedResourceOperationResponse, err error) {
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

	var model beta.LinkedResource
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
