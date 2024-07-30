package devicecustomattributeshellscriptuserrunstate

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

type ListDeviceCustomAttributeShellScriptUserRunStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementScriptUserState
}

type ListDeviceCustomAttributeShellScriptUserRunStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementScriptUserState
}

type ListDeviceCustomAttributeShellScriptUserRunStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCustomAttributeShellScriptUserRunStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCustomAttributeShellScriptUserRunStates ...
func (c DeviceCustomAttributeShellScriptUserRunStateClient) ListDeviceCustomAttributeShellScriptUserRunStates(ctx context.Context, id DeviceManagementDeviceCustomAttributeShellScriptId) (result ListDeviceCustomAttributeShellScriptUserRunStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceCustomAttributeShellScriptUserRunStatesCustomPager{},
		Path:       fmt.Sprintf("%s/userRunStates", id.ID()),
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
		Values *[]beta.DeviceManagementScriptUserState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCustomAttributeShellScriptUserRunStatesComplete retrieves all the results into a single object
func (c DeviceCustomAttributeShellScriptUserRunStateClient) ListDeviceCustomAttributeShellScriptUserRunStatesComplete(ctx context.Context, id DeviceManagementDeviceCustomAttributeShellScriptId) (ListDeviceCustomAttributeShellScriptUserRunStatesCompleteResult, error) {
	return c.ListDeviceCustomAttributeShellScriptUserRunStatesCompleteMatchingPredicate(ctx, id, DeviceManagementScriptUserStateOperationPredicate{})
}

// ListDeviceCustomAttributeShellScriptUserRunStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCustomAttributeShellScriptUserRunStateClient) ListDeviceCustomAttributeShellScriptUserRunStatesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDeviceCustomAttributeShellScriptId, predicate DeviceManagementScriptUserStateOperationPredicate) (result ListDeviceCustomAttributeShellScriptUserRunStatesCompleteResult, err error) {
	items := make([]beta.DeviceManagementScriptUserState, 0)

	resp, err := c.ListDeviceCustomAttributeShellScriptUserRunStates(ctx, id)
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

	result = ListDeviceCustomAttributeShellScriptUserRunStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
