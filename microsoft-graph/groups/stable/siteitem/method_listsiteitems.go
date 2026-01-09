package siteitem

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSiteItemsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.BaseItem
}

type ListSiteItemsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.BaseItem
}

type ListSiteItemsOperationOptions struct {
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

func DefaultListSiteItemsOperationOptions() ListSiteItemsOperationOptions {
	return ListSiteItemsOperationOptions{}
}

func (o ListSiteItemsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteItemsOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteItemsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteItemsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteItemsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteItems - Get items from groups. Used to address any item contained in this site. This collection can't be
// enumerated.
func (c SiteItemClient) ListSiteItems(ctx context.Context, id stable.GroupIdSiteId, options ListSiteItemsOperationOptions) (result ListSiteItemsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteItemsCustomPager{},
		Path:          fmt.Sprintf("%s/items", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]stable.BaseItem, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalBaseItemImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.BaseItem (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListSiteItemsComplete retrieves all the results into a single object
func (c SiteItemClient) ListSiteItemsComplete(ctx context.Context, id stable.GroupIdSiteId, options ListSiteItemsOperationOptions) (ListSiteItemsCompleteResult, error) {
	return c.ListSiteItemsCompleteMatchingPredicate(ctx, id, options, BaseItemOperationPredicate{})
}

// ListSiteItemsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteItemClient) ListSiteItemsCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdSiteId, options ListSiteItemsOperationOptions, predicate BaseItemOperationPredicate) (result ListSiteItemsCompleteResult, err error) {
	items := make([]stable.BaseItem, 0)

	resp, err := c.ListSiteItems(ctx, id, options)
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

	result = ListSiteItemsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
