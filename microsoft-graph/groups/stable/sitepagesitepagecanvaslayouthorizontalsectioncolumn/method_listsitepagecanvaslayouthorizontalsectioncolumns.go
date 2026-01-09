package sitepagesitepagecanvaslayouthorizontalsectioncolumn

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

type ListSitePageCanvasLayoutHorizontalSectionColumnsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.HorizontalSectionColumn
}

type ListSitePageCanvasLayoutHorizontalSectionColumnsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.HorizontalSectionColumn
}

type ListSitePageCanvasLayoutHorizontalSectionColumnsOperationOptions struct {
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

func DefaultListSitePageCanvasLayoutHorizontalSectionColumnsOperationOptions() ListSitePageCanvasLayoutHorizontalSectionColumnsOperationOptions {
	return ListSitePageCanvasLayoutHorizontalSectionColumnsOperationOptions{}
}

func (o ListSitePageCanvasLayoutHorizontalSectionColumnsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSitePageCanvasLayoutHorizontalSectionColumnsOperationOptions) ToOData() *odata.Query {
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

func (o ListSitePageCanvasLayoutHorizontalSectionColumnsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSitePageCanvasLayoutHorizontalSectionColumnsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSitePageCanvasLayoutHorizontalSectionColumnsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSitePageCanvasLayoutHorizontalSectionColumns - Get columns from groups. The set of vertical columns in this
// section.
func (c SitePageSitePageCanvasLayoutHorizontalSectionColumnClient) ListSitePageCanvasLayoutHorizontalSectionColumns(ctx context.Context, id stable.GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId, options ListSitePageCanvasLayoutHorizontalSectionColumnsOperationOptions) (result ListSitePageCanvasLayoutHorizontalSectionColumnsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSitePageCanvasLayoutHorizontalSectionColumnsCustomPager{},
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
		Values *[]stable.HorizontalSectionColumn `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSitePageCanvasLayoutHorizontalSectionColumnsComplete retrieves all the results into a single object
func (c SitePageSitePageCanvasLayoutHorizontalSectionColumnClient) ListSitePageCanvasLayoutHorizontalSectionColumnsComplete(ctx context.Context, id stable.GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId, options ListSitePageCanvasLayoutHorizontalSectionColumnsOperationOptions) (ListSitePageCanvasLayoutHorizontalSectionColumnsCompleteResult, error) {
	return c.ListSitePageCanvasLayoutHorizontalSectionColumnsCompleteMatchingPredicate(ctx, id, options, HorizontalSectionColumnOperationPredicate{})
}

// ListSitePageCanvasLayoutHorizontalSectionColumnsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SitePageSitePageCanvasLayoutHorizontalSectionColumnClient) ListSitePageCanvasLayoutHorizontalSectionColumnsCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId, options ListSitePageCanvasLayoutHorizontalSectionColumnsOperationOptions, predicate HorizontalSectionColumnOperationPredicate) (result ListSitePageCanvasLayoutHorizontalSectionColumnsCompleteResult, err error) {
	items := make([]stable.HorizontalSectionColumn, 0)

	resp, err := c.ListSitePageCanvasLayoutHorizontalSectionColumns(ctx, id, options)
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

	result = ListSitePageCanvasLayoutHorizontalSectionColumnsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
