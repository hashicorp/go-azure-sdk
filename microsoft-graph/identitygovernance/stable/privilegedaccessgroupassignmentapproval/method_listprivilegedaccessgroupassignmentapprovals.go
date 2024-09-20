package privilegedaccessgroupassignmentapproval

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

type ListPrivilegedAccessGroupAssignmentApprovalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Approval
}

type ListPrivilegedAccessGroupAssignmentApprovalsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Approval
}

type ListPrivilegedAccessGroupAssignmentApprovalsOperationOptions struct {
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

func DefaultListPrivilegedAccessGroupAssignmentApprovalsOperationOptions() ListPrivilegedAccessGroupAssignmentApprovalsOperationOptions {
	return ListPrivilegedAccessGroupAssignmentApprovalsOperationOptions{}
}

func (o ListPrivilegedAccessGroupAssignmentApprovalsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPrivilegedAccessGroupAssignmentApprovalsOperationOptions) ToOData() *odata.Query {
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

func (o ListPrivilegedAccessGroupAssignmentApprovalsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPrivilegedAccessGroupAssignmentApprovalsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPrivilegedAccessGroupAssignmentApprovalsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivilegedAccessGroupAssignmentApprovals - Get assignmentApprovals from identityGovernance
func (c PrivilegedAccessGroupAssignmentApprovalClient) ListPrivilegedAccessGroupAssignmentApprovals(ctx context.Context, options ListPrivilegedAccessGroupAssignmentApprovalsOperationOptions) (result ListPrivilegedAccessGroupAssignmentApprovalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPrivilegedAccessGroupAssignmentApprovalsCustomPager{},
		Path:          "/identityGovernance/privilegedAccess/group/assignmentApprovals",
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
		Values *[]stable.Approval `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPrivilegedAccessGroupAssignmentApprovalsComplete retrieves all the results into a single object
func (c PrivilegedAccessGroupAssignmentApprovalClient) ListPrivilegedAccessGroupAssignmentApprovalsComplete(ctx context.Context, options ListPrivilegedAccessGroupAssignmentApprovalsOperationOptions) (ListPrivilegedAccessGroupAssignmentApprovalsCompleteResult, error) {
	return c.ListPrivilegedAccessGroupAssignmentApprovalsCompleteMatchingPredicate(ctx, options, ApprovalOperationPredicate{})
}

// ListPrivilegedAccessGroupAssignmentApprovalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivilegedAccessGroupAssignmentApprovalClient) ListPrivilegedAccessGroupAssignmentApprovalsCompleteMatchingPredicate(ctx context.Context, options ListPrivilegedAccessGroupAssignmentApprovalsOperationOptions, predicate ApprovalOperationPredicate) (result ListPrivilegedAccessGroupAssignmentApprovalsCompleteResult, err error) {
	items := make([]stable.Approval, 0)

	resp, err := c.ListPrivilegedAccessGroupAssignmentApprovals(ctx, options)
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

	result = ListPrivilegedAccessGroupAssignmentApprovalsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
