package entitlementmanagementaccesspackageassignmentincompatiblegroup

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationOptions struct {
	Count   *bool
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Skip    *int64
	Top     *int64
}

func DefaultListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationOptions() ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationOptions {
	return ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefs - Get ref of incompatibleGroups from
// identityGovernance. The groups whose members are ineligible to be assigned this access package.
func (c EntitlementManagementAccessPackageAssignmentIncompatibleGroupClient) ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefs(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, options ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationOptions) (result ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsCustomPager{},
		Path:          fmt.Sprintf("%s/accessPackage/incompatibleGroups/$ref", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.DirectoryObject, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalDirectoryObjectImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.DirectoryObject (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = temp

	return
}

// ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAssignmentIncompatibleGroupClient) ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, options ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationOptions) (ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAssignmentIncompatibleGroupClient) ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, options ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefs(ctx, id, options)
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

	result = ListEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
