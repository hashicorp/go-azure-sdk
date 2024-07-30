package comanagementeligibledevice

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

type ListComanagementEligibleDevicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ComanagementEligibleDevice
}

type ListComanagementEligibleDevicesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ComanagementEligibleDevice
}

type ListComanagementEligibleDevicesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListComanagementEligibleDevicesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListComanagementEligibleDevices ...
func (c ComanagementEligibleDeviceClient) ListComanagementEligibleDevices(ctx context.Context) (result ListComanagementEligibleDevicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListComanagementEligibleDevicesCustomPager{},
		Path:       "/deviceManagement/comanagementEligibleDevices",
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
		Values *[]beta.ComanagementEligibleDevice `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComanagementEligibleDevicesComplete retrieves all the results into a single object
func (c ComanagementEligibleDeviceClient) ListComanagementEligibleDevicesComplete(ctx context.Context) (ListComanagementEligibleDevicesCompleteResult, error) {
	return c.ListComanagementEligibleDevicesCompleteMatchingPredicate(ctx, ComanagementEligibleDeviceOperationPredicate{})
}

// ListComanagementEligibleDevicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComanagementEligibleDeviceClient) ListComanagementEligibleDevicesCompleteMatchingPredicate(ctx context.Context, predicate ComanagementEligibleDeviceOperationPredicate) (result ListComanagementEligibleDevicesCompleteResult, err error) {
	items := make([]beta.ComanagementEligibleDevice, 0)

	resp, err := c.ListComanagementEligibleDevices(ctx)
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

	result = ListComanagementEligibleDevicesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
