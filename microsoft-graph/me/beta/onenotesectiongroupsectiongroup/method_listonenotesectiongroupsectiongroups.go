package onenotesectiongroupsectiongroup

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

type ListOnenoteSectionGroupSectionGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SectionGroup
}

type ListOnenoteSectionGroupSectionGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SectionGroup
}

type ListOnenoteSectionGroupSectionGroupsOperationOptions struct {
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

func DefaultListOnenoteSectionGroupSectionGroupsOperationOptions() ListOnenoteSectionGroupSectionGroupsOperationOptions {
	return ListOnenoteSectionGroupSectionGroupsOperationOptions{}
}

func (o ListOnenoteSectionGroupSectionGroupsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOnenoteSectionGroupSectionGroupsOperationOptions) ToOData() *odata.Query {
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

func (o ListOnenoteSectionGroupSectionGroupsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOnenoteSectionGroupSectionGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnenoteSectionGroupSectionGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnenoteSectionGroupSectionGroups - List sectionGroups. Retrieve a list of section groups from the specified
// section group.
func (c OnenoteSectionGroupSectionGroupClient) ListOnenoteSectionGroupSectionGroups(ctx context.Context, id beta.MeOnenoteSectionGroupId, options ListOnenoteSectionGroupSectionGroupsOperationOptions) (result ListOnenoteSectionGroupSectionGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOnenoteSectionGroupSectionGroupsCustomPager{},
		Path:          fmt.Sprintf("%s/sectionGroups", id.ID()),
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
		Values *[]beta.SectionGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOnenoteSectionGroupSectionGroupsComplete retrieves all the results into a single object
func (c OnenoteSectionGroupSectionGroupClient) ListOnenoteSectionGroupSectionGroupsComplete(ctx context.Context, id beta.MeOnenoteSectionGroupId, options ListOnenoteSectionGroupSectionGroupsOperationOptions) (ListOnenoteSectionGroupSectionGroupsCompleteResult, error) {
	return c.ListOnenoteSectionGroupSectionGroupsCompleteMatchingPredicate(ctx, id, options, SectionGroupOperationPredicate{})
}

// ListOnenoteSectionGroupSectionGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnenoteSectionGroupSectionGroupClient) ListOnenoteSectionGroupSectionGroupsCompleteMatchingPredicate(ctx context.Context, id beta.MeOnenoteSectionGroupId, options ListOnenoteSectionGroupSectionGroupsOperationOptions, predicate SectionGroupOperationPredicate) (result ListOnenoteSectionGroupSectionGroupsCompleteResult, err error) {
	items := make([]beta.SectionGroup, 0)

	resp, err := c.ListOnenoteSectionGroupSectionGroups(ctx, id, options)
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

	result = ListOnenoteSectionGroupSectionGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
