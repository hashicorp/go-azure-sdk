package usermanageddevicesecuritybaselinestate

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

type ListUserByIdManagedDeviceByIdSecurityBaselineStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SecurityBaselineStateCollectionResponse
}

type ListUserByIdManagedDeviceByIdSecurityBaselineStatesCompleteResult struct {
	Items []models.SecurityBaselineStateCollectionResponse
}

// ListUserByIdManagedDeviceByIdSecurityBaselineStates ...
func (c UserManagedDeviceSecurityBaselineStateClient) ListUserByIdManagedDeviceByIdSecurityBaselineStates(ctx context.Context, id UserManagedDeviceId) (result ListUserByIdManagedDeviceByIdSecurityBaselineStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/securityBaselineStates", id.ID()),
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
		Values *[]models.SecurityBaselineStateCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdManagedDeviceByIdSecurityBaselineStatesComplete retrieves all the results into a single object
func (c UserManagedDeviceSecurityBaselineStateClient) ListUserByIdManagedDeviceByIdSecurityBaselineStatesComplete(ctx context.Context, id UserManagedDeviceId) (ListUserByIdManagedDeviceByIdSecurityBaselineStatesCompleteResult, error) {
	return c.ListUserByIdManagedDeviceByIdSecurityBaselineStatesCompleteMatchingPredicate(ctx, id, models.SecurityBaselineStateCollectionResponseOperationPredicate{})
}

// ListUserByIdManagedDeviceByIdSecurityBaselineStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserManagedDeviceSecurityBaselineStateClient) ListUserByIdManagedDeviceByIdSecurityBaselineStatesCompleteMatchingPredicate(ctx context.Context, id UserManagedDeviceId, predicate models.SecurityBaselineStateCollectionResponseOperationPredicate) (result ListUserByIdManagedDeviceByIdSecurityBaselineStatesCompleteResult, err error) {
	items := make([]models.SecurityBaselineStateCollectionResponse, 0)

	resp, err := c.ListUserByIdManagedDeviceByIdSecurityBaselineStates(ctx, id)
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

	result = ListUserByIdManagedDeviceByIdSecurityBaselineStatesCompleteResult{
		Items: items,
	}
	return
}
