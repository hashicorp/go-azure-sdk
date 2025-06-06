package dnsresolverpolicies

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

type ListByVirtualNetworkOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SubResource
}

type ListByVirtualNetworkCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SubResource
}

type ListByVirtualNetworkCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByVirtualNetworkCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByVirtualNetwork ...
func (c DnsResolverPoliciesClient) ListByVirtualNetwork(ctx context.Context, id commonids.VirtualNetworkId) (result ListByVirtualNetworkOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &ListByVirtualNetworkCustomPager{},
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

// ListByVirtualNetworkComplete retrieves all the results into a single object
func (c DnsResolverPoliciesClient) ListByVirtualNetworkComplete(ctx context.Context, id commonids.VirtualNetworkId) (ListByVirtualNetworkCompleteResult, error) {
	return c.ListByVirtualNetworkCompleteMatchingPredicate(ctx, id, SubResourceOperationPredicate{})
}

// ListByVirtualNetworkCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DnsResolverPoliciesClient) ListByVirtualNetworkCompleteMatchingPredicate(ctx context.Context, id commonids.VirtualNetworkId, predicate SubResourceOperationPredicate) (result ListByVirtualNetworkCompleteResult, err error) {
	items := make([]SubResource, 0)

	resp, err := c.ListByVirtualNetwork(ctx, id)
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

	result = ListByVirtualNetworkCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
