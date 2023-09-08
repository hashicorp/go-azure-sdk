package policymobiledevicemanagementpolicyincludedgroup

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

type ListPolicyMobileDeviceManagementPolicyByIdIncludedGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.GroupCollectionResponse
}

type ListPolicyMobileDeviceManagementPolicyByIdIncludedGroupsCompleteResult struct {
	Items []models.GroupCollectionResponse
}

// ListPolicyMobileDeviceManagementPolicyByIdIncludedGroups ...
func (c PolicyMobileDeviceManagementPolicyIncludedGroupClient) ListPolicyMobileDeviceManagementPolicyByIdIncludedGroups(ctx context.Context, id PolicyMobileDeviceManagementPolicyId) (result ListPolicyMobileDeviceManagementPolicyByIdIncludedGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/includedGroups", id.ID()),
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
		Values *[]models.GroupCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyMobileDeviceManagementPolicyByIdIncludedGroupsComplete retrieves all the results into a single object
func (c PolicyMobileDeviceManagementPolicyIncludedGroupClient) ListPolicyMobileDeviceManagementPolicyByIdIncludedGroupsComplete(ctx context.Context, id PolicyMobileDeviceManagementPolicyId) (ListPolicyMobileDeviceManagementPolicyByIdIncludedGroupsCompleteResult, error) {
	return c.ListPolicyMobileDeviceManagementPolicyByIdIncludedGroupsCompleteMatchingPredicate(ctx, id, models.GroupCollectionResponseOperationPredicate{})
}

// ListPolicyMobileDeviceManagementPolicyByIdIncludedGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyMobileDeviceManagementPolicyIncludedGroupClient) ListPolicyMobileDeviceManagementPolicyByIdIncludedGroupsCompleteMatchingPredicate(ctx context.Context, id PolicyMobileDeviceManagementPolicyId, predicate models.GroupCollectionResponseOperationPredicate) (result ListPolicyMobileDeviceManagementPolicyByIdIncludedGroupsCompleteResult, err error) {
	items := make([]models.GroupCollectionResponse, 0)

	resp, err := c.ListPolicyMobileDeviceManagementPolicyByIdIncludedGroups(ctx, id)
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

	result = ListPolicyMobileDeviceManagementPolicyByIdIncludedGroupsCompleteResult{
		Items: items,
	}
	return
}
