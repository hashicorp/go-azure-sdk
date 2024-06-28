package jitnetworkaccesspolicies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByRegionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]JitNetworkAccessPolicy
}

type ListByRegionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []JitNetworkAccessPolicy
}

type ListByRegionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByRegionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByRegion ...
func (c JitNetworkAccessPoliciesClient) ListByRegion(ctx context.Context, id LocationId) (result ListByRegionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByRegionCustomPager{},
		Path:       fmt.Sprintf("%s/jitNetworkAccessPolicies", id.ID()),
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
		Values *[]JitNetworkAccessPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByRegionComplete retrieves all the results into a single object
func (c JitNetworkAccessPoliciesClient) ListByRegionComplete(ctx context.Context, id LocationId) (ListByRegionCompleteResult, error) {
	return c.ListByRegionCompleteMatchingPredicate(ctx, id, JitNetworkAccessPolicyOperationPredicate{})
}

// ListByRegionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JitNetworkAccessPoliciesClient) ListByRegionCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate JitNetworkAccessPolicyOperationPredicate) (result ListByRegionCompleteResult, err error) {
	items := make([]JitNetworkAccessPolicy, 0)

	resp, err := c.ListByRegion(ctx, id)
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

	result = ListByRegionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
