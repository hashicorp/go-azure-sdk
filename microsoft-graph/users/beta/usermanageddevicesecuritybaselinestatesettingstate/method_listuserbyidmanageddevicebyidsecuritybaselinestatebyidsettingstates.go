package usermanageddevicesecuritybaselinestatesettingstate

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

type ListUserByIdManagedDeviceByIdSecurityBaselineStateByIdSettingStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SecurityBaselineSettingStateCollectionResponse
}

type ListUserByIdManagedDeviceByIdSecurityBaselineStateByIdSettingStatesCompleteResult struct {
	Items []models.SecurityBaselineSettingStateCollectionResponse
}

// ListUserByIdManagedDeviceByIdSecurityBaselineStateByIdSettingStates ...
func (c UserManagedDeviceSecurityBaselineStateSettingStateClient) ListUserByIdManagedDeviceByIdSecurityBaselineStateByIdSettingStates(ctx context.Context, id UserManagedDeviceSecurityBaselineStateId) (result ListUserByIdManagedDeviceByIdSecurityBaselineStateByIdSettingStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/settingStates", id.ID()),
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
		Values *[]models.SecurityBaselineSettingStateCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdManagedDeviceByIdSecurityBaselineStateByIdSettingStatesComplete retrieves all the results into a single object
func (c UserManagedDeviceSecurityBaselineStateSettingStateClient) ListUserByIdManagedDeviceByIdSecurityBaselineStateByIdSettingStatesComplete(ctx context.Context, id UserManagedDeviceSecurityBaselineStateId) (ListUserByIdManagedDeviceByIdSecurityBaselineStateByIdSettingStatesCompleteResult, error) {
	return c.ListUserByIdManagedDeviceByIdSecurityBaselineStateByIdSettingStatesCompleteMatchingPredicate(ctx, id, models.SecurityBaselineSettingStateCollectionResponseOperationPredicate{})
}

// ListUserByIdManagedDeviceByIdSecurityBaselineStateByIdSettingStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserManagedDeviceSecurityBaselineStateSettingStateClient) ListUserByIdManagedDeviceByIdSecurityBaselineStateByIdSettingStatesCompleteMatchingPredicate(ctx context.Context, id UserManagedDeviceSecurityBaselineStateId, predicate models.SecurityBaselineSettingStateCollectionResponseOperationPredicate) (result ListUserByIdManagedDeviceByIdSecurityBaselineStateByIdSettingStatesCompleteResult, err error) {
	items := make([]models.SecurityBaselineSettingStateCollectionResponse, 0)

	resp, err := c.ListUserByIdManagedDeviceByIdSecurityBaselineStateByIdSettingStates(ctx, id)
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

	result = ListUserByIdManagedDeviceByIdSecurityBaselineStateByIdSettingStatesCompleteResult{
		Items: items,
	}
	return
}
