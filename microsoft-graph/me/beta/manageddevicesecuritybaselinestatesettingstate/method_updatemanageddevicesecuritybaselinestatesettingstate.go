package manageddevicesecuritybaselinestatesettingstate

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateManagedDeviceSecurityBaselineStateSettingStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateManagedDeviceSecurityBaselineStateSettingStateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateManagedDeviceSecurityBaselineStateSettingStateOperationOptions() UpdateManagedDeviceSecurityBaselineStateSettingStateOperationOptions {
	return UpdateManagedDeviceSecurityBaselineStateSettingStateOperationOptions{}
}

func (o UpdateManagedDeviceSecurityBaselineStateSettingStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateManagedDeviceSecurityBaselineStateSettingStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateManagedDeviceSecurityBaselineStateSettingStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateManagedDeviceSecurityBaselineStateSettingState - Update the navigation property settingStates in me
func (c ManagedDeviceSecurityBaselineStateSettingStateClient) UpdateManagedDeviceSecurityBaselineStateSettingState(ctx context.Context, id beta.MeManagedDeviceIdSecurityBaselineStateIdSettingStateId, input beta.SecurityBaselineSettingState, options UpdateManagedDeviceSecurityBaselineStateSettingStateOperationOptions) (result UpdateManagedDeviceSecurityBaselineStateSettingStateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
