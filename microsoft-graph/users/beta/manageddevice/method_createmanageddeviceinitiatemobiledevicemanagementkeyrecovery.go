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

type CreateManagedDeviceInitiateMobileDeviceManagementKeyRecoveryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateManagedDeviceInitiateMobileDeviceManagementKeyRecoveryOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateManagedDeviceInitiateMobileDeviceManagementKeyRecoveryOperationOptions() CreateManagedDeviceInitiateMobileDeviceManagementKeyRecoveryOperationOptions {
	return CreateManagedDeviceInitiateMobileDeviceManagementKeyRecoveryOperationOptions{}
}

func (o CreateManagedDeviceInitiateMobileDeviceManagementKeyRecoveryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateManagedDeviceInitiateMobileDeviceManagementKeyRecoveryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateManagedDeviceInitiateMobileDeviceManagementKeyRecoveryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateManagedDeviceInitiateMobileDeviceManagementKeyRecovery - Invoke action
// initiateMobileDeviceManagementKeyRecovery. Perform MDM key recovery and TPM attestation
func (c ManagedDeviceClient) CreateManagedDeviceInitiateMobileDeviceManagementKeyRecovery(ctx context.Context, id beta.UserIdManagedDeviceId, options CreateManagedDeviceInitiateMobileDeviceManagementKeyRecoveryOperationOptions) (result CreateManagedDeviceInitiateMobileDeviceManagementKeyRecoveryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/initiateMobileDeviceManagementKeyRecovery", id.ID()),
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
