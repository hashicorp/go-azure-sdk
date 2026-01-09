package privilegedaccessgroupassignmentapprovalstage

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

type ListPrivilegedAccessGroupAssignmentApprovalStagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ApprovalStage
}

type ListPrivilegedAccessGroupAssignmentApprovalStagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ApprovalStage
}

type ListPrivilegedAccessGroupAssignmentApprovalStagesOperationOptions struct {
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

func DefaultListPrivilegedAccessGroupAssignmentApprovalStagesOperationOptions() ListPrivilegedAccessGroupAssignmentApprovalStagesOperationOptions {
	return ListPrivilegedAccessGroupAssignmentApprovalStagesOperationOptions{}
}

func (o ListPrivilegedAccessGroupAssignmentApprovalStagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPrivilegedAccessGroupAssignmentApprovalStagesOperationOptions) ToOData() *odata.Query {
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

func (o ListPrivilegedAccessGroupAssignmentApprovalStagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPrivilegedAccessGroupAssignmentApprovalStagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPrivilegedAccessGroupAssignmentApprovalStagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivilegedAccessGroupAssignmentApprovalStages - Get stages from identityGovernance. A collection of stages in the
// approval decision.
func (c PrivilegedAccessGroupAssignmentApprovalStageClient) ListPrivilegedAccessGroupAssignmentApprovalStages(ctx context.Context, id stable.IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId, options ListPrivilegedAccessGroupAssignmentApprovalStagesOperationOptions) (result ListPrivilegedAccessGroupAssignmentApprovalStagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPrivilegedAccessGroupAssignmentApprovalStagesCustomPager{},
		Path:          fmt.Sprintf("%s/stages", id.ID()),
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
		Values *[]stable.ApprovalStage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPrivilegedAccessGroupAssignmentApprovalStagesComplete retrieves all the results into a single object
func (c PrivilegedAccessGroupAssignmentApprovalStageClient) ListPrivilegedAccessGroupAssignmentApprovalStagesComplete(ctx context.Context, id stable.IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId, options ListPrivilegedAccessGroupAssignmentApprovalStagesOperationOptions) (ListPrivilegedAccessGroupAssignmentApprovalStagesCompleteResult, error) {
	return c.ListPrivilegedAccessGroupAssignmentApprovalStagesCompleteMatchingPredicate(ctx, id, options, ApprovalStageOperationPredicate{})
}

// ListPrivilegedAccessGroupAssignmentApprovalStagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivilegedAccessGroupAssignmentApprovalStageClient) ListPrivilegedAccessGroupAssignmentApprovalStagesCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId, options ListPrivilegedAccessGroupAssignmentApprovalStagesOperationOptions, predicate ApprovalStageOperationPredicate) (result ListPrivilegedAccessGroupAssignmentApprovalStagesCompleteResult, err error) {
	items := make([]stable.ApprovalStage, 0)

	resp, err := c.ListPrivilegedAccessGroupAssignmentApprovalStages(ctx, id, options)
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

	result = ListPrivilegedAccessGroupAssignmentApprovalStagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
