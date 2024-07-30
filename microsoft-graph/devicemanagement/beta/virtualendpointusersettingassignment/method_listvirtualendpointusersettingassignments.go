package virtualendpointusersettingassignment

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

type ListVirtualEndpointUserSettingAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCUserSettingAssignment
}

type ListVirtualEndpointUserSettingAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCUserSettingAssignment
}

type ListVirtualEndpointUserSettingAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointUserSettingAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointUserSettingAssignments ...
func (c VirtualEndpointUserSettingAssignmentClient) ListVirtualEndpointUserSettingAssignments(ctx context.Context, id DeviceManagementVirtualEndpointUserSettingId) (result ListVirtualEndpointUserSettingAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListVirtualEndpointUserSettingAssignmentsCustomPager{},
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
		Values *[]beta.CloudPCUserSettingAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointUserSettingAssignmentsComplete retrieves all the results into a single object
func (c VirtualEndpointUserSettingAssignmentClient) ListVirtualEndpointUserSettingAssignmentsComplete(ctx context.Context, id DeviceManagementVirtualEndpointUserSettingId) (ListVirtualEndpointUserSettingAssignmentsCompleteResult, error) {
	return c.ListVirtualEndpointUserSettingAssignmentsCompleteMatchingPredicate(ctx, id, CloudPCUserSettingAssignmentOperationPredicate{})
}

// ListVirtualEndpointUserSettingAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointUserSettingAssignmentClient) ListVirtualEndpointUserSettingAssignmentsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementVirtualEndpointUserSettingId, predicate CloudPCUserSettingAssignmentOperationPredicate) (result ListVirtualEndpointUserSettingAssignmentsCompleteResult, err error) {
	items := make([]beta.CloudPCUserSettingAssignment, 0)

	resp, err := c.ListVirtualEndpointUserSettingAssignments(ctx, id)
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

	result = ListVirtualEndpointUserSettingAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
