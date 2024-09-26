package virtualendpointonpremisesconnection

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateVirtualEndpointOnPremisesConnectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CloudPCOnPremisesConnection
}

type CreateVirtualEndpointOnPremisesConnectionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateVirtualEndpointOnPremisesConnectionOperationOptions() CreateVirtualEndpointOnPremisesConnectionOperationOptions {
	return CreateVirtualEndpointOnPremisesConnectionOperationOptions{}
}

func (o CreateVirtualEndpointOnPremisesConnectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateVirtualEndpointOnPremisesConnectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateVirtualEndpointOnPremisesConnectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateVirtualEndpointOnPremisesConnection - Create cloudPcOnPremisesConnection. Create a new
// cloudPcOnPremisesConnection object for provisioning Cloud PCs.
func (c VirtualEndpointOnPremisesConnectionClient) CreateVirtualEndpointOnPremisesConnection(ctx context.Context, input beta.CloudPCOnPremisesConnection, options CreateVirtualEndpointOnPremisesConnectionOperationOptions) (result CreateVirtualEndpointOnPremisesConnectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/virtualEndpoint/onPremisesConnections",
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

	var model beta.CloudPCOnPremisesConnection
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
