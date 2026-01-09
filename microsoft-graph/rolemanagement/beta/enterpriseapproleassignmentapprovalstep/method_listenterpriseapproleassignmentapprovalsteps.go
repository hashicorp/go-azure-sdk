package enterpriseapproleassignmentapprovalstep

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

type ListEnterpriseAppRoleAssignmentApprovalStepsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ApprovalStep
}

type ListEnterpriseAppRoleAssignmentApprovalStepsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ApprovalStep
}

type ListEnterpriseAppRoleAssignmentApprovalStepsOperationOptions struct {
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

func DefaultListEnterpriseAppRoleAssignmentApprovalStepsOperationOptions() ListEnterpriseAppRoleAssignmentApprovalStepsOperationOptions {
	return ListEnterpriseAppRoleAssignmentApprovalStepsOperationOptions{}
}

func (o ListEnterpriseAppRoleAssignmentApprovalStepsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEnterpriseAppRoleAssignmentApprovalStepsOperationOptions) ToOData() *odata.Query {
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

func (o ListEnterpriseAppRoleAssignmentApprovalStepsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEnterpriseAppRoleAssignmentApprovalStepsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEnterpriseAppRoleAssignmentApprovalStepsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEnterpriseAppRoleAssignmentApprovalSteps - Get steps from roleManagement. Used to represent the decision
// associated with a single step in the approval process configured in approvalStage.
func (c EnterpriseAppRoleAssignmentApprovalStepClient) ListEnterpriseAppRoleAssignmentApprovalSteps(ctx context.Context, id beta.RoleManagementEnterpriseAppIdRoleAssignmentApprovalId, options ListEnterpriseAppRoleAssignmentApprovalStepsOperationOptions) (result ListEnterpriseAppRoleAssignmentApprovalStepsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEnterpriseAppRoleAssignmentApprovalStepsCustomPager{},
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

// ListEnterpriseAppRoleAssignmentApprovalStepsComplete retrieves all the results into a single object
func (c EnterpriseAppRoleAssignmentApprovalStepClient) ListEnterpriseAppRoleAssignmentApprovalStepsComplete(ctx context.Context, id beta.RoleManagementEnterpriseAppIdRoleAssignmentApprovalId, options ListEnterpriseAppRoleAssignmentApprovalStepsOperationOptions) (ListEnterpriseAppRoleAssignmentApprovalStepsCompleteResult, error) {
	return c.ListEnterpriseAppRoleAssignmentApprovalStepsCompleteMatchingPredicate(ctx, id, options, ApprovalStepOperationPredicate{})
}

// ListEnterpriseAppRoleAssignmentApprovalStepsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EnterpriseAppRoleAssignmentApprovalStepClient) ListEnterpriseAppRoleAssignmentApprovalStepsCompleteMatchingPredicate(ctx context.Context, id beta.RoleManagementEnterpriseAppIdRoleAssignmentApprovalId, options ListEnterpriseAppRoleAssignmentApprovalStepsOperationOptions, predicate ApprovalStepOperationPredicate) (result ListEnterpriseAppRoleAssignmentApprovalStepsCompleteResult, err error) {
	items := make([]beta.ApprovalStep, 0)

	resp, err := c.ListEnterpriseAppRoleAssignmentApprovalSteps(ctx, id, options)
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

	result = ListEnterpriseAppRoleAssignmentApprovalStepsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
