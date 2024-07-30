package mobileappmanagementpolicy

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

type ListMobileAppManagementPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MobilityManagementPolicy
}

type ListMobileAppManagementPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MobilityManagementPolicy
}

type ListMobileAppManagementPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMobileAppManagementPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMobileAppManagementPolicies ...
func (c MobileAppManagementPolicyClient) ListMobileAppManagementPolicies(ctx context.Context) (result ListMobileAppManagementPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMobileAppManagementPoliciesCustomPager{},
		Path:       "/policies/mobileAppManagementPolicies",
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

// ListMobileAppManagementPoliciesComplete retrieves all the results into a single object
func (c MobileAppManagementPolicyClient) ListMobileAppManagementPoliciesComplete(ctx context.Context) (ListMobileAppManagementPoliciesCompleteResult, error) {
	return c.ListMobileAppManagementPoliciesCompleteMatchingPredicate(ctx, MobilityManagementPolicyOperationPredicate{})
}

// ListMobileAppManagementPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MobileAppManagementPolicyClient) ListMobileAppManagementPoliciesCompleteMatchingPredicate(ctx context.Context, predicate MobilityManagementPolicyOperationPredicate) (result ListMobileAppManagementPoliciesCompleteResult, err error) {
	items := make([]beta.MobilityManagementPolicy, 0)

	resp, err := c.ListMobileAppManagementPolicies(ctx)
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

	result = ListMobileAppManagementPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
