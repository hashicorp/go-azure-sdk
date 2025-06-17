package authenticationeventsflowexternalusersselfservicesignupeventsflowonauthenticationmethodloadstart

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.OnAuthenticationMethodLoadStartHandler
}

type GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartOperationOptions() GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartOperationOptions {
	return GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartOperationOptions{}
}

func (o GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartOperationOptions) ToOData() *odata.Query {
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

func (o GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStart - Get
// onAuthenticationMethodLoadStart property value. Required. The configuration for what to invoke when authentication
// methods are ready to be presented to the user. Must have at least one identity provider linked. Supports $filter
// (eq). See support for filtering on user flows for syntax information.
func (c AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartClient) GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStart(ctx context.Context, id stable.IdentityAuthenticationEventsFlowId, options GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartOperationOptions) (result GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalOnAuthenticationMethodLoadStartHandlerImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
