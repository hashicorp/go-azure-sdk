package permissiongrantpreapprovalpolicy

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

type ListPermissionGrantPreApprovalPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PermissionGrantPreApprovalPolicy
}

type ListPermissionGrantPreApprovalPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PermissionGrantPreApprovalPolicy
}

type ListPermissionGrantPreApprovalPoliciesOperationOptions struct {
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

func DefaultListPermissionGrantPreApprovalPoliciesOperationOptions() ListPermissionGrantPreApprovalPoliciesOperationOptions {
	return ListPermissionGrantPreApprovalPoliciesOperationOptions{}
}

func (o ListPermissionGrantPreApprovalPoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPermissionGrantPreApprovalPoliciesOperationOptions) ToOData() *odata.Query {
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

func (o ListPermissionGrantPreApprovalPoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPermissionGrantPreApprovalPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPermissionGrantPreApprovalPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPermissionGrantPreApprovalPolicies - List permissionGrantPreApprovalPolicies. Retrieve the list of
// permissionGrantPreApprovalPolicy objects in the tenant.
func (c PermissionGrantPreApprovalPolicyClient) ListPermissionGrantPreApprovalPolicies(ctx context.Context, options ListPermissionGrantPreApprovalPoliciesOperationOptions) (result ListPermissionGrantPreApprovalPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPermissionGrantPreApprovalPoliciesCustomPager{},
		Path:          "/policies/permissionGrantPreApprovalPolicies",
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
		Values *[]beta.PermissionGrantPreApprovalPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPermissionGrantPreApprovalPoliciesComplete retrieves all the results into a single object
func (c PermissionGrantPreApprovalPolicyClient) ListPermissionGrantPreApprovalPoliciesComplete(ctx context.Context, options ListPermissionGrantPreApprovalPoliciesOperationOptions) (ListPermissionGrantPreApprovalPoliciesCompleteResult, error) {
	return c.ListPermissionGrantPreApprovalPoliciesCompleteMatchingPredicate(ctx, options, PermissionGrantPreApprovalPolicyOperationPredicate{})
}

// ListPermissionGrantPreApprovalPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PermissionGrantPreApprovalPolicyClient) ListPermissionGrantPreApprovalPoliciesCompleteMatchingPredicate(ctx context.Context, options ListPermissionGrantPreApprovalPoliciesOperationOptions, predicate PermissionGrantPreApprovalPolicyOperationPredicate) (result ListPermissionGrantPreApprovalPoliciesCompleteResult, err error) {
	items := make([]beta.PermissionGrantPreApprovalPolicy, 0)

	resp, err := c.ListPermissionGrantPreApprovalPolicies(ctx, options)
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

	result = ListPermissionGrantPreApprovalPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
