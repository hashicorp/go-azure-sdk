package authenticationhardwareoathmethod

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignAuthenticationHardwareOathMethodsAndActivateBySerialNumberOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AssignAuthenticationHardwareOathMethodsAndActivateBySerialNumberOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAssignAuthenticationHardwareOathMethodsAndActivateBySerialNumberOperationOptions() AssignAuthenticationHardwareOathMethodsAndActivateBySerialNumberOperationOptions {
	return AssignAuthenticationHardwareOathMethodsAndActivateBySerialNumberOperationOptions{}
}

func (o AssignAuthenticationHardwareOathMethodsAndActivateBySerialNumberOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AssignAuthenticationHardwareOathMethodsAndActivateBySerialNumberOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AssignAuthenticationHardwareOathMethodsAndActivateBySerialNumberOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AssignAuthenticationHardwareOathMethodsAndActivateBySerialNumber - Invoke action assignAndActivateBySerialNumber.
// Assign and activate a hardware token at the same time by hardware token serial number.
func (c AuthenticationHardwareOathMethodClient) AssignAuthenticationHardwareOathMethodsAndActivateBySerialNumber(ctx context.Context, input AssignAuthenticationHardwareOathMethodsAndActivateBySerialNumberRequest, options AssignAuthenticationHardwareOathMethodsAndActivateBySerialNumberOperationOptions) (result AssignAuthenticationHardwareOathMethodsAndActivateBySerialNumberOperationResponse, err error) {
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
		Path:          "/me/authentication/hardwareOathMethods/assignAndActivateBySerialNumber",
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
