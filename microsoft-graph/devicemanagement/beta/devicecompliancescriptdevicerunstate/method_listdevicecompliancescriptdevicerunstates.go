package devicecompliancescriptdevicerunstate

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

type ListDeviceComplianceScriptDeviceRunStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceComplianceScriptDeviceState
}

type ListDeviceComplianceScriptDeviceRunStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceComplianceScriptDeviceState
}

type ListDeviceComplianceScriptDeviceRunStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceComplianceScriptDeviceRunStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceComplianceScriptDeviceRunStates ...
func (c DeviceComplianceScriptDeviceRunStateClient) ListDeviceComplianceScriptDeviceRunStates(ctx context.Context, id DeviceManagementDeviceComplianceScriptId) (result ListDeviceComplianceScriptDeviceRunStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceComplianceScriptDeviceRunStatesCustomPager{},
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
		Values *[]beta.DeviceComplianceScriptDeviceState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceComplianceScriptDeviceRunStatesComplete retrieves all the results into a single object
func (c DeviceComplianceScriptDeviceRunStateClient) ListDeviceComplianceScriptDeviceRunStatesComplete(ctx context.Context, id DeviceManagementDeviceComplianceScriptId) (ListDeviceComplianceScriptDeviceRunStatesCompleteResult, error) {
	return c.ListDeviceComplianceScriptDeviceRunStatesCompleteMatchingPredicate(ctx, id, DeviceComplianceScriptDeviceStateOperationPredicate{})
}

// ListDeviceComplianceScriptDeviceRunStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceComplianceScriptDeviceRunStateClient) ListDeviceComplianceScriptDeviceRunStatesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDeviceComplianceScriptId, predicate DeviceComplianceScriptDeviceStateOperationPredicate) (result ListDeviceComplianceScriptDeviceRunStatesCompleteResult, err error) {
	items := make([]beta.DeviceComplianceScriptDeviceState, 0)

	resp, err := c.ListDeviceComplianceScriptDeviceRunStates(ctx, id)
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

	result = ListDeviceComplianceScriptDeviceRunStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
