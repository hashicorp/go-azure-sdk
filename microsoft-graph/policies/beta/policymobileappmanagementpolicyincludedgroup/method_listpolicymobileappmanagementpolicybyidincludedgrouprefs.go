package policymobileappmanagementpolicyincludedgroup

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

type ListPolicyMobileAppManagementPolicyByIdIncludedGroupRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.StringCollectionResponse
}

type ListPolicyMobileAppManagementPolicyByIdIncludedGroupRefsCompleteResult struct {
	Items []models.StringCollectionResponse
}

// ListPolicyMobileAppManagementPolicyByIdIncludedGroupRefs ...
func (c PolicyMobileAppManagementPolicyIncludedGroupClient) ListPolicyMobileAppManagementPolicyByIdIncludedGroupRefs(ctx context.Context, id PolicyMobileAppManagementPolicyId) (result ListPolicyMobileAppManagementPolicyByIdIncludedGroupRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/includedGroups/$ref", id.ID()),
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
		Values *[]models.StringCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyMobileAppManagementPolicyByIdIncludedGroupRefsComplete retrieves all the results into a single object
func (c PolicyMobileAppManagementPolicyIncludedGroupClient) ListPolicyMobileAppManagementPolicyByIdIncludedGroupRefsComplete(ctx context.Context, id PolicyMobileAppManagementPolicyId) (ListPolicyMobileAppManagementPolicyByIdIncludedGroupRefsCompleteResult, error) {
	return c.ListPolicyMobileAppManagementPolicyByIdIncludedGroupRefsCompleteMatchingPredicate(ctx, id, models.StringCollectionResponseOperationPredicate{})
}

// ListPolicyMobileAppManagementPolicyByIdIncludedGroupRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyMobileAppManagementPolicyIncludedGroupClient) ListPolicyMobileAppManagementPolicyByIdIncludedGroupRefsCompleteMatchingPredicate(ctx context.Context, id PolicyMobileAppManagementPolicyId, predicate models.StringCollectionResponseOperationPredicate) (result ListPolicyMobileAppManagementPolicyByIdIncludedGroupRefsCompleteResult, err error) {
	items := make([]models.StringCollectionResponse, 0)

	resp, err := c.ListPolicyMobileAppManagementPolicyByIdIncludedGroupRefs(ctx, id)
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

	result = ListPolicyMobileAppManagementPolicyByIdIncludedGroupRefsCompleteResult{
		Items: items,
	}
	return
}
