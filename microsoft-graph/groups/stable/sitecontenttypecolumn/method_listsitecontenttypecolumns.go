package sitecontenttypecolumn

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

type ListSiteContentTypeColumnsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ColumnDefinition
}

type ListSiteContentTypeColumnsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ColumnDefinition
}

type ListSiteContentTypeColumnsOperationOptions struct {
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

func DefaultListSiteContentTypeColumnsOperationOptions() ListSiteContentTypeColumnsOperationOptions {
	return ListSiteContentTypeColumnsOperationOptions{}
}

func (o ListSiteContentTypeColumnsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteContentTypeColumnsOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteContentTypeColumnsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteContentTypeColumnsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteContentTypeColumnsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteContentTypeColumns - Get columns from groups. The collection of column definitions for this content type.
func (c SiteContentTypeColumnClient) ListSiteContentTypeColumns(ctx context.Context, id stable.GroupIdSiteIdContentTypeId, options ListSiteContentTypeColumnsOperationOptions) (result ListSiteContentTypeColumnsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteContentTypeColumnsCustomPager{},
		Path:          fmt.Sprintf("%s/columns", id.ID()),
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
		Values *[]stable.ColumnDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteContentTypeColumnsComplete retrieves all the results into a single object
func (c SiteContentTypeColumnClient) ListSiteContentTypeColumnsComplete(ctx context.Context, id stable.GroupIdSiteIdContentTypeId, options ListSiteContentTypeColumnsOperationOptions) (ListSiteContentTypeColumnsCompleteResult, error) {
	return c.ListSiteContentTypeColumnsCompleteMatchingPredicate(ctx, id, options, ColumnDefinitionOperationPredicate{})
}

// ListSiteContentTypeColumnsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteContentTypeColumnClient) ListSiteContentTypeColumnsCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdSiteIdContentTypeId, options ListSiteContentTypeColumnsOperationOptions, predicate ColumnDefinitionOperationPredicate) (result ListSiteContentTypeColumnsCompleteResult, err error) {
	items := make([]stable.ColumnDefinition, 0)

	resp, err := c.ListSiteContentTypeColumns(ctx, id, options)
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

	result = ListSiteContentTypeColumnsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
