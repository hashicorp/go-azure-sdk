package chromeosonboardingsetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateChromeOSOnboardingSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateChromeOSOnboardingSettingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateChromeOSOnboardingSettingOperationOptions() UpdateChromeOSOnboardingSettingOperationOptions {
	return UpdateChromeOSOnboardingSettingOperationOptions{}
}

func (o UpdateChromeOSOnboardingSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateChromeOSOnboardingSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateChromeOSOnboardingSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateChromeOSOnboardingSetting - Update the navigation property chromeOSOnboardingSettings in deviceManagement
func (c ChromeOSOnboardingSettingClient) UpdateChromeOSOnboardingSetting(ctx context.Context, id beta.DeviceManagementChromeOSOnboardingSettingId, input beta.ChromeOSOnboardingSettings, options UpdateChromeOSOnboardingSettingOperationOptions) (result UpdateChromeOSOnboardingSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
