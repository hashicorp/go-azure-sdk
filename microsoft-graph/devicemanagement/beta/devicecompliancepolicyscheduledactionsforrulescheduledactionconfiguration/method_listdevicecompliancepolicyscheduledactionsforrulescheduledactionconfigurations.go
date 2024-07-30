package devicecompliancepolicyscheduledactionsforrulescheduledactionconfiguration

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

type ListDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceComplianceActionItem
}

type ListDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceComplianceActionItem
}

type ListDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurations ...
func (c DeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient) ListDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurations(ctx context.Context, id DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId) (result ListDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCustomPager{},
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
		Values *[]beta.DeviceComplianceActionItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsComplete retrieves all the results into a single object
func (c DeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient) ListDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsComplete(ctx context.Context, id DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId) (ListDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCompleteResult, error) {
	return c.ListDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCompleteMatchingPredicate(ctx, id, DeviceComplianceActionItemOperationPredicate{})
}

// ListDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient) ListDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId, predicate DeviceComplianceActionItemOperationPredicate) (result ListDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCompleteResult, err error) {
	items := make([]beta.DeviceComplianceActionItem, 0)

	resp, err := c.ListDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurations(ctx, id)
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

	result = ListDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
