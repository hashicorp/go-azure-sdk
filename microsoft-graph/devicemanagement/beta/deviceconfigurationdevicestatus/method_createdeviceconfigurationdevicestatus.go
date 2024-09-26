package deviceconfigurationdevicestatus

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

type CreateDeviceConfigurationDeviceStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceConfigurationDeviceStatus
}

type CreateDeviceConfigurationDeviceStatusOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
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

// CreateDeviceConfigurationDeviceStatus - Create new navigation property to deviceStatuses for deviceManagement
func (c DeviceConfigurationDeviceStatusClient) CreateDeviceConfigurationDeviceStatus(ctx context.Context, id beta.DeviceManagementDeviceConfigurationId, input beta.DeviceConfigurationDeviceStatus, options CreateDeviceConfigurationDeviceStatusOperationOptions) (result CreateDeviceConfigurationDeviceStatusOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/deviceStatuses", id.ID()),
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

	var model beta.DeviceConfigurationDeviceStatus
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
