package serviceprincipalcreationpolicy

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

type ListServicePrincipalCreationPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ServicePrincipalCreationPolicy
}

type ListServicePrincipalCreationPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ServicePrincipalCreationPolicy
}

type ListServicePrincipalCreationPoliciesOperationOptions struct {
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

func DefaultListServicePrincipalCreationPoliciesOperationOptions() ListServicePrincipalCreationPoliciesOperationOptions {
	return ListServicePrincipalCreationPoliciesOperationOptions{}
}

func (o ListServicePrincipalCreationPoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListServicePrincipalCreationPoliciesOperationOptions) ToOData() *odata.Query {
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

func (o ListServicePrincipalCreationPoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListServicePrincipalCreationPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListServicePrincipalCreationPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListServicePrincipalCreationPolicies - Get servicePrincipalCreationPolicies from policies
func (c ServicePrincipalCreationPolicyClient) ListServicePrincipalCreationPolicies(ctx context.Context, options ListServicePrincipalCreationPoliciesOperationOptions) (result ListServicePrincipalCreationPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListServicePrincipalCreationPoliciesCustomPager{},
		Path:          "/policies/servicePrincipalCreationPolicies",
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
		Values *[]beta.ServicePrincipalCreationPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListServicePrincipalCreationPoliciesComplete retrieves all the results into a single object
func (c ServicePrincipalCreationPolicyClient) ListServicePrincipalCreationPoliciesComplete(ctx context.Context, options ListServicePrincipalCreationPoliciesOperationOptions) (ListServicePrincipalCreationPoliciesCompleteResult, error) {
	return c.ListServicePrincipalCreationPoliciesCompleteMatchingPredicate(ctx, options, ServicePrincipalCreationPolicyOperationPredicate{})
}

// ListServicePrincipalCreationPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalCreationPolicyClient) ListServicePrincipalCreationPoliciesCompleteMatchingPredicate(ctx context.Context, options ListServicePrincipalCreationPoliciesOperationOptions, predicate ServicePrincipalCreationPolicyOperationPredicate) (result ListServicePrincipalCreationPoliciesCompleteResult, err error) {
	items := make([]beta.ServicePrincipalCreationPolicy, 0)

	resp, err := c.ListServicePrincipalCreationPolicies(ctx, options)
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

	result = ListServicePrincipalCreationPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
