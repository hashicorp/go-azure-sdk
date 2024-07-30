package comanageddevice

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

type ListComanagedDevicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ManagedDevice
}

type ListComanagedDevicesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ManagedDevice
}

type ListComanagedDevicesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListComanagedDevicesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListComanagedDevices ...
func (c ComanagedDeviceClient) ListComanagedDevices(ctx context.Context) (result ListComanagedDevicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListComanagedDevicesCustomPager{},
		Path:       "/deviceManagement/comanagedDevices",
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
		Values *[]beta.ManagedDevice `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComanagedDevicesComplete retrieves all the results into a single object
func (c ComanagedDeviceClient) ListComanagedDevicesComplete(ctx context.Context) (ListComanagedDevicesCompleteResult, error) {
	return c.ListComanagedDevicesCompleteMatchingPredicate(ctx, ManagedDeviceOperationPredicate{})
}

// ListComanagedDevicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComanagedDeviceClient) ListComanagedDevicesCompleteMatchingPredicate(ctx context.Context, predicate ManagedDeviceOperationPredicate) (result ListComanagedDevicesCompleteResult, err error) {
	items := make([]beta.ManagedDevice, 0)

	resp, err := c.ListComanagedDevices(ctx)
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

	result = ListComanagedDevicesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
