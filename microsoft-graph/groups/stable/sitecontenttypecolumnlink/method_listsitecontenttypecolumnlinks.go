package sitecontenttypecolumnlink

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSiteContentTypeColumnLinksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ColumnLink
}

type ListSiteContentTypeColumnLinksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ColumnLink
}

type ListSiteContentTypeColumnLinksOperationOptions struct {
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

func DefaultListSiteContentTypeColumnLinksOperationOptions() ListSiteContentTypeColumnLinksOperationOptions {
	return ListSiteContentTypeColumnLinksOperationOptions{}
}

func (o ListSiteContentTypeColumnLinksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteContentTypeColumnLinksOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteContentTypeColumnLinksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListSiteContentTypeColumnLinks - Get columnLinks from groups. The collection of columns that are required by this
// content type.
func (c SiteContentTypeColumnLinkClient) ListSiteContentTypeColumnLinks(ctx context.Context, id stable.GroupIdSiteIdContentTypeId, options ListSiteContentTypeColumnLinksOperationOptions) (result ListSiteContentTypeColumnLinksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteContentTypeColumnLinksCustomPager{},
		Path:          fmt.Sprintf("%s/columnLinks", id.ID()),
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
		Values *[]stable.ColumnLink `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteContentTypeColumnLinksComplete retrieves all the results into a single object
func (c SiteContentTypeColumnLinkClient) ListSiteContentTypeColumnLinksComplete(ctx context.Context, id stable.GroupIdSiteIdContentTypeId, options ListSiteContentTypeColumnLinksOperationOptions) (ListSiteContentTypeColumnLinksCompleteResult, error) {
	return c.ListSiteContentTypeColumnLinksCompleteMatchingPredicate(ctx, id, options, ColumnLinkOperationPredicate{})
}

// ListSiteContentTypeColumnLinksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteContentTypeColumnLinkClient) ListSiteContentTypeColumnLinksCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdSiteIdContentTypeId, options ListSiteContentTypeColumnLinksOperationOptions, predicate ColumnLinkOperationPredicate) (result ListSiteContentTypeColumnLinksCompleteResult, err error) {
	items := make([]stable.ColumnLink, 0)

	resp, err := c.ListSiteContentTypeColumnLinks(ctx, id, options)
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
