package authenticationeventsflowexternalusersselfservicesignupeventsflowonauthenticationmethodloadstartonauthenticationmethodloadstartexternalusersselfservicesignupidentityprovider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefOperationOptions() AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefOperationOptions {
	return AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefOperationOptions{}
}

func (o AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRef -
// Add identityProvider (to a user flow). Add an identity provider to an external identities self-service user flow
// represented by an externalUsersSelfServiceSignupEventsFlow object type. The identity provider must first be
// configured in the tenant.
func (c AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderClient) AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRef(ctx context.Context, id stable.IdentityAuthenticationEventsFlowId, input stable.ReferenceCreate, options AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefOperationOptions) (result AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart/onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp/identityProviders/$ref", id.ID()),
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
