package comanageddevice

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

type CreateComanagedDeviceInitiateDeviceAttestationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateComanagedDeviceInitiateDeviceAttestationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateComanagedDeviceInitiateDeviceAttestationOperationOptions() CreateComanagedDeviceInitiateDeviceAttestationOperationOptions {
	return CreateComanagedDeviceInitiateDeviceAttestationOperationOptions{}
}

func (o CreateComanagedDeviceInitiateDeviceAttestationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateComanagedDeviceInitiateDeviceAttestationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateComanagedDeviceInitiateDeviceAttestationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateComanagedDeviceInitiateDeviceAttestation - Invoke action initiateDeviceAttestation. Perform Device Attestation
func (c ComanagedDeviceClient) CreateComanagedDeviceInitiateDeviceAttestation(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options CreateComanagedDeviceInitiateDeviceAttestationOperationOptions) (result CreateComanagedDeviceInitiateDeviceAttestationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/initiateDeviceAttestation", id.ID()),
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
