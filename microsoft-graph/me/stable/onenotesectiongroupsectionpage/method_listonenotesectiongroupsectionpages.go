package onenotesectiongroupsectionpage

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

type ListOnenoteSectionGroupSectionPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.OnenotePage
}

type ListOnenoteSectionGroupSectionPagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.OnenotePage
}

type ListOnenoteSectionGroupSectionPagesOperationOptions struct {
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

func DefaultListOnenoteSectionGroupSectionPagesOperationOptions() ListOnenoteSectionGroupSectionPagesOperationOptions {
	return ListOnenoteSectionGroupSectionPagesOperationOptions{}
}

func (o ListOnenoteSectionGroupSectionPagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOnenoteSectionGroupSectionPagesOperationOptions) ToOData() *odata.Query {
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

func (o ListOnenoteSectionGroupSectionPagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOnenoteSectionGroupSectionPagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnenoteSectionGroupSectionPagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnenoteSectionGroupSectionPages - Get pages from me. The collection of pages in the section. Read-only. Nullable.
func (c OnenoteSectionGroupSectionPageClient) ListOnenoteSectionGroupSectionPages(ctx context.Context, id stable.MeOnenoteSectionGroupIdSectionId, options ListOnenoteSectionGroupSectionPagesOperationOptions) (result ListOnenoteSectionGroupSectionPagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOnenoteSectionGroupSectionPagesCustomPager{},
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

// ListOnenoteSectionGroupSectionPagesComplete retrieves all the results into a single object
func (c OnenoteSectionGroupSectionPageClient) ListOnenoteSectionGroupSectionPagesComplete(ctx context.Context, id stable.MeOnenoteSectionGroupIdSectionId, options ListOnenoteSectionGroupSectionPagesOperationOptions) (ListOnenoteSectionGroupSectionPagesCompleteResult, error) {
	return c.ListOnenoteSectionGroupSectionPagesCompleteMatchingPredicate(ctx, id, options, OnenotePageOperationPredicate{})
}

// ListOnenoteSectionGroupSectionPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnenoteSectionGroupSectionPageClient) ListOnenoteSectionGroupSectionPagesCompleteMatchingPredicate(ctx context.Context, id stable.MeOnenoteSectionGroupIdSectionId, options ListOnenoteSectionGroupSectionPagesOperationOptions, predicate OnenotePageOperationPredicate) (result ListOnenoteSectionGroupSectionPagesCompleteResult, err error) {
	items := make([]stable.OnenotePage, 0)

	resp, err := c.ListOnenoteSectionGroupSectionPages(ctx, id, options)
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

	result = ListOnenoteSectionGroupSectionPagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
