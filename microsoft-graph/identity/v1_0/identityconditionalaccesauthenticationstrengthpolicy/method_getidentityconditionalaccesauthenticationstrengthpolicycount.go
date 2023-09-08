package identityconditionalaccesauthenticationstrengthpolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetIdentityConditionalAccesAuthenticationStrengthPolicyCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// GetIdentityConditionalAccesAuthenticationStrengthPolicyCount ...
func (c IdentityConditionalAccesAuthenticationStrengthPolicyClient) GetIdentityConditionalAccesAuthenticationStrengthPolicyCount(ctx context.Context) (result GetIdentityConditionalAccesAuthenticationStrengthPolicyCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/identity/conditionalAccess/authenticationStrength/policies/$count",
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
