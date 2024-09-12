package permissiongrantpolicyinclude

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

type ListPermissionGrantPolicyIncludesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PermissionGrantConditionSet
}

type ListPermissionGrantPolicyIncludesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PermissionGrantConditionSet
}

type ListPermissionGrantPolicyIncludesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListPermissionGrantPolicyIncludesOperationOptions() ListPermissionGrantPolicyIncludesOperationOptions {
	return ListPermissionGrantPolicyIncludesOperationOptions{}
}

func (o ListPermissionGrantPolicyIncludesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPermissionGrantPolicyIncludesOperationOptions) ToOData() *odata.Query {
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

func (o ListPermissionGrantPolicyIncludesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPermissionGrantPolicyIncludesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPermissionGrantPolicyIncludesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPermissionGrantPolicyIncludes - List includes collection of permissionGrantPolicy. Retrieve the condition sets
// which are *included* in a permissionGrantPolicy.
func (c PermissionGrantPolicyIncludeClient) ListPermissionGrantPolicyIncludes(ctx context.Context, id stable.PolicyPermissionGrantPolicyId, options ListPermissionGrantPolicyIncludesOperationOptions) (result ListPermissionGrantPolicyIncludesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPermissionGrantPolicyIncludesCustomPager{},
		Path:          fmt.Sprintf("%s/includes", id.ID()),
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
		Values *[]stable.PermissionGrantConditionSet `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPermissionGrantPolicyIncludesComplete retrieves all the results into a single object
func (c PermissionGrantPolicyIncludeClient) ListPermissionGrantPolicyIncludesComplete(ctx context.Context, id stable.PolicyPermissionGrantPolicyId, options ListPermissionGrantPolicyIncludesOperationOptions) (ListPermissionGrantPolicyIncludesCompleteResult, error) {
	return c.ListPermissionGrantPolicyIncludesCompleteMatchingPredicate(ctx, id, options, PermissionGrantConditionSetOperationPredicate{})
}

// ListPermissionGrantPolicyIncludesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PermissionGrantPolicyIncludeClient) ListPermissionGrantPolicyIncludesCompleteMatchingPredicate(ctx context.Context, id stable.PolicyPermissionGrantPolicyId, options ListPermissionGrantPolicyIncludesOperationOptions, predicate PermissionGrantConditionSetOperationPredicate) (result ListPermissionGrantPolicyIncludesCompleteResult, err error) {
	items := make([]stable.PermissionGrantConditionSet, 0)

	resp, err := c.ListPermissionGrantPolicyIncludes(ctx, id, options)
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

	result = ListPermissionGrantPolicyIncludesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
