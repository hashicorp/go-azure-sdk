package deponboardingsettingimportedappledeviceidentity

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

type GetDepOnboardingSettingImportedAppleDeviceIdentitiesCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetDepOnboardingSettingImportedAppleDeviceIdentitiesCountOperationOptions struct {
	Filter *string
	Search *string
}

func DefaultGetDepOnboardingSettingImportedAppleDeviceIdentitiesCountOperationOptions() GetDepOnboardingSettingImportedAppleDeviceIdentitiesCountOperationOptions {
	return GetDepOnboardingSettingImportedAppleDeviceIdentitiesCountOperationOptions{}
}

func (o GetDepOnboardingSettingImportedAppleDeviceIdentitiesCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDepOnboardingSettingImportedAppleDeviceIdentitiesCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetDepOnboardingSettingImportedAppleDeviceIdentitiesCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDepOnboardingSettingImportedAppleDeviceIdentitiesCount - Get the number of the resource
func (c DepOnboardingSettingImportedAppleDeviceIdentityClient) GetDepOnboardingSettingImportedAppleDeviceIdentitiesCount(ctx context.Context, id beta.DeviceManagementDepOnboardingSettingId, options GetDepOnboardingSettingImportedAppleDeviceIdentitiesCountOperationOptions) (result GetDepOnboardingSettingImportedAppleDeviceIdentitiesCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/importedAppleDeviceIdentities/$count", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
