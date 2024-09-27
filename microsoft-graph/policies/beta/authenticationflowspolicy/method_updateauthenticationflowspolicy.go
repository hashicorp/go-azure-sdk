package authenticationflowspolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAuthenticationFlowsPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAuthenticationFlowsPolicyOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateAuthenticationFlowsPolicyOperationOptions() UpdateAuthenticationFlowsPolicyOperationOptions {
	return UpdateAuthenticationFlowsPolicyOperationOptions{}
}

func (o UpdateAuthenticationFlowsPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAuthenticationFlowsPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAuthenticationFlowsPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAuthenticationFlowsPolicy - Update authenticationFlowsPolicy. Update the Boolean selfServiceSignUp property of
// an authenticationFlowsPolicy object. The properties id, type, and description cannot be modified.
func (c AuthenticationFlowsPolicyClient) UpdateAuthenticationFlowsPolicy(ctx context.Context, input beta.AuthenticationFlowsPolicy, options UpdateAuthenticationFlowsPolicyOperationOptions) (result UpdateAuthenticationFlowsPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/policies/authenticationFlowsPolicy",
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
