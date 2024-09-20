package devicecompliancepolicydevicestatus

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

type CreateDeviceCompliancePolicyDeviceStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.DeviceComplianceDeviceStatus
}

type CreateDeviceCompliancePolicyDeviceStatusOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateDeviceCompliancePolicyDeviceStatusOperationOptions() CreateDeviceCompliancePolicyDeviceStatusOperationOptions {
	return CreateDeviceCompliancePolicyDeviceStatusOperationOptions{}
}

func (o CreateDeviceCompliancePolicyDeviceStatusOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDeviceCompliancePolicyDeviceStatusOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDeviceCompliancePolicyDeviceStatusOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDeviceCompliancePolicyDeviceStatus - Create deviceComplianceDeviceStatus. Create a new
// deviceComplianceDeviceStatus object.
func (c DeviceCompliancePolicyDeviceStatusClient) CreateDeviceCompliancePolicyDeviceStatus(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyId, input stable.DeviceComplianceDeviceStatus, options CreateDeviceCompliancePolicyDeviceStatusOperationOptions) (result CreateDeviceCompliancePolicyDeviceStatusOperationResponse, err error) {
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

	var model stable.DeviceComplianceDeviceStatus
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
