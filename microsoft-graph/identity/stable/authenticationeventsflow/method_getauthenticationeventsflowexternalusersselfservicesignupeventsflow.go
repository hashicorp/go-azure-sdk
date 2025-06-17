package authenticationeventsflow

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

type GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ExternalUsersSelfServiceSignUpEventsFlow
}

type GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOperationOptions() GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOperationOptions {
	return GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOperationOptions{}
}

func (o GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOperationOptions) ToOData() *odata.Query {
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

func (o GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlow - Get the item of type
// microsoft.graph.authenticationEventsFlow as microsoft.graph.externalUsersSelfServiceSignUpEventsFlow
func (c AuthenticationEventsFlowClient) GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlow(ctx context.Context, id stable.IdentityAuthenticationEventsFlowId, options GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOperationOptions) (result GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/externalUsersSelfServiceSignUpEventsFlow", id.ID()),
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

	var model stable.ExternalUsersSelfServiceSignUpEventsFlow
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
