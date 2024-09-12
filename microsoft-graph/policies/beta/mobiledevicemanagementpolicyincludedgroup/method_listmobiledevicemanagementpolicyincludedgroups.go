package mobiledevicemanagementpolicyincludedgroup

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

type ListMobileDeviceManagementPolicyIncludedGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Group
}

type ListMobileDeviceManagementPolicyIncludedGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Group
}

type ListMobileDeviceManagementPolicyIncludedGroupsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListMobileDeviceManagementPolicyIncludedGroupsOperationOptions() ListMobileDeviceManagementPolicyIncludedGroupsOperationOptions {
	return ListMobileDeviceManagementPolicyIncludedGroupsOperationOptions{}
}

func (o ListMobileDeviceManagementPolicyIncludedGroupsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMobileDeviceManagementPolicyIncludedGroupsOperationOptions) ToOData() *odata.Query {
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

func (o ListMobileDeviceManagementPolicyIncludedGroupsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMobileDeviceManagementPolicyIncludedGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMobileDeviceManagementPolicyIncludedGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMobileDeviceManagementPolicyIncludedGroups - List includedGroups. Get the list of groups that are included in a
// mobile device management policy.
func (c MobileDeviceManagementPolicyIncludedGroupClient) ListMobileDeviceManagementPolicyIncludedGroups(ctx context.Context, id beta.PolicyMobileDeviceManagementPolicyId, options ListMobileDeviceManagementPolicyIncludedGroupsOperationOptions) (result ListMobileDeviceManagementPolicyIncludedGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMobileDeviceManagementPolicyIncludedGroupsCustomPager{},
		Path:          fmt.Sprintf("%s/includedGroups", id.ID()),
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
		Values *[]beta.Group `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMobileDeviceManagementPolicyIncludedGroupsComplete retrieves all the results into a single object
func (c MobileDeviceManagementPolicyIncludedGroupClient) ListMobileDeviceManagementPolicyIncludedGroupsComplete(ctx context.Context, id beta.PolicyMobileDeviceManagementPolicyId, options ListMobileDeviceManagementPolicyIncludedGroupsOperationOptions) (ListMobileDeviceManagementPolicyIncludedGroupsCompleteResult, error) {
	return c.ListMobileDeviceManagementPolicyIncludedGroupsCompleteMatchingPredicate(ctx, id, options, GroupOperationPredicate{})
}

// ListMobileDeviceManagementPolicyIncludedGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MobileDeviceManagementPolicyIncludedGroupClient) ListMobileDeviceManagementPolicyIncludedGroupsCompleteMatchingPredicate(ctx context.Context, id beta.PolicyMobileDeviceManagementPolicyId, options ListMobileDeviceManagementPolicyIncludedGroupsOperationOptions, predicate GroupOperationPredicate) (result ListMobileDeviceManagementPolicyIncludedGroupsCompleteResult, err error) {
	items := make([]beta.Group, 0)

	resp, err := c.ListMobileDeviceManagementPolicyIncludedGroups(ctx, id, options)
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

	result = ListMobileDeviceManagementPolicyIncludedGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
