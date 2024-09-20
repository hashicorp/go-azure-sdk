package windowsdriverupdateprofile

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

type CreateWindowsDriverUpdateProfileExecuteActionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.BulkDriverActionResult
}

type CreateWindowsDriverUpdateProfileExecuteActionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateWindowsDriverUpdateProfileExecuteActionOperationOptions() CreateWindowsDriverUpdateProfileExecuteActionOperationOptions {
	return CreateWindowsDriverUpdateProfileExecuteActionOperationOptions{}
}

func (o CreateWindowsDriverUpdateProfileExecuteActionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateWindowsDriverUpdateProfileExecuteActionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateWindowsDriverUpdateProfileExecuteActionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateWindowsDriverUpdateProfileExecuteAction - Invoke action executeAction
func (c WindowsDriverUpdateProfileClient) CreateWindowsDriverUpdateProfileExecuteAction(ctx context.Context, id beta.DeviceManagementWindowsDriverUpdateProfileId, input CreateWindowsDriverUpdateProfileExecuteActionRequest, options CreateWindowsDriverUpdateProfileExecuteActionOperationOptions) (result CreateWindowsDriverUpdateProfileExecuteActionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/executeAction", id.ID()),
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

	var model beta.BulkDriverActionResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
