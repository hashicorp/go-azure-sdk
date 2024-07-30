package authenticationstrengthpolicycombinationconfiguration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAuthenticationStrengthPolicyCombinationConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AuthenticationCombinationConfiguration
}

type ListAuthenticationStrengthPolicyCombinationConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AuthenticationCombinationConfiguration
}

type ListAuthenticationStrengthPolicyCombinationConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationStrengthPolicyCombinationConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationStrengthPolicyCombinationConfigurations ...
func (c AuthenticationStrengthPolicyCombinationConfigurationClient) ListAuthenticationStrengthPolicyCombinationConfigurations(ctx context.Context, id PolicyAuthenticationStrengthPolicyId) (result ListAuthenticationStrengthPolicyCombinationConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAuthenticationStrengthPolicyCombinationConfigurationsCustomPager{},
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
		Values *[]beta.AuthenticationCombinationConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationStrengthPolicyCombinationConfigurationsComplete retrieves all the results into a single object
func (c AuthenticationStrengthPolicyCombinationConfigurationClient) ListAuthenticationStrengthPolicyCombinationConfigurationsComplete(ctx context.Context, id PolicyAuthenticationStrengthPolicyId) (ListAuthenticationStrengthPolicyCombinationConfigurationsCompleteResult, error) {
	return c.ListAuthenticationStrengthPolicyCombinationConfigurationsCompleteMatchingPredicate(ctx, id, AuthenticationCombinationConfigurationOperationPredicate{})
}

// ListAuthenticationStrengthPolicyCombinationConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationStrengthPolicyCombinationConfigurationClient) ListAuthenticationStrengthPolicyCombinationConfigurationsCompleteMatchingPredicate(ctx context.Context, id PolicyAuthenticationStrengthPolicyId, predicate AuthenticationCombinationConfigurationOperationPredicate) (result ListAuthenticationStrengthPolicyCombinationConfigurationsCompleteResult, err error) {
	items := make([]beta.AuthenticationCombinationConfiguration, 0)

	resp, err := c.ListAuthenticationStrengthPolicyCombinationConfigurations(ctx, id)
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

	result = ListAuthenticationStrengthPolicyCombinationConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
