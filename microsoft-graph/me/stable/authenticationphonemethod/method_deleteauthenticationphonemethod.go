package authenticationphonemethod

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteAuthenticationPhoneMethodOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteAuthenticationPhoneMethodOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteAuthenticationPhoneMethodOperationOptions() DeleteAuthenticationPhoneMethodOperationOptions {
	return DeleteAuthenticationPhoneMethodOperationOptions{}
}

func (o DeleteAuthenticationPhoneMethodOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteAuthenticationPhoneMethodOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteAuthenticationPhoneMethodOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteAuthenticationPhoneMethod - Delete phoneAuthenticationMethod. Delete a user's phone authentication method. This
// removes the phone number from the user and they'll no longer be able to use the number for authentication, whether
// via SMS or voice calls. A user can't have an alternateMobile number without a mobile number. If you want to remove a
// mobile number from a user that also has an alternateMobile number, first update the mobile number to the new number,
// then delete the alternateMobile number. If the phone number is the user's default Azure multi-factor authentication
// (MFA) authentication method, it can't be deleted. Have the user change their default authentication method, and then
// delete the number.
func (c AuthenticationPhoneMethodClient) DeleteAuthenticationPhoneMethod(ctx context.Context, id stable.MeAuthenticationPhoneMethodId, options DeleteAuthenticationPhoneMethodOperationOptions) (result DeleteAuthenticationPhoneMethodOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
