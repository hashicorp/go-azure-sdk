package devicemanagementscriptgroupassignment

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

type ListDeviceManagementScriptGroupAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementScriptGroupAssignment
}

type ListDeviceManagementScriptGroupAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementScriptGroupAssignment
}

type ListDeviceManagementScriptGroupAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceManagementScriptGroupAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceManagementScriptGroupAssignments ...
func (c DeviceManagementScriptGroupAssignmentClient) ListDeviceManagementScriptGroupAssignments(ctx context.Context, id DeviceManagementDeviceManagementScriptId) (result ListDeviceManagementScriptGroupAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceManagementScriptGroupAssignmentsCustomPager{},
		Path:       fmt.Sprintf("%s/groupAssignments", id.ID()),
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
		Values *[]beta.DeviceManagementScriptGroupAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceManagementScriptGroupAssignmentsComplete retrieves all the results into a single object
func (c DeviceManagementScriptGroupAssignmentClient) ListDeviceManagementScriptGroupAssignmentsComplete(ctx context.Context, id DeviceManagementDeviceManagementScriptId) (ListDeviceManagementScriptGroupAssignmentsCompleteResult, error) {
	return c.ListDeviceManagementScriptGroupAssignmentsCompleteMatchingPredicate(ctx, id, DeviceManagementScriptGroupAssignmentOperationPredicate{})
}

// ListDeviceManagementScriptGroupAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementScriptGroupAssignmentClient) ListDeviceManagementScriptGroupAssignmentsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDeviceManagementScriptId, predicate DeviceManagementScriptGroupAssignmentOperationPredicate) (result ListDeviceManagementScriptGroupAssignmentsCompleteResult, err error) {
	items := make([]beta.DeviceManagementScriptGroupAssignment, 0)

	resp, err := c.ListDeviceManagementScriptGroupAssignments(ctx, id)
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

	result = ListDeviceManagementScriptGroupAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
