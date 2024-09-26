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

type CreateManagedDeviceInitiateOnDemandProactiveRemediationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateManagedDeviceInitiateOnDemandProactiveRemediationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateManagedDeviceInitiateOnDemandProactiveRemediationOperationOptions() CreateManagedDeviceInitiateOnDemandProactiveRemediationOperationOptions {
	return CreateManagedDeviceInitiateOnDemandProactiveRemediationOperationOptions{}
}

func (o CreateManagedDeviceInitiateOnDemandProactiveRemediationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateManagedDeviceInitiateOnDemandProactiveRemediationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateManagedDeviceInitiateOnDemandProactiveRemediationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateManagedDeviceInitiateOnDemandProactiveRemediation - Invoke action initiateOnDemandProactiveRemediation. Perform
// On Demand Proactive Remediation
func (c ManagedDeviceClient) CreateManagedDeviceInitiateOnDemandProactiveRemediation(ctx context.Context, id beta.DeviceManagementManagedDeviceId, input CreateManagedDeviceInitiateOnDemandProactiveRemediationRequest, options CreateManagedDeviceInitiateOnDemandProactiveRemediationOperationOptions) (result CreateManagedDeviceInitiateOnDemandProactiveRemediationOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/initiateOnDemandProactiveRemediation", id.ID()),
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
