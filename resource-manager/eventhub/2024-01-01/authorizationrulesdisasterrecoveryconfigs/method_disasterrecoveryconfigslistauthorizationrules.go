package authorizationrulesdisasterrecoveryconfigs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DisasterRecoveryConfigsListAuthorizationRulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AuthorizationRule
}

type DisasterRecoveryConfigsListAuthorizationRulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AuthorizationRule
}

type DisasterRecoveryConfigsListAuthorizationRulesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DisasterRecoveryConfigsListAuthorizationRulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DisasterRecoveryConfigsListAuthorizationRules ...
func (c AuthorizationRulesDisasterRecoveryConfigsClient) DisasterRecoveryConfigsListAuthorizationRules(ctx context.Context, id DisasterRecoveryConfigId) (result DisasterRecoveryConfigsListAuthorizationRulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DisasterRecoveryConfigsListAuthorizationRulesCustomPager{},
		Path:       fmt.Sprintf("%s/authorizationRules", id.ID()),
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
		Values *[]AuthorizationRule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DisasterRecoveryConfigsListAuthorizationRulesComplete retrieves all the results into a single object
func (c AuthorizationRulesDisasterRecoveryConfigsClient) DisasterRecoveryConfigsListAuthorizationRulesComplete(ctx context.Context, id DisasterRecoveryConfigId) (DisasterRecoveryConfigsListAuthorizationRulesCompleteResult, error) {
	return c.DisasterRecoveryConfigsListAuthorizationRulesCompleteMatchingPredicate(ctx, id, AuthorizationRuleOperationPredicate{})
}

// DisasterRecoveryConfigsListAuthorizationRulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthorizationRulesDisasterRecoveryConfigsClient) DisasterRecoveryConfigsListAuthorizationRulesCompleteMatchingPredicate(ctx context.Context, id DisasterRecoveryConfigId, predicate AuthorizationRuleOperationPredicate) (result DisasterRecoveryConfigsListAuthorizationRulesCompleteResult, err error) {
	items := make([]AuthorizationRule, 0)

	resp, err := c.DisasterRecoveryConfigsListAuthorizationRules(ctx, id)
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

	result = DisasterRecoveryConfigsListAuthorizationRulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
