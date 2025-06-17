package sitepagetemplatecanvaslayouthorizontalsectioncolumn

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

type ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.HorizontalSectionColumn
}

type ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.HorizontalSectionColumn
}

type ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsOperationOptions struct {
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

func DefaultListSitePageTemplateCanvasLayoutHorizontalSectionColumnsOperationOptions() ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsOperationOptions {
	return ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsOperationOptions{}
}

func (o ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsOperationOptions) ToOData() *odata.Query {
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

func (o ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSitePageTemplateCanvasLayoutHorizontalSectionColumns - Get columns from groups. The set of vertical columns in
// this section.
func (c SitePageTemplateCanvasLayoutHorizontalSectionColumnClient) ListSitePageTemplateCanvasLayoutHorizontalSectionColumns(ctx context.Context, id beta.GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId, options ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsOperationOptions) (result ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsCustomPager{},
		Path:          fmt.Sprintf("%s/columns", id.ID()),
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
		Values *[]beta.HorizontalSectionColumn `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsComplete retrieves all the results into a single object
func (c SitePageTemplateCanvasLayoutHorizontalSectionColumnClient) ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsComplete(ctx context.Context, id beta.GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId, options ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsOperationOptions) (ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsCompleteResult, error) {
	return c.ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsCompleteMatchingPredicate(ctx, id, options, HorizontalSectionColumnOperationPredicate{})
}

// ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SitePageTemplateCanvasLayoutHorizontalSectionColumnClient) ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId, options ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsOperationOptions, predicate HorizontalSectionColumnOperationPredicate) (result ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsCompleteResult, err error) {
	items := make([]beta.HorizontalSectionColumn, 0)

	resp, err := c.ListSitePageTemplateCanvasLayoutHorizontalSectionColumns(ctx, id, options)
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

	result = ListSitePageTemplateCanvasLayoutHorizontalSectionColumnsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
