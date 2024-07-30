package virtualendpointprovisioningpolicyassignmentassigneduser

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

type ListVirtualEndpointProvisioningPolicyAssignmentAssignedUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.User
}

type ListVirtualEndpointProvisioningPolicyAssignmentAssignedUsersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.User
}

type ListVirtualEndpointProvisioningPolicyAssignmentAssignedUsersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointProvisioningPolicyAssignmentAssignedUsersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointProvisioningPolicyAssignmentAssignedUsers ...
func (c VirtualEndpointProvisioningPolicyAssignmentAssignedUserClient) ListVirtualEndpointProvisioningPolicyAssignmentAssignedUsers(ctx context.Context, id DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId) (result ListVirtualEndpointProvisioningPolicyAssignmentAssignedUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListVirtualEndpointProvisioningPolicyAssignmentAssignedUsersCustomPager{},
		Path:       fmt.Sprintf("%s/assignedUsers", id.ID()),
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
		Values *[]beta.User `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointProvisioningPolicyAssignmentAssignedUsersComplete retrieves all the results into a single object
func (c VirtualEndpointProvisioningPolicyAssignmentAssignedUserClient) ListVirtualEndpointProvisioningPolicyAssignmentAssignedUsersComplete(ctx context.Context, id DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId) (ListVirtualEndpointProvisioningPolicyAssignmentAssignedUsersCompleteResult, error) {
	return c.ListVirtualEndpointProvisioningPolicyAssignmentAssignedUsersCompleteMatchingPredicate(ctx, id, UserOperationPredicate{})
}

// ListVirtualEndpointProvisioningPolicyAssignmentAssignedUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointProvisioningPolicyAssignmentAssignedUserClient) ListVirtualEndpointProvisioningPolicyAssignmentAssignedUsersCompleteMatchingPredicate(ctx context.Context, id DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId, predicate UserOperationPredicate) (result ListVirtualEndpointProvisioningPolicyAssignmentAssignedUsersCompleteResult, err error) {
	items := make([]beta.User, 0)

	resp, err := c.ListVirtualEndpointProvisioningPolicyAssignmentAssignedUsers(ctx, id)
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

	result = ListVirtualEndpointProvisioningPolicyAssignmentAssignedUsersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
