package reportcachedreportconfiguration

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

type ListReportCachedReportConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementCachedReportConfiguration
}

type ListReportCachedReportConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementCachedReportConfiguration
}

type ListReportCachedReportConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListReportCachedReportConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListReportCachedReportConfigurations ...
func (c ReportCachedReportConfigurationClient) ListReportCachedReportConfigurations(ctx context.Context) (result ListReportCachedReportConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListReportCachedReportConfigurationsCustomPager{},
		Path:       "/deviceManagement/reports/cachedReportConfigurations",
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
		Values *[]beta.DeviceManagementCachedReportConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListReportCachedReportConfigurationsComplete retrieves all the results into a single object
func (c ReportCachedReportConfigurationClient) ListReportCachedReportConfigurationsComplete(ctx context.Context) (ListReportCachedReportConfigurationsCompleteResult, error) {
	return c.ListReportCachedReportConfigurationsCompleteMatchingPredicate(ctx, DeviceManagementCachedReportConfigurationOperationPredicate{})
}

// ListReportCachedReportConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ReportCachedReportConfigurationClient) ListReportCachedReportConfigurationsCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementCachedReportConfigurationOperationPredicate) (result ListReportCachedReportConfigurationsCompleteResult, err error) {
	items := make([]beta.DeviceManagementCachedReportConfiguration, 0)

	resp, err := c.ListReportCachedReportConfigurations(ctx)
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

	result = ListReportCachedReportConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
