package androidforworksetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequestAndroidForWorkSettingsSignupUrlOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *RequestAndroidForWorkSettingsSignupUrlResult
}

type RequestAndroidForWorkSettingsSignupUrlOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRequestAndroidForWorkSettingsSignupUrlOperationOptions() RequestAndroidForWorkSettingsSignupUrlOperationOptions {
	return RequestAndroidForWorkSettingsSignupUrlOperationOptions{}
}

func (o RequestAndroidForWorkSettingsSignupUrlOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RequestAndroidForWorkSettingsSignupUrlOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RequestAndroidForWorkSettingsSignupUrlOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RequestAndroidForWorkSettingsSignupUrl - Invoke action requestSignupUrl
func (c AndroidForWorkSettingClient) RequestAndroidForWorkSettingsSignupUrl(ctx context.Context, input RequestAndroidForWorkSettingsSignupUrlRequest, options RequestAndroidForWorkSettingsSignupUrlOperationOptions) (result RequestAndroidForWorkSettingsSignupUrlOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/androidForWorkSettings/requestSignupUrl",
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

	var model RequestAndroidForWorkSettingsSignupUrlResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
