package insightused

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

type ListInsightUsedsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UsedInsight
}

type ListInsightUsedsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UsedInsight
}

type ListInsightUsedsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListInsightUsedsOperationOptions() ListInsightUsedsOperationOptions {
	return ListInsightUsedsOperationOptions{}
}

func (o ListInsightUsedsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListInsightUsedsOperationOptions) ToOData() *odata.Query {
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

func (o ListInsightUsedsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListInsightUsedsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListInsightUsedsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListInsightUseds - Get used from users. Calculated relationship that identifies the latest documents viewed or
// modified by a user, including OneDrive for work or school and SharePoint documents, ranked by recency of use.
func (c InsightUsedClient) ListInsightUseds(ctx context.Context, id stable.UserId, options ListInsightUsedsOperationOptions) (result ListInsightUsedsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListInsightUsedsCustomPager{},
		Path:          fmt.Sprintf("%s/insights/used", id.ID()),
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
		Values *[]stable.UsedInsight `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListInsightUsedsComplete retrieves all the results into a single object
func (c InsightUsedClient) ListInsightUsedsComplete(ctx context.Context, id stable.UserId, options ListInsightUsedsOperationOptions) (ListInsightUsedsCompleteResult, error) {
	return c.ListInsightUsedsCompleteMatchingPredicate(ctx, id, options, UsedInsightOperationPredicate{})
}

// ListInsightUsedsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InsightUsedClient) ListInsightUsedsCompleteMatchingPredicate(ctx context.Context, id stable.UserId, options ListInsightUsedsOperationOptions, predicate UsedInsightOperationPredicate) (result ListInsightUsedsCompleteResult, err error) {
	items := make([]stable.UsedInsight, 0)

	resp, err := c.ListInsightUseds(ctx, id, options)
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

	result = ListInsightUsedsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
