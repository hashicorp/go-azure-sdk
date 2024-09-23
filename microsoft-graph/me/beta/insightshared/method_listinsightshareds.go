package insightshared

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

type ListInsightSharedsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SharedInsight
}

type ListInsightSharedsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SharedInsight
}

type ListInsightSharedsOperationOptions struct {
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

func DefaultListInsightSharedsOperationOptions() ListInsightSharedsOperationOptions {
	return ListInsightSharedsOperationOptions{}
}

func (o ListInsightSharedsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListInsightSharedsOperationOptions) ToOData() *odata.Query {
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

func (o ListInsightSharedsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListInsightSharedsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListInsightSharedsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListInsightShareds - Get shared from me. Access this property from the derived type itemInsights.
func (c InsightSharedClient) ListInsightShareds(ctx context.Context, options ListInsightSharedsOperationOptions) (result ListInsightSharedsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListInsightSharedsCustomPager{},
		Path:          "/me/insights/shared",
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
		Values *[]beta.SharedInsight `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListInsightSharedsComplete retrieves all the results into a single object
func (c InsightSharedClient) ListInsightSharedsComplete(ctx context.Context, options ListInsightSharedsOperationOptions) (ListInsightSharedsCompleteResult, error) {
	return c.ListInsightSharedsCompleteMatchingPredicate(ctx, options, SharedInsightOperationPredicate{})
}

// ListInsightSharedsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InsightSharedClient) ListInsightSharedsCompleteMatchingPredicate(ctx context.Context, options ListInsightSharedsOperationOptions, predicate SharedInsightOperationPredicate) (result ListInsightSharedsCompleteResult, err error) {
	items := make([]beta.SharedInsight, 0)

	resp, err := c.ListInsightShareds(ctx, options)
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

	result = ListInsightSharedsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
