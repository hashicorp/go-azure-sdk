package siteonenotesectiongroup

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

type ListSiteOnenoteSectionGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.SectionGroup
}

type ListSiteOnenoteSectionGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.SectionGroup
}

type ListSiteOnenoteSectionGroupsOperationOptions struct {
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

func DefaultListSiteOnenoteSectionGroupsOperationOptions() ListSiteOnenoteSectionGroupsOperationOptions {
	return ListSiteOnenoteSectionGroupsOperationOptions{}
}

func (o ListSiteOnenoteSectionGroupsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteOnenoteSectionGroupsOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteOnenoteSectionGroupsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteOnenoteSectionGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteOnenoteSectionGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteOnenoteSectionGroups - Get sectionGroups from groups. The section groups in all OneNote notebooks that are
// owned by the user or group. Read-only. Nullable.
func (c SiteOnenoteSectionGroupClient) ListSiteOnenoteSectionGroups(ctx context.Context, id stable.GroupIdSiteId, options ListSiteOnenoteSectionGroupsOperationOptions) (result ListSiteOnenoteSectionGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteOnenoteSectionGroupsCustomPager{},
		Path:          fmt.Sprintf("%s/onenote/sectionGroups", id.ID()),
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
		Values *[]stable.SectionGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteOnenoteSectionGroupsComplete retrieves all the results into a single object
func (c SiteOnenoteSectionGroupClient) ListSiteOnenoteSectionGroupsComplete(ctx context.Context, id stable.GroupIdSiteId, options ListSiteOnenoteSectionGroupsOperationOptions) (ListSiteOnenoteSectionGroupsCompleteResult, error) {
	return c.ListSiteOnenoteSectionGroupsCompleteMatchingPredicate(ctx, id, options, SectionGroupOperationPredicate{})
}

// ListSiteOnenoteSectionGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteOnenoteSectionGroupClient) ListSiteOnenoteSectionGroupsCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdSiteId, options ListSiteOnenoteSectionGroupsOperationOptions, predicate SectionGroupOperationPredicate) (result ListSiteOnenoteSectionGroupsCompleteResult, err error) {
	items := make([]stable.SectionGroup, 0)

	resp, err := c.ListSiteOnenoteSectionGroups(ctx, id, options)
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

	result = ListSiteOnenoteSectionGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
