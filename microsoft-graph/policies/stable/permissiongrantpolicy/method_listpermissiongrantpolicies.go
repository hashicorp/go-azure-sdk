package permissiongrantpolicy

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

type ListPermissionGrantPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PermissionGrantPolicy
}

type ListPermissionGrantPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PermissionGrantPolicy
}

type ListPermissionGrantPoliciesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListPermissionGrantPoliciesOperationOptions() ListPermissionGrantPoliciesOperationOptions {
	return ListPermissionGrantPoliciesOperationOptions{}
}

func (o ListPermissionGrantPoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPermissionGrantPoliciesOperationOptions) ToOData() *odata.Query {
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

func (o ListPermissionGrantPoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPermissionGrantPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPermissionGrantPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPermissionGrantPolicies - List permissionGrantPolicies. Retrieve the list of permissionGrantPolicy objects.
func (c PermissionGrantPolicyClient) ListPermissionGrantPolicies(ctx context.Context, options ListPermissionGrantPoliciesOperationOptions) (result ListPermissionGrantPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPermissionGrantPoliciesCustomPager{},
		Path:          "/policies/permissionGrantPolicies",
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
		Values *[]stable.PermissionGrantPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPermissionGrantPoliciesComplete retrieves all the results into a single object
func (c PermissionGrantPolicyClient) ListPermissionGrantPoliciesComplete(ctx context.Context, options ListPermissionGrantPoliciesOperationOptions) (ListPermissionGrantPoliciesCompleteResult, error) {
	return c.ListPermissionGrantPoliciesCompleteMatchingPredicate(ctx, options, PermissionGrantPolicyOperationPredicate{})
}

// ListPermissionGrantPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PermissionGrantPolicyClient) ListPermissionGrantPoliciesCompleteMatchingPredicate(ctx context.Context, options ListPermissionGrantPoliciesOperationOptions, predicate PermissionGrantPolicyOperationPredicate) (result ListPermissionGrantPoliciesCompleteResult, err error) {
	items := make([]stable.PermissionGrantPolicy, 0)

	resp, err := c.ListPermissionGrantPolicies(ctx, options)
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

	result = ListPermissionGrantPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
