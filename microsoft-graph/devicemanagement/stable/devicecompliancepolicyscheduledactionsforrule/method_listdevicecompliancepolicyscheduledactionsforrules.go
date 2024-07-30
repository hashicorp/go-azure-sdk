package devicecompliancepolicyscheduledactionsforrule

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

type ListDeviceCompliancePolicyScheduledActionsForRulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DeviceComplianceScheduledActionForRule
}

type ListDeviceCompliancePolicyScheduledActionsForRulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DeviceComplianceScheduledActionForRule
}

type ListDeviceCompliancePolicyScheduledActionsForRulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCompliancePolicyScheduledActionsForRulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCompliancePolicyScheduledActionsForRules ...
func (c DeviceCompliancePolicyScheduledActionsForRuleClient) ListDeviceCompliancePolicyScheduledActionsForRules(ctx context.Context, id DeviceManagementDeviceCompliancePolicyId) (result ListDeviceCompliancePolicyScheduledActionsForRulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceCompliancePolicyScheduledActionsForRulesCustomPager{},
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
		Values *[]stable.DeviceComplianceScheduledActionForRule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCompliancePolicyScheduledActionsForRulesComplete retrieves all the results into a single object
func (c DeviceCompliancePolicyScheduledActionsForRuleClient) ListDeviceCompliancePolicyScheduledActionsForRulesComplete(ctx context.Context, id DeviceManagementDeviceCompliancePolicyId) (ListDeviceCompliancePolicyScheduledActionsForRulesCompleteResult, error) {
	return c.ListDeviceCompliancePolicyScheduledActionsForRulesCompleteMatchingPredicate(ctx, id, DeviceComplianceScheduledActionForRuleOperationPredicate{})
}

// ListDeviceCompliancePolicyScheduledActionsForRulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCompliancePolicyScheduledActionsForRuleClient) ListDeviceCompliancePolicyScheduledActionsForRulesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDeviceCompliancePolicyId, predicate DeviceComplianceScheduledActionForRuleOperationPredicate) (result ListDeviceCompliancePolicyScheduledActionsForRulesCompleteResult, err error) {
	items := make([]stable.DeviceComplianceScheduledActionForRule, 0)

	resp, err := c.ListDeviceCompliancePolicyScheduledActionsForRules(ctx, id)
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

	result = ListDeviceCompliancePolicyScheduledActionsForRulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
