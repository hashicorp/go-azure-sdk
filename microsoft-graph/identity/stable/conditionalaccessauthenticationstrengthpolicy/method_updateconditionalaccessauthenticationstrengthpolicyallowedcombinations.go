package conditionalaccessauthenticationstrengthpolicy

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

type UpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombinationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UpdateAllowedCombinationsResult
}

type UpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombinationsOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombinationsOperationOptions() UpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombinationsOperationOptions {
	return UpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombinationsOperationOptions{}
}

func (o UpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombinationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombinationsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombinationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombinations - Invoke action updateAllowedCombinations.
// Update the allowedCombinations property of an authenticationStrengthPolicy object. To update other properties of an
// authenticationStrengthPolicy object, use the Update authenticationStrengthPolicy method.
func (c ConditionalAccessAuthenticationStrengthPolicyClient) UpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombinations(ctx context.Context, id stable.IdentityConditionalAccessAuthenticationStrengthPolicyId, input UpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombinationsRequest, options UpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombinationsOperationOptions) (result UpdateConditionalAccessAuthenticationStrengthPolicyAllowedCombinationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/updateAllowedCombinations", id.ID()),
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

	var model stable.UpdateAllowedCombinationsResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
