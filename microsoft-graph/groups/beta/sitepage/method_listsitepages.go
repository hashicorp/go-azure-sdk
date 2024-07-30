package sitepage

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

type ListSitePagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.BaseSitePage
}

type ListSitePagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.BaseSitePage
}

type ListSitePagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSitePagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSitePages ...
func (c SitePageClient) ListSitePages(ctx context.Context, id GroupIdSiteId) (result ListSitePagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSitePagesCustomPager{},
		Path:       fmt.Sprintf("%s/pages", id.ID()),
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
		Values *[]beta.BaseSitePage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSitePagesComplete retrieves all the results into a single object
func (c SitePageClient) ListSitePagesComplete(ctx context.Context, id GroupIdSiteId) (ListSitePagesCompleteResult, error) {
	return c.ListSitePagesCompleteMatchingPredicate(ctx, id, BaseSitePageOperationPredicate{})
}

// ListSitePagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SitePageClient) ListSitePagesCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteId, predicate BaseSitePageOperationPredicate) (result ListSitePagesCompleteResult, err error) {
	items := make([]beta.BaseSitePage, 0)

	resp, err := c.ListSitePages(ctx, id)
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

	result = ListSitePagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
