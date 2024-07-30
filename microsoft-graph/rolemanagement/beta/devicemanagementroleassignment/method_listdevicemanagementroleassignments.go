package devicemanagementroleassignment

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

type ListDeviceManagementRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleAssignmentMultiple
}

type ListDeviceManagementRoleAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleAssignmentMultiple
}

type ListDeviceManagementRoleAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceManagementRoleAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceManagementRoleAssignments ...
func (c DeviceManagementRoleAssignmentClient) ListDeviceManagementRoleAssignments(ctx context.Context) (result ListDeviceManagementRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceManagementRoleAssignmentsCustomPager{},
		Path:       "/roleManagement/deviceManagement/roleAssignments",
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
		Values *[]beta.UnifiedRoleAssignmentMultiple `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceManagementRoleAssignmentsComplete retrieves all the results into a single object
func (c DeviceManagementRoleAssignmentClient) ListDeviceManagementRoleAssignmentsComplete(ctx context.Context) (ListDeviceManagementRoleAssignmentsCompleteResult, error) {
	return c.ListDeviceManagementRoleAssignmentsCompleteMatchingPredicate(ctx, UnifiedRoleAssignmentMultipleOperationPredicate{})
}

// ListDeviceManagementRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementRoleAssignmentClient) ListDeviceManagementRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, predicate UnifiedRoleAssignmentMultipleOperationPredicate) (result ListDeviceManagementRoleAssignmentsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleAssignmentMultiple, 0)

	resp, err := c.ListDeviceManagementRoleAssignments(ctx)
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

	result = ListDeviceManagementRoleAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
