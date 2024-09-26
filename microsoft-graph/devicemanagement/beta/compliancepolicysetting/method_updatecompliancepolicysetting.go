package compliancepolicysetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateCompliancePolicySettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateCompliancePolicySettingOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateCompliancePolicySettingOperationOptions() UpdateCompliancePolicySettingOperationOptions {
	return UpdateCompliancePolicySettingOperationOptions{}
}

func (o UpdateCompliancePolicySettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateCompliancePolicySettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateCompliancePolicySettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateCompliancePolicySetting - Update the navigation property settings in deviceManagement
func (c CompliancePolicySettingClient) UpdateCompliancePolicySetting(ctx context.Context, id beta.DeviceManagementCompliancePolicyIdSettingId, input beta.DeviceManagementConfigurationSetting, options UpdateCompliancePolicySettingOperationOptions) (result UpdateCompliancePolicySettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
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
