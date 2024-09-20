package featurerolloutpolicy

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

type ListFeatureRolloutPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.FeatureRolloutPolicy
}

type ListFeatureRolloutPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.FeatureRolloutPolicy
}

type ListFeatureRolloutPoliciesOperationOptions struct {
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

func DefaultListFeatureRolloutPoliciesOperationOptions() ListFeatureRolloutPoliciesOperationOptions {
	return ListFeatureRolloutPoliciesOperationOptions{}
}

func (o ListFeatureRolloutPoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListFeatureRolloutPoliciesOperationOptions) ToOData() *odata.Query {
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

func (o ListFeatureRolloutPoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListFeatureRolloutPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListFeatureRolloutPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListFeatureRolloutPolicies - Get featureRolloutPolicies from directory
func (c FeatureRolloutPolicyClient) ListFeatureRolloutPolicies(ctx context.Context, options ListFeatureRolloutPoliciesOperationOptions) (result ListFeatureRolloutPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListFeatureRolloutPoliciesCustomPager{},
		Path:          "/directory/featureRolloutPolicies",
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
		Values *[]beta.FeatureRolloutPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListFeatureRolloutPoliciesComplete retrieves all the results into a single object
func (c FeatureRolloutPolicyClient) ListFeatureRolloutPoliciesComplete(ctx context.Context, options ListFeatureRolloutPoliciesOperationOptions) (ListFeatureRolloutPoliciesCompleteResult, error) {
	return c.ListFeatureRolloutPoliciesCompleteMatchingPredicate(ctx, options, FeatureRolloutPolicyOperationPredicate{})
}

// ListFeatureRolloutPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c FeatureRolloutPolicyClient) ListFeatureRolloutPoliciesCompleteMatchingPredicate(ctx context.Context, options ListFeatureRolloutPoliciesOperationOptions, predicate FeatureRolloutPolicyOperationPredicate) (result ListFeatureRolloutPoliciesCompleteResult, err error) {
	items := make([]beta.FeatureRolloutPolicy, 0)

	resp, err := c.ListFeatureRolloutPolicies(ctx, options)
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

	result = ListFeatureRolloutPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
