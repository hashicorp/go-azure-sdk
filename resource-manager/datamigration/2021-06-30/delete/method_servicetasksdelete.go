package delete

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceTasksDeleteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ServiceTasksDeleteOperationOptions struct {
	DeleteRunningTasks *bool
}

func DefaultServiceTasksDeleteOperationOptions() ServiceTasksDeleteOperationOptions {
	return ServiceTasksDeleteOperationOptions{}
}

func (o ServiceTasksDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ServiceTasksDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ServiceTasksDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.DeleteRunningTasks != nil {
		out.Append("deleteRunningTasks", fmt.Sprintf("%v", *o.DeleteRunningTasks))
	}
	return &out
}

// ServiceTasksDelete ...
func (c DELETEClient) ServiceTasksDelete(ctx context.Context, id ServiceTaskId, options ServiceTasksDeleteOperationOptions) (result ServiceTasksDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		Path:          id.ID(),
		OptionsObject: options,
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
