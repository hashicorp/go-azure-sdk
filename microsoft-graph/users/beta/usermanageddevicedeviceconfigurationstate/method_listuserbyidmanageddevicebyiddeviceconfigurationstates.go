package usermanageddevicedeviceconfigurationstate

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

type ListUserByIdManagedDeviceByIdDeviceConfigurationStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DeviceConfigurationStateCollectionResponse
}

type ListUserByIdManagedDeviceByIdDeviceConfigurationStatesCompleteResult struct {
	Items []models.DeviceConfigurationStateCollectionResponse
}

// ListUserByIdManagedDeviceByIdDeviceConfigurationStates ...
func (c UserManagedDeviceDeviceConfigurationStateClient) ListUserByIdManagedDeviceByIdDeviceConfigurationStates(ctx context.Context, id UserManagedDeviceId) (result ListUserByIdManagedDeviceByIdDeviceConfigurationStatesOperationResponse, err error) {
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

// ListUserByIdManagedDeviceByIdDeviceConfigurationStatesComplete retrieves all the results into a single object
func (c UserManagedDeviceDeviceConfigurationStateClient) ListUserByIdManagedDeviceByIdDeviceConfigurationStatesComplete(ctx context.Context, id UserManagedDeviceId) (ListUserByIdManagedDeviceByIdDeviceConfigurationStatesCompleteResult, error) {
	return c.ListUserByIdManagedDeviceByIdDeviceConfigurationStatesCompleteMatchingPredicate(ctx, id, models.DeviceConfigurationStateCollectionResponseOperationPredicate{})
}

// ListUserByIdManagedDeviceByIdDeviceConfigurationStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserManagedDeviceDeviceConfigurationStateClient) ListUserByIdManagedDeviceByIdDeviceConfigurationStatesCompleteMatchingPredicate(ctx context.Context, id UserManagedDeviceId, predicate models.DeviceConfigurationStateCollectionResponseOperationPredicate) (result ListUserByIdManagedDeviceByIdDeviceConfigurationStatesCompleteResult, err error) {
	items := make([]models.DeviceConfigurationStateCollectionResponse, 0)

	resp, err := c.ListUserByIdManagedDeviceByIdDeviceConfigurationStates(ctx, id)
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

	result = ListUserByIdManagedDeviceByIdDeviceConfigurationStatesCompleteResult{
		Items: items,
	}
	return
}
