package authenticationphonemethod

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAuthenticationPhoneMethodOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PhoneAuthenticationMethod
}

type CreateAuthenticationPhoneMethodOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAuthenticationPhoneMethodOperationOptions() CreateAuthenticationPhoneMethodOperationOptions {
	return CreateAuthenticationPhoneMethodOperationOptions{}
}

func (o CreateAuthenticationPhoneMethodOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAuthenticationPhoneMethodOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAuthenticationPhoneMethodOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAuthenticationPhoneMethod - Create phoneAuthenticationMethod. Add a new phone authentication method. A user may
// only have one phone of each type, captured in the phoneType property. This means, for example, adding a mobile phone
// to a user with a preexisting mobile phone will fail. Additionally, a user must always have a mobile phone before
// adding an alternateMobile phone. Adding a phone number makes it available for use in both Azure multifactor
// authentication (MFA) and self-service password reset (SSPR), if enabled. Additionally, if a user is enabled by policy
// to use SMS sign-in and a mobile number is added, the system attempts to register the number for use in that system.
func (c AuthenticationPhoneMethodClient) CreateAuthenticationPhoneMethod(ctx context.Context, id beta.UserId, input beta.PhoneAuthenticationMethod, options CreateAuthenticationPhoneMethodOperationOptions) (result CreateAuthenticationPhoneMethodOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/authentication/phoneMethods", id.ID()),
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

	var model beta.PhoneAuthenticationMethod
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
