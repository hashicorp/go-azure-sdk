package userscopedrolememberof

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

type ListUserByIdScopedRoleMemberOfOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ScopedRoleMembershipCollectionResponse
}

type ListUserByIdScopedRoleMemberOfCompleteResult struct {
	Items []models.ScopedRoleMembershipCollectionResponse
}

// ListUserByIdScopedRoleMemberOf ...
func (c UserScopedRoleMemberOfClient) ListUserByIdScopedRoleMemberOf(ctx context.Context, id UserId) (result ListUserByIdScopedRoleMemberOfOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/scopedRoleMemberOf", id.ID()),
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
		Values *[]models.ScopedRoleMembershipCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdScopedRoleMemberOfComplete retrieves all the results into a single object
func (c UserScopedRoleMemberOfClient) ListUserByIdScopedRoleMemberOfComplete(ctx context.Context, id UserId) (ListUserByIdScopedRoleMemberOfCompleteResult, error) {
	return c.ListUserByIdScopedRoleMemberOfCompleteMatchingPredicate(ctx, id, models.ScopedRoleMembershipCollectionResponseOperationPredicate{})
}

// ListUserByIdScopedRoleMemberOfCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserScopedRoleMemberOfClient) ListUserByIdScopedRoleMemberOfCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.ScopedRoleMembershipCollectionResponseOperationPredicate) (result ListUserByIdScopedRoleMemberOfCompleteResult, err error) {
	items := make([]models.ScopedRoleMembershipCollectionResponse, 0)

	resp, err := c.ListUserByIdScopedRoleMemberOf(ctx, id)
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

	result = ListUserByIdScopedRoleMemberOfCompleteResult{
		Items: items,
	}
	return
}
