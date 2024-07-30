package manageddevicedevicehealthscriptstate

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

type ListManagedDeviceDeviceHealthScriptStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceHealthScriptPolicyState
}

type ListManagedDeviceDeviceHealthScriptStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceHealthScriptPolicyState
}

type ListManagedDeviceDeviceHealthScriptStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListManagedDeviceDeviceHealthScriptStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListManagedDeviceDeviceHealthScriptStates ...
func (c ManagedDeviceDeviceHealthScriptStateClient) ListManagedDeviceDeviceHealthScriptStates(ctx context.Context, id DeviceManagementManagedDeviceId) (result ListManagedDeviceDeviceHealthScriptStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListManagedDeviceDeviceHealthScriptStatesCustomPager{},
		Path:       fmt.Sprintf("%s/deviceHealthScriptStates", id.ID()),
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
		Values *[]beta.DeviceHealthScriptPolicyState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListManagedDeviceDeviceHealthScriptStatesComplete retrieves all the results into a single object
func (c ManagedDeviceDeviceHealthScriptStateClient) ListManagedDeviceDeviceHealthScriptStatesComplete(ctx context.Context, id DeviceManagementManagedDeviceId) (ListManagedDeviceDeviceHealthScriptStatesCompleteResult, error) {
	return c.ListManagedDeviceDeviceHealthScriptStatesCompleteMatchingPredicate(ctx, id, DeviceHealthScriptPolicyStateOperationPredicate{})
}

// ListManagedDeviceDeviceHealthScriptStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDeviceDeviceHealthScriptStateClient) ListManagedDeviceDeviceHealthScriptStatesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementManagedDeviceId, predicate DeviceHealthScriptPolicyStateOperationPredicate) (result ListManagedDeviceDeviceHealthScriptStatesCompleteResult, err error) {
	items := make([]beta.DeviceHealthScriptPolicyState, 0)

	resp, err := c.ListManagedDeviceDeviceHealthScriptStates(ctx, id)
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

	result = ListManagedDeviceDeviceHealthScriptStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
