package directoryadministrativeunitscopedrolemember

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

type ListDirectoryAdministrativeUnitByIdScopedRoleMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ScopedRoleMembershipCollectionResponse
}

type ListDirectoryAdministrativeUnitByIdScopedRoleMembersCompleteResult struct {
	Items []models.ScopedRoleMembershipCollectionResponse
}

// ListDirectoryAdministrativeUnitByIdScopedRoleMembers ...
func (c DirectoryAdministrativeUnitScopedRoleMemberClient) ListDirectoryAdministrativeUnitByIdScopedRoleMembers(ctx context.Context, id DirectoryAdministrativeUnitId) (result ListDirectoryAdministrativeUnitByIdScopedRoleMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/scopedRoleMembers", id.ID()),
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

// ListDirectoryAdministrativeUnitByIdScopedRoleMembersComplete retrieves all the results into a single object
func (c DirectoryAdministrativeUnitScopedRoleMemberClient) ListDirectoryAdministrativeUnitByIdScopedRoleMembersComplete(ctx context.Context, id DirectoryAdministrativeUnitId) (ListDirectoryAdministrativeUnitByIdScopedRoleMembersCompleteResult, error) {
	return c.ListDirectoryAdministrativeUnitByIdScopedRoleMembersCompleteMatchingPredicate(ctx, id, models.ScopedRoleMembershipCollectionResponseOperationPredicate{})
}

// ListDirectoryAdministrativeUnitByIdScopedRoleMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryAdministrativeUnitScopedRoleMemberClient) ListDirectoryAdministrativeUnitByIdScopedRoleMembersCompleteMatchingPredicate(ctx context.Context, id DirectoryAdministrativeUnitId, predicate models.ScopedRoleMembershipCollectionResponseOperationPredicate) (result ListDirectoryAdministrativeUnitByIdScopedRoleMembersCompleteResult, err error) {
	items := make([]models.ScopedRoleMembershipCollectionResponse, 0)

	resp, err := c.ListDirectoryAdministrativeUnitByIdScopedRoleMembers(ctx, id)
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

	result = ListDirectoryAdministrativeUnitByIdScopedRoleMembersCompleteResult{
		Items: items,
	}
	return
}
