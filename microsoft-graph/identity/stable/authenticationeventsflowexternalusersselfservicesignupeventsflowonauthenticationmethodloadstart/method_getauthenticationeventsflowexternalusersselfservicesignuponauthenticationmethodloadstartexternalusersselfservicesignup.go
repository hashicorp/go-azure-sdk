package authenticationeventsflowexternalusersselfservicesignupeventsflowonauthenticationmethodloadstart

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

type GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.OnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp
}

type GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpOperationOptions() GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpOperationOptions {
	return GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpOperationOptions{}
}

func (o GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp
// - Get the item of type microsoft.graph.onAuthenticationMethodLoadStartHandler as
// microsoft.graph.onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp
func (c AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartClient) GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp(ctx context.Context, id stable.IdentityAuthenticationEventsFlowId, options GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpOperationOptions) (result GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart/onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp", id.ID()),
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

	var model stable.OnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
