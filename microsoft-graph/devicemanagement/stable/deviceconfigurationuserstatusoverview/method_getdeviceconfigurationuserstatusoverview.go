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

type GetDeviceConfigurationUserStatusOverviewOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.DeviceConfigurationUserOverview
}

type GetDeviceConfigurationUserStatusOverviewOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetDeviceConfigurationUserStatusOverviewOperationOptions() GetDeviceConfigurationUserStatusOverviewOperationOptions {
	return GetDeviceConfigurationUserStatusOverviewOperationOptions{}
}

func (o GetDeviceConfigurationUserStatusOverviewOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceConfigurationUserStatusOverviewOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDeviceConfigurationUserStatusOverviewOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceConfigurationUserStatusOverview - Get deviceConfigurationUserOverview. Read properties and relationships of
// the deviceConfigurationUserOverview object.
func (c DeviceConfigurationUserStatusOverviewClient) GetDeviceConfigurationUserStatusOverview(ctx context.Context, id stable.DeviceManagementDeviceConfigurationId, options GetDeviceConfigurationUserStatusOverviewOperationOptions) (result GetDeviceConfigurationUserStatusOverviewOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/userStatusOverview", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model stable.DeviceConfigurationUserOverview
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
