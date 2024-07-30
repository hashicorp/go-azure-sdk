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

type ListMobileDeviceManagementPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMobileDeviceManagementPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMobileDeviceManagementPolicies ...
func (c MobileDeviceManagementPolicyClient) ListMobileDeviceManagementPolicies(ctx context.Context) (result ListMobileDeviceManagementPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMobileDeviceManagementPoliciesCustomPager{},
		Path:       "/policies/mobileDeviceManagementPolicies",
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
func (c MobileDeviceManagementPolicyClient) ListMobileDeviceManagementPoliciesComplete(ctx context.Context) (ListMobileDeviceManagementPoliciesCompleteResult, error) {
	return c.ListMobileDeviceManagementPoliciesCompleteMatchingPredicate(ctx, MobilityManagementPolicyOperationPredicate{})
}

// ListMobileDeviceManagementPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MobileDeviceManagementPolicyClient) ListMobileDeviceManagementPoliciesCompleteMatchingPredicate(ctx context.Context, predicate MobilityManagementPolicyOperationPredicate) (result ListMobileDeviceManagementPoliciesCompleteResult, err error) {
	items := make([]beta.MobilityManagementPolicy, 0)

	resp, err := c.ListMobileDeviceManagementPolicies(ctx)
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
