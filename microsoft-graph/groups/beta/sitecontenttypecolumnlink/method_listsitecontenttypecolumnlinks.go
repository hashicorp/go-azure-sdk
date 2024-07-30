package sitecontenttypecolumnlink

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

type ListSiteContentTypeColumnLinksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ColumnLink
}

type ListSiteContentTypeColumnLinksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ColumnLink
}

type ListSiteContentTypeColumnLinksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteContentTypeColumnLinksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteContentTypeColumnLinks ...
func (c SiteContentTypeColumnLinkClient) ListSiteContentTypeColumnLinks(ctx context.Context, id GroupIdSiteIdContentTypeId) (result ListSiteContentTypeColumnLinksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteContentTypeColumnLinksCustomPager{},
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
		Values *[]beta.ColumnLink `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteContentTypeColumnLinksComplete retrieves all the results into a single object
func (c SiteContentTypeColumnLinkClient) ListSiteContentTypeColumnLinksComplete(ctx context.Context, id GroupIdSiteIdContentTypeId) (ListSiteContentTypeColumnLinksCompleteResult, error) {
	return c.ListSiteContentTypeColumnLinksCompleteMatchingPredicate(ctx, id, ColumnLinkOperationPredicate{})
}

// ListSiteContentTypeColumnLinksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteContentTypeColumnLinkClient) ListSiteContentTypeColumnLinksCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteIdContentTypeId, predicate ColumnLinkOperationPredicate) (result ListSiteContentTypeColumnLinksCompleteResult, err error) {
	items := make([]beta.ColumnLink, 0)

	resp, err := c.ListSiteContentTypeColumnLinks(ctx, id)
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

	result = ListSiteContentTypeColumnLinksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
