package usermanageddevicedevicecompliancepolicystate

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdManagedDeviceByIdDeviceCompliancePolicyStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DeviceCompliancePolicyStateCollectionResponse
}

type ListUserByIdManagedDeviceByIdDeviceCompliancePolicyStatesCompleteResult struct {
	Items []models.DeviceCompliancePolicyStateCollectionResponse
}

// ListUserByIdManagedDeviceByIdDeviceCompliancePolicyStates ...
func (c UserManagedDeviceDeviceCompliancePolicyStateClient) ListUserByIdManagedDeviceByIdDeviceCompliancePolicyStates(ctx context.Context, id UserManagedDeviceId) (result ListUserByIdManagedDeviceByIdDeviceCompliancePolicyStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/deviceCompliancePolicyStates", id.ID()),
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
		Values *[]models.DeviceCompliancePolicyStateCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdManagedDeviceByIdDeviceCompliancePolicyStatesComplete retrieves all the results into a single object
func (c UserManagedDeviceDeviceCompliancePolicyStateClient) ListUserByIdManagedDeviceByIdDeviceCompliancePolicyStatesComplete(ctx context.Context, id UserManagedDeviceId) (ListUserByIdManagedDeviceByIdDeviceCompliancePolicyStatesCompleteResult, error) {
	return c.ListUserByIdManagedDeviceByIdDeviceCompliancePolicyStatesCompleteMatchingPredicate(ctx, id, models.DeviceCompliancePolicyStateCollectionResponseOperationPredicate{})
}

// ListUserByIdManagedDeviceByIdDeviceCompliancePolicyStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserManagedDeviceDeviceCompliancePolicyStateClient) ListUserByIdManagedDeviceByIdDeviceCompliancePolicyStatesCompleteMatchingPredicate(ctx context.Context, id UserManagedDeviceId, predicate models.DeviceCompliancePolicyStateCollectionResponseOperationPredicate) (result ListUserByIdManagedDeviceByIdDeviceCompliancePolicyStatesCompleteResult, err error) {
	items := make([]models.DeviceCompliancePolicyStateCollectionResponse, 0)

	resp, err := c.ListUserByIdManagedDeviceByIdDeviceCompliancePolicyStates(ctx, id)
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

	result = ListUserByIdManagedDeviceByIdDeviceCompliancePolicyStatesCompleteResult{
		Items: items,
	}
	return
}
