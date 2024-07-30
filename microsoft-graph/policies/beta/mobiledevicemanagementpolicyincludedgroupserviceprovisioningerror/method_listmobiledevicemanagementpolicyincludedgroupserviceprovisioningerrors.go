package mobiledevicemanagementpolicyincludedgroupserviceprovisioningerror

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

type ListMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ServiceProvisioningError
}

type ListMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ServiceProvisioningError
}

type ListMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrors ...
func (c MobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorClient) ListMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrors(ctx context.Context, id PolicyMobileDeviceManagementPolicyIdIncludedGroupId) (result ListMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorsCustomPager{},
		Path:       fmt.Sprintf("%s/serviceProvisioningErrors", id.ID()),
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
		Values *[]beta.ServiceProvisioningError `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorsComplete retrieves all the results into a single object
func (c MobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorClient) ListMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorsComplete(ctx context.Context, id PolicyMobileDeviceManagementPolicyIdIncludedGroupId) (ListMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorsCompleteResult, error) {
	return c.ListMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorsCompleteMatchingPredicate(ctx, id, ServiceProvisioningErrorOperationPredicate{})
}

// ListMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorClient) ListMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorsCompleteMatchingPredicate(ctx context.Context, id PolicyMobileDeviceManagementPolicyIdIncludedGroupId, predicate ServiceProvisioningErrorOperationPredicate) (result ListMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorsCompleteResult, err error) {
	items := make([]beta.ServiceProvisioningError, 0)

	resp, err := c.ListMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrors(ctx, id)
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

	result = ListMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
