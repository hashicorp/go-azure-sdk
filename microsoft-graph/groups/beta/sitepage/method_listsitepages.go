package sitepage

import (
	"context"
	"encoding/json"
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

type ListSitePagesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListSitePagesOperationOptions() ListSitePagesOperationOptions {
	return ListSitePagesOperationOptions{}
}

func (o ListSitePagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSitePagesOperationOptions) ToOData() *odata.Query {
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

func (o ListSitePagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListSitePages - Get pages from groups. The collection of pages in the baseSitePages list on this site.
func (c SitePageClient) ListSitePages(ctx context.Context, id beta.GroupIdSiteId, options ListSitePagesOperationOptions) (result ListSitePagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSitePagesCustomPager{},
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.BaseSitePage, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalBaseSitePageImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.BaseSitePage (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListSitePagesComplete retrieves all the results into a single object
func (c SitePageClient) ListSitePagesComplete(ctx context.Context, id beta.GroupIdSiteId, options ListSitePagesOperationOptions) (ListSitePagesCompleteResult, error) {
	return c.ListSitePagesCompleteMatchingPredicate(ctx, id, options, BaseSitePageOperationPredicate{})
}

// ListSitePagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SitePageClient) ListSitePagesCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteId, options ListSitePagesOperationOptions, predicate BaseSitePageOperationPredicate) (result ListSitePagesCompleteResult, err error) {
	items := make([]beta.BaseSitePage, 0)

	resp, err := c.ListSitePages(ctx, id, options)
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
