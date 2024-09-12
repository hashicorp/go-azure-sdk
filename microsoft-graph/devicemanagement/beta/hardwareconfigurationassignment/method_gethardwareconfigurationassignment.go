package hardwareconfigurationassignment

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetHardwareConfigurationAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.HardwareConfigurationAssignment
}

type GetHardwareConfigurationAssignmentOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetHardwareConfigurationAssignmentOperationOptions() GetHardwareConfigurationAssignmentOperationOptions {
	return GetHardwareConfigurationAssignmentOperationOptions{}
}

func (o GetHardwareConfigurationAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetHardwareConfigurationAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetHardwareConfigurationAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetHardwareConfigurationAssignment - Get assignments from deviceManagement. A list of the Entra user group ids that
// hardware configuration will be applied to. Only security groups and Office 365 Groups are supported. Optional.
func (c HardwareConfigurationAssignmentClient) GetHardwareConfigurationAssignment(ctx context.Context, id beta.DeviceManagementHardwareConfigurationIdAssignmentId, options GetHardwareConfigurationAssignmentOperationOptions) (result GetHardwareConfigurationAssignmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.HardwareConfigurationAssignment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
