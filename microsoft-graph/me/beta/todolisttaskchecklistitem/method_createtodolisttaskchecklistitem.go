package todolisttaskchecklistitem

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

type CreateTodoListTaskChecklistItemOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ChecklistItem
}

type CreateTodoListTaskChecklistItemOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTodoListTaskChecklistItemOperationOptions() CreateTodoListTaskChecklistItemOperationOptions {
	return CreateTodoListTaskChecklistItemOperationOptions{}
}

func (o CreateTodoListTaskChecklistItemOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTodoListTaskChecklistItemOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTodoListTaskChecklistItemOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTodoListTaskChecklistItem - Create checklistItem. Create a new checklistItem object as a subtask in a bigger
// todoTask.
func (c TodoListTaskChecklistItemClient) CreateTodoListTaskChecklistItem(ctx context.Context, id beta.MeTodoListIdTaskId, input beta.ChecklistItem, options CreateTodoListTaskChecklistItemOperationOptions) (result CreateTodoListTaskChecklistItemOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/checklistItems", id.ID()),
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

	var model beta.ChecklistItem
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
