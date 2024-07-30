package authorizationpolicy

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

type ListAuthorizationPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AuthorizationPolicy
}

type ListAuthorizationPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AuthorizationPolicy
}

type ListAuthorizationPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthorizationPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthorizationPolicies ...
func (c AuthorizationPolicyClient) ListAuthorizationPolicies(ctx context.Context) (result ListAuthorizationPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAuthorizationPoliciesCustomPager{},
		Path:       "/policies/authorizationPolicy",
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
		Values *[]beta.AuthorizationPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthorizationPoliciesComplete retrieves all the results into a single object
func (c AuthorizationPolicyClient) ListAuthorizationPoliciesComplete(ctx context.Context) (ListAuthorizationPoliciesCompleteResult, error) {
	return c.ListAuthorizationPoliciesCompleteMatchingPredicate(ctx, AuthorizationPolicyOperationPredicate{})
}

// ListAuthorizationPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthorizationPolicyClient) ListAuthorizationPoliciesCompleteMatchingPredicate(ctx context.Context, predicate AuthorizationPolicyOperationPredicate) (result ListAuthorizationPoliciesCompleteResult, err error) {
	items := make([]beta.AuthorizationPolicy, 0)

	resp, err := c.ListAuthorizationPolicies(ctx)
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

	result = ListAuthorizationPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
