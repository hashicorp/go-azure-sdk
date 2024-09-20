package windowsqualityupdatepolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateWindowsQualityUpdatePolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateWindowsQualityUpdatePolicyOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateWindowsQualityUpdatePolicyOperationOptions() UpdateWindowsQualityUpdatePolicyOperationOptions {
	return UpdateWindowsQualityUpdatePolicyOperationOptions{}
}

func (o UpdateWindowsQualityUpdatePolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateWindowsQualityUpdatePolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateWindowsQualityUpdatePolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateWindowsQualityUpdatePolicy - Update the navigation property windowsQualityUpdatePolicies in deviceManagement
func (c WindowsQualityUpdatePolicyClient) UpdateWindowsQualityUpdatePolicy(ctx context.Context, id beta.DeviceManagementWindowsQualityUpdatePolicyId, input beta.WindowsQualityUpdatePolicy, options UpdateWindowsQualityUpdatePolicyOperationOptions) (result UpdateWindowsQualityUpdatePolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
