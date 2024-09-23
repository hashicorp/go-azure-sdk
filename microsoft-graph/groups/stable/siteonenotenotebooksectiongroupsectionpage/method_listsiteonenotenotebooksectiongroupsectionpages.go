package siteonenotenotebooksectiongroupsectionpage

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

type ListSiteOnenoteNotebookSectionGroupSectionPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.OnenotePage
}

type ListSiteOnenoteNotebookSectionGroupSectionPagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.OnenotePage
}

type ListSiteOnenoteNotebookSectionGroupSectionPagesOperationOptions struct {
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

func DefaultListSiteOnenoteNotebookSectionGroupSectionPagesOperationOptions() ListSiteOnenoteNotebookSectionGroupSectionPagesOperationOptions {
	return ListSiteOnenoteNotebookSectionGroupSectionPagesOperationOptions{}
}

func (o ListSiteOnenoteNotebookSectionGroupSectionPagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteOnenoteNotebookSectionGroupSectionPagesOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteOnenoteNotebookSectionGroupSectionPagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteOnenoteNotebookSectionGroupSectionPagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteOnenoteNotebookSectionGroupSectionPagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteOnenoteNotebookSectionGroupSectionPages - Get pages from groups. The collection of pages in the section.
// Read-only. Nullable.
func (c SiteOnenoteNotebookSectionGroupSectionPageClient) ListSiteOnenoteNotebookSectionGroupSectionPages(ctx context.Context, id stable.GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId, options ListSiteOnenoteNotebookSectionGroupSectionPagesOperationOptions) (result ListSiteOnenoteNotebookSectionGroupSectionPagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteOnenoteNotebookSectionGroupSectionPagesCustomPager{},
		Path:          fmt.Sprintf("%s/pages", id.ID()),
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
		Values *[]stable.OnenotePage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteOnenoteNotebookSectionGroupSectionPagesComplete retrieves all the results into a single object
func (c SiteOnenoteNotebookSectionGroupSectionPageClient) ListSiteOnenoteNotebookSectionGroupSectionPagesComplete(ctx context.Context, id stable.GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId, options ListSiteOnenoteNotebookSectionGroupSectionPagesOperationOptions) (ListSiteOnenoteNotebookSectionGroupSectionPagesCompleteResult, error) {
	return c.ListSiteOnenoteNotebookSectionGroupSectionPagesCompleteMatchingPredicate(ctx, id, options, OnenotePageOperationPredicate{})
}

// ListSiteOnenoteNotebookSectionGroupSectionPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteOnenoteNotebookSectionGroupSectionPageClient) ListSiteOnenoteNotebookSectionGroupSectionPagesCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId, options ListSiteOnenoteNotebookSectionGroupSectionPagesOperationOptions, predicate OnenotePageOperationPredicate) (result ListSiteOnenoteNotebookSectionGroupSectionPagesCompleteResult, err error) {
	items := make([]stable.OnenotePage, 0)

	resp, err := c.ListSiteOnenoteNotebookSectionGroupSectionPages(ctx, id, options)
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

	result = ListSiteOnenoteNotebookSectionGroupSectionPagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
