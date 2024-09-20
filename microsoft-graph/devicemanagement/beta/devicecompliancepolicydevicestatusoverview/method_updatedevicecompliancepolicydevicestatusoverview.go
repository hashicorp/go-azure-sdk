package devicecompliancepolicydevicestatusoverview

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

type UpdateDeviceCompliancePolicyDeviceStatusOverviewOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDeviceCompliancePolicyDeviceStatusOverviewOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateDeviceCompliancePolicyDeviceStatusOverviewOperationOptions() UpdateDeviceCompliancePolicyDeviceStatusOverviewOperationOptions {
	return UpdateDeviceCompliancePolicyDeviceStatusOverviewOperationOptions{}
}

func (o UpdateDeviceCompliancePolicyDeviceStatusOverviewOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDeviceCompliancePolicyDeviceStatusOverviewOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDeviceCompliancePolicyDeviceStatusOverviewOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDeviceCompliancePolicyDeviceStatusOverview - Update the navigation property deviceStatusOverview in
// deviceManagement
func (c DeviceCompliancePolicyDeviceStatusOverviewClient) UpdateDeviceCompliancePolicyDeviceStatusOverview(ctx context.Context, id beta.DeviceManagementDeviceCompliancePolicyId, input beta.DeviceComplianceDeviceOverview, options UpdateDeviceCompliancePolicyDeviceStatusOverviewOperationOptions) (result UpdateDeviceCompliancePolicyDeviceStatusOverviewOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/deviceStatusOverview", id.ID()),
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
