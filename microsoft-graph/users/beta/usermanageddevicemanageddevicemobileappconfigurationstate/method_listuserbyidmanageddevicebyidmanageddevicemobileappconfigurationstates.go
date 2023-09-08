package usermanageddevicemanageddevicemobileappconfigurationstate

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

type ListUserByIdManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ManagedDeviceMobileAppConfigurationStateCollectionResponse
}

type ListUserByIdManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesCompleteResult struct {
	Items []models.ManagedDeviceMobileAppConfigurationStateCollectionResponse
}

// ListUserByIdManagedDeviceByIdManagedDeviceMobileAppConfigurationStates ...
func (c UserManagedDeviceManagedDeviceMobileAppConfigurationStateClient) ListUserByIdManagedDeviceByIdManagedDeviceMobileAppConfigurationStates(ctx context.Context, id UserManagedDeviceId) (result ListUserByIdManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/managedDeviceMobileAppConfigurationStates", id.ID()),
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
		Values *[]models.ManagedDeviceMobileAppConfigurationStateCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesComplete retrieves all the results into a single object
func (c UserManagedDeviceManagedDeviceMobileAppConfigurationStateClient) ListUserByIdManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesComplete(ctx context.Context, id UserManagedDeviceId) (ListUserByIdManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesCompleteResult, error) {
	return c.ListUserByIdManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesCompleteMatchingPredicate(ctx, id, models.ManagedDeviceMobileAppConfigurationStateCollectionResponseOperationPredicate{})
}

// ListUserByIdManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserManagedDeviceManagedDeviceMobileAppConfigurationStateClient) ListUserByIdManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesCompleteMatchingPredicate(ctx context.Context, id UserManagedDeviceId, predicate models.ManagedDeviceMobileAppConfigurationStateCollectionResponseOperationPredicate) (result ListUserByIdManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesCompleteResult, err error) {
	items := make([]models.ManagedDeviceMobileAppConfigurationStateCollectionResponse, 0)

	resp, err := c.ListUserByIdManagedDeviceByIdManagedDeviceMobileAppConfigurationStates(ctx, id)
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

	result = ListUserByIdManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesCompleteResult{
		Items: items,
	}
	return
}
