package memanageddevicemanageddevicemobileappconfigurationstate

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

type ListMeManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ManagedDeviceMobileAppConfigurationStateCollectionResponse
}

type ListMeManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesCompleteResult struct {
	Items []models.ManagedDeviceMobileAppConfigurationStateCollectionResponse
}

// ListMeManagedDeviceByIdManagedDeviceMobileAppConfigurationStates ...
func (c MeManagedDeviceManagedDeviceMobileAppConfigurationStateClient) ListMeManagedDeviceByIdManagedDeviceMobileAppConfigurationStates(ctx context.Context, id MeManagedDeviceId) (result ListMeManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesOperationResponse, err error) {
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

// ListMeManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesComplete retrieves all the results into a single object
func (c MeManagedDeviceManagedDeviceMobileAppConfigurationStateClient) ListMeManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesComplete(ctx context.Context, id MeManagedDeviceId) (ListMeManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesCompleteResult, error) {
	return c.ListMeManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesCompleteMatchingPredicate(ctx, id, models.ManagedDeviceMobileAppConfigurationStateCollectionResponseOperationPredicate{})
}

// ListMeManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeManagedDeviceManagedDeviceMobileAppConfigurationStateClient) ListMeManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesCompleteMatchingPredicate(ctx context.Context, id MeManagedDeviceId, predicate models.ManagedDeviceMobileAppConfigurationStateCollectionResponseOperationPredicate) (result ListMeManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesCompleteResult, err error) {
	items := make([]models.ManagedDeviceMobileAppConfigurationStateCollectionResponse, 0)

	resp, err := c.ListMeManagedDeviceByIdManagedDeviceMobileAppConfigurationStates(ctx, id)
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

	result = ListMeManagedDeviceByIdManagedDeviceMobileAppConfigurationStatesCompleteResult{
		Items: items,
	}
	return
}
