package siteonenotesectiongroupsectionpage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSiteOnenoteSectionGroupSectionPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OnenotePage
}

type ListSiteOnenoteSectionGroupSectionPagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OnenotePage
}

type ListSiteOnenoteSectionGroupSectionPagesOperationOptions struct {
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

func DefaultListSiteOnenoteSectionGroupSectionPagesOperationOptions() ListSiteOnenoteSectionGroupSectionPagesOperationOptions {
	return ListSiteOnenoteSectionGroupSectionPagesOperationOptions{}
}

func (o ListSiteOnenoteSectionGroupSectionPagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteOnenoteSectionGroupSectionPagesOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteOnenoteSectionGroupSectionPagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteOnenoteSectionGroupSectionPagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteOnenoteSectionGroupSectionPagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteOnenoteSectionGroupSectionPages - Get pages from groups. The collection of pages in the section. Read-only.
// Nullable.
func (c SiteOnenoteSectionGroupSectionPageClient) ListSiteOnenoteSectionGroupSectionPages(ctx context.Context, id beta.GroupIdSiteIdOnenoteSectionGroupIdSectionId, options ListSiteOnenoteSectionGroupSectionPagesOperationOptions) (result ListSiteOnenoteSectionGroupSectionPagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteOnenoteSectionGroupSectionPagesCustomPager{},
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
		Values *[]beta.OnenotePage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteOnenoteSectionGroupSectionPagesComplete retrieves all the results into a single object
func (c SiteOnenoteSectionGroupSectionPageClient) ListSiteOnenoteSectionGroupSectionPagesComplete(ctx context.Context, id beta.GroupIdSiteIdOnenoteSectionGroupIdSectionId, options ListSiteOnenoteSectionGroupSectionPagesOperationOptions) (ListSiteOnenoteSectionGroupSectionPagesCompleteResult, error) {
	return c.ListSiteOnenoteSectionGroupSectionPagesCompleteMatchingPredicate(ctx, id, options, OnenotePageOperationPredicate{})
}

// ListSiteOnenoteSectionGroupSectionPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteOnenoteSectionGroupSectionPageClient) ListSiteOnenoteSectionGroupSectionPagesCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteIdOnenoteSectionGroupIdSectionId, options ListSiteOnenoteSectionGroupSectionPagesOperationOptions, predicate OnenotePageOperationPredicate) (result ListSiteOnenoteSectionGroupSectionPagesCompleteResult, err error) {
	items := make([]beta.OnenotePage, 0)

	resp, err := c.ListSiteOnenoteSectionGroupSectionPages(ctx, id, options)
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

	result = ListSiteOnenoteSectionGroupSectionPagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
