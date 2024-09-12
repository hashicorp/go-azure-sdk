package appconsentrequestsforapprovaluserconsentrequestapprovalstep

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

type ListAppConsentRequestsForApprovalUserApprovalStepsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ApprovalStep
}

type ListAppConsentRequestsForApprovalUserApprovalStepsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ApprovalStep
}

type ListAppConsentRequestsForApprovalUserApprovalStepsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListAppConsentRequestsForApprovalUserApprovalStepsOperationOptions() ListAppConsentRequestsForApprovalUserApprovalStepsOperationOptions {
	return ListAppConsentRequestsForApprovalUserApprovalStepsOperationOptions{}
}

func (o ListAppConsentRequestsForApprovalUserApprovalStepsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAppConsentRequestsForApprovalUserApprovalStepsOperationOptions) ToOData() *odata.Query {
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

func (o ListAppConsentRequestsForApprovalUserApprovalStepsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAppConsentRequestsForApprovalUserApprovalStepsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppConsentRequestsForApprovalUserApprovalStepsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppConsentRequestsForApprovalUserApprovalSteps - Get steps from me. Used to represent the decision associated
// with a single step in the approval process configured in approvalStage.
func (c AppConsentRequestsForApprovalUserConsentRequestApprovalStepClient) ListAppConsentRequestsForApprovalUserApprovalSteps(ctx context.Context, id beta.MeAppConsentRequestsForApprovalIdUserConsentRequestId, options ListAppConsentRequestsForApprovalUserApprovalStepsOperationOptions) (result ListAppConsentRequestsForApprovalUserApprovalStepsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAppConsentRequestsForApprovalUserApprovalStepsCustomPager{},
		Path:          fmt.Sprintf("%s/approval/steps", id.ID()),
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

// ListAppConsentRequestsForApprovalUserApprovalStepsComplete retrieves all the results into a single object
func (c AppConsentRequestsForApprovalUserConsentRequestApprovalStepClient) ListAppConsentRequestsForApprovalUserApprovalStepsComplete(ctx context.Context, id beta.MeAppConsentRequestsForApprovalIdUserConsentRequestId, options ListAppConsentRequestsForApprovalUserApprovalStepsOperationOptions) (ListAppConsentRequestsForApprovalUserApprovalStepsCompleteResult, error) {
	return c.ListAppConsentRequestsForApprovalUserApprovalStepsCompleteMatchingPredicate(ctx, id, options, ApprovalStepOperationPredicate{})
}

// ListAppConsentRequestsForApprovalUserApprovalStepsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppConsentRequestsForApprovalUserConsentRequestApprovalStepClient) ListAppConsentRequestsForApprovalUserApprovalStepsCompleteMatchingPredicate(ctx context.Context, id beta.MeAppConsentRequestsForApprovalIdUserConsentRequestId, options ListAppConsentRequestsForApprovalUserApprovalStepsOperationOptions, predicate ApprovalStepOperationPredicate) (result ListAppConsentRequestsForApprovalUserApprovalStepsCompleteResult, err error) {
	items := make([]beta.ApprovalStep, 0)

	resp, err := c.ListAppConsentRequestsForApprovalUserApprovalSteps(ctx, id, options)
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

	result = ListAppConsentRequestsForApprovalUserApprovalStepsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
