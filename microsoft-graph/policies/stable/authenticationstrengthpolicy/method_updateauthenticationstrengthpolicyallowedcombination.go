package authenticationstrengthpolicy

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

type UpdateAuthenticationStrengthPolicyAllowedCombinationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UpdateAllowedCombinationsResult
}

// UpdateAuthenticationStrengthPolicyAllowedCombination - Invoke action updateAllowedCombinations. Update the
// allowedCombinations property of an authenticationStrengthPolicy object. To update other properties of an
// authenticationStrengthPolicy object, use the Update authenticationStrengthPolicy method.
func (c AuthenticationStrengthPolicyClient) UpdateAuthenticationStrengthPolicyAllowedCombination(ctx context.Context, id stable.PolicyAuthenticationStrengthPolicyId, input UpdateAuthenticationStrengthPolicyAllowedCombinationRequest) (result UpdateAuthenticationStrengthPolicyAllowedCombinationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/updateAllowedCombinations", id.ID()),
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
