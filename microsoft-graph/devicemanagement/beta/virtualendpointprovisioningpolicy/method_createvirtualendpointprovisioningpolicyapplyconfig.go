package virtualendpointprovisioningpolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateVirtualEndpointProvisioningPolicyApplyConfigOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateVirtualEndpointProvisioningPolicyApplyConfigOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateVirtualEndpointProvisioningPolicyApplyConfigOperationOptions() CreateVirtualEndpointProvisioningPolicyApplyConfigOperationOptions {
	return CreateVirtualEndpointProvisioningPolicyApplyConfigOperationOptions{}
}

func (o CreateVirtualEndpointProvisioningPolicyApplyConfigOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateVirtualEndpointProvisioningPolicyApplyConfigOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateVirtualEndpointProvisioningPolicyApplyConfigOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateVirtualEndpointProvisioningPolicyApplyConfig - Invoke action applyConfig. Update the provisioning policy
// configuration for a set of Cloud PC devices by their IDs. This method supports retry and allows you to apply the
// configuration to a subset of Cloud PCs initially to test.
func (c VirtualEndpointProvisioningPolicyClient) CreateVirtualEndpointProvisioningPolicyApplyConfig(ctx context.Context, input CreateVirtualEndpointProvisioningPolicyApplyConfigRequest, options CreateVirtualEndpointProvisioningPolicyApplyConfigOperationOptions) (result CreateVirtualEndpointProvisioningPolicyApplyConfigOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/virtualEndpoint/provisioningPolicies/applyConfig",
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
