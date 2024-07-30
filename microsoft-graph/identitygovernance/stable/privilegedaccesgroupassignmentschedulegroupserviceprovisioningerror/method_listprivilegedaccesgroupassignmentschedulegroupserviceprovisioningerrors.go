package privilegedaccesgroupassignmentschedulegroupserviceprovisioningerror

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

type ListPrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ServiceProvisioningError
}

type ListPrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ServiceProvisioningError
}

type ListPrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrors ...
func (c PrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrorClient) ListPrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrors(ctx context.Context, id IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId) (result ListPrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrorsCustomPager{},
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

// ListPrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrorsComplete retrieves all the results into a single object
func (c PrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrorClient) ListPrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrorsComplete(ctx context.Context, id IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId) (ListPrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrorsCompleteResult, error) {
	return c.ListPrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrorsCompleteMatchingPredicate(ctx, id, ServiceProvisioningErrorOperationPredicate{})
}

// ListPrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrorClient) ListPrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrorsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId, predicate ServiceProvisioningErrorOperationPredicate) (result ListPrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrorsCompleteResult, err error) {
	items := make([]stable.ServiceProvisioningError, 0)

	resp, err := c.ListPrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrors(ctx, id)
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

	result = ListPrivilegedAccesGroupAssignmentScheduleGroupServiceProvisioningErrorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
