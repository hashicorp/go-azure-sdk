package appmanagementpolicy

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

type ListAppManagementPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AppManagementPolicy
}

type ListAppManagementPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AppManagementPolicy
}

type ListAppManagementPoliciesOperationOptions struct {
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

func DefaultListAppManagementPoliciesOperationOptions() ListAppManagementPoliciesOperationOptions {
	return ListAppManagementPoliciesOperationOptions{}
}

func (o ListAppManagementPoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAppManagementPoliciesOperationOptions) ToOData() *odata.Query {
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

func (o ListAppManagementPoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAppManagementPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppManagementPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppManagementPolicies - Get appManagementPolicies from applications. The appManagementPolicy applied to this
// application.
func (c AppManagementPolicyClient) ListAppManagementPolicies(ctx context.Context, id beta.ApplicationId, options ListAppManagementPoliciesOperationOptions) (result ListAppManagementPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAppManagementPoliciesCustomPager{},
		Path:          fmt.Sprintf("%s/appManagementPolicies", id.ID()),
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
		Values *[]beta.AppManagementPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppManagementPoliciesComplete retrieves all the results into a single object
func (c AppManagementPolicyClient) ListAppManagementPoliciesComplete(ctx context.Context, id beta.ApplicationId, options ListAppManagementPoliciesOperationOptions) (ListAppManagementPoliciesCompleteResult, error) {
	return c.ListAppManagementPoliciesCompleteMatchingPredicate(ctx, id, options, AppManagementPolicyOperationPredicate{})
}

// ListAppManagementPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppManagementPolicyClient) ListAppManagementPoliciesCompleteMatchingPredicate(ctx context.Context, id beta.ApplicationId, options ListAppManagementPoliciesOperationOptions, predicate AppManagementPolicyOperationPredicate) (result ListAppManagementPoliciesCompleteResult, err error) {
	items := make([]beta.AppManagementPolicy, 0)

	resp, err := c.ListAppManagementPolicies(ctx, id, options)
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

	result = ListAppManagementPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
