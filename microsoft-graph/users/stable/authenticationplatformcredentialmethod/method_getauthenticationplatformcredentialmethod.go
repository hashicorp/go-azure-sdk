package authenticationplatformcredentialmethod

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAuthenticationPlatformCredentialMethodOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.PlatformCredentialAuthenticationMethod
}

type GetAuthenticationPlatformCredentialMethodOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetAuthenticationPlatformCredentialMethodOperationOptions() GetAuthenticationPlatformCredentialMethodOperationOptions {
	return GetAuthenticationPlatformCredentialMethodOperationOptions{}
}

func (o GetAuthenticationPlatformCredentialMethodOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationPlatformCredentialMethodOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetAuthenticationPlatformCredentialMethodOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationPlatformCredentialMethod - Get platformCredentialMethods from users. Represents a platform
// credential instance registered to a user on Mac OS.
func (c AuthenticationPlatformCredentialMethodClient) GetAuthenticationPlatformCredentialMethod(ctx context.Context, id stable.UserIdAuthenticationPlatformCredentialMethodId, options GetAuthenticationPlatformCredentialMethodOperationOptions) (result GetAuthenticationPlatformCredentialMethodOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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

	var model stable.PlatformCredentialAuthenticationMethod
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
