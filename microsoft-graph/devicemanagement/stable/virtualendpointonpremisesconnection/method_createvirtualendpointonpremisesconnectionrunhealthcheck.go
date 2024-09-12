package virtualendpointonpremisesconnection

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateVirtualEndpointOnPremisesConnectionRunHealthCheckOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateVirtualEndpointOnPremisesConnectionRunHealthCheck - Invoke action runHealthChecks. Run health checks on the
// cloudPcOnPremisesConnection object. It triggers a new health check for the cloudPcOnPremisesConnection
// (../resources/cloudpconpremisesconnection.md) object and changes the healthCheckStatus and [healthCheckStatusDetail
// properties when check finished.
func (c VirtualEndpointOnPremisesConnectionClient) CreateVirtualEndpointOnPremisesConnectionRunHealthCheck(ctx context.Context, id stable.DeviceManagementVirtualEndpointOnPremisesConnectionId) (result CreateVirtualEndpointOnPremisesConnectionRunHealthCheckOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/runHealthChecks", id.ID()),
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
