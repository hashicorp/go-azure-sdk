package virtualendpointprovisioningpolicyassignment

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

type ListVirtualEndpointProvisioningPolicyAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCProvisioningPolicyAssignment
}

type ListVirtualEndpointProvisioningPolicyAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCProvisioningPolicyAssignment
}

type ListVirtualEndpointProvisioningPolicyAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointProvisioningPolicyAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointProvisioningPolicyAssignments ...
func (c VirtualEndpointProvisioningPolicyAssignmentClient) ListVirtualEndpointProvisioningPolicyAssignments(ctx context.Context, id DeviceManagementVirtualEndpointProvisioningPolicyId) (result ListVirtualEndpointProvisioningPolicyAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListVirtualEndpointProvisioningPolicyAssignmentsCustomPager{},
		Path:       fmt.Sprintf("%s/assignments", id.ID()),
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
		Values *[]beta.CloudPCProvisioningPolicyAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointProvisioningPolicyAssignmentsComplete retrieves all the results into a single object
func (c VirtualEndpointProvisioningPolicyAssignmentClient) ListVirtualEndpointProvisioningPolicyAssignmentsComplete(ctx context.Context, id DeviceManagementVirtualEndpointProvisioningPolicyId) (ListVirtualEndpointProvisioningPolicyAssignmentsCompleteResult, error) {
	return c.ListVirtualEndpointProvisioningPolicyAssignmentsCompleteMatchingPredicate(ctx, id, CloudPCProvisioningPolicyAssignmentOperationPredicate{})
}

// ListVirtualEndpointProvisioningPolicyAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointProvisioningPolicyAssignmentClient) ListVirtualEndpointProvisioningPolicyAssignmentsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementVirtualEndpointProvisioningPolicyId, predicate CloudPCProvisioningPolicyAssignmentOperationPredicate) (result ListVirtualEndpointProvisioningPolicyAssignmentsCompleteResult, err error) {
	items := make([]beta.CloudPCProvisioningPolicyAssignment, 0)

	resp, err := c.ListVirtualEndpointProvisioningPolicyAssignments(ctx, id)
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

	result = ListVirtualEndpointProvisioningPolicyAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
