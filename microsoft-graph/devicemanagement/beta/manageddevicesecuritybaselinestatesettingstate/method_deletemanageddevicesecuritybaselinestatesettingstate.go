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

type DeleteManagedDeviceSecurityBaselineStateSettingStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteManagedDeviceSecurityBaselineStateSettingStateOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteManagedDeviceSecurityBaselineStateSettingStateOperationOptions() DeleteManagedDeviceSecurityBaselineStateSettingStateOperationOptions {
	return DeleteManagedDeviceSecurityBaselineStateSettingStateOperationOptions{}
}

func (o DeleteManagedDeviceSecurityBaselineStateSettingStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteManagedDeviceSecurityBaselineStateSettingStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteManagedDeviceSecurityBaselineStateSettingStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteManagedDeviceSecurityBaselineStateSettingState - Delete navigation property settingStates for deviceManagement
func (c ManagedDeviceSecurityBaselineStateSettingStateClient) DeleteManagedDeviceSecurityBaselineStateSettingState(ctx context.Context, id beta.DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId, options DeleteManagedDeviceSecurityBaselineStateSettingStateOperationOptions) (result DeleteManagedDeviceSecurityBaselineStateSettingStateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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
