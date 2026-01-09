package entitlementmanagementcatalogresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEntitlementManagementCatalogResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResource
}

type ListEntitlementManagementCatalogResourcesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResource
}

type ListEntitlementManagementCatalogResourcesOperationOptions struct {
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

func DefaultListEntitlementManagementCatalogResourcesOperationOptions() ListEntitlementManagementCatalogResourcesOperationOptions {
	return ListEntitlementManagementCatalogResourcesOperationOptions{}
}

func (o ListEntitlementManagementCatalogResourcesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementCatalogResourcesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementCatalogResourcesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementCatalogResourcesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementCatalogResourcesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementCatalogResources - List resources. Retrieve a list of accessPackageResource objects in an
// accessPackageCatalog.
func (c EntitlementManagementCatalogResourceClient) ListEntitlementManagementCatalogResources(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogId, options ListEntitlementManagementCatalogResourcesOperationOptions) (result ListEntitlementManagementCatalogResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementCatalogResourcesCustomPager{},
		Path:          fmt.Sprintf("%s/resources", id.ID()),
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
		Values *[]stable.AccessPackageResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementCatalogResourcesComplete retrieves all the results into a single object
func (c EntitlementManagementCatalogResourceClient) ListEntitlementManagementCatalogResourcesComplete(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogId, options ListEntitlementManagementCatalogResourcesOperationOptions) (ListEntitlementManagementCatalogResourcesCompleteResult, error) {
	return c.ListEntitlementManagementCatalogResourcesCompleteMatchingPredicate(ctx, id, options, AccessPackageResourceOperationPredicate{})
}

// ListEntitlementManagementCatalogResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementCatalogResourceClient) ListEntitlementManagementCatalogResourcesCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogId, options ListEntitlementManagementCatalogResourcesOperationOptions, predicate AccessPackageResourceOperationPredicate) (result ListEntitlementManagementCatalogResourcesCompleteResult, err error) {
	items := make([]stable.AccessPackageResource, 0)

	resp, err := c.ListEntitlementManagementCatalogResources(ctx, id, options)
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

	result = ListEntitlementManagementCatalogResourcesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
