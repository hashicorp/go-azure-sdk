package onenotesectiongroup

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

type ListOnenoteSectionGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SectionGroup
}

type ListOnenoteSectionGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SectionGroup
}

type ListOnenoteSectionGroupsOperationOptions struct {
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

func DefaultListOnenoteSectionGroupsOperationOptions() ListOnenoteSectionGroupsOperationOptions {
	return ListOnenoteSectionGroupsOperationOptions{}
}

func (o ListOnenoteSectionGroupsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOnenoteSectionGroupsOperationOptions) ToOData() *odata.Query {
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

func (o ListOnenoteSectionGroupsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOnenoteSectionGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnenoteSectionGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnenoteSectionGroups - Get sectionGroups from groups. The section groups in all OneNote notebooks that are owned
// by the user or group. Read-only. Nullable.
func (c OnenoteSectionGroupClient) ListOnenoteSectionGroups(ctx context.Context, id beta.GroupId, options ListOnenoteSectionGroupsOperationOptions) (result ListOnenoteSectionGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOnenoteSectionGroupsCustomPager{},
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
		Values *[]beta.SectionGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOnenoteSectionGroupsComplete retrieves all the results into a single object
func (c OnenoteSectionGroupClient) ListOnenoteSectionGroupsComplete(ctx context.Context, id beta.GroupId, options ListOnenoteSectionGroupsOperationOptions) (ListOnenoteSectionGroupsCompleteResult, error) {
	return c.ListOnenoteSectionGroupsCompleteMatchingPredicate(ctx, id, options, SectionGroupOperationPredicate{})
}

// ListOnenoteSectionGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnenoteSectionGroupClient) ListOnenoteSectionGroupsCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, options ListOnenoteSectionGroupsOperationOptions, predicate SectionGroupOperationPredicate) (result ListOnenoteSectionGroupsCompleteResult, err error) {
	items := make([]beta.SectionGroup, 0)

	resp, err := c.ListOnenoteSectionGroups(ctx, id, options)
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

	result = ListOnenoteSectionGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
