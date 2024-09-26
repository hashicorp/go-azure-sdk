package deviceconfigurationassignment

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

type CreateDeviceConfigurationAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.DeviceConfigurationAssignment
}

type CreateDeviceConfigurationAssignmentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDeviceConfigurationAssignmentOperationOptions() CreateDeviceConfigurationAssignmentOperationOptions {
	return CreateDeviceConfigurationAssignmentOperationOptions{}
}

func (o CreateDeviceConfigurationAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDeviceConfigurationAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDeviceConfigurationAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDeviceConfigurationAssignment - Create deviceConfigurationAssignment. Create a new
// deviceConfigurationAssignment object.
func (c DeviceConfigurationAssignmentClient) CreateDeviceConfigurationAssignment(ctx context.Context, id stable.DeviceManagementDeviceConfigurationId, input stable.DeviceConfigurationAssignment, options CreateDeviceConfigurationAssignmentOperationOptions) (result CreateDeviceConfigurationAssignmentOperationResponse, err error) {
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

	var model stable.DeviceConfigurationAssignment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
