package entitlementmanagementaccesspackageaccesspackagesincompatiblewith

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEntitlementManagementAccessPackageIncompatibleWithsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessPackage
}

type ListEntitlementManagementAccessPackageIncompatibleWithsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessPackage
}

type ListEntitlementManagementAccessPackageIncompatibleWithsOperationOptions struct {
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

func DefaultListEntitlementManagementAccessPackageIncompatibleWithsOperationOptions() ListEntitlementManagementAccessPackageIncompatibleWithsOperationOptions {
	return ListEntitlementManagementAccessPackageIncompatibleWithsOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageIncompatibleWithsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageIncompatibleWithsOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAccessPackageIncompatibleWithsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageIncompatibleWithsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageIncompatibleWithsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageIncompatibleWiths - List accessPackagesIncompatibleWith. Retrieve a list of the
// accessPackage objects marked a specified accessPackage as incompatible.
func (c EntitlementManagementAccessPackageAccessPackagesIncompatibleWithClient) ListEntitlementManagementAccessPackageIncompatibleWiths(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageId, options ListEntitlementManagementAccessPackageIncompatibleWithsOperationOptions) (result ListEntitlementManagementAccessPackageIncompatibleWithsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageIncompatibleWithsCustomPager{},
		Path:          fmt.Sprintf("%s/accessPackagesIncompatibleWith", id.ID()),
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
		Values *[]beta.AccessPackage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageIncompatibleWithsComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAccessPackagesIncompatibleWithClient) ListEntitlementManagementAccessPackageIncompatibleWithsComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageId, options ListEntitlementManagementAccessPackageIncompatibleWithsOperationOptions) (ListEntitlementManagementAccessPackageIncompatibleWithsCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageIncompatibleWithsCompleteMatchingPredicate(ctx, id, options, AccessPackageOperationPredicate{})
}

// ListEntitlementManagementAccessPackageIncompatibleWithsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAccessPackagesIncompatibleWithClient) ListEntitlementManagementAccessPackageIncompatibleWithsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageId, options ListEntitlementManagementAccessPackageIncompatibleWithsOperationOptions, predicate AccessPackageOperationPredicate) (result ListEntitlementManagementAccessPackageIncompatibleWithsCompleteResult, err error) {
	items := make([]beta.AccessPackage, 0)

	resp, err := c.ListEntitlementManagementAccessPackageIncompatibleWiths(ctx, id, options)
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

	result = ListEntitlementManagementAccessPackageIncompatibleWithsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
