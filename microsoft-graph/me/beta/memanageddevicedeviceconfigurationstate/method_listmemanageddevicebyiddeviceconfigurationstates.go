package memanageddevicedeviceconfigurationstate

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

type ListMeManagedDeviceByIdDeviceConfigurationStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DeviceConfigurationStateCollectionResponse
}

type ListMeManagedDeviceByIdDeviceConfigurationStatesCompleteResult struct {
	Items []models.DeviceConfigurationStateCollectionResponse
}

// ListMeManagedDeviceByIdDeviceConfigurationStates ...
func (c MeManagedDeviceDeviceConfigurationStateClient) ListMeManagedDeviceByIdDeviceConfigurationStates(ctx context.Context, id MeManagedDeviceId) (result ListMeManagedDeviceByIdDeviceConfigurationStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/deviceConfigurationStates", id.ID()),
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
		Values *[]models.DeviceConfigurationStateCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeManagedDeviceByIdDeviceConfigurationStatesComplete retrieves all the results into a single object
func (c MeManagedDeviceDeviceConfigurationStateClient) ListMeManagedDeviceByIdDeviceConfigurationStatesComplete(ctx context.Context, id MeManagedDeviceId) (ListMeManagedDeviceByIdDeviceConfigurationStatesCompleteResult, error) {
	return c.ListMeManagedDeviceByIdDeviceConfigurationStatesCompleteMatchingPredicate(ctx, id, models.DeviceConfigurationStateCollectionResponseOperationPredicate{})
}

// ListMeManagedDeviceByIdDeviceConfigurationStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeManagedDeviceDeviceConfigurationStateClient) ListMeManagedDeviceByIdDeviceConfigurationStatesCompleteMatchingPredicate(ctx context.Context, id MeManagedDeviceId, predicate models.DeviceConfigurationStateCollectionResponseOperationPredicate) (result ListMeManagedDeviceByIdDeviceConfigurationStatesCompleteResult, err error) {
	items := make([]models.DeviceConfigurationStateCollectionResponse, 0)

	resp, err := c.ListMeManagedDeviceByIdDeviceConfigurationStates(ctx, id)
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

	result = ListMeManagedDeviceByIdDeviceConfigurationStatesCompleteResult{
		Items: items,
	}
	return
}
