package androidforworksetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAndroidForWorkSettingCompleteSignupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateAndroidForWorkSettingCompleteSignupOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAndroidForWorkSettingCompleteSignupOperationOptions() CreateAndroidForWorkSettingCompleteSignupOperationOptions {
	return CreateAndroidForWorkSettingCompleteSignupOperationOptions{}
}

func (o CreateAndroidForWorkSettingCompleteSignupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAndroidForWorkSettingCompleteSignupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAndroidForWorkSettingCompleteSignupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAndroidForWorkSettingCompleteSignup - Invoke action completeSignup
func (c AndroidForWorkSettingClient) CreateAndroidForWorkSettingCompleteSignup(ctx context.Context, input CreateAndroidForWorkSettingCompleteSignupRequest, options CreateAndroidForWorkSettingCompleteSignupOperationOptions) (result CreateAndroidForWorkSettingCompleteSignupOperationResponse, err error) {
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
		Path:          "/deviceManagement/androidForWorkSettings/completeSignup",
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
