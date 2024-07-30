package entitlementmanagementaccesspackageincompatiblegroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEntitlementManagementAccessPackageIncompatibleGroupRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]string
}

type ListEntitlementManagementAccessPackageIncompatibleGroupRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []string
}

type ListEntitlementManagementAccessPackageIncompatibleGroupRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageIncompatibleGroupRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageIncompatibleGroupRefs ...
func (c EntitlementManagementAccessPackageIncompatibleGroupClient) ListEntitlementManagementAccessPackageIncompatibleGroupRefs(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageId) (result ListEntitlementManagementAccessPackageIncompatibleGroupRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementAccessPackageIncompatibleGroupRefsCustomPager{},
		Path:       fmt.Sprintf("%s/incompatibleGroups/$ref", id.ID()),
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
		Values *[]string `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageIncompatibleGroupRefsComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageIncompatibleGroupClient) ListEntitlementManagementAccessPackageIncompatibleGroupRefsComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageId) (result ListEntitlementManagementAccessPackageIncompatibleGroupRefsCompleteResult, err error) {
	items := make([]string, 0)

	resp, err := c.ListEntitlementManagementAccessPackageIncompatibleGroupRefs(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			items = append(items, v)
		}
	}

	result = ListEntitlementManagementAccessPackageIncompatibleGroupRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
