package virtualendpointonpremisesconnection

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

type UpdateVirtualEndpointOnPremisesConnectionAdDomainPasswordOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateVirtualEndpointOnPremisesConnectionAdDomainPasswordOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateVirtualEndpointOnPremisesConnectionAdDomainPasswordOperationOptions() UpdateVirtualEndpointOnPremisesConnectionAdDomainPasswordOperationOptions {
	return UpdateVirtualEndpointOnPremisesConnectionAdDomainPasswordOperationOptions{}
}

func (o UpdateVirtualEndpointOnPremisesConnectionAdDomainPasswordOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateVirtualEndpointOnPremisesConnectionAdDomainPasswordOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateVirtualEndpointOnPremisesConnectionAdDomainPasswordOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateVirtualEndpointOnPremisesConnectionAdDomainPassword - Invoke action updateAdDomainPassword. Update the Active
// Directory domain password for a cloudPcOnPremisesConnection object. This API is supported when the type of the
// cloudPcOnPremisesConnection object is hybridAzureADJoin.
func (c VirtualEndpointOnPremisesConnectionClient) UpdateVirtualEndpointOnPremisesConnectionAdDomainPassword(ctx context.Context, id beta.DeviceManagementVirtualEndpointOnPremisesConnectionId, input UpdateVirtualEndpointOnPremisesConnectionAdDomainPasswordRequest, options UpdateVirtualEndpointOnPremisesConnectionAdDomainPasswordOperationOptions) (result UpdateVirtualEndpointOnPremisesConnectionAdDomainPasswordOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/updateAdDomainPassword", id.ID()),
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
