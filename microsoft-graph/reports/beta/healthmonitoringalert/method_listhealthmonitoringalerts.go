package healthmonitoringalert

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

type ListHealthMonitoringAlertsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.HealthMonitoringAlert
}

type ListHealthMonitoringAlertsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.HealthMonitoringAlert
}

type ListHealthMonitoringAlertsOperationOptions struct {
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

func DefaultListHealthMonitoringAlertsOperationOptions() ListHealthMonitoringAlertsOperationOptions {
	return ListHealthMonitoringAlertsOperationOptions{}
}

func (o ListHealthMonitoringAlertsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListHealthMonitoringAlertsOperationOptions) ToOData() *odata.Query {
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

func (o ListHealthMonitoringAlertsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListHealthMonitoringAlertsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListHealthMonitoringAlertsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListHealthMonitoringAlerts - List alerts. Get the list of the Microsoft Entra health monitoring alert objects and
// their properties from the past 30 days.
func (c HealthMonitoringAlertClient) ListHealthMonitoringAlerts(ctx context.Context, options ListHealthMonitoringAlertsOperationOptions) (result ListHealthMonitoringAlertsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListHealthMonitoringAlertsCustomPager{},
		Path:          "/reports/healthMonitoring/alerts",
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
		Values *[]beta.HealthMonitoringAlert `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListHealthMonitoringAlertsComplete retrieves all the results into a single object
func (c HealthMonitoringAlertClient) ListHealthMonitoringAlertsComplete(ctx context.Context, options ListHealthMonitoringAlertsOperationOptions) (ListHealthMonitoringAlertsCompleteResult, error) {
	return c.ListHealthMonitoringAlertsCompleteMatchingPredicate(ctx, options, HealthMonitoringAlertOperationPredicate{})
}

// ListHealthMonitoringAlertsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HealthMonitoringAlertClient) ListHealthMonitoringAlertsCompleteMatchingPredicate(ctx context.Context, options ListHealthMonitoringAlertsOperationOptions, predicate HealthMonitoringAlertOperationPredicate) (result ListHealthMonitoringAlertsCompleteResult, err error) {
	items := make([]beta.HealthMonitoringAlert, 0)

	resp, err := c.ListHealthMonitoringAlerts(ctx, options)
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

	result = ListHealthMonitoringAlertsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
