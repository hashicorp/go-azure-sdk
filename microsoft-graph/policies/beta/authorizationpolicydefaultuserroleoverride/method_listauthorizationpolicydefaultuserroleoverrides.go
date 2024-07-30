package authorizationpolicydefaultuserroleoverride

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

type ListAuthorizationPolicyDefaultUserRoleOverridesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DefaultUserRoleOverride
}

type ListAuthorizationPolicyDefaultUserRoleOverridesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DefaultUserRoleOverride
}

type ListAuthorizationPolicyDefaultUserRoleOverridesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthorizationPolicyDefaultUserRoleOverridesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthorizationPolicyDefaultUserRoleOverrides ...
func (c AuthorizationPolicyDefaultUserRoleOverrideClient) ListAuthorizationPolicyDefaultUserRoleOverrides(ctx context.Context, id PolicyAuthorizationPolicyId) (result ListAuthorizationPolicyDefaultUserRoleOverridesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAuthorizationPolicyDefaultUserRoleOverridesCustomPager{},
		Path:       fmt.Sprintf("%s/defaultUserRoleOverrides", id.ID()),
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
		Values *[]beta.DefaultUserRoleOverride `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthorizationPolicyDefaultUserRoleOverridesComplete retrieves all the results into a single object
func (c AuthorizationPolicyDefaultUserRoleOverrideClient) ListAuthorizationPolicyDefaultUserRoleOverridesComplete(ctx context.Context, id PolicyAuthorizationPolicyId) (ListAuthorizationPolicyDefaultUserRoleOverridesCompleteResult, error) {
	return c.ListAuthorizationPolicyDefaultUserRoleOverridesCompleteMatchingPredicate(ctx, id, DefaultUserRoleOverrideOperationPredicate{})
}

// ListAuthorizationPolicyDefaultUserRoleOverridesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthorizationPolicyDefaultUserRoleOverrideClient) ListAuthorizationPolicyDefaultUserRoleOverridesCompleteMatchingPredicate(ctx context.Context, id PolicyAuthorizationPolicyId, predicate DefaultUserRoleOverrideOperationPredicate) (result ListAuthorizationPolicyDefaultUserRoleOverridesCompleteResult, err error) {
	items := make([]beta.DefaultUserRoleOverride, 0)

	resp, err := c.ListAuthorizationPolicyDefaultUserRoleOverrides(ctx, id)
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

	result = ListAuthorizationPolicyDefaultUserRoleOverridesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
