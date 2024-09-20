package manageddevicedeviceconfigurationstate

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

type CreateManagedDeviceConfigurationStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.DeviceConfigurationState
}

type CreateManagedDeviceConfigurationStateOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateManagedDeviceConfigurationStateOperationOptions() CreateManagedDeviceConfigurationStateOperationOptions {
	return CreateManagedDeviceConfigurationStateOperationOptions{}
}

func (o CreateManagedDeviceConfigurationStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateManagedDeviceConfigurationStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateManagedDeviceConfigurationStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateManagedDeviceConfigurationState - Create new navigation property to deviceConfigurationStates for
// deviceManagement
func (c ManagedDeviceDeviceConfigurationStateClient) CreateManagedDeviceConfigurationState(ctx context.Context, id stable.DeviceManagementManagedDeviceId, input stable.DeviceConfigurationState, options CreateManagedDeviceConfigurationStateOperationOptions) (result CreateManagedDeviceConfigurationStateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/deviceConfigurationStates", id.ID()),
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

	var model stable.DeviceConfigurationState
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
