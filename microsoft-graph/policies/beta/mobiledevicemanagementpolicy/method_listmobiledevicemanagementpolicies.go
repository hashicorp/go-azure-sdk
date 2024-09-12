package mobiledevicemanagementpolicy

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

type ListMobileDeviceManagementPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MobilityManagementPolicy
}

type ListMobileDeviceManagementPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MobilityManagementPolicy
}

type ListMobileDeviceManagementPoliciesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListMobileDeviceManagementPoliciesOperationOptions() ListMobileDeviceManagementPoliciesOperationOptions {
	return ListMobileDeviceManagementPoliciesOperationOptions{}
}

func (o ListMobileDeviceManagementPoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMobileDeviceManagementPoliciesOperationOptions) ToOData() *odata.Query {
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

func (o ListMobileDeviceManagementPoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMobileDeviceManagementPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMobileDeviceManagementPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMobileDeviceManagementPolicies - List mobileDeviceManagementPolicies. Get a list of the mobilityManagementPolicy
// objects and their properties.
func (c MobileDeviceManagementPolicyClient) ListMobileDeviceManagementPolicies(ctx context.Context, options ListMobileDeviceManagementPoliciesOperationOptions) (result ListMobileDeviceManagementPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMobileDeviceManagementPoliciesCustomPager{},
		Path:          "/policies/mobileDeviceManagementPolicies",
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
		Values *[]beta.MobilityManagementPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMobileDeviceManagementPoliciesComplete retrieves all the results into a single object
func (c MobileDeviceManagementPolicyClient) ListMobileDeviceManagementPoliciesComplete(ctx context.Context, options ListMobileDeviceManagementPoliciesOperationOptions) (ListMobileDeviceManagementPoliciesCompleteResult, error) {
	return c.ListMobileDeviceManagementPoliciesCompleteMatchingPredicate(ctx, options, MobilityManagementPolicyOperationPredicate{})
}

// ListMobileDeviceManagementPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MobileDeviceManagementPolicyClient) ListMobileDeviceManagementPoliciesCompleteMatchingPredicate(ctx context.Context, options ListMobileDeviceManagementPoliciesOperationOptions, predicate MobilityManagementPolicyOperationPredicate) (result ListMobileDeviceManagementPoliciesCompleteResult, err error) {
	items := make([]beta.MobilityManagementPolicy, 0)

	resp, err := c.ListMobileDeviceManagementPolicies(ctx, options)
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

	result = ListMobileDeviceManagementPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
