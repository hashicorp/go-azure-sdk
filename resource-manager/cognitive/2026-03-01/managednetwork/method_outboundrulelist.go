package managednetwork

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutboundRuleListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]OutboundRuleBasicResource
}

type OutboundRuleListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []OutboundRuleBasicResource
}

type OutboundRuleListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *OutboundRuleListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// OutboundRuleList ...
func (c ManagedNetworkClient) OutboundRuleList(ctx context.Context, id ManagedNetworkId) (result OutboundRuleListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &OutboundRuleListCustomPager{},
		Path:       fmt.Sprintf("%s/outboundRules", id.ID()),
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
		Values *[]OutboundRuleBasicResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// OutboundRuleListComplete retrieves all the results into a single object
func (c ManagedNetworkClient) OutboundRuleListComplete(ctx context.Context, id ManagedNetworkId) (OutboundRuleListCompleteResult, error) {
	return c.OutboundRuleListCompleteMatchingPredicate(ctx, id, OutboundRuleBasicResourceOperationPredicate{})
}

// OutboundRuleListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedNetworkClient) OutboundRuleListCompleteMatchingPredicate(ctx context.Context, id ManagedNetworkId, predicate OutboundRuleBasicResourceOperationPredicate) (result OutboundRuleListCompleteResult, err error) {
	items := make([]OutboundRuleBasicResource, 0)

	resp, err := c.OutboundRuleList(ctx, id)
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

	result = OutboundRuleListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
