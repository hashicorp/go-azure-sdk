package virtualendpointserviceplan

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateVirtualEndpointServicePlanOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateVirtualEndpointServicePlanOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateVirtualEndpointServicePlanOperationOptions() UpdateVirtualEndpointServicePlanOperationOptions {
	return UpdateVirtualEndpointServicePlanOperationOptions{}
}

func (o UpdateVirtualEndpointServicePlanOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateVirtualEndpointServicePlanOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateVirtualEndpointServicePlanOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateVirtualEndpointServicePlan - Update the navigation property servicePlans in deviceManagement
func (c VirtualEndpointServicePlanClient) UpdateVirtualEndpointServicePlan(ctx context.Context, id beta.DeviceManagementVirtualEndpointServicePlanId, input beta.CloudPCServicePlan, options UpdateVirtualEndpointServicePlanOperationOptions) (result UpdateVirtualEndpointServicePlanOperationResponse, err error) {
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
