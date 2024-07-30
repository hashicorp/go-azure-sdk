package sitelistcontenttypecolumnlink

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

type ListSiteListContentTypeColumnLinksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ColumnLink
}

type ListSiteListContentTypeColumnLinksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ColumnLink
}

type ListSiteListContentTypeColumnLinksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteListContentTypeColumnLinksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteListContentTypeColumnLinks ...
func (c SiteListContentTypeColumnLinkClient) ListSiteListContentTypeColumnLinks(ctx context.Context, id GroupIdSiteIdListIdContentTypeId) (result ListSiteListContentTypeColumnLinksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteListContentTypeColumnLinksCustomPager{},
		Path:       fmt.Sprintf("%s/columnLinks", id.ID()),
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
		Values *[]stable.ColumnLink `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteListContentTypeColumnLinksComplete retrieves all the results into a single object
func (c SiteListContentTypeColumnLinkClient) ListSiteListContentTypeColumnLinksComplete(ctx context.Context, id GroupIdSiteIdListIdContentTypeId) (ListSiteListContentTypeColumnLinksCompleteResult, error) {
	return c.ListSiteListContentTypeColumnLinksCompleteMatchingPredicate(ctx, id, ColumnLinkOperationPredicate{})
}

// ListSiteListContentTypeColumnLinksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteListContentTypeColumnLinkClient) ListSiteListContentTypeColumnLinksCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteIdListIdContentTypeId, predicate ColumnLinkOperationPredicate) (result ListSiteListContentTypeColumnLinksCompleteResult, err error) {
	items := make([]stable.ColumnLink, 0)

	resp, err := c.ListSiteListContentTypeColumnLinks(ctx, id)
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

	result = ListSiteListContentTypeColumnLinksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
