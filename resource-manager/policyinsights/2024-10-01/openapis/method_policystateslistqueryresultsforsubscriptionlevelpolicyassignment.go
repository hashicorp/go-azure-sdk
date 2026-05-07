package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PolicyState
}

type PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PolicyState
}

type PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions struct {
	Apply   *string
	Filter  *string
	From    *string
	Orderby *string
	Select  *string
	To      *string
	Top     *int64
}

func DefaultPolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions() PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions {
	return PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions{}
}

func (o PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Apply != nil {
		out.Append("$apply", fmt.Sprintf("%v", *o.Apply))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.From != nil {
		out.Append("$from", fmt.Sprintf("%v", *o.From))
	}
	if o.Orderby != nil {
		out.Append("$orderby", fmt.Sprintf("%v", *o.Orderby))
	}
	if o.Select != nil {
		out.Append("$select", fmt.Sprintf("%v", *o.Select))
	}
	if o.To != nil {
		out.Append("$to", fmt.Sprintf("%v", *o.To))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignment ...
func (c OpenapisClient) PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignment(ctx context.Context, id PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId, options PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions) (result PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentCustomPager{},
		Path:          fmt.Sprintf("%s/queryResults", id.ID()),
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
		Values *[]PolicyState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentComplete retrieves all the results into a single object
func (c OpenapisClient) PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentComplete(ctx context.Context, id PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId, options PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions) (PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentCompleteResult, error) {
	return c.PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentCompleteMatchingPredicate(ctx, id, options, PolicyStateOperationPredicate{})
}

// PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentCompleteMatchingPredicate(ctx context.Context, id PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId, options PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions, predicate PolicyStateOperationPredicate) (result PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentCompleteResult, err error) {
	items := make([]PolicyState, 0)

	resp, err := c.PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignment(ctx, id, options)
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

	result = PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
