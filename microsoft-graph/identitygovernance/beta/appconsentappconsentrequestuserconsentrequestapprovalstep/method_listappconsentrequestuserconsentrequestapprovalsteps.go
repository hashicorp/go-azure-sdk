package appconsentappconsentrequestuserconsentrequestapprovalstep

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

type ListAppConsentRequestUserConsentRequestApprovalStepsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ApprovalStep
}

type ListAppConsentRequestUserConsentRequestApprovalStepsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ApprovalStep
}

type ListAppConsentRequestUserConsentRequestApprovalStepsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListAppConsentRequestUserConsentRequestApprovalStepsOperationOptions() ListAppConsentRequestUserConsentRequestApprovalStepsOperationOptions {
	return ListAppConsentRequestUserConsentRequestApprovalStepsOperationOptions{}
}

func (o ListAppConsentRequestUserConsentRequestApprovalStepsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAppConsentRequestUserConsentRequestApprovalStepsOperationOptions) ToOData() *odata.Query {
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

func (o ListAppConsentRequestUserConsentRequestApprovalStepsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAppConsentRequestUserConsentRequestApprovalStepsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppConsentRequestUserConsentRequestApprovalStepsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppConsentRequestUserConsentRequestApprovalSteps - Get steps from identityGovernance. Used to represent the
// decision associated with a single step in the approval process configured in approvalStage.
func (c AppConsentAppConsentRequestUserConsentRequestApprovalStepClient) ListAppConsentRequestUserConsentRequestApprovalSteps(ctx context.Context, id beta.IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId, options ListAppConsentRequestUserConsentRequestApprovalStepsOperationOptions) (result ListAppConsentRequestUserConsentRequestApprovalStepsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAppConsentRequestUserConsentRequestApprovalStepsCustomPager{},
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

// ListAppConsentRequestUserConsentRequestApprovalStepsComplete retrieves all the results into a single object
func (c AppConsentAppConsentRequestUserConsentRequestApprovalStepClient) ListAppConsentRequestUserConsentRequestApprovalStepsComplete(ctx context.Context, id beta.IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId, options ListAppConsentRequestUserConsentRequestApprovalStepsOperationOptions) (ListAppConsentRequestUserConsentRequestApprovalStepsCompleteResult, error) {
	return c.ListAppConsentRequestUserConsentRequestApprovalStepsCompleteMatchingPredicate(ctx, id, options, ApprovalStepOperationPredicate{})
}

// ListAppConsentRequestUserConsentRequestApprovalStepsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppConsentAppConsentRequestUserConsentRequestApprovalStepClient) ListAppConsentRequestUserConsentRequestApprovalStepsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId, options ListAppConsentRequestUserConsentRequestApprovalStepsOperationOptions, predicate ApprovalStepOperationPredicate) (result ListAppConsentRequestUserConsentRequestApprovalStepsCompleteResult, err error) {
	items := make([]beta.ApprovalStep, 0)

	resp, err := c.ListAppConsentRequestUserConsentRequestApprovalSteps(ctx, id, options)
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

	result = ListAppConsentRequestUserConsentRequestApprovalStepsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
