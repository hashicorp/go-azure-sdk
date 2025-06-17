package onlinemeetingaiinsight

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

type ListOnlineMeetingAiInsightsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CallAiInsight
}

type ListOnlineMeetingAiInsightsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CallAiInsight
}

type ListOnlineMeetingAiInsightsOperationOptions struct {
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

func DefaultListOnlineMeetingAiInsightsOperationOptions() ListOnlineMeetingAiInsightsOperationOptions {
	return ListOnlineMeetingAiInsightsOperationOptions{}
}

func (o ListOnlineMeetingAiInsightsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOnlineMeetingAiInsightsOperationOptions) ToOData() *odata.Query {
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

func (o ListOnlineMeetingAiInsightsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOnlineMeetingAiInsightsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnlineMeetingAiInsightsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnlineMeetingAiInsights - Get callAiInsight. Get a callAiInsight object associated with an onlineMeeting. This
// API returns the metadata and content of the single set of AI insights associated with the online meeting.
func (c OnlineMeetingAiInsightClient) ListOnlineMeetingAiInsights(ctx context.Context, id beta.MeOnlineMeetingId, options ListOnlineMeetingAiInsightsOperationOptions) (result ListOnlineMeetingAiInsightsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOnlineMeetingAiInsightsCustomPager{},
		Path:          fmt.Sprintf("%s/aiInsights", id.ID()),
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
		Values *[]beta.CallAiInsight `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOnlineMeetingAiInsightsComplete retrieves all the results into a single object
func (c OnlineMeetingAiInsightClient) ListOnlineMeetingAiInsightsComplete(ctx context.Context, id beta.MeOnlineMeetingId, options ListOnlineMeetingAiInsightsOperationOptions) (ListOnlineMeetingAiInsightsCompleteResult, error) {
	return c.ListOnlineMeetingAiInsightsCompleteMatchingPredicate(ctx, id, options, CallAiInsightOperationPredicate{})
}

// ListOnlineMeetingAiInsightsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnlineMeetingAiInsightClient) ListOnlineMeetingAiInsightsCompleteMatchingPredicate(ctx context.Context, id beta.MeOnlineMeetingId, options ListOnlineMeetingAiInsightsOperationOptions, predicate CallAiInsightOperationPredicate) (result ListOnlineMeetingAiInsightsCompleteResult, err error) {
	items := make([]beta.CallAiInsight, 0)

	resp, err := c.ListOnlineMeetingAiInsights(ctx, id, options)
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

	result = ListOnlineMeetingAiInsightsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
