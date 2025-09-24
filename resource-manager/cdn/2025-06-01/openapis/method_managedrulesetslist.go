package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedRuleSetsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ManagedRuleSetDefinition
}

type ManagedRuleSetsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ManagedRuleSetDefinition
}

type ManagedRuleSetsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ManagedRuleSetsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ManagedRuleSetsList ...
func (c OpenapisClient) ManagedRuleSetsList(ctx context.Context, id commonids.SubscriptionId) (result ManagedRuleSetsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ManagedRuleSetsListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.Cdn/cdnWebApplicationFirewallManagedRuleSets", id.ID()),
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
		Values *[]ManagedRuleSetDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ManagedRuleSetsListComplete retrieves all the results into a single object
func (c OpenapisClient) ManagedRuleSetsListComplete(ctx context.Context, id commonids.SubscriptionId) (ManagedRuleSetsListCompleteResult, error) {
	return c.ManagedRuleSetsListCompleteMatchingPredicate(ctx, id, ManagedRuleSetDefinitionOperationPredicate{})
}

// ManagedRuleSetsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) ManagedRuleSetsListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate ManagedRuleSetDefinitionOperationPredicate) (result ManagedRuleSetsListCompleteResult, err error) {
	items := make([]ManagedRuleSetDefinition, 0)

	resp, err := c.ManagedRuleSetsList(ctx, id)
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

	result = ManagedRuleSetsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
