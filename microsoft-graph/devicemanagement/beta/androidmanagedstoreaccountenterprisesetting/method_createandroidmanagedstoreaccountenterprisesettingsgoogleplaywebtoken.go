package androidmanagedstoreaccountenterprisesetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebTokenOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *CreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebTokenResult
}

type CreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebTokenOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebTokenOperationOptions() CreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebTokenOperationOptions {
	return CreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebTokenOperationOptions{}
}

func (o CreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebTokenOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebTokenOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebTokenOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebToken - Invoke action createGooglePlayWebToken.
// Generates a web token that is used in an embeddable component.
func (c AndroidManagedStoreAccountEnterpriseSettingClient) CreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebToken(ctx context.Context, input CreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebTokenRequest, options CreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebTokenOperationOptions) (result CreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebTokenOperationResponse, err error) {
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
		Path:          "/deviceManagement/androidManagedStoreAccountEnterpriseSettings/createGooglePlayWebToken",
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

	var model CreateAndroidManagedStoreAccountEnterpriseSettingsGooglePlayWebTokenResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
