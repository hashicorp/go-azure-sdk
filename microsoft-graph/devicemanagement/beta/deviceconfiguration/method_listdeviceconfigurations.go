package deviceconfiguration

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

type ListDeviceConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceConfiguration
}

type ListDeviceConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceConfiguration
}

type ListDeviceConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceConfigurations ...
func (c DeviceConfigurationClient) ListDeviceConfigurations(ctx context.Context) (result ListDeviceConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceConfigurationsCustomPager{},
		Path:       "/deviceManagement/deviceConfigurations",
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
		Values *[]beta.DeviceConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceConfigurationsComplete retrieves all the results into a single object
func (c DeviceConfigurationClient) ListDeviceConfigurationsComplete(ctx context.Context) (ListDeviceConfigurationsCompleteResult, error) {
	return c.ListDeviceConfigurationsCompleteMatchingPredicate(ctx, DeviceConfigurationOperationPredicate{})
}

// ListDeviceConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceConfigurationClient) ListDeviceConfigurationsCompleteMatchingPredicate(ctx context.Context, predicate DeviceConfigurationOperationPredicate) (result ListDeviceConfigurationsCompleteResult, err error) {
	items := make([]beta.DeviceConfiguration, 0)

	resp, err := c.ListDeviceConfigurations(ctx)
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

	result = ListDeviceConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
