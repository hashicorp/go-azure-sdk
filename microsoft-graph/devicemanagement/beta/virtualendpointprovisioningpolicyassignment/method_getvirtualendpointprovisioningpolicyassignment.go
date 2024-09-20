package virtualendpointprovisioningpolicyassignment

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVirtualEndpointProvisioningPolicyAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CloudPCProvisioningPolicyAssignment
}

type GetVirtualEndpointProvisioningPolicyAssignmentOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetVirtualEndpointProvisioningPolicyAssignmentOperationOptions() GetVirtualEndpointProvisioningPolicyAssignmentOperationOptions {
	return GetVirtualEndpointProvisioningPolicyAssignmentOperationOptions{}
}

func (o GetVirtualEndpointProvisioningPolicyAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetVirtualEndpointProvisioningPolicyAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetVirtualEndpointProvisioningPolicyAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetVirtualEndpointProvisioningPolicyAssignment - Get assignments from deviceManagement. A defined collection of
// provisioning policy assignments. Represents the set of Microsoft 365 groups and security groups in Microsoft Entra ID
// that have provisioning policy assigned. Returned only on $expand. For an example about how to get the assignments
// relationship, see Get cloudPcProvisioningPolicy.
func (c VirtualEndpointProvisioningPolicyAssignmentClient) GetVirtualEndpointProvisioningPolicyAssignment(ctx context.Context, id beta.DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId, options GetVirtualEndpointProvisioningPolicyAssignmentOperationOptions) (result GetVirtualEndpointProvisioningPolicyAssignmentOperationResponse, err error) {
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

	var model beta.CloudPCProvisioningPolicyAssignment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
