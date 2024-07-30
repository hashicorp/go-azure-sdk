package siteitem

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

type ListSiteItemsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.BaseItem
}

type ListSiteItemsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.BaseItem
}

type ListSiteItemsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteItemsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteItems ...
func (c SiteItemClient) ListSiteItems(ctx context.Context, id GroupIdSiteId) (result ListSiteItemsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteItemsCustomPager{},
		Path:       fmt.Sprintf("%s/items", id.ID()),
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
		Values *[]stable.BaseItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteItemsComplete retrieves all the results into a single object
func (c SiteItemClient) ListSiteItemsComplete(ctx context.Context, id GroupIdSiteId) (ListSiteItemsCompleteResult, error) {
	return c.ListSiteItemsCompleteMatchingPredicate(ctx, id, BaseItemOperationPredicate{})
}

// ListSiteItemsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteItemClient) ListSiteItemsCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteId, predicate BaseItemOperationPredicate) (result ListSiteItemsCompleteResult, err error) {
	items := make([]stable.BaseItem, 0)

	resp, err := c.ListSiteItems(ctx, id)
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

	result = ListSiteItemsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
