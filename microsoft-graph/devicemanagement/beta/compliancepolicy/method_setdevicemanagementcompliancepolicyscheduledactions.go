package compliancepolicy

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

type SetDeviceManagementCompliancePolicyScheduledActionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementComplianceScheduledActionForRule
}

type SetDeviceManagementCompliancePolicyScheduledActionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementComplianceScheduledActionForRule
}

type SetDeviceManagementCompliancePolicyScheduledActionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *SetDeviceManagementCompliancePolicyScheduledActionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SetDeviceManagementCompliancePolicyScheduledActions ...
func (c CompliancePolicyClient) SetDeviceManagementCompliancePolicyScheduledActions(ctx context.Context, id DeviceManagementCompliancePolicyId, input SetDeviceManagementCompliancePolicyScheduledActionsRequest) (result SetDeviceManagementCompliancePolicyScheduledActionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &SetDeviceManagementCompliancePolicyScheduledActionsCustomPager{},
		Path:       fmt.Sprintf("%s/setScheduledActions", id.ID()),
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

// SetDeviceManagementCompliancePolicyScheduledActionsComplete retrieves all the results into a single object
func (c CompliancePolicyClient) SetDeviceManagementCompliancePolicyScheduledActionsComplete(ctx context.Context, id DeviceManagementCompliancePolicyId, input SetDeviceManagementCompliancePolicyScheduledActionsRequest) (SetDeviceManagementCompliancePolicyScheduledActionsCompleteResult, error) {
	return c.SetDeviceManagementCompliancePolicyScheduledActionsCompleteMatchingPredicate(ctx, id, input, DeviceManagementComplianceScheduledActionForRuleOperationPredicate{})
}

// SetDeviceManagementCompliancePolicyScheduledActionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CompliancePolicyClient) SetDeviceManagementCompliancePolicyScheduledActionsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementCompliancePolicyId, input SetDeviceManagementCompliancePolicyScheduledActionsRequest, predicate DeviceManagementComplianceScheduledActionForRuleOperationPredicate) (result SetDeviceManagementCompliancePolicyScheduledActionsCompleteResult, err error) {
	items := make([]beta.DeviceManagementComplianceScheduledActionForRule, 0)

	resp, err := c.SetDeviceManagementCompliancePolicyScheduledActions(ctx, id, input)
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

	result = SetDeviceManagementCompliancePolicyScheduledActionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
