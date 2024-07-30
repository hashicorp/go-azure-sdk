package siterecyclebinitem

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

type ListSiteRecycleBinItemsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.RecycleBinItem
}

type ListSiteRecycleBinItemsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.RecycleBinItem
}

type ListSiteRecycleBinItemsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteRecycleBinItemsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteRecycleBinItems ...
func (c SiteRecycleBinItemClient) ListSiteRecycleBinItems(ctx context.Context, id GroupIdSiteId) (result ListSiteRecycleBinItemsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteRecycleBinItemsCustomPager{},
		Path:       fmt.Sprintf("%s/recycleBin/items", id.ID()),
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
		Values *[]beta.RecycleBinItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteRecycleBinItemsComplete retrieves all the results into a single object
func (c SiteRecycleBinItemClient) ListSiteRecycleBinItemsComplete(ctx context.Context, id GroupIdSiteId) (ListSiteRecycleBinItemsCompleteResult, error) {
	return c.ListSiteRecycleBinItemsCompleteMatchingPredicate(ctx, id, RecycleBinItemOperationPredicate{})
}

// ListSiteRecycleBinItemsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteRecycleBinItemClient) ListSiteRecycleBinItemsCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteId, predicate RecycleBinItemOperationPredicate) (result ListSiteRecycleBinItemsCompleteResult, err error) {
	items := make([]beta.RecycleBinItem, 0)

	resp, err := c.ListSiteRecycleBinItems(ctx, id)
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

	result = ListSiteRecycleBinItemsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
