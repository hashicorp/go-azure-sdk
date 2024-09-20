package siteonenotepage

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

type ListSiteOnenotePagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OnenotePage
}

type ListSiteOnenotePagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OnenotePage
}

type ListSiteOnenotePagesOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListSiteOnenotePagesOperationOptions() ListSiteOnenotePagesOperationOptions {
	return ListSiteOnenotePagesOperationOptions{}
}

func (o ListSiteOnenotePagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteOnenotePagesOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteOnenotePagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteOnenotePagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteOnenotePagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteOnenotePages - Get pages from groups. The pages in all OneNote notebooks that are owned by the user or group.
// Read-only. Nullable.
func (c SiteOnenotePageClient) ListSiteOnenotePages(ctx context.Context, id beta.GroupIdSiteId, options ListSiteOnenotePagesOperationOptions) (result ListSiteOnenotePagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteOnenotePagesCustomPager{},
		Path:          fmt.Sprintf("%s/onenote/pages", id.ID()),
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
		Values *[]beta.OnenotePage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteOnenotePagesComplete retrieves all the results into a single object
func (c SiteOnenotePageClient) ListSiteOnenotePagesComplete(ctx context.Context, id beta.GroupIdSiteId, options ListSiteOnenotePagesOperationOptions) (ListSiteOnenotePagesCompleteResult, error) {
	return c.ListSiteOnenotePagesCompleteMatchingPredicate(ctx, id, options, OnenotePageOperationPredicate{})
}

// ListSiteOnenotePagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteOnenotePageClient) ListSiteOnenotePagesCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteId, options ListSiteOnenotePagesOperationOptions, predicate OnenotePageOperationPredicate) (result ListSiteOnenotePagesCompleteResult, err error) {
	items := make([]beta.OnenotePage, 0)

	resp, err := c.ListSiteOnenotePages(ctx, id, options)
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

	result = ListSiteOnenotePagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
