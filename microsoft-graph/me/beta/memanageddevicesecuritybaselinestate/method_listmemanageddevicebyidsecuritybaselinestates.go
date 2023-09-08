package memanageddevicesecuritybaselinestate

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

type ListMeManagedDeviceByIdSecurityBaselineStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SecurityBaselineStateCollectionResponse
}

type ListMeManagedDeviceByIdSecurityBaselineStatesCompleteResult struct {
	Items []models.SecurityBaselineStateCollectionResponse
}

// ListMeManagedDeviceByIdSecurityBaselineStates ...
func (c MeManagedDeviceSecurityBaselineStateClient) ListMeManagedDeviceByIdSecurityBaselineStates(ctx context.Context, id MeManagedDeviceId) (result ListMeManagedDeviceByIdSecurityBaselineStatesOperationResponse, err error) {
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

// ListMeManagedDeviceByIdSecurityBaselineStatesComplete retrieves all the results into a single object
func (c MeManagedDeviceSecurityBaselineStateClient) ListMeManagedDeviceByIdSecurityBaselineStatesComplete(ctx context.Context, id MeManagedDeviceId) (ListMeManagedDeviceByIdSecurityBaselineStatesCompleteResult, error) {
	return c.ListMeManagedDeviceByIdSecurityBaselineStatesCompleteMatchingPredicate(ctx, id, models.SecurityBaselineStateCollectionResponseOperationPredicate{})
}

// ListMeManagedDeviceByIdSecurityBaselineStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeManagedDeviceSecurityBaselineStateClient) ListMeManagedDeviceByIdSecurityBaselineStatesCompleteMatchingPredicate(ctx context.Context, id MeManagedDeviceId, predicate models.SecurityBaselineStateCollectionResponseOperationPredicate) (result ListMeManagedDeviceByIdSecurityBaselineStatesCompleteResult, err error) {
	items := make([]models.SecurityBaselineStateCollectionResponse, 0)

	resp, err := c.ListMeManagedDeviceByIdSecurityBaselineStates(ctx, id)
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

	result = ListMeManagedDeviceByIdSecurityBaselineStatesCompleteResult{
		Items: items,
	}
	return
}
