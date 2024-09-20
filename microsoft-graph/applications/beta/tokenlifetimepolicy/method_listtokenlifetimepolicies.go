package tokenlifetimepolicy

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

type ListTokenLifetimePoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.TokenLifetimePolicy
}

type ListTokenLifetimePoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.TokenLifetimePolicy
}

type ListTokenLifetimePoliciesOperationOptions struct {
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

func DefaultListTokenLifetimePoliciesOperationOptions() ListTokenLifetimePoliciesOperationOptions {
	return ListTokenLifetimePoliciesOperationOptions{}
}

func (o ListTokenLifetimePoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTokenLifetimePoliciesOperationOptions) ToOData() *odata.Query {
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

func (o ListTokenLifetimePoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTokenLifetimePoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTokenLifetimePoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTokenLifetimePolicies - List assigned tokenLifetimePolicy. List the tokenLifetimePolicy objects that are assigned
// to an application or servicePrincipal. Only one object is returned in the collection because only one
// tokenLifetimePolicy can be assigned to an application.
func (c TokenLifetimePolicyClient) ListTokenLifetimePolicies(ctx context.Context, id beta.ApplicationId, options ListTokenLifetimePoliciesOperationOptions) (result ListTokenLifetimePoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTokenLifetimePoliciesCustomPager{},
		Path:          fmt.Sprintf("%s/tokenLifetimePolicies", id.ID()),
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
		Values *[]beta.TokenLifetimePolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTokenLifetimePoliciesComplete retrieves all the results into a single object
func (c TokenLifetimePolicyClient) ListTokenLifetimePoliciesComplete(ctx context.Context, id beta.ApplicationId, options ListTokenLifetimePoliciesOperationOptions) (ListTokenLifetimePoliciesCompleteResult, error) {
	return c.ListTokenLifetimePoliciesCompleteMatchingPredicate(ctx, id, options, TokenLifetimePolicyOperationPredicate{})
}

// ListTokenLifetimePoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TokenLifetimePolicyClient) ListTokenLifetimePoliciesCompleteMatchingPredicate(ctx context.Context, id beta.ApplicationId, options ListTokenLifetimePoliciesOperationOptions, predicate TokenLifetimePolicyOperationPredicate) (result ListTokenLifetimePoliciesCompleteResult, err error) {
	items := make([]beta.TokenLifetimePolicy, 0)

	resp, err := c.ListTokenLifetimePolicies(ctx, id, options)
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

	result = ListTokenLifetimePoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
