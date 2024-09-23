package virtualendpointcrosscloudgovernmentorganizationmapping

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationOptions() UpdateVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationOptions {
	return UpdateVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationOptions{}
}

func (o UpdateVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateVirtualEndpointCrossCloudGovernmentOrganizationMapping - Update the navigation property
// crossCloudGovernmentOrganizationMapping in deviceManagement
func (c VirtualEndpointCrossCloudGovernmentOrganizationMappingClient) UpdateVirtualEndpointCrossCloudGovernmentOrganizationMapping(ctx context.Context, input beta.CloudPCCrossCloudGovernmentOrganizationMapping, options UpdateVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationOptions) (result UpdateVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/deviceManagement/virtualEndpoint/crossCloudGovernmentOrganizationMapping",
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
