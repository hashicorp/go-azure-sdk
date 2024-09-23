package enterpriseapproleassignmentapproval

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

type ListEnterpriseAppRoleAssignmentApprovalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Approval
}

type ListEnterpriseAppRoleAssignmentApprovalsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Approval
}

type ListEnterpriseAppRoleAssignmentApprovalsOperationOptions struct {
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

func DefaultListEnterpriseAppRoleAssignmentApprovalsOperationOptions() ListEnterpriseAppRoleAssignmentApprovalsOperationOptions {
	return ListEnterpriseAppRoleAssignmentApprovalsOperationOptions{}
}

func (o ListEnterpriseAppRoleAssignmentApprovalsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEnterpriseAppRoleAssignmentApprovalsOperationOptions) ToOData() *odata.Query {
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

func (o ListEnterpriseAppRoleAssignmentApprovalsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEnterpriseAppRoleAssignmentApprovalsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEnterpriseAppRoleAssignmentApprovalsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEnterpriseAppRoleAssignmentApprovals - Get roleAssignmentApprovals from roleManagement
func (c EnterpriseAppRoleAssignmentApprovalClient) ListEnterpriseAppRoleAssignmentApprovals(ctx context.Context, id beta.RoleManagementEnterpriseAppId, options ListEnterpriseAppRoleAssignmentApprovalsOperationOptions) (result ListEnterpriseAppRoleAssignmentApprovalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEnterpriseAppRoleAssignmentApprovalsCustomPager{},
		Path:          fmt.Sprintf("%s/roleAssignmentApprovals", id.ID()),
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
		Values *[]beta.Approval `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEnterpriseAppRoleAssignmentApprovalsComplete retrieves all the results into a single object
func (c EnterpriseAppRoleAssignmentApprovalClient) ListEnterpriseAppRoleAssignmentApprovalsComplete(ctx context.Context, id beta.RoleManagementEnterpriseAppId, options ListEnterpriseAppRoleAssignmentApprovalsOperationOptions) (ListEnterpriseAppRoleAssignmentApprovalsCompleteResult, error) {
	return c.ListEnterpriseAppRoleAssignmentApprovalsCompleteMatchingPredicate(ctx, id, options, ApprovalOperationPredicate{})
}

// ListEnterpriseAppRoleAssignmentApprovalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EnterpriseAppRoleAssignmentApprovalClient) ListEnterpriseAppRoleAssignmentApprovalsCompleteMatchingPredicate(ctx context.Context, id beta.RoleManagementEnterpriseAppId, options ListEnterpriseAppRoleAssignmentApprovalsOperationOptions, predicate ApprovalOperationPredicate) (result ListEnterpriseAppRoleAssignmentApprovalsCompleteResult, err error) {
	items := make([]beta.Approval, 0)

	resp, err := c.ListEnterpriseAppRoleAssignmentApprovals(ctx, id, options)
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

	result = ListEnterpriseAppRoleAssignmentApprovalsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
