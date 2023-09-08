package usermanageddevicedevicehealthscriptstate

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

type ListUserByIdManagedDeviceByIdDeviceHealthScriptStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DeviceHealthScriptPolicyStateCollectionResponse
}

type ListUserByIdManagedDeviceByIdDeviceHealthScriptStatesCompleteResult struct {
	Items []models.DeviceHealthScriptPolicyStateCollectionResponse
}

// ListUserByIdManagedDeviceByIdDeviceHealthScriptStates ...
func (c UserManagedDeviceDeviceHealthScriptStateClient) ListUserByIdManagedDeviceByIdDeviceHealthScriptStates(ctx context.Context, id UserManagedDeviceId) (result ListUserByIdManagedDeviceByIdDeviceHealthScriptStatesOperationResponse, err error) {
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

// ListUserByIdManagedDeviceByIdDeviceHealthScriptStatesComplete retrieves all the results into a single object
func (c UserManagedDeviceDeviceHealthScriptStateClient) ListUserByIdManagedDeviceByIdDeviceHealthScriptStatesComplete(ctx context.Context, id UserManagedDeviceId) (ListUserByIdManagedDeviceByIdDeviceHealthScriptStatesCompleteResult, error) {
	return c.ListUserByIdManagedDeviceByIdDeviceHealthScriptStatesCompleteMatchingPredicate(ctx, id, models.DeviceHealthScriptPolicyStateCollectionResponseOperationPredicate{})
}

// ListUserByIdManagedDeviceByIdDeviceHealthScriptStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserManagedDeviceDeviceHealthScriptStateClient) ListUserByIdManagedDeviceByIdDeviceHealthScriptStatesCompleteMatchingPredicate(ctx context.Context, id UserManagedDeviceId, predicate models.DeviceHealthScriptPolicyStateCollectionResponseOperationPredicate) (result ListUserByIdManagedDeviceByIdDeviceHealthScriptStatesCompleteResult, err error) {
	items := make([]models.DeviceHealthScriptPolicyStateCollectionResponse, 0)

	resp, err := c.ListUserByIdManagedDeviceByIdDeviceHealthScriptStates(ctx, id)
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

	result = ListUserByIdManagedDeviceByIdDeviceHealthScriptStatesCompleteResult{
		Items: items,
	}
	return
}
