package inferenceclassificationoverride

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

type ListInferenceClassificationOverridesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.InferenceClassificationOverride
}

type ListInferenceClassificationOverridesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.InferenceClassificationOverride
}

type ListInferenceClassificationOverridesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListInferenceClassificationOverridesOperationOptions() ListInferenceClassificationOverridesOperationOptions {
	return ListInferenceClassificationOverridesOperationOptions{}
}

func (o ListInferenceClassificationOverridesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListInferenceClassificationOverridesOperationOptions) ToOData() *odata.Query {
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

func (o ListInferenceClassificationOverridesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListInferenceClassificationOverridesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListInferenceClassificationOverridesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListInferenceClassificationOverrides - List overrides. Get the Focused Inbox overrides that a user has set up to
// always classify messages from certain senders in specific ways. Each override corresponds to an SMTP address of a
// sender. Initially, a user does not have any overrides.
func (c InferenceClassificationOverrideClient) ListInferenceClassificationOverrides(ctx context.Context, options ListInferenceClassificationOverridesOperationOptions) (result ListInferenceClassificationOverridesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListInferenceClassificationOverridesCustomPager{},
		Path:          "/me/inferenceClassification/overrides",
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
		Values *[]beta.InferenceClassificationOverride `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListInferenceClassificationOverridesComplete retrieves all the results into a single object
func (c InferenceClassificationOverrideClient) ListInferenceClassificationOverridesComplete(ctx context.Context, options ListInferenceClassificationOverridesOperationOptions) (ListInferenceClassificationOverridesCompleteResult, error) {
	return c.ListInferenceClassificationOverridesCompleteMatchingPredicate(ctx, options, InferenceClassificationOverrideOperationPredicate{})
}

// ListInferenceClassificationOverridesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InferenceClassificationOverrideClient) ListInferenceClassificationOverridesCompleteMatchingPredicate(ctx context.Context, options ListInferenceClassificationOverridesOperationOptions, predicate InferenceClassificationOverrideOperationPredicate) (result ListInferenceClassificationOverridesCompleteResult, err error) {
	items := make([]beta.InferenceClassificationOverride, 0)

	resp, err := c.ListInferenceClassificationOverrides(ctx, options)
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

	result = ListInferenceClassificationOverridesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
