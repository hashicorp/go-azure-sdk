package authenticationphonemethod

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAuthenticationPhoneMethodOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// UpdateAuthenticationPhoneMethod - Update phoneAuthenticationMethod. Update a user's phone number associated with a
// phone authentication method object. You can't change a phone's type. To change a phone's type, add a new number of
// the desired type and then delete the object with the original type. If a user is enabled by policy to use SMS to sign
// in and the mobile number is changed, the system will attempt to register the number for use in that system.
func (c AuthenticationPhoneMethodClient) UpdateAuthenticationPhoneMethod(ctx context.Context, id stable.UserIdAuthenticationPhoneMethodId, input stable.PhoneAuthenticationMethod) (result UpdateAuthenticationPhoneMethodOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPatch,
		Path:       id.ID(),
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
