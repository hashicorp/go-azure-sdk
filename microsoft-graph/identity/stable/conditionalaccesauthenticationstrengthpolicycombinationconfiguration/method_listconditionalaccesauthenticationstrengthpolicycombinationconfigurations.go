package conditionalaccesauthenticationstrengthpolicycombinationconfiguration

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

type ListConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AuthenticationCombinationConfiguration
}

type ListConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AuthenticationCombinationConfiguration
}

type ListConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConditionalAccesAuthenticationStrengthPolicyCombinationConfigurations ...
func (c ConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClient) ListConditionalAccesAuthenticationStrengthPolicyCombinationConfigurations(ctx context.Context, id IdentityConditionalAccesAuthenticationStrengthPolicyId) (result ListConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationsCustomPager{},
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
		Values *[]stable.AuthenticationCombinationConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationsComplete retrieves all the results into a single object
func (c ConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClient) ListConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationsComplete(ctx context.Context, id IdentityConditionalAccesAuthenticationStrengthPolicyId) (ListConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationsCompleteResult, error) {
	return c.ListConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationsCompleteMatchingPredicate(ctx, id, AuthenticationCombinationConfigurationOperationPredicate{})
}

// ListConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClient) ListConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationsCompleteMatchingPredicate(ctx context.Context, id IdentityConditionalAccesAuthenticationStrengthPolicyId, predicate AuthenticationCombinationConfigurationOperationPredicate) (result ListConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationsCompleteResult, err error) {
	items := make([]stable.AuthenticationCombinationConfiguration, 0)

	resp, err := c.ListConditionalAccesAuthenticationStrengthPolicyCombinationConfigurations(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
