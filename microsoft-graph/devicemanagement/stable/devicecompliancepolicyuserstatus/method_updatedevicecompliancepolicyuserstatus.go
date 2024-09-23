package devicecompliancepolicyuserstatus

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDeviceCompliancePolicyUserStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDeviceCompliancePolicyUserStatusOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateDeviceCompliancePolicyUserStatusOperationOptions() UpdateDeviceCompliancePolicyUserStatusOperationOptions {
	return UpdateDeviceCompliancePolicyUserStatusOperationOptions{}
}

func (o UpdateDeviceCompliancePolicyUserStatusOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDeviceCompliancePolicyUserStatusOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDeviceCompliancePolicyUserStatusOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDeviceCompliancePolicyUserStatus - Update deviceComplianceUserStatus. Update the properties of a
// deviceComplianceUserStatus object.
func (c DeviceCompliancePolicyUserStatusClient) UpdateDeviceCompliancePolicyUserStatus(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyIdUserStatusId, input stable.DeviceComplianceUserStatus, options UpdateDeviceCompliancePolicyUserStatusOperationOptions) (result UpdateDeviceCompliancePolicyUserStatusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
