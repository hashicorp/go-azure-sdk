package entitlementmanagementresourcerequestcatalogaccesspackage

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

type ListEntitlementManagementResourceRequestCatalogAccessPackagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackage
}

type ListEntitlementManagementResourceRequestCatalogAccessPackagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackage
}

type ListEntitlementManagementResourceRequestCatalogAccessPackagesOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListEntitlementManagementResourceRequestCatalogAccessPackagesOperationOptions() ListEntitlementManagementResourceRequestCatalogAccessPackagesOperationOptions {
	return ListEntitlementManagementResourceRequestCatalogAccessPackagesOperationOptions{}
}

func (o ListEntitlementManagementResourceRequestCatalogAccessPackagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementResourceRequestCatalogAccessPackagesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementResourceRequestCatalogAccessPackagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementResourceRequestCatalogAccessPackagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourceRequestCatalogAccessPackagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResourceRequestCatalogAccessPackages - Get accessPackages from identityGovernance. The
// access packages in this catalog. Read-only. Nullable.
func (c EntitlementManagementResourceRequestCatalogAccessPackageClient) ListEntitlementManagementResourceRequestCatalogAccessPackages(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRequestId, options ListEntitlementManagementResourceRequestCatalogAccessPackagesOperationOptions) (result ListEntitlementManagementResourceRequestCatalogAccessPackagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementResourceRequestCatalogAccessPackagesCustomPager{},
		Path:          fmt.Sprintf("%s/catalog/accessPackages", id.ID()),
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
		Values *[]stable.AccessPackage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementResourceRequestCatalogAccessPackagesComplete retrieves all the results into a single object
func (c EntitlementManagementResourceRequestCatalogAccessPackageClient) ListEntitlementManagementResourceRequestCatalogAccessPackagesComplete(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRequestId, options ListEntitlementManagementResourceRequestCatalogAccessPackagesOperationOptions) (ListEntitlementManagementResourceRequestCatalogAccessPackagesCompleteResult, error) {
	return c.ListEntitlementManagementResourceRequestCatalogAccessPackagesCompleteMatchingPredicate(ctx, id, options, AccessPackageOperationPredicate{})
}

// ListEntitlementManagementResourceRequestCatalogAccessPackagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceRequestCatalogAccessPackageClient) ListEntitlementManagementResourceRequestCatalogAccessPackagesCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRequestId, options ListEntitlementManagementResourceRequestCatalogAccessPackagesOperationOptions, predicate AccessPackageOperationPredicate) (result ListEntitlementManagementResourceRequestCatalogAccessPackagesCompleteResult, err error) {
	items := make([]stable.AccessPackage, 0)

	resp, err := c.ListEntitlementManagementResourceRequestCatalogAccessPackages(ctx, id, options)
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

	result = ListEntitlementManagementResourceRequestCatalogAccessPackagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
