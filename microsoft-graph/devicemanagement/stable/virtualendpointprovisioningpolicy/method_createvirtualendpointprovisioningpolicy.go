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

type CreateVirtualEndpointProvisioningPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.CloudPCProvisioningPolicy
}

type CreateVirtualEndpointProvisioningPolicyOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateVirtualEndpointProvisioningPolicyOperationOptions() CreateVirtualEndpointProvisioningPolicyOperationOptions {
	return CreateVirtualEndpointProvisioningPolicyOperationOptions{}
}

func (o CreateVirtualEndpointProvisioningPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateVirtualEndpointProvisioningPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateVirtualEndpointProvisioningPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateVirtualEndpointProvisioningPolicy - Create cloudPcProvisioningPolicy. Create a new cloudPcProvisioningPolicy
// object.
func (c VirtualEndpointProvisioningPolicyClient) CreateVirtualEndpointProvisioningPolicy(ctx context.Context, input stable.CloudPCProvisioningPolicy, options CreateVirtualEndpointProvisioningPolicyOperationOptions) (result CreateVirtualEndpointProvisioningPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/virtualEndpoint/provisioningPolicies",
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

	var model stable.CloudPCProvisioningPolicy
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
