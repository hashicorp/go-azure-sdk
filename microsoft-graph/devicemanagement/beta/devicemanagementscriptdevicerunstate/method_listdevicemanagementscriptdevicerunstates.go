package devicemanagementscriptdevicerunstate

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

type ListDeviceManagementScriptDeviceRunStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementScriptDeviceState
}

type ListDeviceManagementScriptDeviceRunStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementScriptDeviceState
}

type ListDeviceManagementScriptDeviceRunStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceManagementScriptDeviceRunStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceManagementScriptDeviceRunStates ...
func (c DeviceManagementScriptDeviceRunStateClient) ListDeviceManagementScriptDeviceRunStates(ctx context.Context, id DeviceManagementDeviceManagementScriptId) (result ListDeviceManagementScriptDeviceRunStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceManagementScriptDeviceRunStatesCustomPager{},
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

// ListDeviceManagementScriptDeviceRunStatesComplete retrieves all the results into a single object
func (c DeviceManagementScriptDeviceRunStateClient) ListDeviceManagementScriptDeviceRunStatesComplete(ctx context.Context, id DeviceManagementDeviceManagementScriptId) (ListDeviceManagementScriptDeviceRunStatesCompleteResult, error) {
	return c.ListDeviceManagementScriptDeviceRunStatesCompleteMatchingPredicate(ctx, id, DeviceManagementScriptDeviceStateOperationPredicate{})
}

// ListDeviceManagementScriptDeviceRunStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementScriptDeviceRunStateClient) ListDeviceManagementScriptDeviceRunStatesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDeviceManagementScriptId, predicate DeviceManagementScriptDeviceStateOperationPredicate) (result ListDeviceManagementScriptDeviceRunStatesCompleteResult, err error) {
	items := make([]beta.DeviceManagementScriptDeviceState, 0)

	resp, err := c.ListDeviceManagementScriptDeviceRunStates(ctx, id)
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

	result = ListDeviceManagementScriptDeviceRunStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
