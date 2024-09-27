package virtualendpointprovisioningpolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateVirtualEndpointProvisioningPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateVirtualEndpointProvisioningPolicyOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateVirtualEndpointProvisioningPolicyOperationOptions() UpdateVirtualEndpointProvisioningPolicyOperationOptions {
	return UpdateVirtualEndpointProvisioningPolicyOperationOptions{}
}

func (o UpdateVirtualEndpointProvisioningPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateVirtualEndpointProvisioningPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateVirtualEndpointProvisioningPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateVirtualEndpointProvisioningPolicy - Update cloudPcProvisioningPolicy. Update the properties of a
// cloudPcProvisioningPolicy object.
func (c VirtualEndpointProvisioningPolicyClient) UpdateVirtualEndpointProvisioningPolicy(ctx context.Context, id stable.DeviceManagementVirtualEndpointProvisioningPolicyId, input stable.CloudPCProvisioningPolicy, options UpdateVirtualEndpointProvisioningPolicyOperationOptions) (result UpdateVirtualEndpointProvisioningPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
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
