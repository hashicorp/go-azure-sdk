package namespacesprivatelinkresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkResourcesGetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateLinkResource
}

type PrivateLinkResourcesGetCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateLinkResource
}

type PrivateLinkResourcesGetCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *PrivateLinkResourcesGetCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// PrivateLinkResourcesGet ...
func (c NamespacesPrivateLinkResourcesClient) PrivateLinkResourcesGet(ctx context.Context, id NamespaceId) (result PrivateLinkResourcesGetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &PrivateLinkResourcesGetCustomPager{},
		Path:       fmt.Sprintf("%s/privateLinkResources", id.ID()),
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
		Values *[]PrivateLinkResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PrivateLinkResourcesGetComplete retrieves all the results into a single object
func (c NamespacesPrivateLinkResourcesClient) PrivateLinkResourcesGetComplete(ctx context.Context, id NamespaceId) (PrivateLinkResourcesGetCompleteResult, error) {
	return c.PrivateLinkResourcesGetCompleteMatchingPredicate(ctx, id, PrivateLinkResourceOperationPredicate{})
}

// PrivateLinkResourcesGetCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NamespacesPrivateLinkResourcesClient) PrivateLinkResourcesGetCompleteMatchingPredicate(ctx context.Context, id NamespaceId, predicate PrivateLinkResourceOperationPredicate) (result PrivateLinkResourcesGetCompleteResult, err error) {
	items := make([]PrivateLinkResource, 0)

	resp, err := c.PrivateLinkResourcesGet(ctx, id)
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

	result = PrivateLinkResourcesGetCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
