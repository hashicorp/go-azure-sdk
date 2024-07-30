package mobileappmanagementpolicyincludedgroupserviceprovisioningerror

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

type ListMobileAppManagementPolicyIncludedGroupServiceProvisioningErrorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ServiceProvisioningError
}

type ListMobileAppManagementPolicyIncludedGroupServiceProvisioningErrorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ServiceProvisioningError
}

type ListMobileAppManagementPolicyIncludedGroupServiceProvisioningErrorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMobileAppManagementPolicyIncludedGroupServiceProvisioningErrorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMobileAppManagementPolicyIncludedGroupServiceProvisioningErrors ...
func (c MobileAppManagementPolicyIncludedGroupServiceProvisioningErrorClient) ListMobileAppManagementPolicyIncludedGroupServiceProvisioningErrors(ctx context.Context, id PolicyMobileAppManagementPolicyIdIncludedGroupId) (result ListMobileAppManagementPolicyIncludedGroupServiceProvisioningErrorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMobileAppManagementPolicyIncludedGroupServiceProvisioningErrorsCustomPager{},
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

// ListMobileAppManagementPolicyIncludedGroupServiceProvisioningErrorsComplete retrieves all the results into a single object
func (c MobileAppManagementPolicyIncludedGroupServiceProvisioningErrorClient) ListMobileAppManagementPolicyIncludedGroupServiceProvisioningErrorsComplete(ctx context.Context, id PolicyMobileAppManagementPolicyIdIncludedGroupId) (ListMobileAppManagementPolicyIncludedGroupServiceProvisioningErrorsCompleteResult, error) {
	return c.ListMobileAppManagementPolicyIncludedGroupServiceProvisioningErrorsCompleteMatchingPredicate(ctx, id, ServiceProvisioningErrorOperationPredicate{})
}

// ListMobileAppManagementPolicyIncludedGroupServiceProvisioningErrorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MobileAppManagementPolicyIncludedGroupServiceProvisioningErrorClient) ListMobileAppManagementPolicyIncludedGroupServiceProvisioningErrorsCompleteMatchingPredicate(ctx context.Context, id PolicyMobileAppManagementPolicyIdIncludedGroupId, predicate ServiceProvisioningErrorOperationPredicate) (result ListMobileAppManagementPolicyIncludedGroupServiceProvisioningErrorsCompleteResult, err error) {
	items := make([]beta.ServiceProvisioningError, 0)

	resp, err := c.ListMobileAppManagementPolicyIncludedGroupServiceProvisioningErrors(ctx, id)
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

	result = ListMobileAppManagementPolicyIncludedGroupServiceProvisioningErrorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
