package healthmonitoringalertconfiguration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListHealthMonitoringAlertConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.HealthMonitoringAlertConfiguration
}

type ListHealthMonitoringAlertConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.HealthMonitoringAlertConfiguration
}

type ListHealthMonitoringAlertConfigurationsOperationOptions struct {
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

func DefaultListHealthMonitoringAlertConfigurationsOperationOptions() ListHealthMonitoringAlertConfigurationsOperationOptions {
	return ListHealthMonitoringAlertConfigurationsOperationOptions{}
}

func (o ListHealthMonitoringAlertConfigurationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListHealthMonitoringAlertConfigurationsOperationOptions) ToOData() *odata.Query {
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

func (o ListHealthMonitoringAlertConfigurationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListHealthMonitoringAlertConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListHealthMonitoringAlertConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListHealthMonitoringAlertConfigurations - List alertConfigurations. Get a list of the Microsoft Entra health
// monitoring alertConfiguration objects and their properties.
func (c HealthMonitoringAlertConfigurationClient) ListHealthMonitoringAlertConfigurations(ctx context.Context, options ListHealthMonitoringAlertConfigurationsOperationOptions) (result ListHealthMonitoringAlertConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListHealthMonitoringAlertConfigurationsCustomPager{},
		Path:          "/reports/healthMonitoring/alertConfigurations",
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
		Values *[]beta.HealthMonitoringAlertConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListHealthMonitoringAlertConfigurationsComplete retrieves all the results into a single object
func (c HealthMonitoringAlertConfigurationClient) ListHealthMonitoringAlertConfigurationsComplete(ctx context.Context, options ListHealthMonitoringAlertConfigurationsOperationOptions) (ListHealthMonitoringAlertConfigurationsCompleteResult, error) {
	return c.ListHealthMonitoringAlertConfigurationsCompleteMatchingPredicate(ctx, options, HealthMonitoringAlertConfigurationOperationPredicate{})
}

// ListHealthMonitoringAlertConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HealthMonitoringAlertConfigurationClient) ListHealthMonitoringAlertConfigurationsCompleteMatchingPredicate(ctx context.Context, options ListHealthMonitoringAlertConfigurationsOperationOptions, predicate HealthMonitoringAlertConfigurationOperationPredicate) (result ListHealthMonitoringAlertConfigurationsCompleteResult, err error) {
	items := make([]beta.HealthMonitoringAlertConfiguration, 0)

	resp, err := c.ListHealthMonitoringAlertConfigurations(ctx, options)
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

	result = ListHealthMonitoringAlertConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
