package compliancepolicyscheduledactionsforrule

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

type ListCompliancePolicyScheduledActionsForRulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementComplianceScheduledActionForRule
}

type ListCompliancePolicyScheduledActionsForRulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementComplianceScheduledActionForRule
}

type ListCompliancePolicyScheduledActionsForRulesOperationOptions struct {
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

func DefaultListCompliancePolicyScheduledActionsForRulesOperationOptions() ListCompliancePolicyScheduledActionsForRulesOperationOptions {
	return ListCompliancePolicyScheduledActionsForRulesOperationOptions{}
}

func (o ListCompliancePolicyScheduledActionsForRulesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCompliancePolicyScheduledActionsForRulesOperationOptions) ToOData() *odata.Query {
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

func (o ListCompliancePolicyScheduledActionsForRulesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCompliancePolicyScheduledActionsForRulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCompliancePolicyScheduledActionsForRulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCompliancePolicyScheduledActionsForRules - Get scheduledActionsForRule from deviceManagement. The list of
// scheduled action for this rule
func (c CompliancePolicyScheduledActionsForRuleClient) ListCompliancePolicyScheduledActionsForRules(ctx context.Context, id beta.DeviceManagementCompliancePolicyId, options ListCompliancePolicyScheduledActionsForRulesOperationOptions) (result ListCompliancePolicyScheduledActionsForRulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCompliancePolicyScheduledActionsForRulesCustomPager{},
		Path:          fmt.Sprintf("%s/scheduledActionsForRule", id.ID()),
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
		Values *[]beta.DeviceManagementComplianceScheduledActionForRule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCompliancePolicyScheduledActionsForRulesComplete retrieves all the results into a single object
func (c CompliancePolicyScheduledActionsForRuleClient) ListCompliancePolicyScheduledActionsForRulesComplete(ctx context.Context, id beta.DeviceManagementCompliancePolicyId, options ListCompliancePolicyScheduledActionsForRulesOperationOptions) (ListCompliancePolicyScheduledActionsForRulesCompleteResult, error) {
	return c.ListCompliancePolicyScheduledActionsForRulesCompleteMatchingPredicate(ctx, id, options, DeviceManagementComplianceScheduledActionForRuleOperationPredicate{})
}

// ListCompliancePolicyScheduledActionsForRulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CompliancePolicyScheduledActionsForRuleClient) ListCompliancePolicyScheduledActionsForRulesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementCompliancePolicyId, options ListCompliancePolicyScheduledActionsForRulesOperationOptions, predicate DeviceManagementComplianceScheduledActionForRuleOperationPredicate) (result ListCompliancePolicyScheduledActionsForRulesCompleteResult, err error) {
	items := make([]beta.DeviceManagementComplianceScheduledActionForRule, 0)

	resp, err := c.ListCompliancePolicyScheduledActionsForRules(ctx, id, options)
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

	result = ListCompliancePolicyScheduledActionsForRulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
