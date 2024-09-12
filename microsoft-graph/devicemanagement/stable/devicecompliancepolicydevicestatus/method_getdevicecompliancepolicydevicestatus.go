package devicecompliancepolicydevicestatus

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceCompliancePolicyDeviceStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.DeviceComplianceDeviceStatus
}

type GetDeviceCompliancePolicyDeviceStatusOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetDeviceCompliancePolicyDeviceStatusOperationOptions() GetDeviceCompliancePolicyDeviceStatusOperationOptions {
	return GetDeviceCompliancePolicyDeviceStatusOperationOptions{}
}

func (o GetDeviceCompliancePolicyDeviceStatusOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceCompliancePolicyDeviceStatusOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDeviceCompliancePolicyDeviceStatusOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceCompliancePolicyDeviceStatus - Get deviceComplianceDeviceStatus. Read properties and relationships of the
// deviceComplianceDeviceStatus object.
func (c DeviceCompliancePolicyDeviceStatusClient) GetDeviceCompliancePolicyDeviceStatus(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyIdDeviceStatusId, options GetDeviceCompliancePolicyDeviceStatusOperationOptions) (result GetDeviceCompliancePolicyDeviceStatusOperationResponse, err error) {
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

	var model stable.DeviceComplianceDeviceStatus
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
