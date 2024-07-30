package sharedemaildomain

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

type ListSharedEmailDomainsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SharedEmailDomain
}

type ListSharedEmailDomainsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SharedEmailDomain
}

type ListSharedEmailDomainsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSharedEmailDomainsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSharedEmailDomains ...
func (c SharedEmailDomainClient) ListSharedEmailDomains(ctx context.Context) (result ListSharedEmailDomainsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSharedEmailDomainsCustomPager{},
		Path:       "/directory/sharedEmailDomains",
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
		Values *[]beta.SharedEmailDomain `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSharedEmailDomainsComplete retrieves all the results into a single object
func (c SharedEmailDomainClient) ListSharedEmailDomainsComplete(ctx context.Context) (ListSharedEmailDomainsCompleteResult, error) {
	return c.ListSharedEmailDomainsCompleteMatchingPredicate(ctx, SharedEmailDomainOperationPredicate{})
}

// ListSharedEmailDomainsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SharedEmailDomainClient) ListSharedEmailDomainsCompleteMatchingPredicate(ctx context.Context, predicate SharedEmailDomainOperationPredicate) (result ListSharedEmailDomainsCompleteResult, err error) {
	items := make([]beta.SharedEmailDomain, 0)

	resp, err := c.ListSharedEmailDomains(ctx)
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

	result = ListSharedEmailDomainsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
