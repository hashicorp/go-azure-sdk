package entitlementmanagementconnectedorganization

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

type ListEntitlementManagementConnectedOrganizationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ConnectedOrganization
}

type ListEntitlementManagementConnectedOrganizationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ConnectedOrganization
}

type ListEntitlementManagementConnectedOrganizationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementConnectedOrganizationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementConnectedOrganizations ...
func (c EntitlementManagementConnectedOrganizationClient) ListEntitlementManagementConnectedOrganizations(ctx context.Context) (result ListEntitlementManagementConnectedOrganizationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementConnectedOrganizationsCustomPager{},
		Path:       "/identityGovernance/entitlementManagement/connectedOrganizations",
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
		Values *[]stable.ConnectedOrganization `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementConnectedOrganizationsComplete retrieves all the results into a single object
func (c EntitlementManagementConnectedOrganizationClient) ListEntitlementManagementConnectedOrganizationsComplete(ctx context.Context) (ListEntitlementManagementConnectedOrganizationsCompleteResult, error) {
	return c.ListEntitlementManagementConnectedOrganizationsCompleteMatchingPredicate(ctx, ConnectedOrganizationOperationPredicate{})
}

// ListEntitlementManagementConnectedOrganizationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementConnectedOrganizationClient) ListEntitlementManagementConnectedOrganizationsCompleteMatchingPredicate(ctx context.Context, predicate ConnectedOrganizationOperationPredicate) (result ListEntitlementManagementConnectedOrganizationsCompleteResult, err error) {
	items := make([]stable.ConnectedOrganization, 0)

	resp, err := c.ListEntitlementManagementConnectedOrganizations(ctx)
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

	result = ListEntitlementManagementConnectedOrganizationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
