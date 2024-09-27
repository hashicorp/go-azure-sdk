package comanageddevicesecuritybaselinestate

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

type CreateComanagedDeviceSecurityBaselineStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.SecurityBaselineState
}

type CreateComanagedDeviceSecurityBaselineStateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateComanagedDeviceSecurityBaselineStateOperationOptions() CreateComanagedDeviceSecurityBaselineStateOperationOptions {
	return CreateComanagedDeviceSecurityBaselineStateOperationOptions{}
}

func (o CreateComanagedDeviceSecurityBaselineStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateComanagedDeviceSecurityBaselineStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateComanagedDeviceSecurityBaselineStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateComanagedDeviceSecurityBaselineState - Create new navigation property to securityBaselineStates for
// deviceManagement
func (c ComanagedDeviceSecurityBaselineStateClient) CreateComanagedDeviceSecurityBaselineState(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, input beta.SecurityBaselineState, options CreateComanagedDeviceSecurityBaselineStateOperationOptions) (result CreateComanagedDeviceSecurityBaselineStateOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/securityBaselineStates", id.ID()),
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

	var model beta.SecurityBaselineState
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
