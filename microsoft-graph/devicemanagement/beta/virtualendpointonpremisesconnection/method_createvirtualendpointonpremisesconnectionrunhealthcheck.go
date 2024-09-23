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

type CreateVirtualEndpointOnPremisesConnectionRunHealthCheckOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateVirtualEndpointOnPremisesConnectionRunHealthCheckOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateVirtualEndpointOnPremisesConnectionRunHealthCheckOperationOptions() CreateVirtualEndpointOnPremisesConnectionRunHealthCheckOperationOptions {
	return CreateVirtualEndpointOnPremisesConnectionRunHealthCheckOperationOptions{}
}

func (o CreateVirtualEndpointOnPremisesConnectionRunHealthCheckOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateVirtualEndpointOnPremisesConnectionRunHealthCheckOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateVirtualEndpointOnPremisesConnectionRunHealthCheckOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateVirtualEndpointOnPremisesConnectionRunHealthCheck - Invoke action runHealthChecks. Run health checks on the
// cloudPcOnPremisesConnection object. It triggers a new health check for this cloudPcOnPremisesConnection object and
// change the healthCheckStatus and healthCheckStatusDetails properties when check finished.
func (c VirtualEndpointOnPremisesConnectionClient) CreateVirtualEndpointOnPremisesConnectionRunHealthCheck(ctx context.Context, id beta.DeviceManagementVirtualEndpointOnPremisesConnectionId, options CreateVirtualEndpointOnPremisesConnectionRunHealthCheckOperationOptions) (result CreateVirtualEndpointOnPremisesConnectionRunHealthCheckOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/runHealthChecks", id.ID()),
		RetryFunc:     options.RetryFunc,
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
