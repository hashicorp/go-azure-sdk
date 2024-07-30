package monitoringalertrule

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

type ListMonitoringAlertRulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementAlertRule
}

type ListMonitoringAlertRulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementAlertRule
}

type ListMonitoringAlertRulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMonitoringAlertRulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMonitoringAlertRules ...
func (c MonitoringAlertRuleClient) ListMonitoringAlertRules(ctx context.Context) (result ListMonitoringAlertRulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMonitoringAlertRulesCustomPager{},
		Path:       "/deviceManagement/monitoring/alertRules",
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
		Values *[]beta.DeviceManagementAlertRule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMonitoringAlertRulesComplete retrieves all the results into a single object
func (c MonitoringAlertRuleClient) ListMonitoringAlertRulesComplete(ctx context.Context) (ListMonitoringAlertRulesCompleteResult, error) {
	return c.ListMonitoringAlertRulesCompleteMatchingPredicate(ctx, DeviceManagementAlertRuleOperationPredicate{})
}

// ListMonitoringAlertRulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MonitoringAlertRuleClient) ListMonitoringAlertRulesCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementAlertRuleOperationPredicate) (result ListMonitoringAlertRulesCompleteResult, err error) {
	items := make([]beta.DeviceManagementAlertRule, 0)

	resp, err := c.ListMonitoringAlertRules(ctx)
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

	result = ListMonitoringAlertRulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
