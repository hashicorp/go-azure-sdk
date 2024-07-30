package devicecustomattributeshellscriptuserrunstatedevicerunstate

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

type ListDeviceCustomAttributeShellScriptUserRunStateDeviceRunStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementScriptDeviceState
}

type ListDeviceCustomAttributeShellScriptUserRunStateDeviceRunStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementScriptDeviceState
}

type ListDeviceCustomAttributeShellScriptUserRunStateDeviceRunStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCustomAttributeShellScriptUserRunStateDeviceRunStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCustomAttributeShellScriptUserRunStateDeviceRunStates ...
func (c DeviceCustomAttributeShellScriptUserRunStateDeviceRunStateClient) ListDeviceCustomAttributeShellScriptUserRunStateDeviceRunStates(ctx context.Context, id DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId) (result ListDeviceCustomAttributeShellScriptUserRunStateDeviceRunStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceCustomAttributeShellScriptUserRunStateDeviceRunStatesCustomPager{},
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

// ListDeviceCustomAttributeShellScriptUserRunStateDeviceRunStatesComplete retrieves all the results into a single object
func (c DeviceCustomAttributeShellScriptUserRunStateDeviceRunStateClient) ListDeviceCustomAttributeShellScriptUserRunStateDeviceRunStatesComplete(ctx context.Context, id DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId) (ListDeviceCustomAttributeShellScriptUserRunStateDeviceRunStatesCompleteResult, error) {
	return c.ListDeviceCustomAttributeShellScriptUserRunStateDeviceRunStatesCompleteMatchingPredicate(ctx, id, DeviceManagementScriptDeviceStateOperationPredicate{})
}

// ListDeviceCustomAttributeShellScriptUserRunStateDeviceRunStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCustomAttributeShellScriptUserRunStateDeviceRunStateClient) ListDeviceCustomAttributeShellScriptUserRunStateDeviceRunStatesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId, predicate DeviceManagementScriptDeviceStateOperationPredicate) (result ListDeviceCustomAttributeShellScriptUserRunStateDeviceRunStatesCompleteResult, err error) {
	items := make([]beta.DeviceManagementScriptDeviceState, 0)

	resp, err := c.ListDeviceCustomAttributeShellScriptUserRunStateDeviceRunStates(ctx, id)
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

	result = ListDeviceCustomAttributeShellScriptUserRunStateDeviceRunStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
