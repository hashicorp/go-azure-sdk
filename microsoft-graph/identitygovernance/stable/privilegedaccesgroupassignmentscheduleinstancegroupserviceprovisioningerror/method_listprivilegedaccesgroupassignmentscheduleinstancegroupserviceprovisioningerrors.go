package privilegedaccesgroupassignmentscheduleinstancegroupserviceprovisioningerror

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

type ListPrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ServiceProvisioningError
}

type ListPrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ServiceProvisioningError
}

type ListPrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrors ...
func (c PrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorClient) ListPrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrors(ctx context.Context, id IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId) (result ListPrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorsCustomPager{},
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

// ListPrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorsComplete retrieves all the results into a single object
func (c PrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorClient) ListPrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorsComplete(ctx context.Context, id IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId) (ListPrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorsCompleteResult, error) {
	return c.ListPrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorsCompleteMatchingPredicate(ctx, id, ServiceProvisioningErrorOperationPredicate{})
}

// ListPrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorClient) ListPrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId, predicate ServiceProvisioningErrorOperationPredicate) (result ListPrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorsCompleteResult, err error) {
	items := make([]stable.ServiceProvisioningError, 0)

	resp, err := c.ListPrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrors(ctx, id)
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

	result = ListPrivilegedAccesGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
