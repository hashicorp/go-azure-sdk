package entitlementmanagementaccesspackageaccesspackageassignmentpolicycustomextensionhandler

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

type ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CustomExtensionHandler
}

type ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CustomExtensionHandler
}

type ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersOperationOptions() ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersOperationOptions {
	return ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlers - Get customExtensionHandlers from
// identityGovernance. The collection of stages when to execute one or more custom access package workflow extensions.
// Supports $expand.
func (c EntitlementManagementAccessPackageAccessPackageAssignmentPolicyCustomExtensionHandlerClient) ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlers(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId, options ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersOperationOptions) (result ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersCustomPager{},
		Path:          fmt.Sprintf("%s/customExtensionHandlers", id.ID()),
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
		Values *[]beta.CustomExtensionHandler `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAccessPackageAssignmentPolicyCustomExtensionHandlerClient) ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId, options ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersOperationOptions) (ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersCompleteMatchingPredicate(ctx, id, options, CustomExtensionHandlerOperationPredicate{})
}

// ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAccessPackageAssignmentPolicyCustomExtensionHandlerClient) ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId, options ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersOperationOptions, predicate CustomExtensionHandlerOperationPredicate) (result ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersCompleteResult, err error) {
	items := make([]beta.CustomExtensionHandler, 0)

	resp, err := c.ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlers(ctx, id, options)
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

	result = ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
