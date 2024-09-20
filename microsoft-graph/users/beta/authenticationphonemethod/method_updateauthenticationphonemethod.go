package authenticationphonemethod

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAuthenticationPhoneMethodOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAuthenticationPhoneMethodOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateAuthenticationPhoneMethodOperationOptions() UpdateAuthenticationPhoneMethodOperationOptions {
	return UpdateAuthenticationPhoneMethodOperationOptions{}
}

func (o UpdateAuthenticationPhoneMethodOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAuthenticationPhoneMethodOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAuthenticationPhoneMethodOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAuthenticationPhoneMethod - Update phoneAuthenticationMethod. Update the phone number associated with a phone
// authentication method. You can't change a phone's type. To change a phone's type, add a new number of the desired
// type and then delete the object with the original type. If a user is enabled by policy to use SMS to sign in and the
// mobile number is changed, the system attempts to register the number for use in that system.
func (c AuthenticationPhoneMethodClient) UpdateAuthenticationPhoneMethod(ctx context.Context, id beta.UserIdAuthenticationPhoneMethodId, input beta.PhoneAuthenticationMethod, options UpdateAuthenticationPhoneMethodOperationOptions) (result UpdateAuthenticationPhoneMethodOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
