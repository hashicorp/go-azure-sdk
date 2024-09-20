package deviceenrollmentconfigurationassignment

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceEnrollmentConfigurationAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.EnrollmentConfigurationAssignment
}

type GetDeviceEnrollmentConfigurationAssignmentOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetDeviceEnrollmentConfigurationAssignmentOperationOptions() GetDeviceEnrollmentConfigurationAssignmentOperationOptions {
	return GetDeviceEnrollmentConfigurationAssignmentOperationOptions{}
}

func (o GetDeviceEnrollmentConfigurationAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceEnrollmentConfigurationAssignmentOperationOptions) ToOData() *odata.Query {
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

func (o GetDeviceEnrollmentConfigurationAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceEnrollmentConfigurationAssignment - Get assignments from deviceManagement. The list of group assignments for
// the device configuration profile
func (c DeviceEnrollmentConfigurationAssignmentClient) GetDeviceEnrollmentConfigurationAssignment(ctx context.Context, id beta.DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId, options GetDeviceEnrollmentConfigurationAssignmentOperationOptions) (result GetDeviceEnrollmentConfigurationAssignmentOperationResponse, err error) {
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

	var model beta.EnrollmentConfigurationAssignment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
