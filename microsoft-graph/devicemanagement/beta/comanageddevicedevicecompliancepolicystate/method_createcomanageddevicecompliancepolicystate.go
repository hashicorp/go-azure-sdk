package comanageddevicedevicecompliancepolicystate

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

type CreateComanagedDeviceCompliancePolicyStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceCompliancePolicyState
}

type CreateComanagedDeviceCompliancePolicyStateOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateComanagedDeviceCompliancePolicyStateOperationOptions() CreateComanagedDeviceCompliancePolicyStateOperationOptions {
	return CreateComanagedDeviceCompliancePolicyStateOperationOptions{}
}

func (o CreateComanagedDeviceCompliancePolicyStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateComanagedDeviceCompliancePolicyStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateComanagedDeviceCompliancePolicyStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateComanagedDeviceCompliancePolicyState - Create new navigation property to deviceCompliancePolicyStates for
// deviceManagement
func (c ComanagedDeviceDeviceCompliancePolicyStateClient) CreateComanagedDeviceCompliancePolicyState(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, input beta.DeviceCompliancePolicyState, options CreateComanagedDeviceCompliancePolicyStateOperationOptions) (result CreateComanagedDeviceCompliancePolicyStateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/deviceCompliancePolicyStates", id.ID()),
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

	var model beta.DeviceCompliancePolicyState
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
