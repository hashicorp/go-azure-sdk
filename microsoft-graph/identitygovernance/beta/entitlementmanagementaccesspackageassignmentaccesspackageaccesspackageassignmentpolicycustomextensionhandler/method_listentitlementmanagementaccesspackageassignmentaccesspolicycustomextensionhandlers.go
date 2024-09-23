package entitlementmanagementaccesspackageassignmentaccesspackageaccesspackageassignmentpolicycustomextensionhandler

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

type ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CustomExtensionHandler
}

type ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CustomExtensionHandler
}

type ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersOperationOptions struct {
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

func DefaultListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersOperationOptions() ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersOperationOptions {
	return ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlers - Get customExtensionHandlers
// from identityGovernance. The collection of stages when to execute one or more custom access package workflow
// extensions. Supports $expand.
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyCustomExtensionHandlerClient) ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlers(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyId, options ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersOperationOptions) (result ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersCustomPager{},
		Path:          fmt.Sprintf("%s/customExtensionHandlers", id.ID()),
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
		Values *[]beta.CustomExtensionHandler `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyCustomExtensionHandlerClient) ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyId, options ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersOperationOptions) (ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersCompleteMatchingPredicate(ctx, id, options, CustomExtensionHandlerOperationPredicate{})
}

// ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyCustomExtensionHandlerClient) ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyId, options ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersOperationOptions, predicate CustomExtensionHandlerOperationPredicate) (result ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersCompleteResult, err error) {
	items := make([]beta.CustomExtensionHandler, 0)

	resp, err := c.ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlers(ctx, id, options)
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

	result = ListEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionHandlersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
