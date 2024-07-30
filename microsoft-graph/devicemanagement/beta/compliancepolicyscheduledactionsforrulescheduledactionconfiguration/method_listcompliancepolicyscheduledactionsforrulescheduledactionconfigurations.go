package compliancepolicyscheduledactionsforrulescheduledactionconfiguration

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

type ListCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementComplianceActionItem
}

type ListCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementComplianceActionItem
}

type ListCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCompliancePolicyScheduledActionsForRuleScheduledActionConfigurations ...
func (c CompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient) ListCompliancePolicyScheduledActionsForRuleScheduledActionConfigurations(ctx context.Context, id DeviceManagementCompliancePolicyIdScheduledActionsForRuleId) (result ListCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCustomPager{},
		Path:       fmt.Sprintf("%s/scheduledActionConfigurations", id.ID()),
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
		Values *[]beta.DeviceManagementComplianceActionItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsComplete retrieves all the results into a single object
func (c CompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient) ListCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsComplete(ctx context.Context, id DeviceManagementCompliancePolicyIdScheduledActionsForRuleId) (ListCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCompleteResult, error) {
	return c.ListCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCompleteMatchingPredicate(ctx, id, DeviceManagementComplianceActionItemOperationPredicate{})
}

// ListCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient) ListCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementCompliancePolicyIdScheduledActionsForRuleId, predicate DeviceManagementComplianceActionItemOperationPredicate) (result ListCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCompleteResult, err error) {
	items := make([]beta.DeviceManagementComplianceActionItem, 0)

	resp, err := c.ListCompliancePolicyScheduledActionsForRuleScheduledActionConfigurations(ctx, id)
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

	result = ListCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
