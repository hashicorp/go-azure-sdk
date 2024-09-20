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

type DeleteVirtualEndpointOnPremisesConnectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteVirtualEndpointOnPremisesConnectionOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultDeleteVirtualEndpointOnPremisesConnectionOperationOptions() DeleteVirtualEndpointOnPremisesConnectionOperationOptions {
	return DeleteVirtualEndpointOnPremisesConnectionOperationOptions{}
}

func (o DeleteVirtualEndpointOnPremisesConnectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteVirtualEndpointOnPremisesConnectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteVirtualEndpointOnPremisesConnectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteVirtualEndpointOnPremisesConnection - Delete cloudPcOnPremisesConnection. Delete a specific
// cloudPcOnPremisesConnection object. When you delete an Azure network connection, permissions to the service are
// removed from the specified Azure resources. You can't delete an Azure network connection when it's in use, as
// indicated by the inUse property.
func (c VirtualEndpointOnPremisesConnectionClient) DeleteVirtualEndpointOnPremisesConnection(ctx context.Context, id beta.DeviceManagementVirtualEndpointOnPremisesConnectionId, options DeleteVirtualEndpointOnPremisesConnectionOperationOptions) (result DeleteVirtualEndpointOnPremisesConnectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
