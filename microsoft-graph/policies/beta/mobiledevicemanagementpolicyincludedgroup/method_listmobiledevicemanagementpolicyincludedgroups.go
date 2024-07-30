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

type ListMobileDeviceManagementPolicyIncludedGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMobileDeviceManagementPolicyIncludedGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMobileDeviceManagementPolicyIncludedGroups ...
func (c MobileDeviceManagementPolicyIncludedGroupClient) ListMobileDeviceManagementPolicyIncludedGroups(ctx context.Context, id PolicyMobileDeviceManagementPolicyId) (result ListMobileDeviceManagementPolicyIncludedGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMobileDeviceManagementPolicyIncludedGroupsCustomPager{},
		Path:       fmt.Sprintf("%s/includedGroups", id.ID()),
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
func (c MobileDeviceManagementPolicyIncludedGroupClient) ListMobileDeviceManagementPolicyIncludedGroupsComplete(ctx context.Context, id PolicyMobileDeviceManagementPolicyId) (ListMobileDeviceManagementPolicyIncludedGroupsCompleteResult, error) {
	return c.ListMobileDeviceManagementPolicyIncludedGroupsCompleteMatchingPredicate(ctx, id, GroupOperationPredicate{})
}

// ListMobileDeviceManagementPolicyIncludedGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MobileDeviceManagementPolicyIncludedGroupClient) ListMobileDeviceManagementPolicyIncludedGroupsCompleteMatchingPredicate(ctx context.Context, id PolicyMobileDeviceManagementPolicyId, predicate GroupOperationPredicate) (result ListMobileDeviceManagementPolicyIncludedGroupsCompleteResult, err error) {
	items := make([]beta.Group, 0)

	resp, err := c.ListMobileDeviceManagementPolicyIncludedGroups(ctx, id)
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
