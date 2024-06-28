package privatelinkresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPrivateLinkResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateLinkResource
}

type GetPrivateLinkResourcesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateLinkResource
}

type GetPrivateLinkResourcesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GetPrivateLinkResourcesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetPrivateLinkResources ...
func (c PrivateLinkResourcesClient) GetPrivateLinkResources(ctx context.Context, id MasterSiteId) (result GetPrivateLinkResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &GetPrivateLinkResourcesCustomPager{},
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

// GetPrivateLinkResourcesComplete retrieves all the results into a single object
func (c PrivateLinkResourcesClient) GetPrivateLinkResourcesComplete(ctx context.Context, id MasterSiteId) (GetPrivateLinkResourcesCompleteResult, error) {
	return c.GetPrivateLinkResourcesCompleteMatchingPredicate(ctx, id, PrivateLinkResourceOperationPredicate{})
}

// GetPrivateLinkResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateLinkResourcesClient) GetPrivateLinkResourcesCompleteMatchingPredicate(ctx context.Context, id MasterSiteId, predicate PrivateLinkResourceOperationPredicate) (result GetPrivateLinkResourcesCompleteResult, err error) {
	items := make([]PrivateLinkResource, 0)

	resp, err := c.GetPrivateLinkResources(ctx, id)
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

	result = GetPrivateLinkResourcesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
