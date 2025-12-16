package costallocationruledefinitions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CostAllocationRulesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CostAllocationRuleDefinition
}

type CostAllocationRulesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []CostAllocationRuleDefinition
}

type CostAllocationRulesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *CostAllocationRulesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// CostAllocationRulesList ...
func (c CostAllocationRuleDefinitionsClient) CostAllocationRulesList(ctx context.Context, id BillingAccountId) (result CostAllocationRulesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &CostAllocationRulesListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.CostManagement/costAllocationRules", id.ID()),
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
		Values *[]CostAllocationRuleDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// CostAllocationRulesListComplete retrieves all the results into a single object
func (c CostAllocationRuleDefinitionsClient) CostAllocationRulesListComplete(ctx context.Context, id BillingAccountId) (CostAllocationRulesListCompleteResult, error) {
	return c.CostAllocationRulesListCompleteMatchingPredicate(ctx, id, CostAllocationRuleDefinitionOperationPredicate{})
}

// CostAllocationRulesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CostAllocationRuleDefinitionsClient) CostAllocationRulesListCompleteMatchingPredicate(ctx context.Context, id BillingAccountId, predicate CostAllocationRuleDefinitionOperationPredicate) (result CostAllocationRulesListCompleteResult, err error) {
	items := make([]CostAllocationRuleDefinition, 0)

	resp, err := c.CostAllocationRulesList(ctx, id)
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

	result = CostAllocationRulesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
