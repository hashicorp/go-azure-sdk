package conditionalaccessauthenticationstrengthpolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListConditionalAccessAuthenticationStrengthPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AuthenticationStrengthPolicy
}

type ListConditionalAccessAuthenticationStrengthPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AuthenticationStrengthPolicy
}

type ListConditionalAccessAuthenticationStrengthPoliciesOperationOptions struct {
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

func DefaultListConditionalAccessAuthenticationStrengthPoliciesOperationOptions() ListConditionalAccessAuthenticationStrengthPoliciesOperationOptions {
	return ListConditionalAccessAuthenticationStrengthPoliciesOperationOptions{}
}

func (o ListConditionalAccessAuthenticationStrengthPoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListConditionalAccessAuthenticationStrengthPoliciesOperationOptions) ToOData() *odata.Query {
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

func (o ListConditionalAccessAuthenticationStrengthPoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListConditionalAccessAuthenticationStrengthPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConditionalAccessAuthenticationStrengthPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConditionalAccessAuthenticationStrengthPolicies - Get policies from identity. A collection of authentication
// strength policies that exist for this tenant, including both built-in and custom policies.
func (c ConditionalAccessAuthenticationStrengthPolicyClient) ListConditionalAccessAuthenticationStrengthPolicies(ctx context.Context, options ListConditionalAccessAuthenticationStrengthPoliciesOperationOptions) (result ListConditionalAccessAuthenticationStrengthPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListConditionalAccessAuthenticationStrengthPoliciesCustomPager{},
		Path:          "/identity/conditionalAccess/authenticationStrength/policies",
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
		Values *[]stable.AuthenticationStrengthPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConditionalAccessAuthenticationStrengthPoliciesComplete retrieves all the results into a single object
func (c ConditionalAccessAuthenticationStrengthPolicyClient) ListConditionalAccessAuthenticationStrengthPoliciesComplete(ctx context.Context, options ListConditionalAccessAuthenticationStrengthPoliciesOperationOptions) (ListConditionalAccessAuthenticationStrengthPoliciesCompleteResult, error) {
	return c.ListConditionalAccessAuthenticationStrengthPoliciesCompleteMatchingPredicate(ctx, options, AuthenticationStrengthPolicyOperationPredicate{})
}

// ListConditionalAccessAuthenticationStrengthPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConditionalAccessAuthenticationStrengthPolicyClient) ListConditionalAccessAuthenticationStrengthPoliciesCompleteMatchingPredicate(ctx context.Context, options ListConditionalAccessAuthenticationStrengthPoliciesOperationOptions, predicate AuthenticationStrengthPolicyOperationPredicate) (result ListConditionalAccessAuthenticationStrengthPoliciesCompleteResult, err error) {
	items := make([]stable.AuthenticationStrengthPolicy, 0)

	resp, err := c.ListConditionalAccessAuthenticationStrengthPolicies(ctx, options)
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

	result = ListConditionalAccessAuthenticationStrengthPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
