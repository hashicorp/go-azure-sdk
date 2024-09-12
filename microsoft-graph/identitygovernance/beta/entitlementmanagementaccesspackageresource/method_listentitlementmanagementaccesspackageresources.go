package entitlementmanagementaccesspackageresource

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

type ListEntitlementManagementAccessPackageResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessPackageResource
}

type ListEntitlementManagementAccessPackageResourcesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessPackageResource
}

type ListEntitlementManagementAccessPackageResourcesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListEntitlementManagementAccessPackageResourcesOperationOptions() ListEntitlementManagementAccessPackageResourcesOperationOptions {
	return ListEntitlementManagementAccessPackageResourcesOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageResourcesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageResourcesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAccessPackageResourcesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageResourcesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageResourcesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageResources - Get accessPackageResources from identityGovernance. A reference to
// a resource associated with an access package catalog.
func (c EntitlementManagementAccessPackageResourceClient) ListEntitlementManagementAccessPackageResources(ctx context.Context, options ListEntitlementManagementAccessPackageResourcesOperationOptions) (result ListEntitlementManagementAccessPackageResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageResourcesCustomPager{},
		Path:          "/identityGovernance/entitlementManagement/accessPackageResources",
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
		Values *[]beta.AccessPackageResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageResourcesComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageResourceClient) ListEntitlementManagementAccessPackageResourcesComplete(ctx context.Context, options ListEntitlementManagementAccessPackageResourcesOperationOptions) (ListEntitlementManagementAccessPackageResourcesCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageResourcesCompleteMatchingPredicate(ctx, options, AccessPackageResourceOperationPredicate{})
}

// ListEntitlementManagementAccessPackageResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageResourceClient) ListEntitlementManagementAccessPackageResourcesCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementAccessPackageResourcesOperationOptions, predicate AccessPackageResourceOperationPredicate) (result ListEntitlementManagementAccessPackageResourcesCompleteResult, err error) {
	items := make([]beta.AccessPackageResource, 0)

	resp, err := c.ListEntitlementManagementAccessPackageResources(ctx, options)
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

	result = ListEntitlementManagementAccessPackageResourcesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
