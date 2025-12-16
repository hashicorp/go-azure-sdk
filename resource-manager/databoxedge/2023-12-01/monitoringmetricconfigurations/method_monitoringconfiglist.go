package monitoringmetricconfigurations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitoringConfigListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]MonitoringMetricConfiguration
}

type MonitoringConfigListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []MonitoringMetricConfiguration
}

type MonitoringConfigListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *MonitoringConfigListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// MonitoringConfigList ...
func (c MonitoringMetricConfigurationsClient) MonitoringConfigList(ctx context.Context, id RoleId) (result MonitoringConfigListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &MonitoringConfigListCustomPager{},
		Path:       fmt.Sprintf("%s/monitoringConfig", id.ID()),
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
		Values *[]MonitoringMetricConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// MonitoringConfigListComplete retrieves all the results into a single object
func (c MonitoringMetricConfigurationsClient) MonitoringConfigListComplete(ctx context.Context, id RoleId) (MonitoringConfigListCompleteResult, error) {
	return c.MonitoringConfigListCompleteMatchingPredicate(ctx, id, MonitoringMetricConfigurationOperationPredicate{})
}

// MonitoringConfigListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MonitoringMetricConfigurationsClient) MonitoringConfigListCompleteMatchingPredicate(ctx context.Context, id RoleId, predicate MonitoringMetricConfigurationOperationPredicate) (result MonitoringConfigListCompleteResult, err error) {
	items := make([]MonitoringMetricConfiguration, 0)

	resp, err := c.MonitoringConfigList(ctx, id)
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

	result = MonitoringConfigListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
