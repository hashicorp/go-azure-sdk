package deviceenrollmentconfiguration

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

type ListDeviceEnrollmentConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceEnrollmentConfiguration
}

type ListDeviceEnrollmentConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceEnrollmentConfiguration
}

type ListDeviceEnrollmentConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceEnrollmentConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceEnrollmentConfigurations ...
func (c DeviceEnrollmentConfigurationClient) ListDeviceEnrollmentConfigurations(ctx context.Context) (result ListDeviceEnrollmentConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceEnrollmentConfigurationsCustomPager{},
		Path:       "/me/deviceEnrollmentConfigurations",
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
		Values *[]beta.DeviceEnrollmentConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceEnrollmentConfigurationsComplete retrieves all the results into a single object
func (c DeviceEnrollmentConfigurationClient) ListDeviceEnrollmentConfigurationsComplete(ctx context.Context) (ListDeviceEnrollmentConfigurationsCompleteResult, error) {
	return c.ListDeviceEnrollmentConfigurationsCompleteMatchingPredicate(ctx, DeviceEnrollmentConfigurationOperationPredicate{})
}

// ListDeviceEnrollmentConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceEnrollmentConfigurationClient) ListDeviceEnrollmentConfigurationsCompleteMatchingPredicate(ctx context.Context, predicate DeviceEnrollmentConfigurationOperationPredicate) (result ListDeviceEnrollmentConfigurationsCompleteResult, err error) {
	items := make([]beta.DeviceEnrollmentConfiguration, 0)

	resp, err := c.ListDeviceEnrollmentConfigurations(ctx)
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

	result = ListDeviceEnrollmentConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
