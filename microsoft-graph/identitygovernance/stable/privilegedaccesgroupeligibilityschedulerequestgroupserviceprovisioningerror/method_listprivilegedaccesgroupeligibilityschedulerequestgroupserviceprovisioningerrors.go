package privilegedaccesgroupeligibilityschedulerequestgroupserviceprovisioningerror

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ServiceProvisioningError
}

type ListPrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ServiceProvisioningError
}

type ListPrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrors ...
func (c PrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrorClient) ListPrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrors(ctx context.Context, id IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId) (result ListPrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrorsCustomPager{},
		Path:       fmt.Sprintf("%s/group/serviceProvisioningErrors", id.ID()),
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
		Values *[]stable.ServiceProvisioningError `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrorsComplete retrieves all the results into a single object
func (c PrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrorClient) ListPrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrorsComplete(ctx context.Context, id IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId) (ListPrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrorsCompleteResult, error) {
	return c.ListPrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrorsCompleteMatchingPredicate(ctx, id, ServiceProvisioningErrorOperationPredicate{})
}

// ListPrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrorClient) ListPrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrorsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId, predicate ServiceProvisioningErrorOperationPredicate) (result ListPrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrorsCompleteResult, err error) {
	items := make([]stable.ServiceProvisioningError, 0)

	resp, err := c.ListPrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrors(ctx, id)
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

	result = ListPrivilegedAccesGroupEligibilityScheduleRequestGroupServiceProvisioningErrorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
