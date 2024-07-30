package conditionalaccespolicy

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

type ListConditionalAccesPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ConditionalAccessPolicy
}

type ListConditionalAccesPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ConditionalAccessPolicy
}

type ListConditionalAccesPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConditionalAccesPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConditionalAccesPolicies ...
func (c ConditionalAccesPolicyClient) ListConditionalAccesPolicies(ctx context.Context) (result ListConditionalAccesPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConditionalAccesPoliciesCustomPager{},
		Path:       "/identity/conditionalAccess/policies",
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
		Values *[]stable.ConditionalAccessPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConditionalAccesPoliciesComplete retrieves all the results into a single object
func (c ConditionalAccesPolicyClient) ListConditionalAccesPoliciesComplete(ctx context.Context) (ListConditionalAccesPoliciesCompleteResult, error) {
	return c.ListConditionalAccesPoliciesCompleteMatchingPredicate(ctx, ConditionalAccessPolicyOperationPredicate{})
}

// ListConditionalAccesPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConditionalAccesPolicyClient) ListConditionalAccesPoliciesCompleteMatchingPredicate(ctx context.Context, predicate ConditionalAccessPolicyOperationPredicate) (result ListConditionalAccesPoliciesCompleteResult, err error) {
	items := make([]stable.ConditionalAccessPolicy, 0)

	resp, err := c.ListConditionalAccesPolicies(ctx)
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

	result = ListConditionalAccesPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
