package scopedrolemember

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListScopedRoleMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ScopedRoleMembership
}

type ListScopedRoleMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ScopedRoleMembership
}

type ListScopedRoleMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListScopedRoleMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListScopedRoleMembers ...
func (c ScopedRoleMemberClient) ListScopedRoleMembers(ctx context.Context, id AdministrativeUnitId) (result ListScopedRoleMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListScopedRoleMembersCustomPager{},
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
		Values *[]beta.ScopedRoleMembership `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListScopedRoleMembersComplete retrieves all the results into a single object
func (c ScopedRoleMemberClient) ListScopedRoleMembersComplete(ctx context.Context, id AdministrativeUnitId) (ListScopedRoleMembersCompleteResult, error) {
	return c.ListScopedRoleMembersCompleteMatchingPredicate(ctx, id, ScopedRoleMembershipOperationPredicate{})
}

// ListScopedRoleMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ScopedRoleMemberClient) ListScopedRoleMembersCompleteMatchingPredicate(ctx context.Context, id AdministrativeUnitId, predicate ScopedRoleMembershipOperationPredicate) (result ListScopedRoleMembersCompleteResult, err error) {
	items := make([]beta.ScopedRoleMembership, 0)

	resp, err := c.ListScopedRoleMembers(ctx, id)
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

	result = ListScopedRoleMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
