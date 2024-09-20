package virtualeventwebinar

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

type ListVirtualEventWebinarsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.VirtualEventWebinar
}

type ListVirtualEventWebinarsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.VirtualEventWebinar
}

type ListVirtualEventWebinarsOperationOptions struct {
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

func DefaultListVirtualEventWebinarsOperationOptions() ListVirtualEventWebinarsOperationOptions {
	return ListVirtualEventWebinarsOperationOptions{}
}

func (o ListVirtualEventWebinarsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListVirtualEventWebinarsOperationOptions) ToOData() *odata.Query {
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

func (o ListVirtualEventWebinarsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListVirtualEventWebinarsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEventWebinarsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEventWebinars - Get webinars from me
func (c VirtualEventWebinarClient) ListVirtualEventWebinars(ctx context.Context, options ListVirtualEventWebinarsOperationOptions) (result ListVirtualEventWebinarsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListVirtualEventWebinarsCustomPager{},
		Path:          "/me/virtualEvents/webinars",
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
		Values *[]beta.VirtualEventWebinar `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEventWebinarsComplete retrieves all the results into a single object
func (c VirtualEventWebinarClient) ListVirtualEventWebinarsComplete(ctx context.Context, options ListVirtualEventWebinarsOperationOptions) (ListVirtualEventWebinarsCompleteResult, error) {
	return c.ListVirtualEventWebinarsCompleteMatchingPredicate(ctx, options, VirtualEventWebinarOperationPredicate{})
}

// ListVirtualEventWebinarsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEventWebinarClient) ListVirtualEventWebinarsCompleteMatchingPredicate(ctx context.Context, options ListVirtualEventWebinarsOperationOptions, predicate VirtualEventWebinarOperationPredicate) (result ListVirtualEventWebinarsCompleteResult, err error) {
	items := make([]beta.VirtualEventWebinar, 0)

	resp, err := c.ListVirtualEventWebinars(ctx, options)
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

	result = ListVirtualEventWebinarsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
