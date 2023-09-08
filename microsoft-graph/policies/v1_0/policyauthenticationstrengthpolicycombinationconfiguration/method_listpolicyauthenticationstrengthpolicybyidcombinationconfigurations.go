package policyauthenticationstrengthpolicycombinationconfiguration

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

type ListPolicyAuthenticationStrengthPolicyByIdCombinationConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AuthenticationCombinationConfigurationCollectionResponse
}

type ListPolicyAuthenticationStrengthPolicyByIdCombinationConfigurationsCompleteResult struct {
	Items []models.AuthenticationCombinationConfigurationCollectionResponse
}

// ListPolicyAuthenticationStrengthPolicyByIdCombinationConfigurations ...
func (c PolicyAuthenticationStrengthPolicyCombinationConfigurationClient) ListPolicyAuthenticationStrengthPolicyByIdCombinationConfigurations(ctx context.Context, id PolicyAuthenticationStrengthPolicyId) (result ListPolicyAuthenticationStrengthPolicyByIdCombinationConfigurationsOperationResponse, err error) {
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

// ListPolicyAuthenticationStrengthPolicyByIdCombinationConfigurationsComplete retrieves all the results into a single object
func (c PolicyAuthenticationStrengthPolicyCombinationConfigurationClient) ListPolicyAuthenticationStrengthPolicyByIdCombinationConfigurationsComplete(ctx context.Context, id PolicyAuthenticationStrengthPolicyId) (ListPolicyAuthenticationStrengthPolicyByIdCombinationConfigurationsCompleteResult, error) {
	return c.ListPolicyAuthenticationStrengthPolicyByIdCombinationConfigurationsCompleteMatchingPredicate(ctx, id, models.AuthenticationCombinationConfigurationCollectionResponseOperationPredicate{})
}

// ListPolicyAuthenticationStrengthPolicyByIdCombinationConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyAuthenticationStrengthPolicyCombinationConfigurationClient) ListPolicyAuthenticationStrengthPolicyByIdCombinationConfigurationsCompleteMatchingPredicate(ctx context.Context, id PolicyAuthenticationStrengthPolicyId, predicate models.AuthenticationCombinationConfigurationCollectionResponseOperationPredicate) (result ListPolicyAuthenticationStrengthPolicyByIdCombinationConfigurationsCompleteResult, err error) {
	items := make([]models.AuthenticationCombinationConfigurationCollectionResponse, 0)

	resp, err := c.ListPolicyAuthenticationStrengthPolicyByIdCombinationConfigurations(ctx, id)
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

	result = ListPolicyAuthenticationStrengthPolicyByIdCombinationConfigurationsCompleteResult{
		Items: items,
	}
	return
}
