package permissiongrantpolicyexclude

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

type ListPermissionGrantPolicyExcludesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PermissionGrantConditionSet
}

type ListPermissionGrantPolicyExcludesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PermissionGrantConditionSet
}

type ListPermissionGrantPolicyExcludesOperationOptions struct {
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

func DefaultListPermissionGrantPolicyExcludesOperationOptions() ListPermissionGrantPolicyExcludesOperationOptions {
	return ListPermissionGrantPolicyExcludesOperationOptions{}
}

func (o ListPermissionGrantPolicyExcludesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPermissionGrantPolicyExcludesOperationOptions) ToOData() *odata.Query {
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

func (o ListPermissionGrantPolicyExcludesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPermissionGrantPolicyExcludesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPermissionGrantPolicyExcludesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPermissionGrantPolicyExcludes - List excludes collection of permissionGrantPolicy. Retrieve the condition sets
// that are *excluded* in a permissionGrantPolicy.
func (c PermissionGrantPolicyExcludeClient) ListPermissionGrantPolicyExcludes(ctx context.Context, id beta.PolicyPermissionGrantPolicyId, options ListPermissionGrantPolicyExcludesOperationOptions) (result ListPermissionGrantPolicyExcludesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPermissionGrantPolicyExcludesCustomPager{},
		Path:          fmt.Sprintf("%s/excludes", id.ID()),
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
		Values *[]beta.PermissionGrantConditionSet `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPermissionGrantPolicyExcludesComplete retrieves all the results into a single object
func (c PermissionGrantPolicyExcludeClient) ListPermissionGrantPolicyExcludesComplete(ctx context.Context, id beta.PolicyPermissionGrantPolicyId, options ListPermissionGrantPolicyExcludesOperationOptions) (ListPermissionGrantPolicyExcludesCompleteResult, error) {
	return c.ListPermissionGrantPolicyExcludesCompleteMatchingPredicate(ctx, id, options, PermissionGrantConditionSetOperationPredicate{})
}

// ListPermissionGrantPolicyExcludesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PermissionGrantPolicyExcludeClient) ListPermissionGrantPolicyExcludesCompleteMatchingPredicate(ctx context.Context, id beta.PolicyPermissionGrantPolicyId, options ListPermissionGrantPolicyExcludesOperationOptions, predicate PermissionGrantConditionSetOperationPredicate) (result ListPermissionGrantPolicyExcludesCompleteResult, err error) {
	items := make([]beta.PermissionGrantConditionSet, 0)

	resp, err := c.ListPermissionGrantPolicyExcludes(ctx, id, options)
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

	result = ListPermissionGrantPolicyExcludesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
