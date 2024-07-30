package administrativeunitscopedrolemember

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAdministrativeUnitScopedRoleMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ScopedRoleMembership
}

type ListAdministrativeUnitScopedRoleMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ScopedRoleMembership
}

type ListAdministrativeUnitScopedRoleMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAdministrativeUnitScopedRoleMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAdministrativeUnitScopedRoleMembers ...
func (c AdministrativeUnitScopedRoleMemberClient) ListAdministrativeUnitScopedRoleMembers(ctx context.Context, id DirectoryAdministrativeUnitId) (result ListAdministrativeUnitScopedRoleMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAdministrativeUnitScopedRoleMembersCustomPager{},
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
		Values *[]stable.ScopedRoleMembership `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAdministrativeUnitScopedRoleMembersComplete retrieves all the results into a single object
func (c AdministrativeUnitScopedRoleMemberClient) ListAdministrativeUnitScopedRoleMembersComplete(ctx context.Context, id DirectoryAdministrativeUnitId) (ListAdministrativeUnitScopedRoleMembersCompleteResult, error) {
	return c.ListAdministrativeUnitScopedRoleMembersCompleteMatchingPredicate(ctx, id, ScopedRoleMembershipOperationPredicate{})
}

// ListAdministrativeUnitScopedRoleMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AdministrativeUnitScopedRoleMemberClient) ListAdministrativeUnitScopedRoleMembersCompleteMatchingPredicate(ctx context.Context, id DirectoryAdministrativeUnitId, predicate ScopedRoleMembershipOperationPredicate) (result ListAdministrativeUnitScopedRoleMembersCompleteResult, err error) {
	items := make([]stable.ScopedRoleMembership, 0)

	resp, err := c.ListAdministrativeUnitScopedRoleMembers(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListAdministrativeUnitScopedRoleMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
