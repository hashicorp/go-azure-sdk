package entitlementmanagementaccesspackage

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

type ListEntitlementManagementAccessPackagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackage
}

type ListEntitlementManagementAccessPackagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackage
}

type ListEntitlementManagementAccessPackagesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListEntitlementManagementAccessPackagesOperationOptions() ListEntitlementManagementAccessPackagesOperationOptions {
	return ListEntitlementManagementAccessPackagesOperationOptions{}
}

func (o ListEntitlementManagementAccessPackagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackagesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAccessPackagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackages - List accessPackages. Retrieve a list of accessPackage objects. The
// resulting list includes all the access packages that the caller has access to read, across all catalogs.
func (c EntitlementManagementAccessPackageClient) ListEntitlementManagementAccessPackages(ctx context.Context, options ListEntitlementManagementAccessPackagesOperationOptions) (result ListEntitlementManagementAccessPackagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackagesCustomPager{},
		Path:          "/identityGovernance/entitlementManagement/accessPackages",
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
		Values *[]stable.AccessPackage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackagesComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageClient) ListEntitlementManagementAccessPackagesComplete(ctx context.Context, options ListEntitlementManagementAccessPackagesOperationOptions) (ListEntitlementManagementAccessPackagesCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackagesCompleteMatchingPredicate(ctx, options, AccessPackageOperationPredicate{})
}

// ListEntitlementManagementAccessPackagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageClient) ListEntitlementManagementAccessPackagesCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementAccessPackagesOperationOptions, predicate AccessPackageOperationPredicate) (result ListEntitlementManagementAccessPackagesCompleteResult, err error) {
	items := make([]stable.AccessPackage, 0)

	resp, err := c.ListEntitlementManagementAccessPackages(ctx, options)
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

	result = ListEntitlementManagementAccessPackagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
