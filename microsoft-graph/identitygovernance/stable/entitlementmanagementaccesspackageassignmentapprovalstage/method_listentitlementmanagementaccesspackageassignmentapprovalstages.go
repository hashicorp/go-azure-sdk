package entitlementmanagementaccesspackageassignmentapprovalstage

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

type ListEntitlementManagementAccessPackageAssignmentApprovalStagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ApprovalStage
}

type ListEntitlementManagementAccessPackageAssignmentApprovalStagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ApprovalStage
}

type ListEntitlementManagementAccessPackageAssignmentApprovalStagesOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListEntitlementManagementAccessPackageAssignmentApprovalStagesOperationOptions() ListEntitlementManagementAccessPackageAssignmentApprovalStagesOperationOptions {
	return ListEntitlementManagementAccessPackageAssignmentApprovalStagesOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageAssignmentApprovalStagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageAssignmentApprovalStagesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAccessPackageAssignmentApprovalStagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageAssignmentApprovalStagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageAssignmentApprovalStagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageAssignmentApprovalStages - List approval stages. List the approvalStage objects
// associated with an approval. This API request is made by an approver in the following scenarios: In Microsoft Entra
// entitlement management, providing the identifier of the access package assignment request. In PIM for groups,
// providing the identifier of the assignment schedule request.
func (c EntitlementManagementAccessPackageAssignmentApprovalStageClient) ListEntitlementManagementAccessPackageAssignmentApprovalStages(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId, options ListEntitlementManagementAccessPackageAssignmentApprovalStagesOperationOptions) (result ListEntitlementManagementAccessPackageAssignmentApprovalStagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageAssignmentApprovalStagesCustomPager{},
		Path:          fmt.Sprintf("%s/stages", id.ID()),
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
		Values *[]stable.ApprovalStage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageAssignmentApprovalStagesComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAssignmentApprovalStageClient) ListEntitlementManagementAccessPackageAssignmentApprovalStagesComplete(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId, options ListEntitlementManagementAccessPackageAssignmentApprovalStagesOperationOptions) (ListEntitlementManagementAccessPackageAssignmentApprovalStagesCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageAssignmentApprovalStagesCompleteMatchingPredicate(ctx, id, options, ApprovalStageOperationPredicate{})
}

// ListEntitlementManagementAccessPackageAssignmentApprovalStagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAssignmentApprovalStageClient) ListEntitlementManagementAccessPackageAssignmentApprovalStagesCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId, options ListEntitlementManagementAccessPackageAssignmentApprovalStagesOperationOptions, predicate ApprovalStageOperationPredicate) (result ListEntitlementManagementAccessPackageAssignmentApprovalStagesCompleteResult, err error) {
	items := make([]stable.ApprovalStage, 0)

	resp, err := c.ListEntitlementManagementAccessPackageAssignmentApprovalStages(ctx, id, options)
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

	result = ListEntitlementManagementAccessPackageAssignmentApprovalStagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
