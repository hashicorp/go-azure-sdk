package entitlementmanagementaccesspackageassignmentaccesspackageincompatibleaccesspackage

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

type ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsOperationOptions struct {
	Count     *bool
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Skip      *int64
	Top       *int64
}

func DefaultListEntitlementManagementAccessPackageAssignmentIncompatibleRefsOperationOptions() ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsOperationOptions {
	return ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
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
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageAssignmentIncompatibleRefs - Get ref of incompatibleAccessPackages from
// identityGovernance. The access packages whose assigned users are ineligible to be assigned this access package.
func (c EntitlementManagementAccessPackageAssignmentAccessPackageIncompatibleAccessPackageClient) ListEntitlementManagementAccessPackageAssignmentIncompatibleRefs(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, options ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsOperationOptions) (result ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsCustomPager{},
		Path:          fmt.Sprintf("%s/accessPackage/incompatibleAccessPackages/$ref", id.ID()),
		RetryFunc:     options.RetryFunc,
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
	result.Model = &temp

	return
}

// ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAssignmentAccessPackageIncompatibleAccessPackageClient) ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, options ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsOperationOptions) (ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAssignmentAccessPackageIncompatibleAccessPackageClient) ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, options ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListEntitlementManagementAccessPackageAssignmentIncompatibleRefs(ctx, id, options)
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

	result = ListEntitlementManagementAccessPackageAssignmentIncompatibleRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
