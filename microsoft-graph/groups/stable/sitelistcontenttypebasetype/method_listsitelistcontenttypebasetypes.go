package sitelistcontenttypebasetype

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

type ListSiteListContentTypeBaseTypesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ContentType
}

type ListSiteListContentTypeBaseTypesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ContentType
}

type ListSiteListContentTypeBaseTypesOperationOptions struct {
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

func DefaultListSiteListContentTypeBaseTypesOperationOptions() ListSiteListContentTypeBaseTypesOperationOptions {
	return ListSiteListContentTypeBaseTypesOperationOptions{}
}

func (o ListSiteListContentTypeBaseTypesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteListContentTypeBaseTypesOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteListContentTypeBaseTypesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteListContentTypeBaseTypesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteListContentTypeBaseTypesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteListContentTypeBaseTypes - Get baseTypes from groups. The collection of content types that are ancestors of
// this content type.
func (c SiteListContentTypeBaseTypeClient) ListSiteListContentTypeBaseTypes(ctx context.Context, id stable.GroupIdSiteIdListIdContentTypeId, options ListSiteListContentTypeBaseTypesOperationOptions) (result ListSiteListContentTypeBaseTypesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteListContentTypeBaseTypesCustomPager{},
		Path:          fmt.Sprintf("%s/baseTypes", id.ID()),
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
		Values *[]stable.ContentType `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteListContentTypeBaseTypesComplete retrieves all the results into a single object
func (c SiteListContentTypeBaseTypeClient) ListSiteListContentTypeBaseTypesComplete(ctx context.Context, id stable.GroupIdSiteIdListIdContentTypeId, options ListSiteListContentTypeBaseTypesOperationOptions) (ListSiteListContentTypeBaseTypesCompleteResult, error) {
	return c.ListSiteListContentTypeBaseTypesCompleteMatchingPredicate(ctx, id, options, ContentTypeOperationPredicate{})
}

// ListSiteListContentTypeBaseTypesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteListContentTypeBaseTypeClient) ListSiteListContentTypeBaseTypesCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdSiteIdListIdContentTypeId, options ListSiteListContentTypeBaseTypesOperationOptions, predicate ContentTypeOperationPredicate) (result ListSiteListContentTypeBaseTypesCompleteResult, err error) {
	items := make([]stable.ContentType, 0)

	resp, err := c.ListSiteListContentTypeBaseTypes(ctx, id, options)
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

	result = ListSiteListContentTypeBaseTypesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
