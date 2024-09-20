package sitesite

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

type ListSiteSitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Site
}

type ListSiteSitesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Site
}

type ListSiteSitesOperationOptions struct {
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

func DefaultListSiteSitesOperationOptions() ListSiteSitesOperationOptions {
	return ListSiteSitesOperationOptions{}
}

func (o ListSiteSitesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteSitesOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteSitesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteSitesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteSitesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteSites - Get sites from groups. The collection of the sub-sites under this site.
func (c SiteSiteClient) ListSiteSites(ctx context.Context, id stable.GroupIdSiteId, options ListSiteSitesOperationOptions) (result ListSiteSitesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteSitesCustomPager{},
		Path:          fmt.Sprintf("%s/sites", id.ID()),
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
		Values *[]stable.Site `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteSitesComplete retrieves all the results into a single object
func (c SiteSiteClient) ListSiteSitesComplete(ctx context.Context, id stable.GroupIdSiteId, options ListSiteSitesOperationOptions) (ListSiteSitesCompleteResult, error) {
	return c.ListSiteSitesCompleteMatchingPredicate(ctx, id, options, SiteOperationPredicate{})
}

// ListSiteSitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteSiteClient) ListSiteSitesCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdSiteId, options ListSiteSitesOperationOptions, predicate SiteOperationPredicate) (result ListSiteSitesCompleteResult, err error) {
	items := make([]stable.Site, 0)

	resp, err := c.ListSiteSites(ctx, id, options)
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

	result = ListSiteSitesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
