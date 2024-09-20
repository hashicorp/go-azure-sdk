package sitelistcolumn

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

type ListSiteListColumnsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ColumnDefinition
}

type ListSiteListColumnsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ColumnDefinition
}

type ListSiteListColumnsOperationOptions struct {
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

func DefaultListSiteListColumnsOperationOptions() ListSiteListColumnsOperationOptions {
	return ListSiteListColumnsOperationOptions{}
}

func (o ListSiteListColumnsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteListColumnsOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteListColumnsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteListColumnsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteListColumnsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteListColumns - Get columns from groups. The collection of field definitions for this list.
func (c SiteListColumnClient) ListSiteListColumns(ctx context.Context, id stable.GroupIdSiteIdListId, options ListSiteListColumnsOperationOptions) (result ListSiteListColumnsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteListColumnsCustomPager{},
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

// ListSiteListColumnsComplete retrieves all the results into a single object
func (c SiteListColumnClient) ListSiteListColumnsComplete(ctx context.Context, id stable.GroupIdSiteIdListId, options ListSiteListColumnsOperationOptions) (ListSiteListColumnsCompleteResult, error) {
	return c.ListSiteListColumnsCompleteMatchingPredicate(ctx, id, options, ColumnDefinitionOperationPredicate{})
}

// ListSiteListColumnsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteListColumnClient) ListSiteListColumnsCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdSiteIdListId, options ListSiteListColumnsOperationOptions, predicate ColumnDefinitionOperationPredicate) (result ListSiteListColumnsCompleteResult, err error) {
	items := make([]stable.ColumnDefinition, 0)

	resp, err := c.ListSiteListColumns(ctx, id, options)
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

	result = ListSiteListColumnsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
