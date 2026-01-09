package sitecontenttypecolumnposition

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

type ListSiteContentTypeColumnPositionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ColumnDefinition
}

type ListSiteContentTypeColumnPositionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ColumnDefinition
}

type ListSiteContentTypeColumnPositionsOperationOptions struct {
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

func DefaultListSiteContentTypeColumnPositionsOperationOptions() ListSiteContentTypeColumnPositionsOperationOptions {
	return ListSiteContentTypeColumnPositionsOperationOptions{}
}

func (o ListSiteContentTypeColumnPositionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteContentTypeColumnPositionsOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteContentTypeColumnPositionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteContentTypeColumnPositionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteContentTypeColumnPositionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteContentTypeColumnPositions - Get columnPositions from groups. Column order information in a content type.
func (c SiteContentTypeColumnPositionClient) ListSiteContentTypeColumnPositions(ctx context.Context, id stable.GroupIdSiteIdContentTypeId, options ListSiteContentTypeColumnPositionsOperationOptions) (result ListSiteContentTypeColumnPositionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteContentTypeColumnPositionsCustomPager{},
		Path:          fmt.Sprintf("%s/columnPositions", id.ID()),
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
		Values *[]stable.ColumnDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteContentTypeColumnPositionsComplete retrieves all the results into a single object
func (c SiteContentTypeColumnPositionClient) ListSiteContentTypeColumnPositionsComplete(ctx context.Context, id stable.GroupIdSiteIdContentTypeId, options ListSiteContentTypeColumnPositionsOperationOptions) (ListSiteContentTypeColumnPositionsCompleteResult, error) {
	return c.ListSiteContentTypeColumnPositionsCompleteMatchingPredicate(ctx, id, options, ColumnDefinitionOperationPredicate{})
}

// ListSiteContentTypeColumnPositionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteContentTypeColumnPositionClient) ListSiteContentTypeColumnPositionsCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdSiteIdContentTypeId, options ListSiteContentTypeColumnPositionsOperationOptions, predicate ColumnDefinitionOperationPredicate) (result ListSiteContentTypeColumnPositionsCompleteResult, err error) {
	items := make([]stable.ColumnDefinition, 0)

	resp, err := c.ListSiteContentTypeColumnPositions(ctx, id, options)
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

	result = ListSiteContentTypeColumnPositionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
