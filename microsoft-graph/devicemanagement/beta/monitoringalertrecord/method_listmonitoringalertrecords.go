package monitoringalertrecord

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

type ListMonitoringAlertRecordsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementAlertRecord
}

type ListMonitoringAlertRecordsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementAlertRecord
}

type ListMonitoringAlertRecordsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMonitoringAlertRecordsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMonitoringAlertRecords ...
func (c MonitoringAlertRecordClient) ListMonitoringAlertRecords(ctx context.Context) (result ListMonitoringAlertRecordsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMonitoringAlertRecordsCustomPager{},
		Path:       "/deviceManagement/monitoring/alertRecords",
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
		Values *[]beta.DeviceManagementAlertRecord `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMonitoringAlertRecordsComplete retrieves all the results into a single object
func (c MonitoringAlertRecordClient) ListMonitoringAlertRecordsComplete(ctx context.Context) (ListMonitoringAlertRecordsCompleteResult, error) {
	return c.ListMonitoringAlertRecordsCompleteMatchingPredicate(ctx, DeviceManagementAlertRecordOperationPredicate{})
}

// ListMonitoringAlertRecordsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MonitoringAlertRecordClient) ListMonitoringAlertRecordsCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementAlertRecordOperationPredicate) (result ListMonitoringAlertRecordsCompleteResult, err error) {
	items := make([]beta.DeviceManagementAlertRecord, 0)

	resp, err := c.ListMonitoringAlertRecords(ctx)
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

	result = ListMonitoringAlertRecordsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
