package memanageddevicedevicehealthscriptstate

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMeManagedDeviceByIdDeviceHealthScriptStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DeviceHealthScriptPolicyStateCollectionResponse
}

type ListMeManagedDeviceByIdDeviceHealthScriptStatesCompleteResult struct {
	Items []models.DeviceHealthScriptPolicyStateCollectionResponse
}

// ListMeManagedDeviceByIdDeviceHealthScriptStates ...
func (c MeManagedDeviceDeviceHealthScriptStateClient) ListMeManagedDeviceByIdDeviceHealthScriptStates(ctx context.Context, id MeManagedDeviceId) (result ListMeManagedDeviceByIdDeviceHealthScriptStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.DeviceHealthScriptPolicyStateCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeManagedDeviceByIdDeviceHealthScriptStatesComplete retrieves all the results into a single object
func (c MeManagedDeviceDeviceHealthScriptStateClient) ListMeManagedDeviceByIdDeviceHealthScriptStatesComplete(ctx context.Context, id MeManagedDeviceId) (ListMeManagedDeviceByIdDeviceHealthScriptStatesCompleteResult, error) {
	return c.ListMeManagedDeviceByIdDeviceHealthScriptStatesCompleteMatchingPredicate(ctx, id, models.DeviceHealthScriptPolicyStateCollectionResponseOperationPredicate{})
}

// ListMeManagedDeviceByIdDeviceHealthScriptStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeManagedDeviceDeviceHealthScriptStateClient) ListMeManagedDeviceByIdDeviceHealthScriptStatesCompleteMatchingPredicate(ctx context.Context, id MeManagedDeviceId, predicate models.DeviceHealthScriptPolicyStateCollectionResponseOperationPredicate) (result ListMeManagedDeviceByIdDeviceHealthScriptStatesCompleteResult, err error) {
	items := make([]models.DeviceHealthScriptPolicyStateCollectionResponse, 0)

	resp, err := c.ListMeManagedDeviceByIdDeviceHealthScriptStates(ctx, id)
	if err != nil {
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

	result = ListMeManagedDeviceByIdDeviceHealthScriptStatesCompleteResult{
		Items: items,
	}
	return
}
