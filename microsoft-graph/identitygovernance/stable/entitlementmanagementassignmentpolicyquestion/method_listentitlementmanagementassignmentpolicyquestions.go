package entitlementmanagementassignmentpolicyquestion

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEntitlementManagementAssignmentPolicyQuestionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageQuestion
}

type ListEntitlementManagementAssignmentPolicyQuestionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageQuestion
}

type ListEntitlementManagementAssignmentPolicyQuestionsOperationOptions struct {
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

func DefaultListEntitlementManagementAssignmentPolicyQuestionsOperationOptions() ListEntitlementManagementAssignmentPolicyQuestionsOperationOptions {
	return ListEntitlementManagementAssignmentPolicyQuestionsOperationOptions{}
}

func (o ListEntitlementManagementAssignmentPolicyQuestionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAssignmentPolicyQuestionsOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAssignmentPolicyQuestionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAssignmentPolicyQuestionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAssignmentPolicyQuestionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAssignmentPolicyQuestions - Get questions from identityGovernance. Questions that are posed
// to the requestor.
func (c EntitlementManagementAssignmentPolicyQuestionClient) ListEntitlementManagementAssignmentPolicyQuestions(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAssignmentPolicyId, options ListEntitlementManagementAssignmentPolicyQuestionsOperationOptions) (result ListEntitlementManagementAssignmentPolicyQuestionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAssignmentPolicyQuestionsCustomPager{},
		Path:          fmt.Sprintf("%s/questions", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]stable.AccessPackageQuestion, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalAccessPackageQuestionImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.AccessPackageQuestion (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListEntitlementManagementAssignmentPolicyQuestionsComplete retrieves all the results into a single object
func (c EntitlementManagementAssignmentPolicyQuestionClient) ListEntitlementManagementAssignmentPolicyQuestionsComplete(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAssignmentPolicyId, options ListEntitlementManagementAssignmentPolicyQuestionsOperationOptions) (ListEntitlementManagementAssignmentPolicyQuestionsCompleteResult, error) {
	return c.ListEntitlementManagementAssignmentPolicyQuestionsCompleteMatchingPredicate(ctx, id, options, AccessPackageQuestionOperationPredicate{})
}

// ListEntitlementManagementAssignmentPolicyQuestionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAssignmentPolicyQuestionClient) ListEntitlementManagementAssignmentPolicyQuestionsCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAssignmentPolicyId, options ListEntitlementManagementAssignmentPolicyQuestionsOperationOptions, predicate AccessPackageQuestionOperationPredicate) (result ListEntitlementManagementAssignmentPolicyQuestionsCompleteResult, err error) {
	items := make([]stable.AccessPackageQuestion, 0)

	resp, err := c.ListEntitlementManagementAssignmentPolicyQuestions(ctx, id, options)
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

	result = ListEntitlementManagementAssignmentPolicyQuestionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
