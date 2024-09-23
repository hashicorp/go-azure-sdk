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

type ListMonitoringAlertRecordsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListMonitoringAlertRecordsOperationOptions() ListMonitoringAlertRecordsOperationOptions {
	return ListMonitoringAlertRecordsOperationOptions{}
}

func (o ListMonitoringAlertRecordsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMonitoringAlertRecordsOperationOptions) ToOData() *odata.Query {
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

func (o ListMonitoringAlertRecordsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListMonitoringAlertRecords - List alertRecords. Get a list of the alertRecord objects and their properties.
func (c MonitoringAlertRecordClient) ListMonitoringAlertRecords(ctx context.Context, options ListMonitoringAlertRecordsOperationOptions) (result ListMonitoringAlertRecordsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMonitoringAlertRecordsCustomPager{},
		Path:          "/deviceManagement/monitoring/alertRecords",
		RetryFunc:     options.RetryFunc,
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
func (c MonitoringAlertRecordClient) ListMonitoringAlertRecordsComplete(ctx context.Context, options ListMonitoringAlertRecordsOperationOptions) (ListMonitoringAlertRecordsCompleteResult, error) {
	return c.ListMonitoringAlertRecordsCompleteMatchingPredicate(ctx, options, DeviceManagementAlertRecordOperationPredicate{})
}

// ListMonitoringAlertRecordsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MonitoringAlertRecordClient) ListMonitoringAlertRecordsCompleteMatchingPredicate(ctx context.Context, options ListMonitoringAlertRecordsOperationOptions, predicate DeviceManagementAlertRecordOperationPredicate) (result ListMonitoringAlertRecordsCompleteResult, err error) {
	items := make([]beta.DeviceManagementAlertRecord, 0)

	resp, err := c.ListMonitoringAlertRecords(ctx, options)
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
