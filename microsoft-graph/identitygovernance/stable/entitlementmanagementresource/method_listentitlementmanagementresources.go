package entitlementmanagementresource

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

type ListEntitlementManagementResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResource
}

type ListEntitlementManagementResourcesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResource
}

type ListEntitlementManagementResourcesOperationOptions struct {
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

func DefaultListEntitlementManagementResourcesOperationOptions() ListEntitlementManagementResourcesOperationOptions {
	return ListEntitlementManagementResourcesOperationOptions{}
}

func (o ListEntitlementManagementResourcesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementResourcesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementResourcesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementResourcesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourcesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResources - Get resources from identityGovernance. The resources associated with the
// catalogs.
func (c EntitlementManagementResourceClient) ListEntitlementManagementResources(ctx context.Context, options ListEntitlementManagementResourcesOperationOptions) (result ListEntitlementManagementResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementResourcesCustomPager{},
		Path:          "/identityGovernance/entitlementManagement/resources",
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

// ListEntitlementManagementResourcesComplete retrieves all the results into a single object
func (c EntitlementManagementResourceClient) ListEntitlementManagementResourcesComplete(ctx context.Context, options ListEntitlementManagementResourcesOperationOptions) (ListEntitlementManagementResourcesCompleteResult, error) {
	return c.ListEntitlementManagementResourcesCompleteMatchingPredicate(ctx, options, AccessPackageResourceOperationPredicate{})
}

// ListEntitlementManagementResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceClient) ListEntitlementManagementResourcesCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementResourcesOperationOptions, predicate AccessPackageResourceOperationPredicate) (result ListEntitlementManagementResourcesCompleteResult, err error) {
	items := make([]stable.AccessPackageResource, 0)

	resp, err := c.ListEntitlementManagementResources(ctx, options)
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

	result = ListEntitlementManagementResourcesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
