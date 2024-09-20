package deviceconfigurationdevicestatus

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

type CreateDeviceConfigurationDeviceStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.DeviceConfigurationDeviceStatus
}

type CreateDeviceConfigurationDeviceStatusOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateDeviceConfigurationDeviceStatusOperationOptions() CreateDeviceConfigurationDeviceStatusOperationOptions {
	return CreateDeviceConfigurationDeviceStatusOperationOptions{}
}

func (o CreateDeviceConfigurationDeviceStatusOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDeviceConfigurationDeviceStatusOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDeviceConfigurationDeviceStatusOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDeviceConfigurationDeviceStatus - Create deviceConfigurationDeviceStatus. Create a new
// deviceConfigurationDeviceStatus object.
func (c DeviceConfigurationDeviceStatusClient) CreateDeviceConfigurationDeviceStatus(ctx context.Context, id stable.DeviceManagementDeviceConfigurationId, input stable.DeviceConfigurationDeviceStatus, options CreateDeviceConfigurationDeviceStatusOperationOptions) (result CreateDeviceConfigurationDeviceStatusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/deviceStatuses", id.ID()),
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

	var model stable.DeviceConfigurationDeviceStatus
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
