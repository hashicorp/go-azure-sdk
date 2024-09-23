package androidmanagedstoreaccountenterprisesetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAndroidManagedStoreAccountEnterpriseSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAndroidManagedStoreAccountEnterpriseSettingOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateAndroidManagedStoreAccountEnterpriseSettingOperationOptions() UpdateAndroidManagedStoreAccountEnterpriseSettingOperationOptions {
	return UpdateAndroidManagedStoreAccountEnterpriseSettingOperationOptions{}
}

func (o UpdateAndroidManagedStoreAccountEnterpriseSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAndroidManagedStoreAccountEnterpriseSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAndroidManagedStoreAccountEnterpriseSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAndroidManagedStoreAccountEnterpriseSetting - Update the navigation property
// androidManagedStoreAccountEnterpriseSettings in deviceManagement
func (c AndroidManagedStoreAccountEnterpriseSettingClient) UpdateAndroidManagedStoreAccountEnterpriseSetting(ctx context.Context, input beta.AndroidManagedStoreAccountEnterpriseSettings, options UpdateAndroidManagedStoreAccountEnterpriseSettingOperationOptions) (result UpdateAndroidManagedStoreAccountEnterpriseSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/deviceManagement/androidManagedStoreAccountEnterpriseSettings",
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
