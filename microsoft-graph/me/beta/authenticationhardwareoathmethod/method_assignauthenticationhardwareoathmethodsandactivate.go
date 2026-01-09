package authenticationhardwareoathmethod

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignAuthenticationHardwareOathMethodsAndActivateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AssignAuthenticationHardwareOathMethodsAndActivateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAssignAuthenticationHardwareOathMethodsAndActivateOperationOptions() AssignAuthenticationHardwareOathMethodsAndActivateOperationOptions {
	return AssignAuthenticationHardwareOathMethodsAndActivateOperationOptions{}
}

func (o AssignAuthenticationHardwareOathMethodsAndActivateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AssignAuthenticationHardwareOathMethodsAndActivateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AssignAuthenticationHardwareOathMethodsAndActivateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AssignAuthenticationHardwareOathMethodsAndActivate - Invoke action assignAndActivate. Assign and activate a hardware
// token at the same time. This operation requires the device ID to activate it.
func (c AuthenticationHardwareOathMethodClient) AssignAuthenticationHardwareOathMethodsAndActivate(ctx context.Context, input AssignAuthenticationHardwareOathMethodsAndActivateRequest, options AssignAuthenticationHardwareOathMethodsAndActivateOperationOptions) (result AssignAuthenticationHardwareOathMethodsAndActivateOperationResponse, err error) {
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
		Path:          "/me/authentication/hardwareOathMethods/assignAndActivate",
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
