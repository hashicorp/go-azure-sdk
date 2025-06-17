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

type ListSitePageSitePagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SitePage
}

type ListSitePageSitePagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SitePage
}

type ListSitePageSitePagesOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListSitePageSitePagesOperationOptions() ListSitePageSitePagesOperationOptions {
	return ListSitePageSitePagesOperationOptions{}
}

func (o ListSitePageSitePagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSitePageSitePagesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListSitePageSitePagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSitePageSitePagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSitePageSitePagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSitePageSitePages - Get the items of type microsoft.graph.sitePage in the microsoft.graph.baseSitePage collection
func (c SitePageClient) ListSitePageSitePages(ctx context.Context, id beta.GroupIdSiteId, options ListSitePageSitePagesOperationOptions) (result ListSitePageSitePagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSitePageSitePagesCustomPager{},
		Path:          fmt.Sprintf("%s/pages/sitePage", id.ID()),
		RetryFunc:     options.RetryFunc,
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
		Values *[]beta.SitePage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSitePageSitePagesComplete retrieves all the results into a single object
func (c SitePageClient) ListSitePageSitePagesComplete(ctx context.Context, id beta.GroupIdSiteId, options ListSitePageSitePagesOperationOptions) (ListSitePageSitePagesCompleteResult, error) {
	return c.ListSitePageSitePagesCompleteMatchingPredicate(ctx, id, options, SitePageOperationPredicate{})
}

// ListSitePageSitePagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SitePageClient) ListSitePageSitePagesCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteId, options ListSitePageSitePagesOperationOptions, predicate SitePageOperationPredicate) (result ListSitePageSitePagesCompleteResult, err error) {
	items := make([]beta.SitePage, 0)

	resp, err := c.ListSitePageSitePages(ctx, id, options)
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

	result = ListSitePageSitePagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
