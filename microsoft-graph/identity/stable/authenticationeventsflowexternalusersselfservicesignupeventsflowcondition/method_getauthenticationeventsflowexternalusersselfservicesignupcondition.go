package authenticationeventsflowexternalusersselfservicesignupeventsflowcondition

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

type GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpConditionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AuthenticationConditions
}

type GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpConditionOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetAuthenticationEventsFlowExternalUsersSelfServiceSignUpConditionOperationOptions() GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpConditionOperationOptions {
	return GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpConditionOperationOptions{}
}

func (o GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpConditionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpConditionOperationOptions) ToOData() *odata.Query {
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

func (o GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpConditionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpCondition - Get conditions property value. The conditions
// representing the context of the authentication request that's used to decide whether the events policy is invoked.
// Supports $filter (eq). See support for filtering on user flows for syntax information.
func (c AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionClient) GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpCondition(ctx context.Context, id stable.IdentityAuthenticationEventsFlowId, options GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpConditionOperationOptions) (result GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpConditionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/externalUsersSelfServiceSignUpEventsFlow/conditions", id.ID()),
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

	var model stable.AuthenticationConditions
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
