package authenticationeventsflow

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAuthenticationEventsFlowOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAuthenticationEventsFlowOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateAuthenticationEventsFlowOperationOptions() UpdateAuthenticationEventsFlowOperationOptions {
	return UpdateAuthenticationEventsFlowOperationOptions{}
}

func (o UpdateAuthenticationEventsFlowOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAuthenticationEventsFlowOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAuthenticationEventsFlowOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAuthenticationEventsFlow - Update authenticationEventsFlow. Update the properties of an
// authenticationEventsFlow object by ID. You must specify the @odata.type property and the value of the
// authenticationEventsFlow object type to update. The following derived subtypes are supported: -
// externalUsersSelfServiceSignupEventsFlow
func (c AuthenticationEventsFlowClient) UpdateAuthenticationEventsFlow(ctx context.Context, id stable.IdentityAuthenticationEventsFlowId, input stable.AuthenticationEventsFlow, options UpdateAuthenticationEventsFlowOperationOptions) (result UpdateAuthenticationEventsFlowOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
