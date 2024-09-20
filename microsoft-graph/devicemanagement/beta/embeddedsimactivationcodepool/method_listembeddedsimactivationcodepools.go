package embeddedsimactivationcodepool

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

type ListEmbeddedSIMActivationCodePoolsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.EmbeddedSIMActivationCodePool
}

type ListEmbeddedSIMActivationCodePoolsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.EmbeddedSIMActivationCodePool
}

type ListEmbeddedSIMActivationCodePoolsOperationOptions struct {
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

func DefaultListEmbeddedSIMActivationCodePoolsOperationOptions() ListEmbeddedSIMActivationCodePoolsOperationOptions {
	return ListEmbeddedSIMActivationCodePoolsOperationOptions{}
}

func (o ListEmbeddedSIMActivationCodePoolsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEmbeddedSIMActivationCodePoolsOperationOptions) ToOData() *odata.Query {
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

func (o ListEmbeddedSIMActivationCodePoolsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEmbeddedSIMActivationCodePoolsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEmbeddedSIMActivationCodePoolsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEmbeddedSIMActivationCodePools - Get embeddedSIMActivationCodePools from deviceManagement. The embedded SIM
// activation code pools created by this account.
func (c EmbeddedSIMActivationCodePoolClient) ListEmbeddedSIMActivationCodePools(ctx context.Context, options ListEmbeddedSIMActivationCodePoolsOperationOptions) (result ListEmbeddedSIMActivationCodePoolsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEmbeddedSIMActivationCodePoolsCustomPager{},
		Path:          "/deviceManagement/embeddedSIMActivationCodePools",
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
		Values *[]beta.EmbeddedSIMActivationCodePool `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEmbeddedSIMActivationCodePoolsComplete retrieves all the results into a single object
func (c EmbeddedSIMActivationCodePoolClient) ListEmbeddedSIMActivationCodePoolsComplete(ctx context.Context, options ListEmbeddedSIMActivationCodePoolsOperationOptions) (ListEmbeddedSIMActivationCodePoolsCompleteResult, error) {
	return c.ListEmbeddedSIMActivationCodePoolsCompleteMatchingPredicate(ctx, options, EmbeddedSIMActivationCodePoolOperationPredicate{})
}

// ListEmbeddedSIMActivationCodePoolsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EmbeddedSIMActivationCodePoolClient) ListEmbeddedSIMActivationCodePoolsCompleteMatchingPredicate(ctx context.Context, options ListEmbeddedSIMActivationCodePoolsOperationOptions, predicate EmbeddedSIMActivationCodePoolOperationPredicate) (result ListEmbeddedSIMActivationCodePoolsCompleteResult, err error) {
	items := make([]beta.EmbeddedSIMActivationCodePool, 0)

	resp, err := c.ListEmbeddedSIMActivationCodePools(ctx, options)
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

	result = ListEmbeddedSIMActivationCodePoolsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
