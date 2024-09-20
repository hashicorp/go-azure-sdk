package androidmanagedstoreaccountenterprisesetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrlOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *RequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrlResult
}

type RequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrlOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultRequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrlOperationOptions() RequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrlOperationOptions {
	return RequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrlOperationOptions{}
}

func (o RequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrlOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrlOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrlOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrl - Invoke action requestSignupUrl
func (c AndroidManagedStoreAccountEnterpriseSettingClient) RequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrl(ctx context.Context, input RequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrlRequest, options RequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrlOperationOptions) (result RequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrlOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/androidManagedStoreAccountEnterpriseSettings/requestSignupUrl",
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

	var model RequestAndroidManagedStoreAccountEnterpriseSettingsSignupUrlResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
