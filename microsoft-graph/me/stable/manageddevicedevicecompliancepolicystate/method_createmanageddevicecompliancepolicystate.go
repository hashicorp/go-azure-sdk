package manageddevicedevicecompliancepolicystate

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

type CreateManagedDeviceCompliancePolicyStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.DeviceCompliancePolicyState
}

type CreateManagedDeviceCompliancePolicyStateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateManagedDeviceCompliancePolicyStateOperationOptions() CreateManagedDeviceCompliancePolicyStateOperationOptions {
	return CreateManagedDeviceCompliancePolicyStateOperationOptions{}
}

func (o CreateManagedDeviceCompliancePolicyStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateManagedDeviceCompliancePolicyStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateManagedDeviceCompliancePolicyStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateManagedDeviceCompliancePolicyState - Create new navigation property to deviceCompliancePolicyStates for me
func (c ManagedDeviceDeviceCompliancePolicyStateClient) CreateManagedDeviceCompliancePolicyState(ctx context.Context, id stable.MeManagedDeviceId, input stable.DeviceCompliancePolicyState, options CreateManagedDeviceCompliancePolicyStateOperationOptions) (result CreateManagedDeviceCompliancePolicyStateOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/deviceCompliancePolicyStates", id.ID()),
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

	var model stable.DeviceCompliancePolicyState
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
