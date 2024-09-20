package profileaddress

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

type ListProfileAddressesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ItemAddress
}

type ListProfileAddressesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ItemAddress
}

type ListProfileAddressesOperationOptions struct {
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

func DefaultListProfileAddressesOperationOptions() ListProfileAddressesOperationOptions {
	return ListProfileAddressesOperationOptions{}
}

func (o ListProfileAddressesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListProfileAddressesOperationOptions) ToOData() *odata.Query {
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

func (o ListProfileAddressesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListProfileAddressesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileAddressesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileAddresses - List addresses. Get the itemAddress resources from the addresses navigation property.
func (c ProfileAddressClient) ListProfileAddresses(ctx context.Context, options ListProfileAddressesOperationOptions) (result ListProfileAddressesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListProfileAddressesCustomPager{},
		Path:          "/me/profile/addresses",
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
		Values *[]beta.ItemAddress `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileAddressesComplete retrieves all the results into a single object
func (c ProfileAddressClient) ListProfileAddressesComplete(ctx context.Context, options ListProfileAddressesOperationOptions) (ListProfileAddressesCompleteResult, error) {
	return c.ListProfileAddressesCompleteMatchingPredicate(ctx, options, ItemAddressOperationPredicate{})
}

// ListProfileAddressesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileAddressClient) ListProfileAddressesCompleteMatchingPredicate(ctx context.Context, options ListProfileAddressesOperationOptions, predicate ItemAddressOperationPredicate) (result ListProfileAddressesCompleteResult, err error) {
	items := make([]beta.ItemAddress, 0)

	resp, err := c.ListProfileAddresses(ctx, options)
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

	result = ListProfileAddressesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
