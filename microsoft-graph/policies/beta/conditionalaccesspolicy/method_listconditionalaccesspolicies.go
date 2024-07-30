package conditionalaccesspolicy

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

type ListConditionalAccessPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ConditionalAccessPolicy
}

type ListConditionalAccessPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ConditionalAccessPolicy
}

type ListConditionalAccessPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConditionalAccessPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConditionalAccessPolicies ...
func (c ConditionalAccessPolicyClient) ListConditionalAccessPolicies(ctx context.Context) (result ListConditionalAccessPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConditionalAccessPoliciesCustomPager{},
		Path:       "/policies/conditionalAccessPolicies",
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
		Values *[]beta.ConditionalAccessPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConditionalAccessPoliciesComplete retrieves all the results into a single object
func (c ConditionalAccessPolicyClient) ListConditionalAccessPoliciesComplete(ctx context.Context) (ListConditionalAccessPoliciesCompleteResult, error) {
	return c.ListConditionalAccessPoliciesCompleteMatchingPredicate(ctx, ConditionalAccessPolicyOperationPredicate{})
}

// ListConditionalAccessPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConditionalAccessPolicyClient) ListConditionalAccessPoliciesCompleteMatchingPredicate(ctx context.Context, predicate ConditionalAccessPolicyOperationPredicate) (result ListConditionalAccessPoliciesCompleteResult, err error) {
	items := make([]beta.ConditionalAccessPolicy, 0)

	resp, err := c.ListConditionalAccessPolicies(ctx)
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

	result = ListConditionalAccessPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
