package entitlementmanagementassignment

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

type ListEntitlementManagementAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageAssignment
}

type ListEntitlementManagementAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageAssignment
}

type ListEntitlementManagementAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAssignments ...
func (c EntitlementManagementAssignmentClient) ListEntitlementManagementAssignments(ctx context.Context) (result ListEntitlementManagementAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementAssignmentsCustomPager{},
		Path:       "/identityGovernance/entitlementManagement/assignments",
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
		Values *[]stable.AccessPackageAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAssignmentsComplete retrieves all the results into a single object
func (c EntitlementManagementAssignmentClient) ListEntitlementManagementAssignmentsComplete(ctx context.Context) (ListEntitlementManagementAssignmentsCompleteResult, error) {
	return c.ListEntitlementManagementAssignmentsCompleteMatchingPredicate(ctx, AccessPackageAssignmentOperationPredicate{})
}

// ListEntitlementManagementAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAssignmentClient) ListEntitlementManagementAssignmentsCompleteMatchingPredicate(ctx context.Context, predicate AccessPackageAssignmentOperationPredicate) (result ListEntitlementManagementAssignmentsCompleteResult, err error) {
	items := make([]stable.AccessPackageAssignment, 0)

	resp, err := c.ListEntitlementManagementAssignments(ctx)
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

	result = ListEntitlementManagementAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
