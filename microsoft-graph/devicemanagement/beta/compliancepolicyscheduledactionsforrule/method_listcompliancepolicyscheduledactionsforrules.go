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

type ListCompliancePolicyScheduledActionsForRulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCompliancePolicyScheduledActionsForRulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCompliancePolicyScheduledActionsForRules ...
func (c CompliancePolicyScheduledActionsForRuleClient) ListCompliancePolicyScheduledActionsForRules(ctx context.Context, id DeviceManagementCompliancePolicyId) (result ListCompliancePolicyScheduledActionsForRulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCompliancePolicyScheduledActionsForRulesCustomPager{},
		Path:       fmt.Sprintf("%s/scheduledActionsForRule", id.ID()),
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
func (c CompliancePolicyScheduledActionsForRuleClient) ListCompliancePolicyScheduledActionsForRulesComplete(ctx context.Context, id DeviceManagementCompliancePolicyId) (ListCompliancePolicyScheduledActionsForRulesCompleteResult, error) {
	return c.ListCompliancePolicyScheduledActionsForRulesCompleteMatchingPredicate(ctx, id, DeviceManagementComplianceScheduledActionForRuleOperationPredicate{})
}

// ListCompliancePolicyScheduledActionsForRulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CompliancePolicyScheduledActionsForRuleClient) ListCompliancePolicyScheduledActionsForRulesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementCompliancePolicyId, predicate DeviceManagementComplianceScheduledActionForRuleOperationPredicate) (result ListCompliancePolicyScheduledActionsForRulesCompleteResult, err error) {
	items := make([]beta.DeviceManagementComplianceScheduledActionForRule, 0)

	resp, err := c.ListCompliancePolicyScheduledActionsForRules(ctx, id)
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
