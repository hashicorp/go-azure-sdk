package administrativeunitscopedrolemember

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

type ListAdministrativeUnitByIdScopedRoleMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ScopedRoleMembershipCollectionResponse
}

type ListAdministrativeUnitByIdScopedRoleMembersCompleteResult struct {
	Items []models.ScopedRoleMembershipCollectionResponse
}

// ListAdministrativeUnitByIdScopedRoleMembers ...
func (c AdministrativeUnitScopedRoleMemberClient) ListAdministrativeUnitByIdScopedRoleMembers(ctx context.Context, id AdministrativeUnitId) (result ListAdministrativeUnitByIdScopedRoleMembersOperationResponse, err error) {
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

// ListAdministrativeUnitByIdScopedRoleMembersComplete retrieves all the results into a single object
func (c AdministrativeUnitScopedRoleMemberClient) ListAdministrativeUnitByIdScopedRoleMembersComplete(ctx context.Context, id AdministrativeUnitId) (ListAdministrativeUnitByIdScopedRoleMembersCompleteResult, error) {
	return c.ListAdministrativeUnitByIdScopedRoleMembersCompleteMatchingPredicate(ctx, id, models.ScopedRoleMembershipCollectionResponseOperationPredicate{})
}

// ListAdministrativeUnitByIdScopedRoleMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AdministrativeUnitScopedRoleMemberClient) ListAdministrativeUnitByIdScopedRoleMembersCompleteMatchingPredicate(ctx context.Context, id AdministrativeUnitId, predicate models.ScopedRoleMembershipCollectionResponseOperationPredicate) (result ListAdministrativeUnitByIdScopedRoleMembersCompleteResult, err error) {
	items := make([]models.ScopedRoleMembershipCollectionResponse, 0)

	resp, err := c.ListAdministrativeUnitByIdScopedRoleMembers(ctx, id)
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

	result = ListAdministrativeUnitByIdScopedRoleMembersCompleteResult{
		Items: items,
	}
	return
}
