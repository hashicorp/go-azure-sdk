package permissiongrantpolicy

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

type ListPermissionGrantPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PermissionGrantPolicy
}

type ListPermissionGrantPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PermissionGrantPolicy
}

type ListPermissionGrantPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPermissionGrantPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPermissionGrantPolicies ...
func (c PermissionGrantPolicyClient) ListPermissionGrantPolicies(ctx context.Context) (result ListPermissionGrantPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPermissionGrantPoliciesCustomPager{},
		Path:       "/policies/permissionGrantPolicies",
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
		Values *[]stable.PermissionGrantPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPermissionGrantPoliciesComplete retrieves all the results into a single object
func (c PermissionGrantPolicyClient) ListPermissionGrantPoliciesComplete(ctx context.Context) (ListPermissionGrantPoliciesCompleteResult, error) {
	return c.ListPermissionGrantPoliciesCompleteMatchingPredicate(ctx, PermissionGrantPolicyOperationPredicate{})
}

// ListPermissionGrantPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PermissionGrantPolicyClient) ListPermissionGrantPoliciesCompleteMatchingPredicate(ctx context.Context, predicate PermissionGrantPolicyOperationPredicate) (result ListPermissionGrantPoliciesCompleteResult, err error) {
	items := make([]stable.PermissionGrantPolicy, 0)

	resp, err := c.ListPermissionGrantPolicies(ctx)
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

	result = ListPermissionGrantPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
