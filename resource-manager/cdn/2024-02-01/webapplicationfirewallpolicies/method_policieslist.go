package webapplicationfirewallpolicies

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

type PoliciesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CdnWebApplicationFirewallPolicy
}

type PoliciesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []CdnWebApplicationFirewallPolicy
}

type PoliciesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *PoliciesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// PoliciesList ...
func (c WebApplicationFirewallPoliciesClient) PoliciesList(ctx context.Context, id commonids.ResourceGroupId) (result PoliciesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &PoliciesListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.CDN/cdnWebApplicationFirewallPolicies", id.ID()),
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
		Values *[]CdnWebApplicationFirewallPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PoliciesListComplete retrieves all the results into a single object
func (c WebApplicationFirewallPoliciesClient) PoliciesListComplete(ctx context.Context, id commonids.ResourceGroupId) (PoliciesListCompleteResult, error) {
	return c.PoliciesListCompleteMatchingPredicate(ctx, id, CdnWebApplicationFirewallPolicyOperationPredicate{})
}

// PoliciesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WebApplicationFirewallPoliciesClient) PoliciesListCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate CdnWebApplicationFirewallPolicyOperationPredicate) (result PoliciesListCompleteResult, err error) {
	items := make([]CdnWebApplicationFirewallPolicy, 0)

	resp, err := c.PoliciesList(ctx, id)
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

	result = PoliciesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
