package devicecompliancepolicydevicestatusoverview

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

type DeleteDeviceCompliancePolicyDeviceStatusOverviewOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteDeviceCompliancePolicyDeviceStatusOverviewOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteDeviceCompliancePolicyDeviceStatusOverviewOperationOptions() DeleteDeviceCompliancePolicyDeviceStatusOverviewOperationOptions {
	return DeleteDeviceCompliancePolicyDeviceStatusOverviewOperationOptions{}
}

func (o DeleteDeviceCompliancePolicyDeviceStatusOverviewOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteDeviceCompliancePolicyDeviceStatusOverviewOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteDeviceCompliancePolicyDeviceStatusOverviewOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteDeviceCompliancePolicyDeviceStatusOverview - Delete navigation property deviceStatusOverview for
// deviceManagement
func (c DeviceCompliancePolicyDeviceStatusOverviewClient) DeleteDeviceCompliancePolicyDeviceStatusOverview(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyId, options DeleteDeviceCompliancePolicyDeviceStatusOverviewOperationOptions) (result DeleteDeviceCompliancePolicyDeviceStatusOverviewOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/deviceStatusOverview", id.ID()),
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

	return
}
