package entitlementmanagementaccesspackageassignmentapprovalstep

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

type ListEntitlementManagementAccessPackageAssignmentApprovalStepsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ApprovalStep
}

type ListEntitlementManagementAccessPackageAssignmentApprovalStepsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ApprovalStep
}

type ListEntitlementManagementAccessPackageAssignmentApprovalStepsOperationOptions struct {
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

func DefaultListEntitlementManagementAccessPackageAssignmentApprovalStepsOperationOptions() ListEntitlementManagementAccessPackageAssignmentApprovalStepsOperationOptions {
	return ListEntitlementManagementAccessPackageAssignmentApprovalStepsOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageAssignmentApprovalStepsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageAssignmentApprovalStepsOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAccessPackageAssignmentApprovalStepsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageAssignmentApprovalStepsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageAssignmentApprovalStepsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageAssignmentApprovalSteps - List approvalSteps. List the approvalStep objects
// associated with an approval. This API request is made by an approver in the following scenarios
func (c EntitlementManagementAccessPackageAssignmentApprovalStepClient) ListEntitlementManagementAccessPackageAssignmentApprovalSteps(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId, options ListEntitlementManagementAccessPackageAssignmentApprovalStepsOperationOptions) (result ListEntitlementManagementAccessPackageAssignmentApprovalStepsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageAssignmentApprovalStepsCustomPager{},
		Path:          fmt.Sprintf("%s/steps", id.ID()),
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
		Values *[]beta.ApprovalStep `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageAssignmentApprovalStepsComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAssignmentApprovalStepClient) ListEntitlementManagementAccessPackageAssignmentApprovalStepsComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId, options ListEntitlementManagementAccessPackageAssignmentApprovalStepsOperationOptions) (ListEntitlementManagementAccessPackageAssignmentApprovalStepsCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageAssignmentApprovalStepsCompleteMatchingPredicate(ctx, id, options, ApprovalStepOperationPredicate{})
}

// ListEntitlementManagementAccessPackageAssignmentApprovalStepsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAssignmentApprovalStepClient) ListEntitlementManagementAccessPackageAssignmentApprovalStepsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId, options ListEntitlementManagementAccessPackageAssignmentApprovalStepsOperationOptions, predicate ApprovalStepOperationPredicate) (result ListEntitlementManagementAccessPackageAssignmentApprovalStepsCompleteResult, err error) {
	items := make([]beta.ApprovalStep, 0)

	resp, err := c.ListEntitlementManagementAccessPackageAssignmentApprovalSteps(ctx, id, options)
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

	result = ListEntitlementManagementAccessPackageAssignmentApprovalStepsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
