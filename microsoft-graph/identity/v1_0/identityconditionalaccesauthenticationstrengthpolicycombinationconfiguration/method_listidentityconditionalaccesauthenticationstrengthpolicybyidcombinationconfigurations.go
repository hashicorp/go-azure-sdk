package identityconditionalaccesauthenticationstrengthpolicycombinationconfiguration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListIdentityConditionalAccesAuthenticationStrengthPolicyByIdCombinationConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AuthenticationCombinationConfigurationCollectionResponse
}

type ListIdentityConditionalAccesAuthenticationStrengthPolicyByIdCombinationConfigurationsCompleteResult struct {
	Items []models.AuthenticationCombinationConfigurationCollectionResponse
}

// ListIdentityConditionalAccesAuthenticationStrengthPolicyByIdCombinationConfigurations ...
func (c IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClient) ListIdentityConditionalAccesAuthenticationStrengthPolicyByIdCombinationConfigurations(ctx context.Context, id IdentityConditionalAccesAuthenticationStrengthPolicyId) (result ListIdentityConditionalAccesAuthenticationStrengthPolicyByIdCombinationConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/combinationConfigurations", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]models.AuthenticationCombinationConfigurationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIdentityConditionalAccesAuthenticationStrengthPolicyByIdCombinationConfigurationsComplete retrieves all the results into a single object
func (c IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClient) ListIdentityConditionalAccesAuthenticationStrengthPolicyByIdCombinationConfigurationsComplete(ctx context.Context, id IdentityConditionalAccesAuthenticationStrengthPolicyId) (ListIdentityConditionalAccesAuthenticationStrengthPolicyByIdCombinationConfigurationsCompleteResult, error) {
	return c.ListIdentityConditionalAccesAuthenticationStrengthPolicyByIdCombinationConfigurationsCompleteMatchingPredicate(ctx, id, models.AuthenticationCombinationConfigurationCollectionResponseOperationPredicate{})
}

// ListIdentityConditionalAccesAuthenticationStrengthPolicyByIdCombinationConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClient) ListIdentityConditionalAccesAuthenticationStrengthPolicyByIdCombinationConfigurationsCompleteMatchingPredicate(ctx context.Context, id IdentityConditionalAccesAuthenticationStrengthPolicyId, predicate models.AuthenticationCombinationConfigurationCollectionResponseOperationPredicate) (result ListIdentityConditionalAccesAuthenticationStrengthPolicyByIdCombinationConfigurationsCompleteResult, err error) {
	items := make([]models.AuthenticationCombinationConfigurationCollectionResponse, 0)

	resp, err := c.ListIdentityConditionalAccesAuthenticationStrengthPolicyByIdCombinationConfigurations(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListIdentityConditionalAccesAuthenticationStrengthPolicyByIdCombinationConfigurationsCompleteResult{
		Items: items,
	}
	return
}
