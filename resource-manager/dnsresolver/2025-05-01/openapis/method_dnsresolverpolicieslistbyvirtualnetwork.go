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

type DnsResolverPoliciesListByVirtualNetworkOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SubResource
}

type DnsResolverPoliciesListByVirtualNetworkCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SubResource
}

type DnsResolverPoliciesListByVirtualNetworkCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DnsResolverPoliciesListByVirtualNetworkCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DnsResolverPoliciesListByVirtualNetwork ...
func (c OpenapisClient) DnsResolverPoliciesListByVirtualNetwork(ctx context.Context, id commonids.VirtualNetworkId) (result DnsResolverPoliciesListByVirtualNetworkOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &DnsResolverPoliciesListByVirtualNetworkCustomPager{},
		Path:       fmt.Sprintf("%s/listDnsResolverPolicies", id.ID()),
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
		Values *[]SubResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DnsResolverPoliciesListByVirtualNetworkComplete retrieves all the results into a single object
func (c OpenapisClient) DnsResolverPoliciesListByVirtualNetworkComplete(ctx context.Context, id commonids.VirtualNetworkId) (DnsResolverPoliciesListByVirtualNetworkCompleteResult, error) {
	return c.DnsResolverPoliciesListByVirtualNetworkCompleteMatchingPredicate(ctx, id, SubResourceOperationPredicate{})
}

// DnsResolverPoliciesListByVirtualNetworkCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) DnsResolverPoliciesListByVirtualNetworkCompleteMatchingPredicate(ctx context.Context, id commonids.VirtualNetworkId, predicate SubResourceOperationPredicate) (result DnsResolverPoliciesListByVirtualNetworkCompleteResult, err error) {
	items := make([]SubResource, 0)

	resp, err := c.DnsResolverPoliciesListByVirtualNetwork(ctx, id)
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

	result = DnsResolverPoliciesListByVirtualNetworkCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
