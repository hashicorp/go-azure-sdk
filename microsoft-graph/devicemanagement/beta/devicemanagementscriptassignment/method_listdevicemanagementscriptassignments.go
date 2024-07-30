package devicemanagementscriptassignment

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

type ListDeviceManagementScriptAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementScriptAssignment
}

type ListDeviceManagementScriptAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementScriptAssignment
}

type ListDeviceManagementScriptAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceManagementScriptAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceManagementScriptAssignments ...
func (c DeviceManagementScriptAssignmentClient) ListDeviceManagementScriptAssignments(ctx context.Context, id DeviceManagementDeviceManagementScriptId) (result ListDeviceManagementScriptAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceManagementScriptAssignmentsCustomPager{},
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
		Values *[]beta.DeviceManagementScriptAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceManagementScriptAssignmentsComplete retrieves all the results into a single object
func (c DeviceManagementScriptAssignmentClient) ListDeviceManagementScriptAssignmentsComplete(ctx context.Context, id DeviceManagementDeviceManagementScriptId) (ListDeviceManagementScriptAssignmentsCompleteResult, error) {
	return c.ListDeviceManagementScriptAssignmentsCompleteMatchingPredicate(ctx, id, DeviceManagementScriptAssignmentOperationPredicate{})
}

// ListDeviceManagementScriptAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementScriptAssignmentClient) ListDeviceManagementScriptAssignmentsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDeviceManagementScriptId, predicate DeviceManagementScriptAssignmentOperationPredicate) (result ListDeviceManagementScriptAssignmentsCompleteResult, err error) {
	items := make([]beta.DeviceManagementScriptAssignment, 0)

	resp, err := c.ListDeviceManagementScriptAssignments(ctx, id)
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

	result = ListDeviceManagementScriptAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
