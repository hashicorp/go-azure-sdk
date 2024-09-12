package appconsentappconsentrequestuserconsentrequestapprovalstage

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

type ListAppConsentRequestUserConsentRequestApprovalStagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ApprovalStage
}

type ListAppConsentRequestUserConsentRequestApprovalStagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ApprovalStage
}

type ListAppConsentRequestUserConsentRequestApprovalStagesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListAppConsentRequestUserConsentRequestApprovalStagesOperationOptions() ListAppConsentRequestUserConsentRequestApprovalStagesOperationOptions {
	return ListAppConsentRequestUserConsentRequestApprovalStagesOperationOptions{}
}

func (o ListAppConsentRequestUserConsentRequestApprovalStagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAppConsentRequestUserConsentRequestApprovalStagesOperationOptions) ToOData() *odata.Query {
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

func (o ListAppConsentRequestUserConsentRequestApprovalStagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAppConsentRequestUserConsentRequestApprovalStagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppConsentRequestUserConsentRequestApprovalStagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppConsentRequestUserConsentRequestApprovalStages - Get stages from identityGovernance. A collection of stages in
// the approval decision.
func (c AppConsentAppConsentRequestUserConsentRequestApprovalStageClient) ListAppConsentRequestUserConsentRequestApprovalStages(ctx context.Context, id stable.IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId, options ListAppConsentRequestUserConsentRequestApprovalStagesOperationOptions) (result ListAppConsentRequestUserConsentRequestApprovalStagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAppConsentRequestUserConsentRequestApprovalStagesCustomPager{},
		Path:          fmt.Sprintf("%s/approval/stages", id.ID()),
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

// ListAppConsentRequestUserConsentRequestApprovalStagesComplete retrieves all the results into a single object
func (c AppConsentAppConsentRequestUserConsentRequestApprovalStageClient) ListAppConsentRequestUserConsentRequestApprovalStagesComplete(ctx context.Context, id stable.IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId, options ListAppConsentRequestUserConsentRequestApprovalStagesOperationOptions) (ListAppConsentRequestUserConsentRequestApprovalStagesCompleteResult, error) {
	return c.ListAppConsentRequestUserConsentRequestApprovalStagesCompleteMatchingPredicate(ctx, id, options, ApprovalStageOperationPredicate{})
}

// ListAppConsentRequestUserConsentRequestApprovalStagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppConsentAppConsentRequestUserConsentRequestApprovalStageClient) ListAppConsentRequestUserConsentRequestApprovalStagesCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId, options ListAppConsentRequestUserConsentRequestApprovalStagesOperationOptions, predicate ApprovalStageOperationPredicate) (result ListAppConsentRequestUserConsentRequestApprovalStagesCompleteResult, err error) {
	items := make([]stable.ApprovalStage, 0)

	resp, err := c.ListAppConsentRequestUserConsentRequestApprovalStages(ctx, id, options)
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

	result = ListAppConsentRequestUserConsentRequestApprovalStagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
