package entitlementmanagementaccesspackageassignmentaccesspackageaccesspackagesincompatiblewith

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

type ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessPackage
}

type ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessPackage
}

type ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsOperationOptions() ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsOperationOptions {
	return ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWiths - Get accessPackagesIncompatibleWith
// from identityGovernance. The access packages that are incompatible with this package. Read-only.
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackagesIncompatibleWithClient) ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWiths(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, options ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsOperationOptions) (result ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsCustomPager{},
		Path:          fmt.Sprintf("%s/accessPackage/accessPackagesIncompatibleWith", id.ID()),
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
		Values *[]beta.AccessPackage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackagesIncompatibleWithClient) ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, options ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsOperationOptions) (ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsCompleteMatchingPredicate(ctx, id, options, AccessPackageOperationPredicate{})
}

// ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackagesIncompatibleWithClient) ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, options ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsOperationOptions, predicate AccessPackageOperationPredicate) (result ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsCompleteResult, err error) {
	items := make([]beta.AccessPackage, 0)

	resp, err := c.ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWiths(ctx, id, options)
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

	result = ListEntitlementManagementAccessPackageAssignmentAccessPackagesIncompatibleWithsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
