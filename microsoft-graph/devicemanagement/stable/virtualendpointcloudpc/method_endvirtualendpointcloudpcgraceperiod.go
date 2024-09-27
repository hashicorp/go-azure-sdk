package virtualendpointcloudpc

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

type EndVirtualEndpointCloudPCGracePeriodOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type EndVirtualEndpointCloudPCGracePeriodOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultEndVirtualEndpointCloudPCGracePeriodOperationOptions() EndVirtualEndpointCloudPCGracePeriodOperationOptions {
	return EndVirtualEndpointCloudPCGracePeriodOperationOptions{}
}

func (o EndVirtualEndpointCloudPCGracePeriodOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o EndVirtualEndpointCloudPCGracePeriodOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o EndVirtualEndpointCloudPCGracePeriodOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// EndVirtualEndpointCloudPCGracePeriod - Invoke action endGracePeriod. End the grace period for a specific cloudPC
// object. The grace period is triggered when the Cloud PC license is removed or the provisioning policy is unassigned.
// It allows users to access Cloud PCs for up to seven days before deprovisioning occurs. Ending the grace period
// immediately deprovisions the Cloud PC without waiting the seven days.
func (c VirtualEndpointCloudPCClient) EndVirtualEndpointCloudPCGracePeriod(ctx context.Context, id stable.DeviceManagementVirtualEndpointCloudPCId, options EndVirtualEndpointCloudPCGracePeriodOperationOptions) (result EndVirtualEndpointCloudPCGracePeriodOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/endGracePeriod", id.ID()),
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
