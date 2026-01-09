package entitlementmanagementaccesspackageassignmentaccesspackage

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

type GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessPackageAssignmentRequestRequirements
}

type GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessPackageAssignmentRequestRequirements
}

type GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultGetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsOperationOptions() GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsOperationOptions {
	return GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsOperationOptions{}
}

func (o GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirements - Invoke action
// getApplicablePolicyRequirements. In Microsoft Entra entitlement management, this action retrieves a list of
// accessPackageAssignmentRequestRequirements objects that the currently signed-in user can use to create an
// accessPackageAssignmentRequest. Each requirement object corresponds to an access package assignment policy that the
// currently signed-in user is allowed to request an assignment for.
func (c EntitlementManagementAccessPackageAssignmentAccessPackageClient) GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirements(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, options GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsOperationOptions) (result GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsCustomPager{},
		Path:          fmt.Sprintf("%s/accessPackage/getApplicablePolicyRequirements", id.ID()),
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
		Values *[]beta.AccessPackageAssignmentRequestRequirements `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAssignmentAccessPackageClient) GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, options GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsOperationOptions) (GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsCompleteResult, error) {
	return c.GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsCompleteMatchingPredicate(ctx, id, options, AccessPackageAssignmentRequestRequirementsOperationPredicate{})
}

// GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAssignmentAccessPackageClient) GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, options GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsOperationOptions, predicate AccessPackageAssignmentRequestRequirementsOperationPredicate) (result GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsCompleteResult, err error) {
	items := make([]beta.AccessPackageAssignmentRequestRequirements, 0)

	resp, err := c.GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirements(ctx, id, options)
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

	result = GetEntitlementManagementAccessPackageAssignmentApplicablePolicyRequirementsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
