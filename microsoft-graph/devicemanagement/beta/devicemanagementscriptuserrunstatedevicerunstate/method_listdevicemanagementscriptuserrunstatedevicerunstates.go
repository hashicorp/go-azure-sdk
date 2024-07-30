package devicemanagementscriptuserrunstatedevicerunstate

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

type ListDeviceManagementScriptUserRunStateDeviceRunStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementScriptDeviceState
}

type ListDeviceManagementScriptUserRunStateDeviceRunStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementScriptDeviceState
}

type ListDeviceManagementScriptUserRunStateDeviceRunStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceManagementScriptUserRunStateDeviceRunStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceManagementScriptUserRunStateDeviceRunStates ...
func (c DeviceManagementScriptUserRunStateDeviceRunStateClient) ListDeviceManagementScriptUserRunStateDeviceRunStates(ctx context.Context, id DeviceManagementDeviceManagementScriptIdUserRunStateId) (result ListDeviceManagementScriptUserRunStateDeviceRunStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceManagementScriptUserRunStateDeviceRunStatesCustomPager{},
		Path:       fmt.Sprintf("%s/deviceRunStates", id.ID()),
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
		Values *[]beta.DeviceManagementScriptDeviceState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceManagementScriptUserRunStateDeviceRunStatesComplete retrieves all the results into a single object
func (c DeviceManagementScriptUserRunStateDeviceRunStateClient) ListDeviceManagementScriptUserRunStateDeviceRunStatesComplete(ctx context.Context, id DeviceManagementDeviceManagementScriptIdUserRunStateId) (ListDeviceManagementScriptUserRunStateDeviceRunStatesCompleteResult, error) {
	return c.ListDeviceManagementScriptUserRunStateDeviceRunStatesCompleteMatchingPredicate(ctx, id, DeviceManagementScriptDeviceStateOperationPredicate{})
}

// ListDeviceManagementScriptUserRunStateDeviceRunStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementScriptUserRunStateDeviceRunStateClient) ListDeviceManagementScriptUserRunStateDeviceRunStatesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDeviceManagementScriptIdUserRunStateId, predicate DeviceManagementScriptDeviceStateOperationPredicate) (result ListDeviceManagementScriptUserRunStateDeviceRunStatesCompleteResult, err error) {
	items := make([]beta.DeviceManagementScriptDeviceState, 0)

	resp, err := c.ListDeviceManagementScriptUserRunStateDeviceRunStates(ctx, id)
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

	result = ListDeviceManagementScriptUserRunStateDeviceRunStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
