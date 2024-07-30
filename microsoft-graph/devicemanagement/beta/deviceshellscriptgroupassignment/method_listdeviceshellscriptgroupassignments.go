package deviceshellscriptgroupassignment

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

type ListDeviceShellScriptGroupAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementScriptGroupAssignment
}

type ListDeviceShellScriptGroupAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementScriptGroupAssignment
}

type ListDeviceShellScriptGroupAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceShellScriptGroupAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceShellScriptGroupAssignments ...
func (c DeviceShellScriptGroupAssignmentClient) ListDeviceShellScriptGroupAssignments(ctx context.Context, id DeviceManagementDeviceShellScriptId) (result ListDeviceShellScriptGroupAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceShellScriptGroupAssignmentsCustomPager{},
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

// ListDeviceShellScriptGroupAssignmentsComplete retrieves all the results into a single object
func (c DeviceShellScriptGroupAssignmentClient) ListDeviceShellScriptGroupAssignmentsComplete(ctx context.Context, id DeviceManagementDeviceShellScriptId) (ListDeviceShellScriptGroupAssignmentsCompleteResult, error) {
	return c.ListDeviceShellScriptGroupAssignmentsCompleteMatchingPredicate(ctx, id, DeviceManagementScriptGroupAssignmentOperationPredicate{})
}

// ListDeviceShellScriptGroupAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceShellScriptGroupAssignmentClient) ListDeviceShellScriptGroupAssignmentsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDeviceShellScriptId, predicate DeviceManagementScriptGroupAssignmentOperationPredicate) (result ListDeviceShellScriptGroupAssignmentsCompleteResult, err error) {
	items := make([]beta.DeviceManagementScriptGroupAssignment, 0)

	resp, err := c.ListDeviceShellScriptGroupAssignments(ctx, id)
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

	result = ListDeviceShellScriptGroupAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
