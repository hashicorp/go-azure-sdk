package registereddevice

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

type ListRegisteredDevicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListRegisteredDevicesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListRegisteredDevicesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRegisteredDevicesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRegisteredDevices ...
func (c RegisteredDeviceClient) ListRegisteredDevices(ctx context.Context) (result ListRegisteredDevicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListRegisteredDevicesCustomPager{},
		Path:       "/me/registeredDevices",
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
		Values *[]beta.DirectoryObject `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRegisteredDevicesComplete retrieves all the results into a single object
func (c RegisteredDeviceClient) ListRegisteredDevicesComplete(ctx context.Context) (ListRegisteredDevicesCompleteResult, error) {
	return c.ListRegisteredDevicesCompleteMatchingPredicate(ctx, DirectoryObjectOperationPredicate{})
}

// ListRegisteredDevicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RegisteredDeviceClient) ListRegisteredDevicesCompleteMatchingPredicate(ctx context.Context, predicate DirectoryObjectOperationPredicate) (result ListRegisteredDevicesCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListRegisteredDevices(ctx)
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

	result = ListRegisteredDevicesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
