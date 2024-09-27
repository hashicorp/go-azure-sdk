package todolisttaskchecklistitem

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

type CreateTodoListTaskChecklistItemOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ChecklistItem
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

// CreateTodoListTaskChecklistItem - Create new navigation property to checklistItems for users
func (c TodoListTaskChecklistItemClient) CreateTodoListTaskChecklistItem(ctx context.Context, id stable.UserIdTodoListIdTaskId, input stable.ChecklistItem, options CreateTodoListTaskChecklistItemOperationOptions) (result CreateTodoListTaskChecklistItemOperationResponse, err error) {
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

	var model stable.ChecklistItem
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
