package manageddevice

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

type ReprovisionManagedDeviceCloudPCOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ReprovisionManagedDeviceCloudPCOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultReprovisionManagedDeviceCloudPCOperationOptions() ReprovisionManagedDeviceCloudPCOperationOptions {
	return ReprovisionManagedDeviceCloudPCOperationOptions{}
}

func (o ReprovisionManagedDeviceCloudPCOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReprovisionManagedDeviceCloudPCOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ReprovisionManagedDeviceCloudPCOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ReprovisionManagedDeviceCloudPC - Invoke action reprovisionCloudPc. Reprovision a Cloud PC with an Intune managed
// device ID.
func (c ManagedDeviceClient) ReprovisionManagedDeviceCloudPC(ctx context.Context, id beta.DeviceManagementManagedDeviceId, options ReprovisionManagedDeviceCloudPCOperationOptions) (result ReprovisionManagedDeviceCloudPCOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/reprovisionCloudPc", id.ID()),
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
