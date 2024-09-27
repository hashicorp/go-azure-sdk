package manageddevicesecuritybaselinestatesettingstate

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

type CreateManagedDeviceSecurityBaselineStateSettingStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.SecurityBaselineSettingState
}

type CreateManagedDeviceSecurityBaselineStateSettingStateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateManagedDeviceSecurityBaselineStateSettingStateOperationOptions() CreateManagedDeviceSecurityBaselineStateSettingStateOperationOptions {
	return CreateManagedDeviceSecurityBaselineStateSettingStateOperationOptions{}
}

func (o CreateManagedDeviceSecurityBaselineStateSettingStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateManagedDeviceSecurityBaselineStateSettingStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateManagedDeviceSecurityBaselineStateSettingStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateManagedDeviceSecurityBaselineStateSettingState - Create new navigation property to settingStates for me
func (c ManagedDeviceSecurityBaselineStateSettingStateClient) CreateManagedDeviceSecurityBaselineStateSettingState(ctx context.Context, id beta.MeManagedDeviceIdSecurityBaselineStateId, input beta.SecurityBaselineSettingState, options CreateManagedDeviceSecurityBaselineStateSettingStateOperationOptions) (result CreateManagedDeviceSecurityBaselineStateSettingStateOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/settingStates", id.ID()),
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

	var model beta.SecurityBaselineSettingState
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
