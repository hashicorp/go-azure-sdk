package siteonenotesectionpage

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

type ListSiteOnenoteSectionPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.OnenotePage
}

type ListSiteOnenoteSectionPagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.OnenotePage
}

type ListSiteOnenoteSectionPagesOperationOptions struct {
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

func DefaultListSiteOnenoteSectionPagesOperationOptions() ListSiteOnenoteSectionPagesOperationOptions {
	return ListSiteOnenoteSectionPagesOperationOptions{}
}

func (o ListSiteOnenoteSectionPagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteOnenoteSectionPagesOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteOnenoteSectionPagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteOnenoteSectionPagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteOnenoteSectionPagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteOnenoteSectionPages - Get pages from groups. The collection of pages in the section. Read-only. Nullable.
func (c SiteOnenoteSectionPageClient) ListSiteOnenoteSectionPages(ctx context.Context, id stable.GroupIdSiteIdOnenoteSectionId, options ListSiteOnenoteSectionPagesOperationOptions) (result ListSiteOnenoteSectionPagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteOnenoteSectionPagesCustomPager{},
		Path:          fmt.Sprintf("%s/pages", id.ID()),
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
		Values *[]stable.OnenotePage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteOnenoteSectionPagesComplete retrieves all the results into a single object
func (c SiteOnenoteSectionPageClient) ListSiteOnenoteSectionPagesComplete(ctx context.Context, id stable.GroupIdSiteIdOnenoteSectionId, options ListSiteOnenoteSectionPagesOperationOptions) (ListSiteOnenoteSectionPagesCompleteResult, error) {
	return c.ListSiteOnenoteSectionPagesCompleteMatchingPredicate(ctx, id, options, OnenotePageOperationPredicate{})
}

// ListSiteOnenoteSectionPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteOnenoteSectionPageClient) ListSiteOnenoteSectionPagesCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdSiteIdOnenoteSectionId, options ListSiteOnenoteSectionPagesOperationOptions, predicate OnenotePageOperationPredicate) (result ListSiteOnenoteSectionPagesCompleteResult, err error) {
	items := make([]stable.OnenotePage, 0)

	resp, err := c.ListSiteOnenoteSectionPages(ctx, id, options)
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

	result = ListSiteOnenoteSectionPagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
