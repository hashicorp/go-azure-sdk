package devicecustomattributeshellscriptassignment

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

type CreateDeviceCustomAttributeShellScriptAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceManagementScriptAssignment
}

type CreateDeviceCustomAttributeShellScriptAssignmentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDeviceCustomAttributeShellScriptAssignmentOperationOptions() CreateDeviceCustomAttributeShellScriptAssignmentOperationOptions {
	return CreateDeviceCustomAttributeShellScriptAssignmentOperationOptions{}
}

func (o CreateDeviceCustomAttributeShellScriptAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDeviceCustomAttributeShellScriptAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDeviceCustomAttributeShellScriptAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDeviceCustomAttributeShellScriptAssignment - Create new navigation property to assignments for deviceManagement
func (c DeviceCustomAttributeShellScriptAssignmentClient) CreateDeviceCustomAttributeShellScriptAssignment(ctx context.Context, id beta.DeviceManagementDeviceCustomAttributeShellScriptId, input beta.DeviceManagementScriptAssignment, options CreateDeviceCustomAttributeShellScriptAssignmentOperationOptions) (result CreateDeviceCustomAttributeShellScriptAssignmentOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/assignments", id.ID()),
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

	var model beta.DeviceManagementScriptAssignment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
