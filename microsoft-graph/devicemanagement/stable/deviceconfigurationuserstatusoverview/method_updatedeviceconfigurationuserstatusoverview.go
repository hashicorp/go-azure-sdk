package deviceconfigurationuserstatusoverview

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

type UpdateDeviceConfigurationUserStatusOverviewOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDeviceConfigurationUserStatusOverviewOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateDeviceConfigurationUserStatusOverviewOperationOptions() UpdateDeviceConfigurationUserStatusOverviewOperationOptions {
	return UpdateDeviceConfigurationUserStatusOverviewOperationOptions{}
}

func (o UpdateDeviceConfigurationUserStatusOverviewOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDeviceConfigurationUserStatusOverviewOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDeviceConfigurationUserStatusOverviewOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDeviceConfigurationUserStatusOverview - Update deviceConfigurationUserOverview. Update the properties of a
// deviceConfigurationUserOverview object.
func (c DeviceConfigurationUserStatusOverviewClient) UpdateDeviceConfigurationUserStatusOverview(ctx context.Context, id stable.DeviceManagementDeviceConfigurationId, input stable.DeviceConfigurationUserOverview, options UpdateDeviceConfigurationUserStatusOverviewOperationOptions) (result UpdateDeviceConfigurationUserStatusOverviewOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/userStatusOverview", id.ID()),
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

	return
}
